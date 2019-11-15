package merkle

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"crypto"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_merkle "github.com/insolar/insolar/network/merkle"
)

// CalculatorMock implements merkle.Calculator
type CalculatorMock struct {
	t minimock.Tester

	funcGetCloudProof          func(cp1 *mm_merkle.CloudEntry) (o1 mm_merkle.OriginHash, cp2 *mm_merkle.CloudProof, err error)
	inspectFuncGetCloudProof   func(cp1 *mm_merkle.CloudEntry)
	afterGetCloudProofCounter  uint64
	beforeGetCloudProofCounter uint64
	GetCloudProofMock          mCalculatorMockGetCloudProof

	funcGetGlobuleProof          func(gp1 *mm_merkle.GlobuleEntry) (o1 mm_merkle.OriginHash, gp2 *mm_merkle.GlobuleProof, err error)
	inspectFuncGetGlobuleProof   func(gp1 *mm_merkle.GlobuleEntry)
	afterGetGlobuleProofCounter  uint64
	beforeGetGlobuleProofCounter uint64
	GetGlobuleProofMock          mCalculatorMockGetGlobuleProof

	funcGetPulseProof          func(pp1 *mm_merkle.PulseEntry) (o1 mm_merkle.OriginHash, pp2 *mm_merkle.PulseProof, err error)
	inspectFuncGetPulseProof   func(pp1 *mm_merkle.PulseEntry)
	afterGetPulseProofCounter  uint64
	beforeGetPulseProofCounter uint64
	GetPulseProofMock          mCalculatorMockGetPulseProof

	funcIsValid          func(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey) (b1 bool)
	inspectFuncIsValid   func(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey)
	afterIsValidCounter  uint64
	beforeIsValidCounter uint64
	IsValidMock          mCalculatorMockIsValid
}

// NewCalculatorMock returns a mock for merkle.Calculator
func NewCalculatorMock(t minimock.Tester) *CalculatorMock {
	m := &CalculatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetCloudProofMock = mCalculatorMockGetCloudProof{mock: m}
	m.GetCloudProofMock.callArgs = []*CalculatorMockGetCloudProofParams{}

	m.GetGlobuleProofMock = mCalculatorMockGetGlobuleProof{mock: m}
	m.GetGlobuleProofMock.callArgs = []*CalculatorMockGetGlobuleProofParams{}

	m.GetPulseProofMock = mCalculatorMockGetPulseProof{mock: m}
	m.GetPulseProofMock.callArgs = []*CalculatorMockGetPulseProofParams{}

	m.IsValidMock = mCalculatorMockIsValid{mock: m}
	m.IsValidMock.callArgs = []*CalculatorMockIsValidParams{}

	return m
}

type mCalculatorMockGetCloudProof struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockGetCloudProofExpectation
	expectations       []*CalculatorMockGetCloudProofExpectation

	callArgs []*CalculatorMockGetCloudProofParams
	mutex    sync.RWMutex
}

// CalculatorMockGetCloudProofExpectation specifies expectation struct of the Calculator.GetCloudProof
type CalculatorMockGetCloudProofExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockGetCloudProofParams
	results *CalculatorMockGetCloudProofResults
	Counter uint64
}

// CalculatorMockGetCloudProofParams contains parameters of the Calculator.GetCloudProof
type CalculatorMockGetCloudProofParams struct {
	cp1 *mm_merkle.CloudEntry
}

// CalculatorMockGetCloudProofResults contains results of the Calculator.GetCloudProof
type CalculatorMockGetCloudProofResults struct {
	o1  mm_merkle.OriginHash
	cp2 *mm_merkle.CloudProof
	err error
}

// Expect sets up expected params for Calculator.GetCloudProof
func (mmGetCloudProof *mCalculatorMockGetCloudProof) Expect(cp1 *mm_merkle.CloudEntry) *mCalculatorMockGetCloudProof {
	if mmGetCloudProof.mock.funcGetCloudProof != nil {
		mmGetCloudProof.mock.t.Fatalf("CalculatorMock.GetCloudProof mock is already set by Set")
	}

	if mmGetCloudProof.defaultExpectation == nil {
		mmGetCloudProof.defaultExpectation = &CalculatorMockGetCloudProofExpectation{}
	}

	mmGetCloudProof.defaultExpectation.params = &CalculatorMockGetCloudProofParams{cp1}
	for _, e := range mmGetCloudProof.expectations {
		if minimock.Equal(e.params, mmGetCloudProof.defaultExpectation.params) {
			mmGetCloudProof.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCloudProof.defaultExpectation.params)
		}
	}

	return mmGetCloudProof
}

