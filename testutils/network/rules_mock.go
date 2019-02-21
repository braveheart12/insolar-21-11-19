package network

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Rules" can be found in github.com/insolar/insolar/network
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
)

//RulesMock implements github.com/insolar/insolar/network.Rules
type RulesMock struct {
	t minimock.Tester

	CheckMajorityRuleFunc       func() (r bool, r1 int)
	CheckMajorityRuleCounter    uint64
	CheckMajorityRulePreCounter uint64
	CheckMajorityRuleMock       mRulesMockCheckMajorityRule

	CheckMinRoleRuleFunc       func() (r bool)
	CheckMinRoleRuleCounter    uint64
	CheckMinRoleRulePreCounter uint64
	CheckMinRoleRuleMock       mRulesMockCheckMinRoleRule
}

//NewRulesMock returns a mock for github.com/insolar/insolar/network.Rules
func NewRulesMock(t minimock.Tester) *RulesMock {
	m := &RulesMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckMajorityRuleMock = mRulesMockCheckMajorityRule{mock: m}
	m.CheckMinRoleRuleMock = mRulesMockCheckMinRoleRule{mock: m}

	return m
}

type mRulesMockCheckMajorityRule struct {
	mock              *RulesMock
	mainExpectation   *RulesMockCheckMajorityRuleExpectation
	expectationSeries []*RulesMockCheckMajorityRuleExpectation
}

type RulesMockCheckMajorityRuleExpectation struct {
	result *RulesMockCheckMajorityRuleResult
}

type RulesMockCheckMajorityRuleResult struct {
	r  bool
	r1 int
}

//Expect specifies that invocation of Rules.CheckMajorityRule is expected from 1 to Infinity times
func (m *mRulesMockCheckMajorityRule) Expect() *mRulesMockCheckMajorityRule {
	m.mock.CheckMajorityRuleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RulesMockCheckMajorityRuleExpectation{}
	}

	return m
}

//Return specifies results of invocation of Rules.CheckMajorityRule
func (m *mRulesMockCheckMajorityRule) Return(r bool, r1 int) *RulesMock {
	m.mock.CheckMajorityRuleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RulesMockCheckMajorityRuleExpectation{}
	}
	m.mainExpectation.result = &RulesMockCheckMajorityRuleResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of Rules.CheckMajorityRule is expected once
func (m *mRulesMockCheckMajorityRule) ExpectOnce() *RulesMockCheckMajorityRuleExpectation {
	m.mock.CheckMajorityRuleFunc = nil
	m.mainExpectation = nil

	expectation := &RulesMockCheckMajorityRuleExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *RulesMockCheckMajorityRuleExpectation) Return(r bool, r1 int) {
	e.result = &RulesMockCheckMajorityRuleResult{r, r1}
}

//Set uses given function f as a mock of Rules.CheckMajorityRule method
func (m *mRulesMockCheckMajorityRule) Set(f func() (r bool, r1 int)) *RulesMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.CheckMajorityRuleFunc = f
	return m.mock
}

//CheckMajorityRule implements github.com/insolar/insolar/network.Rules interface
func (m *RulesMock) CheckMajorityRule() (r bool, r1 int) {
	counter := atomic.AddUint64(&m.CheckMajorityRulePreCounter, 1)
	defer atomic.AddUint64(&m.CheckMajorityRuleCounter, 1)

	if len(m.CheckMajorityRuleMock.expectationSeries) > 0 {
		if counter > uint64(len(m.CheckMajorityRuleMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RulesMock.CheckMajorityRule.")
			return
		}

		result := m.CheckMajorityRuleMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RulesMock.CheckMajorityRule")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.CheckMajorityRuleMock.mainExpectation != nil {

		result := m.CheckMajorityRuleMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RulesMock.CheckMajorityRule")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.CheckMajorityRuleFunc == nil {
		m.t.Fatalf("Unexpected call to RulesMock.CheckMajorityRule.")
		return
	}

	return m.CheckMajorityRuleFunc()
}

//CheckMajorityRuleMinimockCounter returns a count of RulesMock.CheckMajorityRuleFunc invocations
func (m *RulesMock) CheckMajorityRuleMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CheckMajorityRuleCounter)
}

//CheckMajorityRuleMinimockPreCounter returns the value of RulesMock.CheckMajorityRule invocations
func (m *RulesMock) CheckMajorityRuleMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CheckMajorityRulePreCounter)
}

