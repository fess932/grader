package go_test

import "testing"

func Test(t *testing.T) {
	t.Parallel()
	t.Log("Hello, world.")
}

//
//import (
//	"bytes"
//	"github.com/rs/zerolog/log"
//	"grader/examples/go/task"
//	"strings"
//)
//
//var testData = "1 2"
//var result = "3"
//
//func main() {
//	buf := bytes.NewBufferString("")
//	task.Task(strings.NewReader(testData), buf)
//
//	if !bytes.Equal(buf.Bytes(), []byte(result)) {
//		log.Info().Msgf("Response not equal: %s: %s", buf.String(), result)
//	}
//}
