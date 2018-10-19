/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package artifactmanager

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/insolar/insolar/core/reply"
	"github.com/insolar/insolar/inscontext"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/core/message"
	"github.com/insolar/insolar/ledger/index"
	"github.com/insolar/insolar/ledger/record"
	"github.com/insolar/insolar/ledger/storage"

	"github.com/insolar/insolar/ledger/storage/storagetest"
)

var (
	domainID  = *genRandomID(0)
	domainRef = record.Reference{
		Record: domainID,
		Domain: domainID,
	}
)

func genRandomID(pulse core.PulseNumber) *record.ID {
	id := record.ID{
		Pulse: pulse,
		Hash:  []byte{byte(rand.Int() % 256)},
	}
	zeroFilledID := record.Bytes2ID(id.CoreID()[:]) // Double conversion hack to fill missing length with zeros.
	return &zeroFilledID
}

func genRandomRef(pulse core.PulseNumber) *record.Reference {
	return &record.Reference{
		Record: *genRandomID(pulse),
		Domain: domainID,
	}
}

func genRefWithID(id *record.ID) *core.RecordRef {
	ref := record.Reference{Record: *id, Domain: domainID}
	return ref.CoreRef()
}

type preparedAMTestData struct {
	db         *storage.DB
	manager    *LedgerArtifactManager
	requestRef *record.Reference
}

type messageBusMock struct {
	handlers map[core.MessageType]core.MessageHandler
}

func NewMessageBusMock() *messageBusMock {
	return &messageBusMock{handlers: map[core.MessageType]core.MessageHandler{}}
}

func (mb *messageBusMock) Register(p core.MessageType, handler core.MessageHandler) error {
	_, ok := mb.handlers[p]
	if ok {
		return errors.New("handler for this type already exists")
	}

	mb.handlers[p] = handler
	return nil
}

func (mb *messageBusMock) MustRegister(p core.MessageType, handler core.MessageHandler) {
	err := mb.Register(p, handler)
	if err != nil {
		panic(err)
	}
}

func (mb *messageBusMock) Start(components core.Components) error {
	panic("implement me")
}

func (mb *messageBusMock) Stop() error {
	panic("implement me")
}

func (mb *messageBusMock) Send(m core.Message) (core.Reply, error) {
	typ := m.Type()
	handler, ok := mb.handlers[typ]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no handler for this message type %s", typ))
	}
	return handler(m)
}

func (mb *messageBusMock) SendAsync(m core.Message) {
	panic("implement me")
}

func prepareAMTestData(t *testing.T) (preparedAMTestData, func()) {
	db, cleaner := storagetest.TmpDB(t, "")

	mb := NewMessageBusMock()
	components := core.Components{MessageBus: mb}
	handler := MessageHandler{db: db}
	handler.Link(components)

	return preparedAMTestData{
		db: db,
		manager: &LedgerArtifactManager{
			db:                   db,
			messageBus:           mb,
			getChildrenChunkSize: 100,
		},
		requestRef: genRandomRef(0),
	}, cleaner
}

func TestLedgerArtifactManager_RegisterRequest(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	msg := message.BootstrapRequest{Name: "my little message"}
	coreID, err := td.manager.RegisterRequest(ctx, &msg)
	assert.NoError(t, err)
	id := record.Bytes2ID(coreID[:])
	rec, err := td.db.GetRecord(&id)
	assert.NoError(t, err)
	assert.Equal(t, message.MustSerializeBytes(&msg), rec.(*record.CallRequest).Payload)
}

func TestLedgerArtifactManager_DeclareType(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	typeDec := []byte{1, 2, 3}
	coreID, err := td.manager.DeclareType(ctx, *domainRef.CoreRef(), *td.requestRef.CoreRef(), typeDec)
	assert.NoError(t, err)
	id := record.Bytes2ID(coreID[:])
	typeRec, err := td.db.GetRecord(&id)
	assert.NoError(t, err)
	assert.Equal(t, &record.TypeRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		TypeDeclaration: typeDec,
	}, typeRec)
}

func TestLedgerArtifactManager_DeployCode_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	coreID, err := td.manager.DeployCode(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), []byte{1, 2, 3}, core.MachineTypeBuiltin,
	)
	assert.NoError(t, err)
	id := record.Bytes2ID(coreID[:])
	codeRec, err := td.db.GetRecord(&id)
	assert.NoError(t, err)
	assert.Equal(t, codeRec, &record.CodeRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		Code:        []byte{1, 2, 3},
		MachineType: core.MachineTypeBuiltin,
	})
}

