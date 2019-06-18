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

package deposit

import (
	"fmt"
	"math/big"
	"time"

	"github.com/insolar/insolar/insolar"

	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

type DepositStatus string

const (
	DepositConfirms uint = 3

	Open    DepositStatus = "Open"
	Holding DepositStatus = "Holding"
	Close   DepositStatus = "Close"
)

type Deposit struct {
	foundation.BaseContract
	Timestamp               time.Time
	HoldReleaseDate         time.Time
	MigrationDaemonConfirms map[insolar.Reference]bool
	Confirms                uint
	Amount                  string
	Bonus                   string
	TxHash                  string
	Status                  DepositStatus
}

func (d *Deposit) GetTxHash() (string, error) {
	return d.TxHash, nil
}

func (d *Deposit) GetAmount() (string, error) {
	return d.Amount, nil
}

func New(migrationDaemonConfirms map[insolar.Reference]bool, txHash string, amount string, holdReleaseDate time.Time) (*Deposit, error) {
	return &Deposit{

		MigrationDaemonConfirms: migrationDaemonConfirms,
		Confirms:                0,
		TxHash:                  txHash,
		HoldReleaseDate:         holdReleaseDate,
		Amount:                  amount,
		Status:                  Open,
	}, nil
}

func (d *Deposit) MapMarshal() (map[string]string, error) {
	return map[string]string{
		"timestamp":       d.Timestamp.String(),
		"holdReleaseDate": d.HoldReleaseDate.String(),
		"amount":          d.Amount,
		"bonus":           d.Bonus,
		"txId":            d.TxHash,
	}, nil
}

func (d *Deposit) Confirm(migrationDaemon insolar.Reference, txHash string, amountStr string) (uint, error) {
	if txHash != d.TxHash {
		return 0, fmt.Errorf("[ Confirm ] Transaction hash is incorrect")
	}

	inputAmount := new(big.Int)
	inputAmount, ok := inputAmount.SetString(amountStr, 10)
	if !ok {
		return 0, fmt.Errorf("[ Confirm ] can't parse input amount")
	}
	depositAmount := new(big.Int)
	depositAmount, ok = depositAmount.SetString(d.Amount, 10)
	if !ok {
		return 0, fmt.Errorf("[ Confirm ] can't parse Deposit amount")
	}

	if (inputAmount).Cmp(depositAmount) != 0 {
		return 0, fmt.Errorf("[ Confirm ] Amount is incorrect")
	}

	if confirm, ok := d.MigrationDaemonConfirms[migrationDaemon]; ok {
		if confirm {
			return 0, fmt.Errorf("[ Confirm ] Confirm from the migration daemon " + migrationDaemon.String() + " already exists")
		} else {
			d.MigrationDaemonConfirms[migrationDaemon] = true
			d.Confirms++
			if d.Confirms == DepositConfirms {
				d.Status = Holding
			}
			return d.Confirms, nil
		}
	} else {
		return 0, fmt.Errorf("[ Confirm ] Migration daemon name is incorrect")
	}
}
