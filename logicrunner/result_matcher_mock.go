package logicrunner

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/payload"
)

// ResultMatcherMock implements ResultMatcher
type ResultMatcherMock struct {
	t minimock.Tester

	funcAddStillExecution          func(ctx context.Context, msg *payload.StillExecuting)
	inspectFuncAddStillExecution   func(ctx context.Context, msg *payload.StillExecuting)
	afterAddStillExecutionCounter  uint64
	beforeAddStillExecutionCounter uint64
	AddStillExecutionMock          mResultMatcherMockAddStillExecution

	funcAddUnwantedResponse          func(ctx context.Context, msg *message.ReturnResults) (err error)
	inspectFuncAddUnwantedResponse   func(ctx context.Context, msg *message.ReturnResults)
	afterAddUnwantedResponseCounter  uint64
	beforeAddUnwantedResponseCounter uint64
	AddUnwantedResponseMock          mResultMatcherMockAddUnwantedResponse

	funcClear          func()
	inspectFuncClear   func()
	afterClearCounter  uint64
	beforeClearCounter uint64
	ClearMock          mResultMatcherMockClear
}

// NewResultMatcherMock returns a mock for ResultMatcher
func NewResultMatcherMock(t minimock.Tester) *ResultMatcherMock {
	m := &ResultMatcherMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddStillExecutionMock = mResultMatcherMockAddStillExecution{mock: m}
	m.AddStillExecutionMock.callArgs = []*ResultMatcherMockAddStillExecutionParams{}

	m.AddUnwantedResponseMock = mResultMatcherMockAddUnwantedResponse{mock: m}
	m.AddUnwantedResponseMock.callArgs = []*ResultMatcherMockAddUnwantedResponseParams{}

	m.ClearMock = mResultMatcherMockClear{mock: m}

	return m
}

type mResultMatcherMockAddStillExecution struct {
	mock               *ResultMatcherMock
	defaultExpectation *ResultMatcherMockAddStillExecutionExpectation
	expectations       []*ResultMatcherMockAddStillExecutionExpectation

	callArgs []*ResultMatcherMockAddStillExecutionParams
	mutex    sync.RWMutex
}

// ResultMatcherMockAddStillExecutionExpectation specifies expectation struct of the ResultMatcher.AddStillExecution
type ResultMatcherMockAddStillExecutionExpectation struct {
	mock   *ResultMatcherMock
	params *ResultMatcherMockAddStillExecutionParams

	Counter uint64
}

// ResultMatcherMockAddStillExecutionParams contains parameters of the ResultMatcher.AddStillExecution
type ResultMatcherMockAddStillExecutionParams struct {
	ctx context.Context
	msg *payload.StillExecuting
}

// Expect sets up expected params for ResultMatcher.AddStillExecution
func (mmAddStillExecution *mResultMatcherMockAddStillExecution) Expect(ctx context.Context, msg *payload.StillExecuting) *mResultMatcherMockAddStillExecution {
	if mmAddStillExecution.mock.funcAddStillExecution != nil {
		mmAddStillExecution.mock.t.Fatalf("ResultMatcherMock.AddStillExecution mock is already set by Set")
	}

	if mmAddStillExecution.defaultExpectation == nil {
		mmAddStillExecution.defaultExpectation = &ResultMatcherMockAddStillExecutionExpectation{}
	}

	mmAddStillExecution.defaultExpectation.params = &ResultMatcherMockAddStillExecutionParams{ctx, msg}
	for _, e := range mmAddStillExecution.expectations {
		if minimock.Equal(e.params, mmAddStillExecution.defaultExpectation.params) {
			mmAddStillExecution.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddStillExecution.defaultExpectation.params)
		}
	}

	return mmAddStillExecution
}

// Inspect accepts an inspector function that has same arguments as the ResultMatcher.AddStillExecution
func (mmAddStillExecution *mResultMatcherMockAddStillExecution) Inspect(f func(ctx context.Context, msg *payload.StillExecuting)) *mResultMatcherMockAddStillExecution {
	if mmAddStillExecution.mock.inspectFuncAddStillExecution != nil {
		mmAddStillExecution.mock.t.Fatalf("Inspect function is already set for ResultMatcherMock.AddStillExecution")
	}

	mmAddStillExecution.mock.inspectFuncAddStillExecution = f

	return mmAddStillExecution
}

