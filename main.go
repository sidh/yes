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
	for {
		buf = append(buf, []byte(arg)...)
		if len(buf)+len(arg) > 4096*16 {
			break
		}
	}

	for {
		os.Stdout.Write(buf)
	}
}