// Inspect accepts an inspector function that has same arguments as the Calculator.GetCloudProof
func (mmGetCloudProof *mCalculatorMockGetCloudProof) Inspect(f func(cp1 *mm_merkle.CloudEntry)) *mCalculatorMockGetCloudProof {
	if mmGetCloudProof.mock.inspectFuncGetCloudProof != nil {
		mmGetCloudProof.mock.t.Fatalf("Inspect function is already set for CalculatorMock.GetCloudProof")
	}

	mmGetCloudProof.mock.inspectFuncGetCloudProof = f

	return mmGetCloudProof
}

// Return sets up results that will be returned by Calculator.GetCloudProof
func (mmGetCloudProof *mCalculatorMockGetCloudProof) Return(o1 mm_merkle.OriginHash, cp2 *mm_merkle.CloudProof, err error) *CalculatorMock {
	if mmGetCloudProof.mock.funcGetCloudProof != nil {
		mmGetCloudProof.mock.t.Fatalf("CalculatorMock.GetCloudProof mock is already set by Set")
	}

	if mmGetCloudProof.defaultExpectation == nil {
		mmGetCloudProof.defaultExpectation = &CalculatorMockGetCloudProofExpectation{mock: mmGetCloudProof.mock}
	}
	mmGetCloudProof.defaultExpectation.results = &CalculatorMockGetCloudProofResults{o1, cp2, err}
	return mmGetCloudProof.mock
}

//Set uses given function f to mock the Calculator.GetCloudProof method
func (mmGetCloudProof *mCalculatorMockGetCloudProof) Set(f func(cp1 *mm_merkle.CloudEntry) (o1 mm_merkle.OriginHash, cp2 *mm_merkle.CloudProof, err error)) *CalculatorMock {
	if mmGetCloudProof.defaultExpectation != nil {
		mmGetCloudProof.mock.t.Fatalf("Default expectation is already set for the Calculator.GetCloudProof method")
	}

	if len(mmGetCloudProof.expectations) > 0 {
		mmGetCloudProof.mock.t.Fatalf("Some expectations are already set for the Calculator.GetCloudProof method")
	}

	mmGetCloudProof.mock.funcGetCloudProof = f
	return mmGetCloudProof.mock
}

// When sets expectation for the Calculator.GetCloudProof which will trigger the result defined by the following
// Then helper
func (mmGetCloudProof *mCalculatorMockGetCloudProof) When(cp1 *mm_merkle.CloudEntry) *CalculatorMockGetCloudProofExpectation {
	if mmGetCloudProof.mock.funcGetCloudProof != nil {
		mmGetCloudProof.mock.t.Fatalf("CalculatorMock.GetCloudProof mock is already set by Set")
	}

	expectation := &CalculatorMockGetCloudProofExpectation{
		mock:   mmGetCloudProof.mock,
		params: &CalculatorMockGetCloudProofParams{cp1},
	}
	mmGetCloudProof.expectations = append(mmGetCloudProof.expectations, expectation)
	return expectation
}

// Then sets up Calculator.GetCloudProof return parameters for the expectation previously defined by the When method
func (e *CalculatorMockGetCloudProofExpectation) Then(o1 mm_merkle.OriginHash, cp2 *mm_merkle.CloudProof, err error) *CalculatorMock {
	e.results = &CalculatorMockGetCloudProofResults{o1, cp2, err}
	return e.mock
}

// GetCloudProof implements merkle.Calculator
func (mmGetCloudProof *CalculatorMock) GetCloudProof(cp1 *mm_merkle.CloudEntry) (o1 mm_merkle.OriginHash, cp2 *mm_merkle.CloudProof, err error) {
	mm_atomic.AddUint64(&mmGetCloudProof.beforeGetCloudProofCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCloudProof.afterGetCloudProofCounter, 1)

	if mmGetCloudProof.inspectFuncGetCloudProof != nil {
		mmGetCloudProof.inspectFuncGetCloudProof(cp1)
	}

	mm_params := &CalculatorMockGetCloudProofParams{cp1}

	// Record call args
	mmGetCloudProof.GetCloudProofMock.mutex.Lock()
	mmGetCloudProof.GetCloudProofMock.callArgs = append(mmGetCloudProof.GetCloudProofMock.callArgs, mm_params)
	mmGetCloudProof.GetCloudProofMock.mutex.Unlock()

	for _, e := range mmGetCloudProof.GetCloudProofMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.o1, e.results.cp2, e.results.err
		}
	}

	if mmGetCloudProof.GetCloudProofMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCloudProof.GetCloudProofMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCloudProof.GetCloudProofMock.defaultExpectation.params
		mm_got := CalculatorMockGetCloudProofParams{cp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCloudProof.t.Errorf("CalculatorMock.GetCloudProof got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCloudProof.GetCloudProofMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCloudProof.t.Fatal("No results are set for the CalculatorMock.GetCloudProof")
		}
		return (*mm_results).o1, (*mm_results).cp2, (*mm_results).err
	}
	if mmGetCloudProof.funcGetCloudProof != nil {
		return mmGetCloudProof.funcGetCloudProof(cp1)
	}
	mmGetCloudProof.t.Fatalf("Unexpected call to CalculatorMock.GetCloudProof. %v", cp1)
	return
}

