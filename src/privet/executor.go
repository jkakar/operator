package privet

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type TestBatchExecutor func(commandPath string, worker int, batchIndex int, batch *PlanTestBatch) (bool, error)

var execTestExecutor = func(commandPath string, worker int, batchIndex int, batch *PlanTestBatch) (bool, error) {
	testFiles := []string{}
	testCaseNames := []string{}
	for _, execution := range batch.TestExecutions {
		testFiles = append(testFiles, execution.File)
		testCaseNames = append(testCaseNames, execution.TestCaseNames...)
	}

	cmd := &exec.Cmd{
		Path: commandPath,
		Env: []string{
			fmt.Sprintf("PRIVET_TEST_FILES=%s", strings.Join(testFiles, "\x00")),
			fmt.Sprintf("PRIVET_TEST_CASE_NAMES=%s", strings.Join(testCaseNames, "\x00")),
			fmt.Sprintf("PRIVET_TEST_BATCH=%d", batchIndex),
		},
		Stdin:  nil,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		// The command exited non-zero
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

type PlanExecutionOpts struct {
	// CommandPath is the path of the command to run as part of plan
	// execution.
	//
	// The command will receive two environment variables which it should
	// use to execute the requested test case(s):
	// * PRIVET_TEST_FILES: Test files (separated by \0)
	// * PRIVET_TEST_CASE_NAMES: Test cases (separated by \0)
	// * PRIVET_TEST_BATCH: Strictly increasing integer identifier of the
	//                      test batch
	//
	// If PRIVET_TEST_CASE_NAMES is specified, only one test file will be
	// present in PRIVET_TEST_FILES.
	CommandPath string

	Worker int

	TestBatchExecutor TestBatchExecutor
}

// ExecutePlan executes the specified test plan. It returns true if the test
// plan executed successfully; otherwise, false. An error is returned only if an
// egregious error happens while executing the plan. An error will not be present
// if one of the plan command invocation returns non-zero.
func ExecutePlan(plan *Plan, opts *PlanExecutionOpts) (bool, error) {
	executor := opts.TestBatchExecutor
	if executor == nil {
		executor = execTestExecutor
	}

	worker, ok := plan.Workers[opts.Worker]
	if !ok {
		return false, fmt.Errorf("worker %d not found in plan", opts.Worker)
	}

	overallSuccess := true
	for batchIndex, batch := range worker.TestBatches {
		success, err := executor(opts.CommandPath, opts.Worker, batchIndex, batch)
		if err != nil {
			return false, err
		} else if !success {
			overallSuccess = false
		}
	}
	return overallSuccess, nil
}
