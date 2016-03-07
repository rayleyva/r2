package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func echoF(cmd *Cmd, f func(...interface{})) (error, CmdExecCallBack) {
	args := make([]interface{}, len(cmd.args))
	for i := 0; i < len(args); i++ {
		args[i] = interface{}(cmd.args[i])
	}
	f(args...)
	return nil, nil
}

// echo msg1 msg2 ... msgN
func echo(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, func(args ...interface{}) {
		fmt.Println(args...)
	})
}

// red msg1 msg2 ... msgN
func red(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Red)
}

// green msg1 msg2 ... msgN
func green(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Green)
}

// brown msg1 msg2 ... msgN
func brown(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Brown)
}

// blue msg1 msg2 ... msgN
func blue(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Blue)
}

// magenta msg1 msg2 ... msgN
func magenta(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Magenta)
}

// Cyan msg1 msg2 ... msgN
func cyan(cmd *Cmd) (error, CmdExecCallBack) {
	return echoF(cmd, Cyan)
}

// req get https://github.com/
func req(cmd *Cmd) (error, CmdExecCallBack) {
	gReq.Cleanup()
	gRep.Cleanup()
	gReq.Init()

	if err := gReq.SetMethod(strings.ToUpper(cmd.args[0])); err != nil {
		return err, nil
	}

	if err := gReq.SetUrl(cmd.args[1]); err != nil {
		return err, nil
	}
	return nil, nil
}

// header Content-Type application/json
func header(cmd *Cmd) (error, CmdExecCallBack) {
	if !gReq.inited {
		return errors.New("req has not been inited"), nil
	}
	gReq.header[cmd.args[0]] = cmd.args[1]
	return nil, nil
}

// ret 200
func ret(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	statusCode := gRep.rawRep.StatusCode
	expected, _ := strconv.Atoi(cmd.args[0])
	if statusCode != expected {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line),
				"StatusCode:", statusCode,
				"expected:", expected)
		}
	}
	return nil, func(file string, line int) {
		Green("[PASS]", FileLine(file, line), "ret", expected)
	}
}
