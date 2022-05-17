package task

import "io"

func Task(in io.Reader, out io.Writer) {
	io.WriteString(out, "Hello, World!\n")
}
