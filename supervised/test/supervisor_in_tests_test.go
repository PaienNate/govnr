// Copyright 2019 the orbs-network-go authors
// This file is part of the orbs-network-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import "testing"

var expectedLogsOnPanic = []string{
	Failed, BeforeLoggerCreated, LoggedWithLogger, BeforeCallPanic, PanicOhNo,
}
var unexpectedLogsOnPanic = []string{
	Passed, AfterCallPanic, MustNotShow,
}

var expectedLogsOnLogError = []string{
	Failed, BeforeLoggerCreated, LoggedWithLogger, BeforeLoggerError, ErrorWithLogger, AfterLoggerError, MustShow,
}
var unexpectedLogsOnLogError = []string{
	Passed, MustNotShow,
}

func Test_Panics(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnPanic, unexpectedLogsOnPanic)
}

func Test_LogsError(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnLogError, unexpectedLogsOnLogError)
}

func TestGoOnce_Panics(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnPanic, unexpectedLogsOnPanic)
}

func TestGoOnce_LogsError(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnLogError, unexpectedLogsOnLogError)
}

func TestTRun_Panics(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnPanic, unexpectedLogsOnPanic)
}

func TestTRun_LogsError(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnLogError, unexpectedLogsOnLogError)
}

func TestTRun_GoOnce_Panics(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnPanic, unexpectedLogsOnPanic)
}

func TestTRun_GoOnce_LogsError(t *testing.T) {
	executeGoTestRunner(t, expectedLogsOnLogError, unexpectedLogsOnLogError)
}

func TestTRun_GoOnce_PanicsAfterSubTestPasses(t *testing.T) {
	expect := []string{
		Passed, BeforeLoggerCreated, LoggedWithLogger, PanicOhNo,
	}
	expectNot := []string{
		AfterCallPanic, MustNotShow, // TODO explain why Failed is excluded
	}
	executeGoTestRunner(t, expect, expectNot)
}

func TestTRun_GoOnce_LogsErrorAfterSubTestPasses(t *testing.T) {
	expect := []string{
		Passed, BeforeLoggerCreated, LoggedWithLogger, ErrorWithLogger,
	}
	expectNot := []string{
		MustNotShow, // TODO explain why Failed is excluded
	}
	executeGoTestRunner(t, expect, expectNot)
}
