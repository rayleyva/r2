package main

import (
	"fmt"
	"os"
)

var _p func(string, ...interface{}) = func(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func usage() {
	cmd := os.Args[0]
	_p("Usage of %s:\n", cmd)
	_p("    %s <mytest.l> ...\n", cmd)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	for _, script := range os.Args[1:] {
		f, err := os.Open(script)
		if err != nil {
			_p("[Error] Open [%s] error: %s\n",
				script, err.Error())
			continue
		}

		if line, err := Launch(f); err != nil {
			_p("[Error] %s:%d >>> %s\n",
				script, line, err.Error())
		}

		f.Close()
	}
}