// GetCloudProofAfterCounter returns a count of finished CalculatorMock.GetCloudProof invocations
func (mmGetCloudProof *CalculatorMock) GetCloudProofAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCloudProof.afterGetCloudProofCounter)
}

// GetCloudProofBeforeCounter returns a count of CalculatorMock.GetCloudProof invocations
func (mmGetCloudProof *CalculatorMock) GetCloudProofBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCloudProof.beforeGetCloudProofCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.GetCloudProof.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCloudProof *mCalculatorMockGetCloudProof) Calls() []*CalculatorMockGetCloudProofParams {
	mmGetCloudProof.mutex.RLock()

	argCopy := make([]*CalculatorMockGetCloudProofParams, len(mmGetCloudProof.callArgs))
	copy(argCopy, mmGetCloudProof.callArgs)

	mmGetCloudProof.mutex.RUnlock()

	return argCopy
}

// MinimockGetCloudProofDone returns true if the count of the GetCloudProof invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockGetCloudProofDone() bool {
	for _, e := range m.GetCloudProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCloudProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCloudProofCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCloudProof != nil && mm_atomic.LoadUint64(&m.afterGetCloudProofCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCloudProofInspect logs each unmet expectation
func (m *CalculatorMock) MinimockGetCloudProofInspect() {
	for _, e := range m.GetCloudProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.GetCloudProof with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCloudProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCloudProofCounter) < 1 {
		if m.GetCloudProofMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.GetCloudProof")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.GetCloudProof with params: %#v", *m.GetCloudProofMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCloudProof != nil && mm_atomic.LoadUint64(&m.afterGetCloudProofCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.GetCloudProof")
	}
}

type mCalculatorMockGetGlobuleProof struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockGetGlobuleProofExpectation
	expectations       []*CalculatorMockGetGlobuleProofExpectation

	callArgs []*CalculatorMockGetGlobuleProofParams
	mutex    sync.RWMutex
}

// CalculatorMockGetGlobuleProofExpectation specifies expectation struct of the Calculator.GetGlobuleProof
type CalculatorMockGetGlobuleProofExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockGetGlobuleProofParams
	results *CalculatorMockGetGlobuleProofResults
	Counter uint64
}

// CalculatorMockGetGlobuleProofParams contains parameters of the Calculator.GetGlobuleProof
type CalculatorMockGetGlobuleProofParams struct {
	gp1 *mm_merkle.GlobuleEntry
}

// CalculatorMockGetGlobuleProofResults contains results of the Calculator.GetGlobuleProof
type CalculatorMockGetGlobuleProofResults struct {
	o1  mm_merkle.OriginHash
	gp2 *mm_merkle.GlobuleProof
	err error
}

// Expect sets up expected params for Calculator.GetGlobuleProof
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) Expect(gp1 *mm_merkle.GlobuleEntry) *mCalculatorMockGetGlobuleProof {
	if mmGetGlobuleProof.mock.funcGetGlobuleProof != nil {
		mmGetGlobuleProof.mock.t.Fatalf("CalculatorMock.GetGlobuleProof mock is already set by Set")
	}

	if mmGetGlobuleProof.defaultExpectation == nil {
		mmGetGlobuleProof.defaultExpectation = &CalculatorMockGetGlobuleProofExpectation{}
	}

	mmGetGlobuleProof.defaultExpectation.params = &CalculatorMockGetGlobuleProofParams{gp1}
	for _, e := range mmGetGlobuleProof.expectations {
		if minimock.Equal(e.params, mmGetGlobuleProof.defaultExpectation.params) {
			mmGetGlobuleProof.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetGlobuleProof.defaultExpectation.params)
		}
	}

	return mmGetGlobuleProof
}

// Inspect accepts an inspector function that has same arguments as the Calculator.GetGlobuleProof
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) Inspect(f func(gp1 *mm_merkle.GlobuleEntry)) *mCalculatorMockGetGlobuleProof {
	if mmGetGlobuleProof.mock.inspectFuncGetGlobuleProof != nil {
		mmGetGlobuleProof.mock.t.Fatalf("Inspect function is already set for CalculatorMock.GetGlobuleProof")
	}

	mmGetGlobuleProof.mock.inspectFuncGetGlobuleProof = f

	return mmGetGlobuleProof
}