func TestLedgerArtifactManager_ActivateClass_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	codeID, err := td.db.SetRecord(&record.CodeRecord{})
	codeRef := record.Reference{Record: *codeID, Domain: domainID}
	classRef := genRandomRef(0)
	activateCoreID, err := td.manager.ActivateClass(
		ctx,
		*domainRef.CoreRef(), *classRef.CoreRef(), *codeRef.CoreRef(),
	)
	assert.Nil(t, err)
	activateID := record.Bytes2ID(activateCoreID[:])
	activateRec, getErr := td.db.GetRecord(&activateID)
	assert.Nil(t, getErr)
	assert.Equal(t, activateRec, &record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *classRef,
		},
		ClassStateRecord: record.ClassStateRecord{
			Code: codeRef,
		},
	})
	idx, err := td.db.GetClassIndex(&classRef.Record, false)
	assert.NoError(t, err)
	assert.Equal(t, activateID, idx.LatestState)
}

func TestLedgerArtifactManager_DeactivateClass_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	_, err := td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef(),
	)
	assert.NotNil(t, err)

	notClassID, _ := td.db.SetRecord(&record.CodeRecord{})
	_, err = td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notClassID),
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_DeactivateClass_VerifiesClassIsActive(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{})
	deactivateRef, _ := td.db.SetRecord(&record.DeactivationRecord{})
	td.db.SetClassIndex(classID, &index.ClassLifeline{
		LatestState: *deactivateRef,
	})
	_, err := td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(classID),
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_DeactivateClass_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetClassIndex(classID, &index.ClassLifeline{
		LatestState: *classID,
	})

	deactivateCoreID, err := td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(classID),
	)
	assert.NoError(t, err)
	deactivateID := record.Bytes2ID(deactivateCoreID[:])
	deactivateRec, err := td.db.GetRecord(&deactivateID)
	assert.NoError(t, err)
	assert.Equal(t, deactivateRec, &record.DeactivationRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		PrevState: *classID,
	})
}

func TestLedgerArtifactManager_UpdateClass_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	_, err := td.manager.UpdateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef(), *genRandomRef(0).CoreRef(), nil,
	)
	assert.NotNil(t, err)
	notClassID, _ := td.db.SetRecord(&record.CodeRecord{})
	_, err = td.manager.UpdateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notClassID), *genRandomRef(0).CoreRef(), nil,
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_UpdateClass_VerifiesClassIsActive(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{})
	deactivateID, _ := td.db.SetRecord(&record.DeactivationRecord{})
	codeRef, _ := td.db.SetRecord(&record.CodeRecord{})
	td.db.SetClassIndex(classID, &index.ClassLifeline{
		LatestState: *deactivateID,
	})
	_, err := td.manager.UpdateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(classID), *genRefWithID(codeRef), nil)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_UpdateClass_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
	})
	td.db.SetClassIndex(classID, &index.ClassLifeline{
		LatestState: *classID,
	})
	codeID, _ := td.db.SetRecord(&record.CodeRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
		Code: []byte{1},
	})
	migrationID, err := td.db.SetRecord(&record.CodeRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
		Code: []byte{2},
	})
	migrationRefs := []record.Reference{{Domain: domainID, Record: *migrationID}}
	migrationCoreRefs := []core.RecordRef{*migrationRefs[0].CoreRef()}
	updateCoreID, err := td.manager.UpdateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(classID), *genRefWithID(codeID),
		migrationCoreRefs,
	)
	assert.NoError(t, err)
	updateID := record.Bytes2ID(updateCoreID[:])
	updateRec, getErr := td.db.GetRecord(&updateID)
	assert.Nil(t, getErr)
	assert.Equal(t, updateRec, &record.ClassAmendRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		ClassStateRecord: record.ClassStateRecord{
			Code: record.Reference{Domain: td.requestRef.Domain, Record: *codeID},
		},
		PrevState:  *classID,
		Migrations: migrationRefs,
	})
}

