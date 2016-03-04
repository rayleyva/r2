package main

import (
	"os"
)

func usage() {
	cmd := os.Args[0]
	Yellow("Usage of ", cmd)
	Yellow("    ", cmd, "<mytest.l> ...")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	for _, file := range os.Args[1:] {
		if line, err := Launch(file); err != nil {
			Yellow("[Error]", FileLine(file, line), ">>>", err.Error())
		}
	}
}
