///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package smachine

import (
	"context"
	"sync"
)

type DetachableFunc func(DetachableSlotWorker)
type NonDetachableFunc func(FixedSlotWorker)

type WorkerContextMode uint8

const (
	_ WorkerContextMode = iota
	NonDetachableContext
	DetachableContext
	DetachedContext
)

type SlotWorker interface {
	//EnsureMode(expectedMode WorkerContextMode)
	AttachTo(slot *Slot, link SlotLink, wakeUpOnUse bool) (SharedAccessReport, context.CancelFunc)
	ActivateLinkedList(linkedList *Slot, hotWait bool)
	HasSignal() bool
	IsDetached() bool
}

type DetachableSlotWorker interface {
	SlotWorker

	CanLoopOrHasSignal(loopCount uint32) (canLoop, hasSignal bool)
	GetCond() (bool, *sync.Cond)
	// provides temporary protection from detach
	NonDetachableCall(NonDetachableFunc) (wasExecuted bool)
}

type FixedSlotWorker interface {
	SlotWorker
	DetachableCall(DetachableFunc) (wasDetached bool)
}