func TestLedgerArtifactManager_ActivateObject_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	_, err := td.manager.ActivateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef(), *genRandomRef(0).CoreRef(),
		[]byte{},
	)
	assert.NotNil(t, err)
	notClassID, _ := td.db.SetRecord(&record.ObjectActivateRecord{})
	_, err = td.manager.ActivateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notClassID), *genRandomRef(0).CoreRef(), []byte{},
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_ActivateObject_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	memory := []byte{1, 2, 3}
	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetClassIndex(classID, &index.ClassLifeline{
		LatestState: *classID,
	})
	parentID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetObjectIndex(parentID, &index.ObjectLifeline{
		ClassRef:    record.Reference{Record: *classID},
		LatestState: *parentID,
	})

	objRef := *genRandomRef(0)
	activateCoreID, err := td.manager.ActivateObject(
		ctx,
		*domainRef.CoreRef(), *objRef.CoreRef(), *genRefWithID(classID), *genRefWithID(parentID), memory,
	)
	assert.Nil(t, err)
	activateID := record.Bytes2ID(activateCoreID[:])
	activateRec, err := td.db.GetRecord(&activateID)
	assert.Nil(t, err)
	assert.Equal(t, activateRec, &record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: objRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: memory,
		},
		Class:    record.Reference{Domain: td.requestRef.Domain, Record: *classID},
		Parent:   record.Reference{Domain: td.requestRef.Domain, Record: *parentID},
		Delegate: false,
	})

	idx, err := td.db.GetObjectIndex(parentID, false)
	assert.NoError(t, err)
	childRec, err := td.db.GetRecord(idx.LatestChild)
	assert.NoError(t, err)
	assert.Equal(t, objRef, childRec.(*record.ChildRecord).Ref)

	idx, err = td.db.GetObjectIndex(&objRef.Record, false)
	assert.NoError(t, err)
	assert.Equal(t, activateID, idx.LatestState)
}

func TestLedgerArtifactManager_ActivateObjectDelegate_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	_, err := td.manager.ActivateObjectDelegate(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef(), *genRandomRef(0).CoreRef(),
		[]byte{},
	)
	assert.NotNil(t, err)
	notClassID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	_, err = td.manager.ActivateObjectDelegate(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notClassID), *genRefWithID(notClassID),
		[]byte{},
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_ActivateObjectDelegate_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	memory := []byte{1, 2, 3}
	classRef := genRandomRef(0)
	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetClassIndex(&classRef.Record, &index.ClassLifeline{
		LatestState: *classID,
	})
	parentRef := genRandomRef(0)
	parentID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetObjectIndex(&parentRef.Record, &index.ObjectLifeline{
		ClassRef:    record.Reference{Domain: td.requestRef.Domain, Record: *classID},
		LatestState: *parentID,
	})

	delegateRef := genRandomRef(0)
	activateCoreID, err := td.manager.ActivateObjectDelegate(
		ctx,
		*domainRef.CoreRef(), *delegateRef.CoreRef(), *classRef.CoreRef(), *parentRef.CoreRef(), memory,
	)
	activateID := record.Bytes2ID(activateCoreID[:])
	assert.Nil(t, err)
	activateRec, err := td.db.GetRecord(&activateID)
	assert.Nil(t, err)
	assert.Equal(t, activateRec, &record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *delegateRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: memory,
		},
		Class:    *classRef,
		Parent:   *parentRef,
		Delegate: true,
	})

	delegate, err := td.manager.GetDelegate(ctx, *parentRef.CoreRef(), *classRef.CoreRef())
	assert.NoError(t, err)
	assert.Equal(t, *delegateRef.CoreRef(), *delegate)
}

func TestLedgerArtifactManager_DeactivateObject_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	_, err := td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef())
	assert.NotNil(t, err)
	notObjID, _ := td.db.SetRecord(&record.ClassActivateRecord{})
	_, err = td.manager.DeactivateClass(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notObjID),
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_DeactivateObject_VerifiesObjectIsActive(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	objRef, _ := td.db.SetRecord(&record.ObjectActivateRecord{})
	deactivateID, _ := td.db.SetRecord(&record.DeactivationRecord{})
	td.db.SetObjectIndex(objRef, &index.ObjectLifeline{
		LatestState: *deactivateID,
	})
	_, err := td.manager.DeactivateObject(ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(deactivateID))
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_DeactivateObject_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	objID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetObjectIndex(objID, &index.ObjectLifeline{
		LatestState: *objID,
	})
	deactivateCoreID, err := td.manager.DeactivateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(objID),
	)
	assert.Nil(t, err)
	deactivateID := record.Bytes2ID(deactivateCoreID[:])
	deactivateRec, err := td.db.GetRecord(&deactivateID)
	assert.Nil(t, err)
	assert.Equal(t, deactivateRec, &record.DeactivationRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		PrevState: *objID,
	})
}