// Return sets up results that will be returned by ResultMatcher.AddStillExecution
func (mmAddStillExecution *mResultMatcherMockAddStillExecution) Return() *ResultMatcherMock {
	if mmAddStillExecution.mock.funcAddStillExecution != nil {
		mmAddStillExecution.mock.t.Fatalf("ResultMatcherMock.AddStillExecution mock is already set by Set")
	}

	if mmAddStillExecution.defaultExpectation == nil {
		mmAddStillExecution.defaultExpectation = &ResultMatcherMockAddStillExecutionExpectation{mock: mmAddStillExecution.mock}
	}

	return mmAddStillExecution.mock
}

//Set uses given function f to mock the ResultMatcher.AddStillExecution method
func (mmAddStillExecution *mResultMatcherMockAddStillExecution) Set(f func(ctx context.Context, msg *payload.StillExecuting)) *ResultMatcherMock {
	if mmAddStillExecution.defaultExpectation != nil {
		mmAddStillExecution.mock.t.Fatalf("Default expectation is already set for the ResultMatcher.AddStillExecution method")
	}

	if len(mmAddStillExecution.expectations) > 0 {
		mmAddStillExecution.mock.t.Fatalf("Some expectations are already set for the ResultMatcher.AddStillExecution method")
	}

	mmAddStillExecution.mock.funcAddStillExecution = f
	return mmAddStillExecution.mock
}

// AddStillExecution implements ResultMatcher
func (mmAddStillExecution *ResultMatcherMock) AddStillExecution(ctx context.Context, msg *payload.StillExecuting) {
	mm_atomic.AddUint64(&mmAddStillExecution.beforeAddStillExecutionCounter, 1)
	defer mm_atomic.AddUint64(&mmAddStillExecution.afterAddStillExecutionCounter, 1)

	if mmAddStillExecution.inspectFuncAddStillExecution != nil {
		mmAddStillExecution.inspectFuncAddStillExecution(ctx, msg)
	}

	params := &ResultMatcherMockAddStillExecutionParams{ctx, msg}

	// Record call args
	mmAddStillExecution.AddStillExecutionMock.mutex.Lock()
	mmAddStillExecution.AddStillExecutionMock.callArgs = append(mmAddStillExecution.AddStillExecutionMock.callArgs, params)
	mmAddStillExecution.AddStillExecutionMock.mutex.Unlock()

	for _, e := range mmAddStillExecution.AddStillExecutionMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmAddStillExecution.AddStillExecutionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddStillExecution.AddStillExecutionMock.defaultExpectation.Counter, 1)
		want := mmAddStillExecution.AddStillExecutionMock.defaultExpectation.params
		got := ResultMatcherMockAddStillExecutionParams{ctx, msg}
		if want != nil && !minimock.Equal(*want, got) {
			mmAddStillExecution.t.Errorf("ResultMatcherMock.AddStillExecution got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmAddStillExecution.funcAddStillExecution != nil {
		mmAddStillExecution.funcAddStillExecution(ctx, msg)
		return
	}
	mmAddStillExecution.t.Fatalf("Unexpected call to ResultMatcherMock.AddStillExecution. %v %v", ctx, msg)

}

// AddStillExecutionAfterCounter returns a count of finished ResultMatcherMock.AddStillExecution invocations
func (mmAddStillExecution *ResultMatcherMock) AddStillExecutionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddStillExecution.afterAddStillExecutionCounter)
}

// AddStillExecutionBeforeCounter returns a count of ResultMatcherMock.AddStillExecution invocations
func (mmAddStillExecution *ResultMatcherMock) AddStillExecutionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddStillExecution.beforeAddStillExecutionCounter)
}

