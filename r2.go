package main

import (
	"os"
)

func usage() {
	cmd := os.Args[0]
	Brown("Usage of ", cmd)
	Brown("    ", cmd, "<mytest.r2> ...")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	for _, file := range os.Args[1:] {
		if line, err := Launch(file); err != nil {
			Brown("[Error]", FileLine(file, line), ">>>", err.Error())
		}
	}
}
