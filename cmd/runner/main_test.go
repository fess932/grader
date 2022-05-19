package runner_test

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type TC struct {
	Name         string `yaml:"name"`
	InputData    string `yaml:"input"`
	ExpectedData string `yaml:"expected"`

	ActualData string
}

func (tc TC) String() string {
	return fmt.Sprintf("\nInput: %s\nExpected: \n[%s]\nActual: \n[%s]", tc.InputData, tc.ExpectedData, tc.ActualData)
}

func (tc TC) Equal() bool {
	return tc.ExpectedData == tc.ActualData
}

type Exercise struct {
	ID        string `yaml:"id"`
	Lang      string `yaml:"lang"`
	Homedir   string `yaml:"homedir"`
	TestCases []TC   `yaml:"tests"`
}

func Test(t *testing.T) {
	var exc Exercise

	f, err := os.Open("tests.yaml")
	if err != nil {

		t.Fatalf("failed to open tests.yaml: %v", err)
	}

	if err = yaml.NewDecoder(f).Decode(&exc); err != nil {
		t.Fatalf("failed to decode tests.yaml: %v", err)
	}

	log.Printf("Running [%+v]", exc)

	t.Parallel()

	switch exc.Lang {
	case "go":
		t.Log("Testing Go")
		Run(t, RunGo, exc.TestCases, exc.Homedir)
	case "python":
		t.Log("Testing Python")

		t.Log(exc.Lang, exc.Homedir)

		Run(t, RunPython, exc.TestCases, exc.Homedir)
	default:
		t.Fatal("unknown language")
	}
}

func Run(t *testing.T, tFunc func(r io.Reader, w io.Writer, homedir string) error, tcs []TC, homedir string) {
	t.Helper()

	var (
		buf = new(bytes.Buffer)
		r   = new(strings.Reader)
		err error
	)

	for _, v := range tcs {
		t.Run(v.Name, func(t *testing.T) {
			r = strings.NewReader(v.InputData)
			err = tFunc(r, buf, homedir)
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

func RunGo(r io.Reader, w io.Writer, homedir string) error {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = homedir
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// 1/go
	return nil
}

func RunPython(r io.Reader, w io.Writer, homedir string) error {
	cmd := exec.Command("python", "main.py")
	cmd.Dir = homedir
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
