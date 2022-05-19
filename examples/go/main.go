package main

import (
	"io"
	"os"
	"strconv"
)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	num, err := strconv.Atoi(string(buf))
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	os.Stdout.Write([]byte(strconv.Itoa(Fib(num))))
}
