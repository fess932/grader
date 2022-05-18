package go_test

import (
	"bytes"
	"grader/examples/go/task"
	"strings"
	"testing"
)

var testData = "1 2"
var result = "3"

type TC struct {
	Expected string
	Actual   string
}

func (tc TC) String() string {
	return "Expected: " + tc.Expected + "\nActual: " + tc.Actual
}

func (tc TC) Equal() bool {
	if tc.Expected != tc.Actual {
		return false
	}

	return true
}

type Exercise struct {
	ID        string
	Lang      string
	TestCases []struct {
		InputData    string
		ExpectedData string
	}
}

func Test(t *testing.T) {
	exercises := []Exercise{
		{
			ID:   "1",
			Lang: "go",
			TestCases: []struct {
				InputData    string
				ExpectedData string
			}{
				{
					InputData:    testData,
					ExpectedData: result,
				},
			},
		},
		{
			ID:   "1",
			Lang: "python",
			TestCases: []struct {
				InputData    string
				ExpectedData string
			}{
				{
					InputData:    testData,
					ExpectedData: result,
				},
			},
		},
	}

	for _, exrc := range exercises {
		switch exrc.Lang {
		case "go":
			for _, v := range exrc.TestCases {
				t.Run(exrc.ID, func(t *testing.T) {

					var buf bytes.Buffer
					task.Run(strings.NewReader(v.InputData), &buf)

					if !TC{v.ExpectedData, buf.String()}.Equal() {
						t.Error(TC{v.ExpectedData, buf.String()})
					}
				})
			}

			RunGo()
		case "python":
			RunPython()
		}
	}

	{
		buf := bytes.NewBufferString("")
		task.Task(strings.NewReader(testData), buf)
		tc := TC{
			Expected: result,
			Actual:   buf.String(),
		}

		if !tc.Equal() {
			t.Error(tc.String())
		}
	}
}

func Run() {

}

func RunPython() {

	1/python
}

func RunGo() {

}
