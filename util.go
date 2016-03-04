package main

import (
	"fmt"
	"strconv"
)

func Red(args ...interface{}) {
	fmt.Print("\033[031m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

func Green(args ...interface{}) {
	fmt.Print("\033[032m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

func Brown(args ...interface{}) {
	fmt.Print("\033[033m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

func Blue(args ...interface{}) {
	fmt.Print("\033[034m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

func Magenta(args ...interface{}) {
	fmt.Print("\033[035m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

func Cyan(args ...interface{}) {
	fmt.Print("\033[036m")
	fmt.Println(args...)
	fmt.Print("\033[0m")
}

type CmdExecCallBack func(file string, line int)

func FileLine(file string, line int) string {
	return "<" + file + ":" + strconv.Itoa(line) + ">"
}
