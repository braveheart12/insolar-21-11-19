package executor

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MetricsRegistryMock implements MetricsRegistry
type MetricsRegistryMock struct {
	t minimock.Tester

	funcSetOldestAbandonedRequestAge          func(age int)
	inspectFuncSetOldestAbandonedRequestAge   func(age int)
	afterSetOldestAbandonedRequestAgeCounter  uint64
	beforeSetOldestAbandonedRequestAgeCounter uint64
	SetOldestAbandonedRequestAgeMock          mMetricsRegistryMockSetOldestAbandonedRequestAge

	funcUpdateMetrics          func(ctx context.Context)
	inspectFuncUpdateMetrics   func(ctx context.Context)
	afterUpdateMetricsCounter  uint64
	beforeUpdateMetricsCounter uint64
	UpdateMetricsMock          mMetricsRegistryMockUpdateMetrics
}

// NewMetricsRegistryMock returns a mock for MetricsRegistry
func NewMetricsRegistryMock(t minimock.Tester) *MetricsRegistryMock {
	m := &MetricsRegistryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SetOldestAbandonedRequestAgeMock = mMetricsRegistryMockSetOldestAbandonedRequestAge{mock: m}
	m.SetOldestAbandonedRequestAgeMock.callArgs = []*MetricsRegistryMockSetOldestAbandonedRequestAgeParams{}

	m.UpdateMetricsMock = mMetricsRegistryMockUpdateMetrics{mock: m}
	m.UpdateMetricsMock.callArgs = []*MetricsRegistryMockUpdateMetricsParams{}

	return m
}

type mMetricsRegistryMockSetOldestAbandonedRequestAge struct {
	mock               *MetricsRegistryMock
	defaultExpectation *MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation
	expectations       []*MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation

	callArgs []*MetricsRegistryMockSetOldestAbandonedRequestAgeParams
	mutex    sync.RWMutex
}

// MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation specifies expectation struct of the MetricsRegistry.SetOldestAbandonedRequestAge
type MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation struct {
	mock   *MetricsRegistryMock
	params *MetricsRegistryMockSetOldestAbandonedRequestAgeParams

	Counter uint64
}

// MetricsRegistryMockSetOldestAbandonedRequestAgeParams contains parameters of the MetricsRegistry.SetOldestAbandonedRequestAge
type MetricsRegistryMockSetOldestAbandonedRequestAgeParams struct {
	age int
}

// Expect sets up expected params for MetricsRegistry.SetOldestAbandonedRequestAge
func (mmSetOldestAbandonedRequestAge *mMetricsRegistryMockSetOldestAbandonedRequestAge) Expect(age int) *mMetricsRegistryMockSetOldestAbandonedRequestAge {
	if mmSetOldestAbandonedRequestAge.mock.funcSetOldestAbandonedRequestAge != nil {
		mmSetOldestAbandonedRequestAge.mock.t.Fatalf("MetricsRegistryMock.SetOldestAbandonedRequestAge mock is already set by Set")
	}

	if mmSetOldestAbandonedRequestAge.defaultExpectation == nil {
		mmSetOldestAbandonedRequestAge.defaultExpectation = &MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation{}
	}

	mmSetOldestAbandonedRequestAge.defaultExpectation.params = &MetricsRegistryMockSetOldestAbandonedRequestAgeParams{age}
	for _, e := range mmSetOldestAbandonedRequestAge.expectations {
		if minimock.Equal(e.params, mmSetOldestAbandonedRequestAge.defaultExpectation.params) {
			mmSetOldestAbandonedRequestAge.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSetOldestAbandonedRequestAge.defaultExpectation.params)
		}
	}

	return mmSetOldestAbandonedRequestAge
}

// Inspect accepts an inspector function that has same arguments as the MetricsRegistry.SetOldestAbandonedRequestAge
func (mmSetOldestAbandonedRequestAge *mMetricsRegistryMockSetOldestAbandonedRequestAge) Inspect(f func(age int)) *mMetricsRegistryMockSetOldestAbandonedRequestAge {
	if mmSetOldestAbandonedRequestAge.mock.inspectFuncSetOldestAbandonedRequestAge != nil {
		mmSetOldestAbandonedRequestAge.mock.t.Fatalf("Inspect function is already set for MetricsRegistryMock.SetOldestAbandonedRequestAge")
	}

	mmSetOldestAbandonedRequestAge.mock.inspectFuncSetOldestAbandonedRequestAge = f

	return mmSetOldestAbandonedRequestAge
}

