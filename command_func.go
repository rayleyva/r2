package main

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
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

// body "{\"id\":1,\"value\":[1,2,3],\"comment\":\"just for testing\"}"
func body(cmd *Cmd) (error, CmdExecCallBack) {
	if !gReq.inited {
		return errors.New("req has not been inited"), nil
	}
	gReq.body = bytes.NewBuffer([]byte(cmd.args[0]))
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

// header-equal Content-Type application/json
func headerEqual(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	header := gRep.rawRep.Header
	v := strings.TrimSpace(header.Get(cmd.args[0]))
	if len(v) == 0 {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line),
				"missing header:", cmd.args[0])
		}
	}

	if v == cmd.args[1] {
		return nil, func(file string, line int) {
			Green("[PASS]", FileLine(file, line), cmd.args[0], cmd.args[1])
		}
	}

	return nil, func(file string, line int) {
		Red("[FAIL]", FileLine(file, line),
			"header:", cmd.args[0],
			"return:", v,
			"expected:", cmd.args[1])
	}
}

// body-equal {"errno":0,"errmsg":"ok","result":[200]}
func bodyEqual(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	if err := gRep.readBody(); err != nil {
		return err, nil
	}

	expectedBody := []byte(cmd.args[0])
	if !bytes.Equal(expectedBody, gRep.body) {
		return nil, func(file string, line int) {
			Green("[PASS]", FileLine(file, line), "body equal")
		}
	}
	return nil, func(file string, line int) {
		Red("[FAIL]", FileLine(file, line), "body un-equal")
	}
}

// header-match key <regexp>
func headerMatch(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	header := gRep.rawRep.Header
	v := strings.TrimSpace(header.Get(cmd.args[0]))
	if len(v) == 0 {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line),
				"missing header:", cmd.args[0])
		}
	}

	re, err := regexp.Compile(cmd.args[1])
	if err != nil {
		return err, nil
	}

	if !re.MatchString(v) {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line), "header un-match: ", v)
		}
	}
	return nil, func(file string, line int) {
		Green("[PASS]", FileLine(file, line), "header match")
	}
}

// body-match <regexp>
func bodyMatch(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	if err := gRep.readBody(); err != nil {
		return err, nil
	}

	re, err := regexp.Compile(cmd.args[0])
	if err != nil {
		return err, nil
	}

	body := string(gRep.body)
	if !re.Match(gRep.body) {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line), "body un-match")
		}
	}

	// get vars
	for i, n := range re.SubexpNames() {
		if i != 0 && len(strings.TrimSpace(n)) != 0 {
			dst := make([]byte, 0, 1024)
			matches := re.FindStringSubmatchIndex(body)
			v := re.ExpandString(dst, "${"+n+"}", body, matches)
			if len(v) != 0 {
				gRep.vars[n] = string(v)
			}
		}
	}
	return nil, func(file string, line int) {
		Green("[PASS]", FileLine(file, line), "body matched")
	}
}

// var-equal var_name expected_value
func varEqual(cmd *Cmd) (error, CmdExecCallBack) {
	k := cmd.args[0]
	v, ok := gRep.vars[k]
	if !ok {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line), "var:", k, "not exist")
		}
	}

	if v != cmd.args[1] {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line), "var:", k,
				"value:", v, ", expected:", cmd.args[1])
		}
	}
	return nil, func(file string, line int) {
		Green("[PASS]", FileLine(file, line), "var: ", k, "equal")
	}
}

// var-echo var_name
func varEcho(cmd *Cmd) (error, CmdExecCallBack) {
	k := cmd.args[0]
	v, ok := gRep.vars[k]
	if !ok {
		return nil, func(file string, line int) {
			Red("[FAIL]", FileLine(file, line), "var:", k, "not exist")
		}
	}
	return nil, func(file string, line int) {
		Cyan("[VAR]", FileLine(file, line), k, "==>", v)
	}
}

// internal debug command
func bodyEcho(cmd *Cmd) (error, CmdExecCallBack) {
	if err := gReq.Launch(gRep); err != nil {
		return err, nil
	}
	if err := gRep.readBody(); err != nil {
		return err, nil
	}
	fmt.Println(string(gRep.body))
	fmt.Println(len(gRep.body))
	return nil, nil
}