// Return sets up results that will be returned by Calculator.GetGlobuleProof
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) Return(o1 mm_merkle.OriginHash, gp2 *mm_merkle.GlobuleProof, err error) *CalculatorMock {
	if mmGetGlobuleProof.mock.funcGetGlobuleProof != nil {
		mmGetGlobuleProof.mock.t.Fatalf("CalculatorMock.GetGlobuleProof mock is already set by Set")
	}

	if mmGetGlobuleProof.defaultExpectation == nil {
		mmGetGlobuleProof.defaultExpectation = &CalculatorMockGetGlobuleProofExpectation{mock: mmGetGlobuleProof.mock}
	}
	mmGetGlobuleProof.defaultExpectation.results = &CalculatorMockGetGlobuleProofResults{o1, gp2, err}
	return mmGetGlobuleProof.mock
}

//Set uses given function f to mock the Calculator.GetGlobuleProof method
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) Set(f func(gp1 *mm_merkle.GlobuleEntry) (o1 mm_merkle.OriginHash, gp2 *mm_merkle.GlobuleProof, err error)) *CalculatorMock {
	if mmGetGlobuleProof.defaultExpectation != nil {
		mmGetGlobuleProof.mock.t.Fatalf("Default expectation is already set for the Calculator.GetGlobuleProof method")
	}

	if len(mmGetGlobuleProof.expectations) > 0 {
		mmGetGlobuleProof.mock.t.Fatalf("Some expectations are already set for the Calculator.GetGlobuleProof method")
	}

	mmGetGlobuleProof.mock.funcGetGlobuleProof = f
	return mmGetGlobuleProof.mock
}

// When sets expectation for the Calculator.GetGlobuleProof which will trigger the result defined by the following
// Then helper
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) When(gp1 *mm_merkle.GlobuleEntry) *CalculatorMockGetGlobuleProofExpectation {
	if mmGetGlobuleProof.mock.funcGetGlobuleProof != nil {
		mmGetGlobuleProof.mock.t.Fatalf("CalculatorMock.GetGlobuleProof mock is already set by Set")
	}

	expectation := &CalculatorMockGetGlobuleProofExpectation{
		mock:   mmGetGlobuleProof.mock,
		params: &CalculatorMockGetGlobuleProofParams{gp1},
	}
	mmGetGlobuleProof.expectations = append(mmGetGlobuleProof.expectations, expectation)
	return expectation
}

// Then sets up Calculator.GetGlobuleProof return parameters for the expectation previously defined by the When method
func (e *CalculatorMockGetGlobuleProofExpectation) Then(o1 mm_merkle.OriginHash, gp2 *mm_merkle.GlobuleProof, err error) *CalculatorMock {
	e.results = &CalculatorMockGetGlobuleProofResults{o1, gp2, err}
	return e.mock
}

// GetGlobuleProof implements merkle.Calculator
func (mmGetGlobuleProof *CalculatorMock) GetGlobuleProof(gp1 *mm_merkle.GlobuleEntry) (o1 mm_merkle.OriginHash, gp2 *mm_merkle.GlobuleProof, err error) {
	mm_atomic.AddUint64(&mmGetGlobuleProof.beforeGetGlobuleProofCounter, 1)
	defer mm_atomic.AddUint64(&mmGetGlobuleProof.afterGetGlobuleProofCounter, 1)

	if mmGetGlobuleProof.inspectFuncGetGlobuleProof != nil {
		mmGetGlobuleProof.inspectFuncGetGlobuleProof(gp1)
	}

	mm_params := &CalculatorMockGetGlobuleProofParams{gp1}

	// Record call args
	mmGetGlobuleProof.GetGlobuleProofMock.mutex.Lock()
	mmGetGlobuleProof.GetGlobuleProofMock.callArgs = append(mmGetGlobuleProof.GetGlobuleProofMock.callArgs, mm_params)
	mmGetGlobuleProof.GetGlobuleProofMock.mutex.Unlock()

	for _, e := range mmGetGlobuleProof.GetGlobuleProofMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.o1, e.results.gp2, e.results.err
		}
	}

	if mmGetGlobuleProof.GetGlobuleProofMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetGlobuleProof.GetGlobuleProofMock.defaultExpectation.Counter, 1)
		mm_want := mmGetGlobuleProof.GetGlobuleProofMock.defaultExpectation.params
		mm_got := CalculatorMockGetGlobuleProofParams{gp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetGlobuleProof.t.Errorf("CalculatorMock.GetGlobuleProof got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetGlobuleProof.GetGlobuleProofMock.defaultExpectation.results
		if mm_results == nil {
			mmGetGlobuleProof.t.Fatal("No results are set for the CalculatorMock.GetGlobuleProof")
		}
		return (*mm_results).o1, (*mm_results).gp2, (*mm_results).err
	}
	if mmGetGlobuleProof.funcGetGlobuleProof != nil {
		return mmGetGlobuleProof.funcGetGlobuleProof(gp1)
	}
	mmGetGlobuleProof.t.Fatalf("Unexpected call to CalculatorMock.GetGlobuleProof. %v", gp1)
	return
}