// Calls returns a list of arguments used in each call to ResultMatcherMock.AddStillExecution.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddStillExecution *mResultMatcherMockAddStillExecution) Calls() []*ResultMatcherMockAddStillExecutionParams {
	mmAddStillExecution.mutex.RLock()

	argCopy := make([]*ResultMatcherMockAddStillExecutionParams, len(mmAddStillExecution.callArgs))
	copy(argCopy, mmAddStillExecution.callArgs)

	mmAddStillExecution.mutex.RUnlock()

	return argCopy
}

// MinimockAddStillExecutionDone returns true if the count of the AddStillExecution invocations corresponds
// the number of defined expectations
func (m *ResultMatcherMock) MinimockAddStillExecutionDone() bool {
	for _, e := range m.AddStillExecutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddStillExecutionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddStillExecutionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddStillExecution != nil && mm_atomic.LoadUint64(&m.afterAddStillExecutionCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddStillExecutionInspect logs each unmet expectation
func (m *ResultMatcherMock) MinimockAddStillExecutionInspect() {
	for _, e := range m.AddStillExecutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ResultMatcherMock.AddStillExecution with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddStillExecutionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddStillExecutionCounter) < 1 {
		if m.AddStillExecutionMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ResultMatcherMock.AddStillExecution")
		} else {
			m.t.Errorf("Expected call to ResultMatcherMock.AddStillExecution with params: %#v", *m.AddStillExecutionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddStillExecution != nil && mm_atomic.LoadUint64(&m.afterAddStillExecutionCounter) < 1 {
		m.t.Error("Expected call to ResultMatcherMock.AddStillExecution")
	}
}

type mResultMatcherMockAddUnwantedResponse struct {
	mock               *ResultMatcherMock
	defaultExpectation *ResultMatcherMockAddUnwantedResponseExpectation
	expectations       []*ResultMatcherMockAddUnwantedResponseExpectation

	callArgs []*ResultMatcherMockAddUnwantedResponseParams
	mutex    sync.RWMutex
}

// ResultMatcherMockAddUnwantedResponseExpectation specifies expectation struct of the ResultMatcher.AddUnwantedResponse
type ResultMatcherMockAddUnwantedResponseExpectation struct {
	mock    *ResultMatcherMock
	params  *ResultMatcherMockAddUnwantedResponseParams
	results *ResultMatcherMockAddUnwantedResponseResults
	Counter uint64
}

// ResultMatcherMockAddUnwantedResponseParams contains parameters of the ResultMatcher.AddUnwantedResponse
type ResultMatcherMockAddUnwantedResponseParams struct {
	ctx context.Context
	msg *message.ReturnResults
}

// ResultMatcherMockAddUnwantedResponseResults contains results of the ResultMatcher.AddUnwantedResponse
type ResultMatcherMockAddUnwantedResponseResults struct {
	err error
}

// Expect sets up expected params for ResultMatcher.AddUnwantedResponse
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) Expect(ctx context.Context, msg *message.ReturnResults) *mResultMatcherMockAddUnwantedResponse {
	if mmAddUnwantedResponse.mock.funcAddUnwantedResponse != nil {
		mmAddUnwantedResponse.mock.t.Fatalf("ResultMatcherMock.AddUnwantedResponse mock is already set by Set")
	}

	if mmAddUnwantedResponse.defaultExpectation == nil {
		mmAddUnwantedResponse.defaultExpectation = &ResultMatcherMockAddUnwantedResponseExpectation{}
	}

	mmAddUnwantedResponse.defaultExpectation.params = &ResultMatcherMockAddUnwantedResponseParams{ctx, msg}
	for _, e := range mmAddUnwantedResponse.expectations {
		if minimock.Equal(e.params, mmAddUnwantedResponse.defaultExpectation.params) {
			mmAddUnwantedResponse.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddUnwantedResponse.defaultExpectation.params)
		}
	}

	return mmAddUnwantedResponse
}

// Inspect accepts an inspector function that has same arguments as the ResultMatcher.AddUnwantedResponse
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) Inspect(f func(ctx context.Context, msg *message.ReturnResults)) *mResultMatcherMockAddUnwantedResponse {
	if mmAddUnwantedResponse.mock.inspectFuncAddUnwantedResponse != nil {
		mmAddUnwantedResponse.mock.t.Fatalf("Inspect function is already set for ResultMatcherMock.AddUnwantedResponse")
	}

	mmAddUnwantedResponse.mock.inspectFuncAddUnwantedResponse = f

	return mmAddUnwantedResponse
}

