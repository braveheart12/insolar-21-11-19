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

package record

import (
	"io"

	"github.com/insolar/insolar/core"
)

// ClassState is common class state record.
type ClassState interface {
	// IsDeactivation determines if current state is deactivation.
	IsDeactivation() bool
	// GetCode returns state code.
	GetCode() *Reference
	// GetMachineType returns state code machine type.
	GetMachineType() core.MachineType
}

// ObjectState is common object state record.
type ObjectState interface {
	// IsDeactivation determines if current state is deactivation.
	IsDeactivation() bool
	// GetMemory returns state memory.
	GetMemory() []byte
}

// ResultRecord is a record which is created in response to a request.
type ResultRecord struct {
	Domain  Reference
	Request Reference
}

// TypeRecord is a code interface declaration.
type TypeRecord struct {
	ResultRecord

	TypeDeclaration []byte
}

// Type implementation of Record interface.
func (r *TypeRecord) Type() TypeID { return typeType }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *TypeRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// CodeRecord is a code storage record.
type CodeRecord struct {
	ResultRecord

	Code        []byte
	MachineType core.MachineType
}

// Type implementation of Record interface.
func (r *CodeRecord) Type() TypeID { return typeCode }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *CodeRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// ClassStateRecord is a record containing data for a class state.
type ClassStateRecord struct {
	Code        Reference
	MachineType core.MachineType
}

// GetMachineType returns state code machine type.
func (r *ClassStateRecord) GetMachineType() core.MachineType {
	return r.MachineType
}

// GetCode returns state code.
func (r *ClassStateRecord) GetCode() *Reference {
	return &r.Code
}

// IsDeactivation determines if current state is deactivation.
func (r *ClassStateRecord) IsDeactivation() bool {
	return false
}

// ClassActivateRecord is produced when we "activate" new contract class.
type ClassActivateRecord struct {
	ResultRecord
	ClassStateRecord
}

// Type implementation of Record interface.
func (r *ClassActivateRecord) Type() TypeID { return typeClassActivate }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *ClassActivateRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// ClassAmendRecord is an amendment record for classes.
type ClassAmendRecord struct {
	ResultRecord
	ClassStateRecord

	PrevState ID
}

// Type implementation of Record interface.
func (r *ClassAmendRecord) Type() TypeID { return typeClassAmend }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *ClassAmendRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// ObjectStateRecord is a record containing data for an object state.
type ObjectStateRecord struct {
	Memory []byte
}

// IsDeactivation determines if current state is deactivation.
func (r *ObjectStateRecord) IsDeactivation() bool {
	return false
}

// GetMemory returns state memory.
func (r *ObjectStateRecord) GetMemory() []byte {
	return r.Memory
}

// ObjectActivateRecord is produced when we instantiate new object from an available class.
type ObjectActivateRecord struct {
	ResultRecord
	ObjectStateRecord

	Class    Reference
	Parent   Reference
	Delegate bool
}

// Type implementation of Record interface.
func (r *ObjectActivateRecord) Type() TypeID { return typeObjectActivate }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *ObjectActivateRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// ObjectAmendRecord is an amendment record for objects.
type ObjectAmendRecord struct {
	ResultRecord
	ObjectStateRecord

	PrevState ID
}

// Type implementation of Record interface.
func (r *ObjectAmendRecord) Type() TypeID { return typeObjectAmend }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *ObjectAmendRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// DeactivationRecord marks targeted object as disabled.
type DeactivationRecord struct {
	ResultRecord
	PrevState ID
}

// Type implementation of Record interface.
func (r *DeactivationRecord) Type() TypeID { return typeDeactivate }

// WriteHashData writes record data to provided writer. This data is used to calculate record's hash.
func (r *DeactivationRecord) WriteHashData(w io.Writer) (int, error) {
	return w.Write(SerializeRecord(r))
}

// GetMachineType returns state code machine type.
func (*DeactivationRecord) GetMachineType() core.MachineType {
	return core.MachineTypeNotExist
}

// IsDeactivation determines if current state is deactivation.
func (*DeactivationRecord) IsDeactivation() bool {
	return true
}

// IsAmend determines if current state is amend.
func (*DeactivationRecord) IsAmend() bool {
	return false
}

// GetMemory returns state memory.
func (*DeactivationRecord) GetMemory() []byte {
	return nil
}

// GetCode returns state code.
func (*DeactivationRecord) GetCode() *Reference {
	return nil
}