// GetGlobuleProofAfterCounter returns a count of finished CalculatorMock.GetGlobuleProof invocations
func (mmGetGlobuleProof *CalculatorMock) GetGlobuleProofAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetGlobuleProof.afterGetGlobuleProofCounter)
}

// GetGlobuleProofBeforeCounter returns a count of CalculatorMock.GetGlobuleProof invocations
func (mmGetGlobuleProof *CalculatorMock) GetGlobuleProofBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetGlobuleProof.beforeGetGlobuleProofCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.GetGlobuleProof.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetGlobuleProof *mCalculatorMockGetGlobuleProof) Calls() []*CalculatorMockGetGlobuleProofParams {
	mmGetGlobuleProof.mutex.RLock()

	argCopy := make([]*CalculatorMockGetGlobuleProofParams, len(mmGetGlobuleProof.callArgs))
	copy(argCopy, mmGetGlobuleProof.callArgs)

	mmGetGlobuleProof.mutex.RUnlock()

	return argCopy
}

// MinimockGetGlobuleProofDone returns true if the count of the GetGlobuleProof invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockGetGlobuleProofDone() bool {
	for _, e := range m.GetGlobuleProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetGlobuleProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetGlobuleProofCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetGlobuleProof != nil && mm_atomic.LoadUint64(&m.afterGetGlobuleProofCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetGlobuleProofInspect logs each unmet expectation
func (m *CalculatorMock) MinimockGetGlobuleProofInspect() {
	for _, e := range m.GetGlobuleProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.GetGlobuleProof with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetGlobuleProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetGlobuleProofCounter) < 1 {
		if m.GetGlobuleProofMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.GetGlobuleProof")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.GetGlobuleProof with params: %#v", *m.GetGlobuleProofMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetGlobuleProof != nil && mm_atomic.LoadUint64(&m.afterGetGlobuleProofCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.GetGlobuleProof")
	}
}

type mCalculatorMockGetPulseProof struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockGetPulseProofExpectation
	expectations       []*CalculatorMockGetPulseProofExpectation

	callArgs []*CalculatorMockGetPulseProofParams
	mutex    sync.RWMutex
}

// CalculatorMockGetPulseProofExpectation specifies expectation struct of the Calculator.GetPulseProof
type CalculatorMockGetPulseProofExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockGetPulseProofParams
	results *CalculatorMockGetPulseProofResults
	Counter uint64
}

// CalculatorMockGetPulseProofParams contains parameters of the Calculator.GetPulseProof
type CalculatorMockGetPulseProofParams struct {
	pp1 *mm_merkle.PulseEntry
}

// CalculatorMockGetPulseProofResults contains results of the Calculator.GetPulseProof
type CalculatorMockGetPulseProofResults struct {
	o1  mm_merkle.OriginHash
	pp2 *mm_merkle.PulseProof
	err error
}

// Expect sets up expected params for Calculator.GetPulseProof
func (mmGetPulseProof *mCalculatorMockGetPulseProof) Expect(pp1 *mm_merkle.PulseEntry) *mCalculatorMockGetPulseProof {
	if mmGetPulseProof.mock.funcGetPulseProof != nil {
		mmGetPulseProof.mock.t.Fatalf("CalculatorMock.GetPulseProof mock is already set by Set")
	}

	if mmGetPulseProof.defaultExpectation == nil {
		mmGetPulseProof.defaultExpectation = &CalculatorMockGetPulseProofExpectation{}
	}

	mmGetPulseProof.defaultExpectation.params = &CalculatorMockGetPulseProofParams{pp1}
	for _, e := range mmGetPulseProof.expectations {
		if minimock.Equal(e.params, mmGetPulseProof.defaultExpectation.params) {
			mmGetPulseProof.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetPulseProof.defaultExpectation.params)
		}
	}

	return mmGetPulseProof
}

// Inspect accepts an inspector function that has same arguments as the Calculator.GetPulseProof
func (mmGetPulseProof *mCalculatorMockGetPulseProof) Inspect(f func(pp1 *mm_merkle.PulseEntry)) *mCalculatorMockGetPulseProof {
	if mmGetPulseProof.mock.inspectFuncGetPulseProof != nil {
		mmGetPulseProof.mock.t.Fatalf("Inspect function is already set for CalculatorMock.GetPulseProof")
	}

	mmGetPulseProof.mock.inspectFuncGetPulseProof = f

	return mmGetPulseProof
}