// Return sets up results that will be returned by ResultMatcher.AddUnwantedResponse
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) Return(err error) *ResultMatcherMock {
	if mmAddUnwantedResponse.mock.funcAddUnwantedResponse != nil {
		mmAddUnwantedResponse.mock.t.Fatalf("ResultMatcherMock.AddUnwantedResponse mock is already set by Set")
	}

	if mmAddUnwantedResponse.defaultExpectation == nil {
		mmAddUnwantedResponse.defaultExpectation = &ResultMatcherMockAddUnwantedResponseExpectation{mock: mmAddUnwantedResponse.mock}
	}
	mmAddUnwantedResponse.defaultExpectation.results = &ResultMatcherMockAddUnwantedResponseResults{err}
	return mmAddUnwantedResponse.mock
}

//Set uses given function f to mock the ResultMatcher.AddUnwantedResponse method
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) Set(f func(ctx context.Context, msg *message.ReturnResults) (err error)) *ResultMatcherMock {
	if mmAddUnwantedResponse.defaultExpectation != nil {
		mmAddUnwantedResponse.mock.t.Fatalf("Default expectation is already set for the ResultMatcher.AddUnwantedResponse method")
	}

	if len(mmAddUnwantedResponse.expectations) > 0 {
		mmAddUnwantedResponse.mock.t.Fatalf("Some expectations are already set for the ResultMatcher.AddUnwantedResponse method")
	}

	mmAddUnwantedResponse.mock.funcAddUnwantedResponse = f
	return mmAddUnwantedResponse.mock
}

// When sets expectation for the ResultMatcher.AddUnwantedResponse which will trigger the result defined by the following
// Then helper
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) When(ctx context.Context, msg *message.ReturnResults) *ResultMatcherMockAddUnwantedResponseExpectation {
	if mmAddUnwantedResponse.mock.funcAddUnwantedResponse != nil {
		mmAddUnwantedResponse.mock.t.Fatalf("ResultMatcherMock.AddUnwantedResponse mock is already set by Set")
	}

	expectation := &ResultMatcherMockAddUnwantedResponseExpectation{
		mock:   mmAddUnwantedResponse.mock,
		params: &ResultMatcherMockAddUnwantedResponseParams{ctx, msg},
	}
	mmAddUnwantedResponse.expectations = append(mmAddUnwantedResponse.expectations, expectation)
	return expectation
}

// Then sets up ResultMatcher.AddUnwantedResponse return parameters for the expectation previously defined by the When method
func (e *ResultMatcherMockAddUnwantedResponseExpectation) Then(err error) *ResultMatcherMock {
	e.results = &ResultMatcherMockAddUnwantedResponseResults{err}
	return e.mock
}

