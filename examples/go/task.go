package main

import "io"

func Task(in io.Reader, out io.Writer) {
	io.WriteString(out, "Hello, World!\n")
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