// Return sets up results that will be returned by MetricsRegistry.SetOldestAbandonedRequestAge
func (mmSetOldestAbandonedRequestAge *mMetricsRegistryMockSetOldestAbandonedRequestAge) Return() *MetricsRegistryMock {
	if mmSetOldestAbandonedRequestAge.mock.funcSetOldestAbandonedRequestAge != nil {
		mmSetOldestAbandonedRequestAge.mock.t.Fatalf("MetricsRegistryMock.SetOldestAbandonedRequestAge mock is already set by Set")
	}

	if mmSetOldestAbandonedRequestAge.defaultExpectation == nil {
		mmSetOldestAbandonedRequestAge.defaultExpectation = &MetricsRegistryMockSetOldestAbandonedRequestAgeExpectation{mock: mmSetOldestAbandonedRequestAge.mock}
	}

	return mmSetOldestAbandonedRequestAge.mock
}

//Set uses given function f to mock the MetricsRegistry.SetOldestAbandonedRequestAge method
func (mmSetOldestAbandonedRequestAge *mMetricsRegistryMockSetOldestAbandonedRequestAge) Set(f func(age int)) *MetricsRegistryMock {
	if mmSetOldestAbandonedRequestAge.defaultExpectation != nil {
		mmSetOldestAbandonedRequestAge.mock.t.Fatalf("Default expectation is already set for the MetricsRegistry.SetOldestAbandonedRequestAge method")
	}

	if len(mmSetOldestAbandonedRequestAge.expectations) > 0 {
		mmSetOldestAbandonedRequestAge.mock.t.Fatalf("Some expectations are already set for the MetricsRegistry.SetOldestAbandonedRequestAge method")
	}

	mmSetOldestAbandonedRequestAge.mock.funcSetOldestAbandonedRequestAge = f
	return mmSetOldestAbandonedRequestAge.mock
}

// SetOldestAbandonedRequestAge implements MetricsRegistry
func (mmSetOldestAbandonedRequestAge *MetricsRegistryMock) SetOldestAbandonedRequestAge(age int) {
	mm_atomic.AddUint64(&mmSetOldestAbandonedRequestAge.beforeSetOldestAbandonedRequestAgeCounter, 1)
	defer mm_atomic.AddUint64(&mmSetOldestAbandonedRequestAge.afterSetOldestAbandonedRequestAgeCounter, 1)

	if mmSetOldestAbandonedRequestAge.inspectFuncSetOldestAbandonedRequestAge != nil {
		mmSetOldestAbandonedRequestAge.inspectFuncSetOldestAbandonedRequestAge(age)
	}

	mm_params := &MetricsRegistryMockSetOldestAbandonedRequestAgeParams{age}

	// Record call args
	mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.mutex.Lock()
	mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.callArgs = append(mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.callArgs, mm_params)
	mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.mutex.Unlock()

	for _, e := range mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.defaultExpectation.Counter, 1)
		mm_want := mmSetOldestAbandonedRequestAge.SetOldestAbandonedRequestAgeMock.defaultExpectation.params
		mm_got := MetricsRegistryMockSetOldestAbandonedRequestAgeParams{age}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSetOldestAbandonedRequestAge.t.Errorf("MetricsRegistryMock.SetOldestAbandonedRequestAge got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmSetOldestAbandonedRequestAge.funcSetOldestAbandonedRequestAge != nil {
		mmSetOldestAbandonedRequestAge.funcSetOldestAbandonedRequestAge(age)
		return
	}
	mmSetOldestAbandonedRequestAge.t.Fatalf("Unexpected call to MetricsRegistryMock.SetOldestAbandonedRequestAge. %v", age)

}

// SetOldestAbandonedRequestAgeAfterCounter returns a count of finished MetricsRegistryMock.SetOldestAbandonedRequestAge invocations
func (mmSetOldestAbandonedRequestAge *MetricsRegistryMock) SetOldestAbandonedRequestAgeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetOldestAbandonedRequestAge.afterSetOldestAbandonedRequestAgeCounter)
}

// SetOldestAbandonedRequestAgeBeforeCounter returns a count of MetricsRegistryMock.SetOldestAbandonedRequestAge invocations
func (mmSetOldestAbandonedRequestAge *MetricsRegistryMock) SetOldestAbandonedRequestAgeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetOldestAbandonedRequestAge.beforeSetOldestAbandonedRequestAgeCounter)
}

// Calls returns a list of arguments used in each call to MetricsRegistryMock.SetOldestAbandonedRequestAge.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSetOldestAbandonedRequestAge *mMetricsRegistryMockSetOldestAbandonedRequestAge) Calls() []*MetricsRegistryMockSetOldestAbandonedRequestAgeParams {
	mmSetOldestAbandonedRequestAge.mutex.RLock()

	argCopy := make([]*MetricsRegistryMockSetOldestAbandonedRequestAgeParams, len(mmSetOldestAbandonedRequestAge.callArgs))
	copy(argCopy, mmSetOldestAbandonedRequestAge.callArgs)

	mmSetOldestAbandonedRequestAge.mutex.RUnlock()

	return argCopy
}