// Return sets up results that will be returned by Calculator.GetPulseProof
func (mmGetPulseProof *mCalculatorMockGetPulseProof) Return(o1 mm_merkle.OriginHash, pp2 *mm_merkle.PulseProof, err error) *CalculatorMock {
	if mmGetPulseProof.mock.funcGetPulseProof != nil {
		mmGetPulseProof.mock.t.Fatalf("CalculatorMock.GetPulseProof mock is already set by Set")
	}

	if mmGetPulseProof.defaultExpectation == nil {
		mmGetPulseProof.defaultExpectation = &CalculatorMockGetPulseProofExpectation{mock: mmGetPulseProof.mock}
	}
	mmGetPulseProof.defaultExpectation.results = &CalculatorMockGetPulseProofResults{o1, pp2, err}
	return mmGetPulseProof.mock
}

//Set uses given function f to mock the Calculator.GetPulseProof method
func (mmGetPulseProof *mCalculatorMockGetPulseProof) Set(f func(pp1 *mm_merkle.PulseEntry) (o1 mm_merkle.OriginHash, pp2 *mm_merkle.PulseProof, err error)) *CalculatorMock {
	if mmGetPulseProof.defaultExpectation != nil {
		mmGetPulseProof.mock.t.Fatalf("Default expectation is already set for the Calculator.GetPulseProof method")
	}

	if len(mmGetPulseProof.expectations) > 0 {
		mmGetPulseProof.mock.t.Fatalf("Some expectations are already set for the Calculator.GetPulseProof method")
	}

	mmGetPulseProof.mock.funcGetPulseProof = f
	return mmGetPulseProof.mock
}

// When sets expectation for the Calculator.GetPulseProof which will trigger the result defined by the following
// Then helper
func (mmGetPulseProof *mCalculatorMockGetPulseProof) When(pp1 *mm_merkle.PulseEntry) *CalculatorMockGetPulseProofExpectation {
	if mmGetPulseProof.mock.funcGetPulseProof != nil {
		mmGetPulseProof.mock.t.Fatalf("CalculatorMock.GetPulseProof mock is already set by Set")
	}

	expectation := &CalculatorMockGetPulseProofExpectation{
		mock:   mmGetPulseProof.mock,
		params: &CalculatorMockGetPulseProofParams{pp1},
	}
	mmGetPulseProof.expectations = append(mmGetPulseProof.expectations, expectation)
	return expectation
}

// Then sets up Calculator.GetPulseProof return parameters for the expectation previously defined by the When method
func (e *CalculatorMockGetPulseProofExpectation) Then(o1 mm_merkle.OriginHash, pp2 *mm_merkle.PulseProof, err error) *CalculatorMock {
	e.results = &CalculatorMockGetPulseProofResults{o1, pp2, err}
	return e.mock
}

// GetPulseProof implements merkle.Calculator
func (mmGetPulseProof *CalculatorMock) GetPulseProof(pp1 *mm_merkle.PulseEntry) (o1 mm_merkle.OriginHash, pp2 *mm_merkle.PulseProof, err error) {
	mm_atomic.AddUint64(&mmGetPulseProof.beforeGetPulseProofCounter, 1)
	defer mm_atomic.AddUint64(&mmGetPulseProof.afterGetPulseProofCounter, 1)

	if mmGetPulseProof.inspectFuncGetPulseProof != nil {
		mmGetPulseProof.inspectFuncGetPulseProof(pp1)
	}

	mm_params := &CalculatorMockGetPulseProofParams{pp1}

	// Record call args
	mmGetPulseProof.GetPulseProofMock.mutex.Lock()
	mmGetPulseProof.GetPulseProofMock.callArgs = append(mmGetPulseProof.GetPulseProofMock.callArgs, mm_params)
	mmGetPulseProof.GetPulseProofMock.mutex.Unlock()

	for _, e := range mmGetPulseProof.GetPulseProofMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.o1, e.results.pp2, e.results.err
		}
	}

	if mmGetPulseProof.GetPulseProofMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetPulseProof.GetPulseProofMock.defaultExpectation.Counter, 1)
		mm_want := mmGetPulseProof.GetPulseProofMock.defaultExpectation.params
		mm_got := CalculatorMockGetPulseProofParams{pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetPulseProof.t.Errorf("CalculatorMock.GetPulseProof got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetPulseProof.GetPulseProofMock.defaultExpectation.results
		if mm_results == nil {
			mmGetPulseProof.t.Fatal("No results are set for the CalculatorMock.GetPulseProof")
		}
		return (*mm_results).o1, (*mm_results).pp2, (*mm_results).err
	}
	if mmGetPulseProof.funcGetPulseProof != nil {
		return mmGetPulseProof.funcGetPulseProof(pp1)
	}
	mmGetPulseProof.t.Fatalf("Unexpected call to CalculatorMock.GetPulseProof. %v", pp1)
	return
}