// AddUnwantedResponse implements ResultMatcher
func (mmAddUnwantedResponse *ResultMatcherMock) AddUnwantedResponse(ctx context.Context, msg *message.ReturnResults) (err error) {
	mm_atomic.AddUint64(&mmAddUnwantedResponse.beforeAddUnwantedResponseCounter, 1)
	defer mm_atomic.AddUint64(&mmAddUnwantedResponse.afterAddUnwantedResponseCounter, 1)

	if mmAddUnwantedResponse.inspectFuncAddUnwantedResponse != nil {
		mmAddUnwantedResponse.inspectFuncAddUnwantedResponse(ctx, msg)
	}

	params := &ResultMatcherMockAddUnwantedResponseParams{ctx, msg}

	// Record call args
	mmAddUnwantedResponse.AddUnwantedResponseMock.mutex.Lock()
	mmAddUnwantedResponse.AddUnwantedResponseMock.callArgs = append(mmAddUnwantedResponse.AddUnwantedResponseMock.callArgs, params)
	mmAddUnwantedResponse.AddUnwantedResponseMock.mutex.Unlock()

	for _, e := range mmAddUnwantedResponse.AddUnwantedResponseMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddUnwantedResponse.AddUnwantedResponseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddUnwantedResponse.AddUnwantedResponseMock.defaultExpectation.Counter, 1)
		want := mmAddUnwantedResponse.AddUnwantedResponseMock.defaultExpectation.params
		got := ResultMatcherMockAddUnwantedResponseParams{ctx, msg}
		if want != nil && !minimock.Equal(*want, got) {
			mmAddUnwantedResponse.t.Errorf("ResultMatcherMock.AddUnwantedResponse got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmAddUnwantedResponse.AddUnwantedResponseMock.defaultExpectation.results
		if results == nil {
			mmAddUnwantedResponse.t.Fatal("No results are set for the ResultMatcherMock.AddUnwantedResponse")
		}
		return (*results).err
	}
	if mmAddUnwantedResponse.funcAddUnwantedResponse != nil {
		return mmAddUnwantedResponse.funcAddUnwantedResponse(ctx, msg)
	}
	mmAddUnwantedResponse.t.Fatalf("Unexpected call to ResultMatcherMock.AddUnwantedResponse. %v %v", ctx, msg)
	return
}

// AddUnwantedResponseAfterCounter returns a count of finished ResultMatcherMock.AddUnwantedResponse invocations
func (mmAddUnwantedResponse *ResultMatcherMock) AddUnwantedResponseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddUnwantedResponse.afterAddUnwantedResponseCounter)
}

// AddUnwantedResponseBeforeCounter returns a count of ResultMatcherMock.AddUnwantedResponse invocations
func (mmAddUnwantedResponse *ResultMatcherMock) AddUnwantedResponseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddUnwantedResponse.beforeAddUnwantedResponseCounter)
}

// Calls returns a list of arguments used in each call to ResultMatcherMock.AddUnwantedResponse.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddUnwantedResponse *mResultMatcherMockAddUnwantedResponse) Calls() []*ResultMatcherMockAddUnwantedResponseParams {
	mmAddUnwantedResponse.mutex.RLock()

	argCopy := make([]*ResultMatcherMockAddUnwantedResponseParams, len(mmAddUnwantedResponse.callArgs))
	copy(argCopy, mmAddUnwantedResponse.callArgs)

	mmAddUnwantedResponse.mutex.RUnlock()

	return argCopy
}

