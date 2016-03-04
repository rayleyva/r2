package main

import (
	"fmt"
	"strconv"
	"strings"
)

// echo msg1 msg2 ... msgN
func echo(cmd *Cmd) (error, CmdExecCallBack) {
	args := make([]interface{}, len(cmd.args))
	for i := 0; i < len(args); i++ {
		args[i] = interface{}(cmd.args[i])
	}
	_, err := fmt.Println(args...)
	return err, nil
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