// MinimockSetOldestAbandonedRequestAgeDone returns true if the count of the SetOldestAbandonedRequestAge invocations corresponds
// the number of defined expectations
func (m *MetricsRegistryMock) MinimockSetOldestAbandonedRequestAgeDone() bool {
	for _, e := range m.SetOldestAbandonedRequestAgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetOldestAbandonedRequestAgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetOldestAbandonedRequestAgeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetOldestAbandonedRequestAge != nil && mm_atomic.LoadUint64(&m.afterSetOldestAbandonedRequestAgeCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetOldestAbandonedRequestAgeInspect logs each unmet expectation
func (m *MetricsRegistryMock) MinimockSetOldestAbandonedRequestAgeInspect() {
	for _, e := range m.SetOldestAbandonedRequestAgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MetricsRegistryMock.SetOldestAbandonedRequestAge with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetOldestAbandonedRequestAgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetOldestAbandonedRequestAgeCounter) < 1 {
		if m.SetOldestAbandonedRequestAgeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MetricsRegistryMock.SetOldestAbandonedRequestAge")
		} else {
			m.t.Errorf("Expected call to MetricsRegistryMock.SetOldestAbandonedRequestAge with params: %#v", *m.SetOldestAbandonedRequestAgeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetOldestAbandonedRequestAge != nil && mm_atomic.LoadUint64(&m.afterSetOldestAbandonedRequestAgeCounter) < 1 {
		m.t.Error("Expected call to MetricsRegistryMock.SetOldestAbandonedRequestAge")
	}
}

type mMetricsRegistryMockUpdateMetrics struct {
	mock               *MetricsRegistryMock
	defaultExpectation *MetricsRegistryMockUpdateMetricsExpectation
	expectations       []*MetricsRegistryMockUpdateMetricsExpectation

	callArgs []*MetricsRegistryMockUpdateMetricsParams
	mutex    sync.RWMutex
}

// MetricsRegistryMockUpdateMetricsExpectation specifies expectation struct of the MetricsRegistry.UpdateMetrics
type MetricsRegistryMockUpdateMetricsExpectation struct {
	mock   *MetricsRegistryMock
	params *MetricsRegistryMockUpdateMetricsParams

	Counter uint64
}

// MetricsRegistryMockUpdateMetricsParams contains parameters of the MetricsRegistry.UpdateMetrics
type MetricsRegistryMockUpdateMetricsParams struct {
	ctx context.Context
}

// Expect sets up expected params for MetricsRegistry.UpdateMetrics
func (mmUpdateMetrics *mMetricsRegistryMockUpdateMetrics) Expect(ctx context.Context) *mMetricsRegistryMockUpdateMetrics {
	if mmUpdateMetrics.mock.funcUpdateMetrics != nil {
		mmUpdateMetrics.mock.t.Fatalf("MetricsRegistryMock.UpdateMetrics mock is already set by Set")
	}

	if mmUpdateMetrics.defaultExpectation == nil {
		mmUpdateMetrics.defaultExpectation = &MetricsRegistryMockUpdateMetricsExpectation{}
	}

	mmUpdateMetrics.defaultExpectation.params = &MetricsRegistryMockUpdateMetricsParams{ctx}
	for _, e := range mmUpdateMetrics.expectations {
		if minimock.Equal(e.params, mmUpdateMetrics.defaultExpectation.params) {
			mmUpdateMetrics.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateMetrics.defaultExpectation.params)
		}
	}

	return mmUpdateMetrics
}

// Inspect accepts an inspector function that has same arguments as the MetricsRegistry.UpdateMetrics
func (mmUpdateMetrics *mMetricsRegistryMockUpdateMetrics) Inspect(f func(ctx context.Context)) *mMetricsRegistryMockUpdateMetrics {
	if mmUpdateMetrics.mock.inspectFuncUpdateMetrics != nil {
		mmUpdateMetrics.mock.t.Fatalf("Inspect function is already set for MetricsRegistryMock.UpdateMetrics")
	}

	mmUpdateMetrics.mock.inspectFuncUpdateMetrics = f

	return mmUpdateMetrics
}

// Return sets up results that will be returned by MetricsRegistry.UpdateMetrics
func (mmUpdateMetrics *mMetricsRegistryMockUpdateMetrics) Return() *MetricsRegistryMock {
	if mmUpdateMetrics.mock.funcUpdateMetrics != nil {
		mmUpdateMetrics.mock.t.Fatalf("MetricsRegistryMock.UpdateMetrics mock is already set by Set")
	}

	if mmUpdateMetrics.defaultExpectation == nil {
		mmUpdateMetrics.defaultExpectation = &MetricsRegistryMockUpdateMetricsExpectation{mock: mmUpdateMetrics.mock}
	}

	return mmUpdateMetrics.mock
}

//Set uses given function f to mock the MetricsRegistry.UpdateMetrics method
func (mmUpdateMetrics *mMetricsRegistryMockUpdateMetrics) Set(f func(ctx context.Context)) *MetricsRegistryMock {
	if mmUpdateMetrics.defaultExpectation != nil {
		mmUpdateMetrics.mock.t.Fatalf("Default expectation is already set for the MetricsRegistry.UpdateMetrics method")
	}

	if len(mmUpdateMetrics.expectations) > 0 {
		mmUpdateMetrics.mock.t.Fatalf("Some expectations are already set for the MetricsRegistry.UpdateMetrics method")
	}

	mmUpdateMetrics.mock.funcUpdateMetrics = f
	return mmUpdateMetrics.mock
}

// UpdateMetrics implements MetricsRegistry
func (mmUpdateMetrics *MetricsRegistryMock) UpdateMetrics(ctx context.Context) {
	mm_atomic.AddUint64(&mmUpdateMetrics.beforeUpdateMetricsCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateMetrics.afterUpdateMetricsCounter, 1)

	if mmUpdateMetrics.inspectFuncUpdateMetrics != nil {
		mmUpdateMetrics.inspectFuncUpdateMetrics(ctx)
	}

	mm_params := &MetricsRegistryMockUpdateMetricsParams{ctx}

	// Record call args
	mmUpdateMetrics.UpdateMetricsMock.mutex.Lock()
	mmUpdateMetrics.UpdateMetricsMock.callArgs = append(mmUpdateMetrics.UpdateMetricsMock.callArgs, mm_params)
	mmUpdateMetrics.UpdateMetricsMock.mutex.Unlock()

	for _, e := range mmUpdateMetrics.UpdateMetricsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmUpdateMetrics.UpdateMetricsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateMetrics.UpdateMetricsMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateMetrics.UpdateMetricsMock.defaultExpectation.params
		mm_got := MetricsRegistryMockUpdateMetricsParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateMetrics.t.Errorf("MetricsRegistryMock.UpdateMetrics got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmUpdateMetrics.funcUpdateMetrics != nil {
		mmUpdateMetrics.funcUpdateMetrics(ctx)
		return
	}
	mmUpdateMetrics.t.Fatalf("Unexpected call to MetricsRegistryMock.UpdateMetrics. %v", ctx)

}

// UpdateMetricsAfterCounter returns a count of finished MetricsRegistryMock.UpdateMetrics invocations
func (mmUpdateMetrics *MetricsRegistryMock) UpdateMetricsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateMetrics.afterUpdateMetricsCounter)
}