// MinimockAddUnwantedResponseDone returns true if the count of the AddUnwantedResponse invocations corresponds
// the number of defined expectations
func (m *ResultMatcherMock) MinimockAddUnwantedResponseDone() bool {
	for _, e := range m.AddUnwantedResponseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddUnwantedResponseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddUnwantedResponseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddUnwantedResponse != nil && mm_atomic.LoadUint64(&m.afterAddUnwantedResponseCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddUnwantedResponseInspect logs each unmet expectation
func (m *ResultMatcherMock) MinimockAddUnwantedResponseInspect() {
	for _, e := range m.AddUnwantedResponseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ResultMatcherMock.AddUnwantedResponse with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddUnwantedResponseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddUnwantedResponseCounter) < 1 {
		if m.AddUnwantedResponseMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ResultMatcherMock.AddUnwantedResponse")
		} else {
			m.t.Errorf("Expected call to ResultMatcherMock.AddUnwantedResponse with params: %#v", *m.AddUnwantedResponseMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddUnwantedResponse != nil && mm_atomic.LoadUint64(&m.afterAddUnwantedResponseCounter) < 1 {
		m.t.Error("Expected call to ResultMatcherMock.AddUnwantedResponse")
	}
}

type mResultMatcherMockClear struct {
	mock               *ResultMatcherMock
	defaultExpectation *ResultMatcherMockClearExpectation
	expectations       []*ResultMatcherMockClearExpectation
}

// ResultMatcherMockClearExpectation specifies expectation struct of the ResultMatcher.Clear
type ResultMatcherMockClearExpectation struct {
	mock *ResultMatcherMock

	Counter uint64
}

// Expect sets up expected params for ResultMatcher.Clear
func (mmClear *mResultMatcherMockClear) Expect() *mResultMatcherMockClear {
	if mmClear.mock.funcClear != nil {
		mmClear.mock.t.Fatalf("ResultMatcherMock.Clear mock is already set by Set")
	}

	if mmClear.defaultExpectation == nil {
		mmClear.defaultExpectation = &ResultMatcherMockClearExpectation{}
	}

	return mmClear
}

// Inspect accepts an inspector function that has same arguments as the ResultMatcher.Clear
func (mmClear *mResultMatcherMockClear) Inspect(f func()) *mResultMatcherMockClear {
	if mmClear.mock.inspectFuncClear != nil {
		mmClear.mock.t.Fatalf("Inspect function is already set for ResultMatcherMock.Clear")
	}

	mmClear.mock.inspectFuncClear = f

	return mmClear
}

// Return sets up results that will be returned by ResultMatcher.Clear
func (mmClear *mResultMatcherMockClear) Return() *ResultMatcherMock {
	if mmClear.mock.funcClear != nil {
		mmClear.mock.t.Fatalf("ResultMatcherMock.Clear mock is already set by Set")
	}

	if mmClear.defaultExpectation == nil {
		mmClear.defaultExpectation = &ResultMatcherMockClearExpectation{mock: mmClear.mock}
	}

	return mmClear.mock
}

//Set uses given function f to mock the ResultMatcher.Clear method
func (mmClear *mResultMatcherMockClear) Set(f func()) *ResultMatcherMock {
	if mmClear.defaultExpectation != nil {
		mmClear.mock.t.Fatalf("Default expectation is already set for the ResultMatcher.Clear method")
	}

	if len(mmClear.expectations) > 0 {
		mmClear.mock.t.Fatalf("Some expectations are already set for the ResultMatcher.Clear method")
	}

	mmClear.mock.funcClear = f
	return mmClear.mock
}

// Clear implements ResultMatcher
func (mmClear *ResultMatcherMock) Clear() {
	mm_atomic.AddUint64(&mmClear.beforeClearCounter, 1)
	defer mm_atomic.AddUint64(&mmClear.afterClearCounter, 1)

	if mmClear.inspectFuncClear != nil {
		mmClear.inspectFuncClear()
	}

	if mmClear.ClearMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClear.ClearMock.defaultExpectation.Counter, 1)

		return

	}
	if mmClear.funcClear != nil {
		mmClear.funcClear()
		return
	}
	mmClear.t.Fatalf("Unexpected call to ResultMatcherMock.Clear.")

}

// ClearAfterCounter returns a count of finished ResultMatcherMock.Clear invocations
func (mmClear *ResultMatcherMock) ClearAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClear.afterClearCounter)
}

// ClearBeforeCounter returns a count of ResultMatcherMock.Clear invocations
func (mmClear *ResultMatcherMock) ClearBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClear.beforeClearCounter)
}

// MinimockClearDone returns true if the count of the Clear invocations corresponds
// the number of defined expectations
func (m *ResultMatcherMock) MinimockClearDone() bool {
	for _, e := range m.ClearMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClear != nil && mm_atomic.LoadUint64(&m.afterClearCounter) < 1 {
		return false
	}
	return true
}

// MinimockClearInspect logs each unmet expectation
func (m *ResultMatcherMock) MinimockClearInspect() {
	for _, e := range m.ClearMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ResultMatcherMock.Clear")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearCounter) < 1 {
		m.t.Error("Expected call to ResultMatcherMock.Clear")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClear != nil && mm_atomic.LoadUint64(&m.afterClearCounter) < 1 {
		m.t.Error("Expected call to ResultMatcherMock.Clear")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ResultMatcherMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddStillExecutionInspect()

		m.MinimockAddUnwantedResponseInspect()

		m.MinimockClearInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ResultMatcherMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ResultMatcherMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddStillExecutionDone() &&
		m.MinimockAddUnwantedResponseDone() &&
		m.MinimockClearDone()
}