func TestLedgerArtifactManager_UpdateObject_VerifiesRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	_, err := td.manager.UpdateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRandomRef(0).CoreRef(), nil)
	assert.NotNil(t, err)
	notObjID, _ := td.db.SetRecord(&record.CodeRecord{})
	_, err = td.manager.UpdateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(notObjID), nil,
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_UpdateObject_VerifiesObjectIsActive(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	objRef, _ := td.db.SetRecord(&record.ObjectActivateRecord{})
	deactivateID, _ := td.db.SetRecord(&record.DeactivationRecord{})
	td.db.SetObjectIndex(objRef, &index.ObjectLifeline{
		LatestState: *deactivateID,
	})
	_, err := td.manager.UpdateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(deactivateID), nil,
	)
	assert.NotNil(t, err)
}

func TestLedgerArtifactManager_UpdateObject_CreatesCorrectRecord(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	ctx := inscontext.TODO()
	objID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: *genRandomRef(0),
		},
	})
	td.db.SetObjectIndex(objID, &index.ObjectLifeline{
		LatestState: *objID,
	})
	memory := []byte{1, 2, 3}
	updateCoreID, err := td.manager.UpdateObject(
		ctx,
		*domainRef.CoreRef(), *td.requestRef.CoreRef(), *genRefWithID(objID), memory)
	assert.Nil(t, err)
	updateID := record.Bytes2ID(updateCoreID[:])
	updateRec, err := td.db.GetRecord(&updateID)
	assert.Nil(t, err)
	assert.Equal(t, updateRec, &record.ObjectAmendRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: memory,
		},
		PrevState: *objID,
	})
}

func TestLedgerArtifactManager_GetClass_ReturnsCorrectDescriptors(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	codeRef := *genRandomRef(0)
	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
	})
	classAmendID, _ := td.db.SetRecord(&record.ClassAmendRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
		ClassStateRecord: record.ClassStateRecord{
			Code: codeRef,
		},
	})
	classIndex := index.ClassLifeline{
		LatestState: *classAmendID,
	}
	td.db.SetClassIndex(classID, &classIndex)

	classRef := genRefWithID(classID)
	classDesc, err := td.manager.GetClass(ctx, *classRef, nil)
	assert.NoError(t, err)
	expectedClassDesc := &ClassDescriptor{
		am:    td.manager,
		head:  *classRef,
		state: *classAmendID.CoreID(),
		code:  codeRef.CoreRef(),
	}

	assert.Equal(t, *expectedClassDesc, *classDesc.(*ClassDescriptor))
}

func TestLedgerArtifactManager_GetObject_VerifiesRecords(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	objID := genRandomID(0)
	_, err := td.manager.GetObject(ctx, *genRefWithID(objID), nil)
	assert.NotNil(t, err)

	deactivateID, _ := td.db.SetRecord(&record.DeactivationRecord{})
	objectIndex := index.ObjectLifeline{
		LatestState: *deactivateID,
		ClassRef:    *genRandomRef(0),
	}
	td.db.SetObjectIndex(objID, &objectIndex)

	_, err = td.manager.GetObject(ctx, *genRefWithID(objID), nil)
	assert.Equal(t, core.ErrDeactivated, err)
}

func TestLedgerArtifactManager_GetLatestObj_ReturnsCorrectDescriptors(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	classID, _ := td.db.SetRecord(&record.ClassActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
	})
	classAmendRef, _ := td.db.SetRecord(&record.ClassAmendRecord{
		ResultRecord: record.ResultRecord{
			Domain:  domainRef,
			Request: *td.requestRef,
		},
	})
	classIndex := index.ClassLifeline{
		LatestState: *classAmendRef,
	}
	td.db.SetClassIndex(classID, &classIndex)

	objectID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: []byte{3},
		},
	})
	objectAmendID, _ := td.db.SetRecord(&record.ObjectAmendRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: []byte{4},
		},
	})
	objectIndex := index.ObjectLifeline{
		LatestState: *objectAmendID,
		ClassRef:    record.Reference{Domain: td.requestRef.Domain, Record: *classID},
	}
	td.db.SetObjectIndex(objectID, &objectIndex)

	objDesc, err := td.manager.GetObject(ctx, *genRefWithID(objectID), nil)
	assert.NoError(t, err)
	expectedObjDesc := &ObjectDescriptor{
		am: td.manager,

		head:   *getReference(td.requestRef.CoreRef(), objectID),
		state:  *objectAmendID.CoreID(),
		class:  *getReference(td.requestRef.CoreRef(), classID),
		memory: []byte{4},
	}

	assert.Equal(t, *expectedObjDesc, *objDesc.(*ObjectDescriptor))
}

