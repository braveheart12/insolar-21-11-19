//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package object

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/insolar/store"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/pkg/errors"
)

// RecordDB is a DB storage implementation. It saves records to disk and does not allow removal.
type RecordDB struct {
	batchLock sync.Mutex

	db *store.BadgerDB
}

type recordKey insolar.ID

func (k recordKey) Scope() store.Scope {
	return store.ScopeRecord
}

func (k recordKey) String() string {
	id := insolar.ID(k)
	return "recordKey. " + id.DebugString()
}

func (k recordKey) ID() []byte {
	id := insolar.ID(k)
	return bytes.Join([][]byte{id.Pulse().Bytes(), id.Hash()}, nil)
}

func newRecordKey(raw []byte) recordKey {
	pulse := insolar.NewPulseNumber(raw)
	hash := raw[pulse.Size():]

	return recordKey(*insolar.NewID(pulse, hash))
}

const (
	recordPositionKeyPrefix          = 0x01
	lastKnownRecordPositionKeyPrefix = 0x02
)

type recordPositionKey struct {
	pn     insolar.PulseNumber
	number uint32
}

func newRecordPositionKey(pn insolar.PulseNumber, number uint32) recordPositionKey {
	return recordPositionKey{pn: pn, number: number}
}

func (k recordPositionKey) Scope() store.Scope {
	return store.ScopeRecordPosition
}

func (k recordPositionKey) ID() []byte {
	parsedNum := make([]byte, 4)
	binary.BigEndian.PutUint32(parsedNum, k.number)
	return bytes.Join([][]byte{{recordPositionKeyPrefix}, k.pn.Bytes(), parsedNum}, nil)
}

func newRecordPositionKeyFromBytes(raw []byte) recordPositionKey {
	k := recordPositionKey{}

	k.pn = insolar.NewPulseNumber(raw[1:])
	k.number = binary.BigEndian.Uint32(raw[(k.pn.Size() + 1):])

	return k
}

func (k recordPositionKey) String() string {
	return fmt.Sprintf("recordPositionKey. pulse: %d, number: %d", k.pn, k.number)
}

type lastKnownRecordPositionKey struct {
	pn insolar.PulseNumber
}

func newLastKnownRecordPositionKey(raw []byte) lastKnownRecordPositionKey {
	k := lastKnownRecordPositionKey{}
	k.pn = insolar.NewPulseNumber(raw[1:])
	return k
}

func (k lastKnownRecordPositionKey) String() string {
	return fmt.Sprintf("lastKnownRecordPositionKey. pulse: %d", k.pn)
}

func (k lastKnownRecordPositionKey) Scope() store.Scope {
	return store.ScopeRecordPosition
}

func (k lastKnownRecordPositionKey) ID() []byte {
	return bytes.Join([][]byte{{lastKnownRecordPositionKeyPrefix}, k.pn.Bytes()}, nil)
}

// NewRecordDB creates new DB storage instance.
func NewRecordDB(db *store.BadgerDB) *RecordDB {
	return &RecordDB{db: db}
}

// Set saves new record-value in storage.
func (r *RecordDB) Set(ctx context.Context, rec record.Material) error {
	return r.BatchSet(ctx, []record.Material{rec})
}