// UpdateMetricsBeforeCounter returns a count of MetricsRegistryMock.UpdateMetrics invocations
func (mmUpdateMetrics *MetricsRegistryMock) UpdateMetricsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateMetrics.beforeUpdateMetricsCounter)
}

// Calls returns a list of arguments used in each call to MetricsRegistryMock.UpdateMetrics.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateMetrics *mMetricsRegistryMockUpdateMetrics) Calls() []*MetricsRegistryMockUpdateMetricsParams {
	mmUpdateMetrics.mutex.RLock()

	argCopy := make([]*MetricsRegistryMockUpdateMetricsParams, len(mmUpdateMetrics.callArgs))
	copy(argCopy, mmUpdateMetrics.callArgs)

	mmUpdateMetrics.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateMetricsDone returns true if the count of the UpdateMetrics invocations corresponds
// the number of defined expectations
func (m *MetricsRegistryMock) MinimockUpdateMetricsDone() bool {
	for _, e := range m.UpdateMetricsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMetricsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateMetricsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateMetrics != nil && mm_atomic.LoadUint64(&m.afterUpdateMetricsCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateMetricsInspect logs each unmet expectation
func (m *MetricsRegistryMock) MinimockUpdateMetricsInspect() {
	for _, e := range m.UpdateMetricsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MetricsRegistryMock.UpdateMetrics with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMetricsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateMetricsCounter) < 1 {
		if m.UpdateMetricsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MetricsRegistryMock.UpdateMetrics")
		} else {
			m.t.Errorf("Expected call to MetricsRegistryMock.UpdateMetrics with params: %#v", *m.UpdateMetricsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateMetrics != nil && mm_atomic.LoadUint64(&m.afterUpdateMetricsCounter) < 1 {
		m.t.Error("Expected call to MetricsRegistryMock.UpdateMetrics")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MetricsRegistryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSetOldestAbandonedRequestAgeInspect()

		m.MinimockUpdateMetricsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MetricsRegistryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MetricsRegistryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSetOldestAbandonedRequestAgeDone() &&
		m.MinimockUpdateMetricsDone()
}
