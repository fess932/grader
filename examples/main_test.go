package tests_test

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type TC struct {
	Name         string
	InputData    string
	ExpectedData string

	ActualData string
}

func (tc TC) String() string {
	return fmt.Sprintf("\nInput: %s\nExpected: \n[%s]\nActual: \n[%s]", tc.InputData, tc.ExpectedData, tc.ActualData)
}

func (tc TC) Equal() bool {
	return tc.ExpectedData == tc.ActualData
}

type Exercise struct {
	ID        string
	Lang      string
	TestCases []TC
}

func Test(t *testing.T) {
	tcs := []TC{
		{
			Name:         "zero",
			InputData:    "0",
			ExpectedData: "0",
		},
		{
			Name:         "one",
			InputData:    "1",
			ExpectedData: "1",
		},
		{
			Name:         "11",
			InputData:    "11",
			ExpectedData: "89",
		},
	}

	exercises := []Exercise{
		{
			ID:        "1",
			Lang:      "go",
			TestCases: tcs,
		},
		{
			ID:        "1",
			Lang:      "python",
			TestCases: tcs,
		},
	}

	t.Parallel()

	for _, exrc := range exercises {
		switch exrc.Lang {
		case "go":
			t.Log("Testing Go")
			Run(t, RunGo, exrc.TestCases)
		case "python":
			t.Log("Testing Python")
			Run(t, RunPython, exrc.TestCases)
		default:
			t.Fatal("unknown language")
		}
	}
}

func Run(t *testing.T, tFunc func(r io.Reader, w io.Writer) error, tcs []TC) {
	t.Helper()

	var (
		buf = new(bytes.Buffer)
		r   = new(strings.Reader)
		err error
	)

	for _, v := range tcs {
		t.Run(v.Name, func(t *testing.T) {
			r = strings.NewReader(v.InputData)
			err = tFunc(r, buf)
			v.ActualData = buf.String()
			buf.Reset()

			if err != nil {
				t.Error(err)
				t.Error(v.ActualData)

				return
			}

			if !v.Equal() {
				t.Error(v)

				return
			}
		})
	}
}

func RunGo(r io.Reader, w io.Writer) error {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = "go/task"
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// 1/go
	return nil
}

func RunPython(r io.Reader, w io.Writer) error {
	cmd := exec.Command("python", "main.py")
	cmd.Dir = "python"
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
