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

package configuration

import (
	"time"
)

// Storage configures Ledger's storage.
type Storage struct {
	// DataDirectory is a directory where database's files live.
	DataDirectory string
	// TxRetriesOnConflict defines how many retries on transaction conflicts
	// storage update methods should do.
	TxRetriesOnConflict int
}

// JetSplit holds configuration for jet split.
type JetSplit struct {
	// RecordsCountThreshold is a drop threshold in records to perform split for jet.
	ThresholdRecordsCount int
	// ThresholdOverflowCount is a how many times in row ThresholdRecordsCount should be surpassed.
	ThresholdOverflowCount int
	// DepthLimit limits jet tree depth (maximum possible jets = 2^DepthLimit)
	DepthLimit uint8
}

// Backoff configures retry backoff algorithm
type Backoff struct {
	Factor float64
	// Jitter eases contention by randomizing backoff steps
	Jitter bool
	// Min and Max are the minimum and maximum values of the counter
	Min, Max time.Duration
	// MaxAttempts holds max count of attempts for a instance of Backoff
	MaxAttempts int
}

// Exporter holds configuration of Exporter
type Exporter struct {
	// ExportLag is lag in second before we start to export pulse
	ExportLag uint32
}

// Ledger holds configuration for ledger.
type Ledger struct {
	// Storage defines storage configuration.
	Storage Storage
	// JetSplit holds jet split configuration.
	JetSplit JetSplit

	// common/sharable values:

	// LightChainLimit is maximum pulse difference (NOT number of pulses)
	// between current and the latest replicated on heavy.
	//
	// IMPORTANT: It should be the same on ALL nodes.
	LightChainLimit int

	// Exporter holds configuration of Exporter
	Exporter Exporter

	// Backup holds configuration of BackupMaker
	Backup Backup
}

// Backup holds configuration for backuping.
type Backup struct {
	// Enabled switches on backuping
	Enabled bool

	// TmpDirectory is directory for tmp storage of backup data. Must be created
	TmpDirectory string

	// TargetDirectory is directory where backups will be moved to
	TargetDirectory string

	// MetaInfoFile contains meta info about backup
	MetaInfoFile string

	// ConfirmFile: we wait this file being created when backup was saved on remote host
	ConfirmFile string

	// BackupFile is file with incremental backup data
	BackupFile string

	// DirNameTemplate is template for saving current incremental backup. Should be like "pulse-%d"
	DirNameTemplate string

	// BackupWaitPeriod - how much time we will wait for appearing of file ConfirmFile
	BackupWaitPeriod uint

	// Paths:
	// Every incremental backup live in  TargetDirectory/"DirNameTemplate"%<pulse_number>
	// and it contains:
	// incr.bkp - backup file
	// MetaInfoFile - meta info about current backup
	// ConfirmFile - must be set from inside. When it appear it means that we successfully saved the backup
}

// NewLedger creates new default Ledger configuration.
func NewLedger() Ledger {
	return Ledger{
		Storage: Storage{
			DataDirectory:       "./data",
			TxRetriesOnConflict: 3,
		},

		JetSplit: JetSplit{
			// TODO: find best default values
			ThresholdRecordsCount:  100,
			ThresholdOverflowCount: 3,
			DepthLimit:             10, // limit to 1024 jets
		},
		LightChainLimit: 5, // 5 pulses

		Exporter: Exporter{
			ExportLag: 40, // 40 seconds
		},

		Backup: Backup{
			Enabled: false,
		},
	}
}
