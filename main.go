package main

import (
	"os"
)

func main() {
	arg := "y"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	arg += "\n"

	var buf []byte
	for len(buf)+len(arg) <= 4096*16 {
		buf = append(buf, []byte(arg)...)
	}

	for {
		os.Stdout.Write(buf)
	}
}