// GetPulseProofAfterCounter returns a count of finished CalculatorMock.GetPulseProof invocations
func (mmGetPulseProof *CalculatorMock) GetPulseProofAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPulseProof.afterGetPulseProofCounter)
}

// GetPulseProofBeforeCounter returns a count of CalculatorMock.GetPulseProof invocations
func (mmGetPulseProof *CalculatorMock) GetPulseProofBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPulseProof.beforeGetPulseProofCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.GetPulseProof.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetPulseProof *mCalculatorMockGetPulseProof) Calls() []*CalculatorMockGetPulseProofParams {
	mmGetPulseProof.mutex.RLock()

	argCopy := make([]*CalculatorMockGetPulseProofParams, len(mmGetPulseProof.callArgs))
	copy(argCopy, mmGetPulseProof.callArgs)

	mmGetPulseProof.mutex.RUnlock()

	return argCopy
}

// MinimockGetPulseProofDone returns true if the count of the GetPulseProof invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockGetPulseProofDone() bool {
	for _, e := range m.GetPulseProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPulseProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPulseProofCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPulseProof != nil && mm_atomic.LoadUint64(&m.afterGetPulseProofCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetPulseProofInspect logs each unmet expectation
func (m *CalculatorMock) MinimockGetPulseProofInspect() {
	for _, e := range m.GetPulseProofMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.GetPulseProof with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPulseProofMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPulseProofCounter) < 1 {
		if m.GetPulseProofMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.GetPulseProof")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.GetPulseProof with params: %#v", *m.GetPulseProofMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPulseProof != nil && mm_atomic.LoadUint64(&m.afterGetPulseProofCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.GetPulseProof")
	}
}

type mCalculatorMockIsValid struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockIsValidExpectation
	expectations       []*CalculatorMockIsValidExpectation

	callArgs []*CalculatorMockIsValidParams
	mutex    sync.RWMutex
}

// CalculatorMockIsValidExpectation specifies expectation struct of the Calculator.IsValid
type CalculatorMockIsValidExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockIsValidParams
	results *CalculatorMockIsValidResults
	Counter uint64
}

// CalculatorMockIsValidParams contains parameters of the Calculator.IsValid
type CalculatorMockIsValidParams struct {
	p1 mm_merkle.Proof
	o1 mm_merkle.OriginHash
	p2 crypto.PublicKey
}

// CalculatorMockIsValidResults contains results of the Calculator.IsValid
type CalculatorMockIsValidResults struct {
	b1 bool
}

// Expect sets up expected params for Calculator.IsValid
func (mmIsValid *mCalculatorMockIsValid) Expect(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey) *mCalculatorMockIsValid {
	if mmIsValid.mock.funcIsValid != nil {
		mmIsValid.mock.t.Fatalf("CalculatorMock.IsValid mock is already set by Set")
	}

	if mmIsValid.defaultExpectation == nil {
		mmIsValid.defaultExpectation = &CalculatorMockIsValidExpectation{}
	}

	mmIsValid.defaultExpectation.params = &CalculatorMockIsValidParams{p1, o1, p2}
	for _, e := range mmIsValid.expectations {
		if minimock.Equal(e.params, mmIsValid.defaultExpectation.params) {
			mmIsValid.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIsValid.defaultExpectation.params)
		}
	}

	return mmIsValid
}

// Inspect accepts an inspector function that has same arguments as the Calculator.IsValid
func (mmIsValid *mCalculatorMockIsValid) Inspect(f func(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey)) *mCalculatorMockIsValid {
	if mmIsValid.mock.inspectFuncIsValid != nil {
		mmIsValid.mock.t.Fatalf("Inspect function is already set for CalculatorMock.IsValid")
	}

	mmIsValid.mock.inspectFuncIsValid = f

	return mmIsValid
}

// Return sets up results that will be returned by Calculator.IsValid
func (mmIsValid *mCalculatorMockIsValid) Return(b1 bool) *CalculatorMock {
	if mmIsValid.mock.funcIsValid != nil {
		mmIsValid.mock.t.Fatalf("CalculatorMock.IsValid mock is already set by Set")
	}

	if mmIsValid.defaultExpectation == nil {
		mmIsValid.defaultExpectation = &CalculatorMockIsValidExpectation{mock: mmIsValid.mock}
	}
	mmIsValid.defaultExpectation.results = &CalculatorMockIsValidResults{b1}
	return mmIsValid.mock
}

//Set uses given function f to mock the Calculator.IsValid method
func (mmIsValid *mCalculatorMockIsValid) Set(f func(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey) (b1 bool)) *CalculatorMock {
	if mmIsValid.defaultExpectation != nil {
		mmIsValid.mock.t.Fatalf("Default expectation is already set for the Calculator.IsValid method")
	}

	if len(mmIsValid.expectations) > 0 {
		mmIsValid.mock.t.Fatalf("Some expectations are already set for the Calculator.IsValid method")
	}

	mmIsValid.mock.funcIsValid = f
	return mmIsValid.mock
}