//CheckMajorityRuleFinished returns true if mock invocations count is ok
func (m *RulesMock) CheckMajorityRuleFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.CheckMajorityRuleMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.CheckMajorityRuleCounter) == uint64(len(m.CheckMajorityRuleMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.CheckMajorityRuleMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.CheckMajorityRuleCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.CheckMajorityRuleFunc != nil {
		return atomic.LoadUint64(&m.CheckMajorityRuleCounter) > 0
	}

	return true
}

type mRulesMockCheckMinRoleRule struct {
	mock              *RulesMock
	mainExpectation   *RulesMockCheckMinRoleRuleExpectation
	expectationSeries []*RulesMockCheckMinRoleRuleExpectation
}

type RulesMockCheckMinRoleRuleExpectation struct {
	result *RulesMockCheckMinRoleRuleResult
}

type RulesMockCheckMinRoleRuleResult struct {
	r bool
}

//Expect specifies that invocation of Rules.CheckMinRoleRule is expected from 1 to Infinity times
func (m *mRulesMockCheckMinRoleRule) Expect() *mRulesMockCheckMinRoleRule {
	m.mock.CheckMinRoleRuleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RulesMockCheckMinRoleRuleExpectation{}
	}

	return m
}

//Return specifies results of invocation of Rules.CheckMinRoleRule
func (m *mRulesMockCheckMinRoleRule) Return(r bool) *RulesMock {
	m.mock.CheckMinRoleRuleFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &RulesMockCheckMinRoleRuleExpectation{}
	}
	m.mainExpectation.result = &RulesMockCheckMinRoleRuleResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of Rules.CheckMinRoleRule is expected once
func (m *mRulesMockCheckMinRoleRule) ExpectOnce() *RulesMockCheckMinRoleRuleExpectation {
	m.mock.CheckMinRoleRuleFunc = nil
	m.mainExpectation = nil

	expectation := &RulesMockCheckMinRoleRuleExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *RulesMockCheckMinRoleRuleExpectation) Return(r bool) {
	e.result = &RulesMockCheckMinRoleRuleResult{r}
}

//Set uses given function f as a mock of Rules.CheckMinRoleRule method
func (m *mRulesMockCheckMinRoleRule) Set(f func() (r bool)) *RulesMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.CheckMinRoleRuleFunc = f
	return m.mock
}

//CheckMinRoleRule implements github.com/insolar/insolar/network.Rules interface
func (m *RulesMock) CheckMinRoleRule() (r bool) {
	counter := atomic.AddUint64(&m.CheckMinRoleRulePreCounter, 1)
	defer atomic.AddUint64(&m.CheckMinRoleRuleCounter, 1)

	if len(m.CheckMinRoleRuleMock.expectationSeries) > 0 {
		if counter > uint64(len(m.CheckMinRoleRuleMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to RulesMock.CheckMinRoleRule.")
			return
		}

		result := m.CheckMinRoleRuleMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the RulesMock.CheckMinRoleRule")
			return
		}

		r = result.r

		return
	}

	if m.CheckMinRoleRuleMock.mainExpectation != nil {

		result := m.CheckMinRoleRuleMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the RulesMock.CheckMinRoleRule")
		}

		r = result.r

		return
	}

	if m.CheckMinRoleRuleFunc == nil {
		m.t.Fatalf("Unexpected call to RulesMock.CheckMinRoleRule.")
		return
	}

	return m.CheckMinRoleRuleFunc()
}

//CheckMinRoleRuleMinimockCounter returns a count of RulesMock.CheckMinRoleRuleFunc invocations
func (m *RulesMock) CheckMinRoleRuleMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CheckMinRoleRuleCounter)
}

//CheckMinRoleRuleMinimockPreCounter returns the value of RulesMock.CheckMinRoleRule invocations
func (m *RulesMock) CheckMinRoleRuleMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CheckMinRoleRulePreCounter)
}

//CheckMinRoleRuleFinished returns true if mock invocations count is ok
func (m *RulesMock) CheckMinRoleRuleFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.CheckMinRoleRuleMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.CheckMinRoleRuleCounter) == uint64(len(m.CheckMinRoleRuleMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.CheckMinRoleRuleMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.CheckMinRoleRuleCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.CheckMinRoleRuleFunc != nil {
		return atomic.LoadUint64(&m.CheckMinRoleRuleCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *RulesMock) ValidateCallCounters() {

	if !m.CheckMajorityRuleFinished() {
		m.t.Fatal("Expected call to RulesMock.CheckMajorityRule")
	}

	if !m.CheckMinRoleRuleFinished() {
		m.t.Fatal("Expected call to RulesMock.CheckMinRoleRule")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *RulesMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *RulesMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *RulesMock) MinimockFinish() {

	if !m.CheckMajorityRuleFinished() {
		m.t.Fatal("Expected call to RulesMock.CheckMajorityRule")
	}

	if !m.CheckMinRoleRuleFinished() {
		m.t.Fatal("Expected call to RulesMock.CheckMinRoleRule")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *RulesMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *RulesMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.CheckMajorityRuleFinished()
		ok = ok && m.CheckMinRoleRuleFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.CheckMajorityRuleFinished() {
				m.t.Error("Expected call to RulesMock.CheckMajorityRule")
			}

			if !m.CheckMinRoleRuleFinished() {
				m.t.Error("Expected call to RulesMock.CheckMinRoleRule")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *RulesMock) AllMocksCalled() bool {

	if !m.CheckMajorityRuleFinished() {
		return false
	}

	if !m.CheckMinRoleRuleFinished() {
		return false
	}

	return true
}
