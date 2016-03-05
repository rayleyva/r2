package main

import (
	"fmt"
	"strconv"
)

const (
	RED     = "\033[031m"
	GREEN   = "\033[032m"
	BROWN   = "\033[033m"
	BLUE    = "\033[034m"
	MAGENTA = "\033[035m"
	CYAN    = "\033[036m"
	RECOVER = "\033[0m"
)

func echoColor(color string, args ...interface{}) {
	fmt.Print(color)
	fmt.Println(args...)
	fmt.Print(RECOVER)
}

func Red(args ...interface{}) {
	echoColor(RED, args...)
}

func Green(args ...interface{}) {
	echoColor(GREEN, args...)
}

func Brown(args ...interface{}) {
	echoColor(BROWN, args...)
}

func Blue(args ...interface{}) {
	echoColor(BLUE, args...)
}

func Magenta(args ...interface{}) {
	echoColor(MAGENTA, args...)
}

func Cyan(args ...interface{}) {
	echoColor(CYAN, args...)
}

type CmdExecCallBack func(file string, line int)

func FileLine(file string, line int) string {
	return "<" + file + ":" + strconv.Itoa(line) + ">"
}