// When sets expectation for the Calculator.IsValid which will trigger the result defined by the following
// Then helper
func (mmIsValid *mCalculatorMockIsValid) When(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey) *CalculatorMockIsValidExpectation {
	if mmIsValid.mock.funcIsValid != nil {
		mmIsValid.mock.t.Fatalf("CalculatorMock.IsValid mock is already set by Set")
	}

	expectation := &CalculatorMockIsValidExpectation{
		mock:   mmIsValid.mock,
		params: &CalculatorMockIsValidParams{p1, o1, p2},
	}
	mmIsValid.expectations = append(mmIsValid.expectations, expectation)
	return expectation
}

// Then sets up Calculator.IsValid return parameters for the expectation previously defined by the When method
func (e *CalculatorMockIsValidExpectation) Then(b1 bool) *CalculatorMock {
	e.results = &CalculatorMockIsValidResults{b1}
	return e.mock
}

// IsValid implements merkle.Calculator
func (mmIsValid *CalculatorMock) IsValid(p1 mm_merkle.Proof, o1 mm_merkle.OriginHash, p2 crypto.PublicKey) (b1 bool) {
	mm_atomic.AddUint64(&mmIsValid.beforeIsValidCounter, 1)
	defer mm_atomic.AddUint64(&mmIsValid.afterIsValidCounter, 1)

	if mmIsValid.inspectFuncIsValid != nil {
		mmIsValid.inspectFuncIsValid(p1, o1, p2)
	}

	mm_params := &CalculatorMockIsValidParams{p1, o1, p2}

	// Record call args
	mmIsValid.IsValidMock.mutex.Lock()
	mmIsValid.IsValidMock.callArgs = append(mmIsValid.IsValidMock.callArgs, mm_params)
	mmIsValid.IsValidMock.mutex.Unlock()

	for _, e := range mmIsValid.IsValidMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1
		}
	}

	if mmIsValid.IsValidMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsValid.IsValidMock.defaultExpectation.Counter, 1)
		mm_want := mmIsValid.IsValidMock.defaultExpectation.params
		mm_got := CalculatorMockIsValidParams{p1, o1, p2}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIsValid.t.Errorf("CalculatorMock.IsValid got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIsValid.IsValidMock.defaultExpectation.results
		if mm_results == nil {
			mmIsValid.t.Fatal("No results are set for the CalculatorMock.IsValid")
		}
		return (*mm_results).b1
	}
	if mmIsValid.funcIsValid != nil {
		return mmIsValid.funcIsValid(p1, o1, p2)
	}
	mmIsValid.t.Fatalf("Unexpected call to CalculatorMock.IsValid. %v %v %v", p1, o1, p2)
	return
}

// IsValidAfterCounter returns a count of finished CalculatorMock.IsValid invocations
func (mmIsValid *CalculatorMock) IsValidAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsValid.afterIsValidCounter)
}

// IsValidBeforeCounter returns a count of CalculatorMock.IsValid invocations
func (mmIsValid *CalculatorMock) IsValidBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsValid.beforeIsValidCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.IsValid.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIsValid *mCalculatorMockIsValid) Calls() []*CalculatorMockIsValidParams {
	mmIsValid.mutex.RLock()

	argCopy := make([]*CalculatorMockIsValidParams, len(mmIsValid.callArgs))
	copy(argCopy, mmIsValid.callArgs)

	mmIsValid.mutex.RUnlock()

	return argCopy
}

// MinimockIsValidDone returns true if the count of the IsValid invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockIsValidDone() bool {
	for _, e := range m.IsValidMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsValidMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsValidCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsValid != nil && mm_atomic.LoadUint64(&m.afterIsValidCounter) < 1 {
		return false
	}
	return true
}

// MinimockIsValidInspect logs each unmet expectation
func (m *CalculatorMock) MinimockIsValidInspect() {
	for _, e := range m.IsValidMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.IsValid with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsValidMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsValidCounter) < 1 {
		if m.IsValidMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.IsValid")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.IsValid with params: %#v", *m.IsValidMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsValid != nil && mm_atomic.LoadUint64(&m.afterIsValidCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.IsValid")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CalculatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetCloudProofInspect()

		m.MinimockGetGlobuleProofInspect()

		m.MinimockGetPulseProofInspect()

		m.MinimockIsValidInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CalculatorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CalculatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetCloudProofDone() &&
		m.MinimockGetGlobuleProofDone() &&
		m.MinimockGetPulseProofDone() &&
		m.MinimockIsValidDone()
}