// BatchSet saves a batch of records to storage with order-processing.
func (r *RecordDB) BatchSet(ctx context.Context, recs []record.Material) error {
	if len(recs) == 0 {
		return nil
	}

	r.batchLock.Lock()
	defer r.batchLock.Unlock()

	// It's possible, that in the batch can be records from different pulses
	// because of that we need to track a current pulse and position
	// for different pulses position is requested from db
	// We can get position on every loop, but we SHOULDN'T do this
	// Because it isn't efficient
	lastKnowPulse := insolar.PulseNumber(0)
	position := uint32(0)

	err := r.db.Backend().Update(func(txn *badger.Txn) error {
		for _, rec := range recs {
			if rec.ID.IsEmpty() {
				return errors.New("id is empty")
			}

			err := setRecord(txn, recordKey(rec.ID), rec)
			if err != nil {
				return err
			}

			// For cross-pulse batches
			if lastKnowPulse != rec.ID.Pulse() {
				// Set last known before changing pulse/position
				err := setLastKnownPosition(txn, lastKnowPulse, position)
				if err != nil {
					return err
				}

				// fetch position for a new pulse
				position, err = getLastKnownPosition(txn, rec.ID.Pulse())
				if err != nil && err != ErrNotFound {
					return err
				}
				lastKnowPulse = rec.ID.Pulse()
			}

			position++

			err = setPosition(txn, rec.ID, position)
			if err != nil {
				return err
			}

		}

		// set position for last record
		err := setLastKnownPosition(txn, lastKnowPulse, position)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// setRecord is a helper method for storaging record to db in scope of txn.
func setRecord(txn *badger.Txn, key store.Key, record record.Material) error {
	data, err := record.Marshal()
	if err != nil {
		return err
	}

	fullKey := append(key.Scope().Bytes(), key.ID()...)

	_, err = txn.Get(fullKey)
	if err != nil && err != badger.ErrKeyNotFound {
		return err
	}
	if err == nil {
		return ErrOverride
	}

	return txn.Set(fullKey, data)
}

// setRecord is a helper method for getting last known position of record to db in scope of txn and pulse.
func getLastKnownPosition(txn *badger.Txn, pn insolar.PulseNumber) (uint32, error) {
	key := lastKnownRecordPositionKey{pn: pn}

	fullKey := append(key.Scope().Bytes(), key.ID()...)

	item, err := txn.Get(fullKey)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return 0, ErrNotFound
		}
		return 0, err
	}

	buff, err := item.ValueCopy(nil)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(buff), nil
}

// setLastKnownPosition is a helper method for setting last known position of record to db in scope of txn and pulse.
func setLastKnownPosition(txn *badger.Txn, pn insolar.PulseNumber, position uint32) error {
	lastPositionKey := lastKnownRecordPositionKey{pn: pn}
	parsedPosition := make([]byte, 4)
	binary.BigEndian.PutUint32(parsedPosition, position)

	fullKey := append(lastPositionKey.Scope().Bytes(), lastPositionKey.ID()...)

	return txn.Set(fullKey, parsedPosition)
}

func setPosition(txn *badger.Txn, recID insolar.ID, position uint32) error {
	positionKey := newRecordPositionKey(recID.Pulse(), position)
	fullKey := append(positionKey.Scope().Bytes(), positionKey.ID()...)

	return txn.Set(fullKey, recID.Bytes())
}

func (r *RecordDB) truncateHead(ctx context.Context, from printableKey, km keyMaker) error {
	it := store.NewReadIterator(r.db.Backend(), from, false)
	defer it.Close()

	var hasKeys bool
	for it.Next() {
		hasKeys = true
		key := km(it.Key())

		err := r.db.Delete(key)
		if err != nil {
			return errors.Wrapf(err, "can't delete key: %s", key.String())
		}

		inslogger.FromContext(ctx).Debugf("Erased key: %s", key.String())
	}

	if !hasKeys {
		inslogger.FromContext(ctx).Infof("No records. Nothing done. Start key: %s", from.String())
	}

	return nil
}

type printableKey interface {
	store.Key
	String() string
}

type keyMaker func([]byte) printableKey

func recordKeyMaker(raw []byte) printableKey {
	return newRecordKey(raw)
}

func positionKeyMaker(raw []byte) printableKey {
	switch raw[0] {
	case recordPositionKeyPrefix:
		return newRecordPositionKeyFromBytes(raw)
	case lastKnownRecordPositionKeyPrefix:
		return newLastKnownRecordPositionKey(raw)
	default:
		panic("unknown prefix: " + string(raw[0]))
	}
}

// TruncateHead remove all records after lastPulse
func (r *RecordDB) TruncateHead(ctx context.Context, from insolar.PulseNumber) error {

	if err := r.truncateHead(ctx, recordKey(*insolar.NewID(from, nil)), recordKeyMaker); err != nil {
		return errors.Wrap(err, "failed to truncate records head")
	}

	if err := r.truncateHead(ctx, recordPositionKey{pn: from}, positionKeyMaker); err != nil {
		return errors.Wrap(err, "failed to truncate record positions head")
	}

	if err := r.truncateHead(ctx, lastKnownRecordPositionKey{pn: from}, positionKeyMaker); err != nil {
		return errors.Wrap(err, "failed to truncate last known record positions head")
	}

	return nil
}

// ForID returns record for provided id.
func (r *RecordDB) ForID(ctx context.Context, id insolar.ID) (record.Material, error) {
	return r.get(id)
}

// get loads record.Material from DB
func (r *RecordDB) get(id insolar.ID) (record.Material, error) {
	var buff []byte
	var err error
	err = r.db.Backend().View(func(txn *badger.Txn) error {
		key := recordKey(id)
		fullKey := append(key.Scope().Bytes(), key.ID()...)

		item, err := txn.Get(fullKey)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrNotFound
			}
			return err
		}

		buff, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return record.Material{}, err
	}

	rec := record.Material{}
	err = rec.Unmarshal(buff)

	return rec, err
}

// LastKnownPosition returns last known position of record in Pulse.
func (r *RecordDB) LastKnownPosition(pn insolar.PulseNumber) (uint32, error) {
	var position uint32
	var err error

	err = r.db.Backend().View(func(txn *badger.Txn) error {
		position, err = getLastKnownPosition(txn, pn)
		return err
	})

	return position, err
}

// AtPosition returns record ID for a specific pulse and a position
func (r *RecordDB) AtPosition(pn insolar.PulseNumber, position uint32) (insolar.ID, error) {
	var recID insolar.ID
	err := r.db.Backend().View(func(txn *badger.Txn) error {
		lastKnownPosition, err := getLastKnownPosition(txn, pn)
		if err != nil {
			return err
		}

		if position > lastKnownPosition {
			return ErrNotFound
		}
		positionKey := newRecordPositionKey(pn, position)
		fullKey := append(positionKey.Scope().Bytes(), positionKey.ID()...)

		item, err := txn.Get(fullKey)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return ErrNotFound
			}
			return err
		}
		rawID, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		recID = *insolar.NewIDFromBytes(rawID)
		return nil
	})
	return recID, err
}