func TestLedgerArtifactManager_GetChildren(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()
	ctx := inscontext.TODO()

	parentID, _ := td.db.SetRecord(&record.ObjectActivateRecord{
		ResultRecord: record.ResultRecord{
			Domain: domainRef,
		},
		ObjectStateRecord: record.ObjectStateRecord{
			Memory: []byte{0},
		},
	})
	child1Ref := genRandomRef(1)
	child2Ref := genRandomRef(1)
	child3Ref := genRandomRef(2)

	childMeta1, _ := td.db.SetRecord(&record.ChildRecord{
		Ref: *child1Ref,
	})
	childMeta2, _ := td.db.SetRecord(&record.ChildRecord{
		PrevChild: childMeta1,
		Ref:       *child2Ref,
	})
	childMeta3, _ := td.db.SetRecord(&record.ChildRecord{
		PrevChild: childMeta2,
		Ref:       *child3Ref,
	})

	parentIndex := index.ObjectLifeline{
		LatestState: *parentID,
		LatestChild: childMeta3,
	}
	td.db.SetObjectIndex(parentID, &parentIndex)

	t.Run("returns correct children without pulse", func(t *testing.T) {
		i, err := td.manager.GetChildren(ctx, *genRefWithID(parentID), nil)
		assert.NoError(t, err)
		child, err := i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child3Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child2Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child1Ref.CoreRef(), *child)
		hasNext := i.HasNext()
		assert.False(t, hasNext)
		_, err = i.Next()
		assert.Error(t, err)
	})

	t.Run("returns correct children with pulse", func(t *testing.T) {
		pn := core.PulseNumber(1)
		i, err := td.manager.GetChildren(ctx, *genRefWithID(parentID), &pn)
		assert.NoError(t, err)
		child, err := i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child2Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child1Ref.CoreRef(), *child)
		hasNext := i.HasNext()
		assert.NoError(t, err)
		assert.False(t, hasNext)
		_, err = i.Next()
		assert.Error(t, err)
	})

	t.Run("returns correct children in many chunks", func(t *testing.T) {
		td.manager.getChildrenChunkSize = 1
		i, err := td.manager.GetChildren(ctx, *genRefWithID(parentID), nil)
		assert.NoError(t, err)
		child, err := i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child3Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child2Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child1Ref.CoreRef(), *child)
		hasNext := i.HasNext()
		assert.NoError(t, err)
		assert.False(t, hasNext)
		_, err = i.Next()
		assert.Error(t, err)
	})

	t.Run("doesn't fail when has no children to return", func(t *testing.T) {
		td.manager.getChildrenChunkSize = 1
		pn := core.PulseNumber(3)
		i, err := td.manager.GetChildren(ctx, *genRefWithID(parentID), &pn)
		assert.NoError(t, err)
		child, err := i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child3Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child2Ref.CoreRef(), *child)
		child, err = i.Next()
		assert.NoError(t, err)
		assert.Equal(t, *child1Ref.CoreRef(), *child)
		hasNext := i.HasNext()
		assert.NoError(t, err)
		assert.False(t, hasNext)
		_, err = i.Next()
		assert.Error(t, err)
	})
}

func TestLedgerArtifactManager_HandleJetDrop(t *testing.T) {
	t.Parallel()
	td, cleaner := prepareAMTestData(t)
	defer cleaner()

	records := []record.ObjectActivateRecord{
		{ObjectStateRecord: record.ObjectStateRecord{Memory: []byte{1}}},
		{ObjectStateRecord: record.ObjectStateRecord{Memory: []byte{2}}},
		{ObjectStateRecord: record.ObjectStateRecord{Memory: []byte{3}}},
	}
	ids := []record.ID{
		{Hash: []byte{4}},
		{Hash: []byte{5}},
		{Hash: []byte{6}},
	}
	recordData := [][2][]byte{
		{record.ID2Bytes(ids[0]), record.SerializeRecord(&records[0])},
		{record.ID2Bytes(ids[1]), record.SerializeRecord(&records[1])},
		{record.ID2Bytes(ids[2]), record.SerializeRecord(&records[2])},
	}

	rep, err := td.manager.messageBus.Send(&message.JetDrop{
		Records: recordData,
	})
	assert.NoError(t, err)
	assert.Equal(t, reply.OK{}, *rep.(*reply.OK))

	for i := 0; i < len(records); i++ {
		rec, err := td.db.GetRecord(&ids[i])
		assert.NoError(t, err)
		assert.Equal(t, records[i], *rec.(*record.ObjectActivateRecord))
	}
}
