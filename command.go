package main

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

type Cmd struct {
	name string
	narg int
	args []string
	exec func(*Cmd) (error, CmdExecCallBack)
}

func (cmd *Cmd) Exec() (error, CmdExecCallBack) {
	return cmd.exec(cmd)
}

type cmdValue struct {
	narg int
	exec func(cmd *Cmd) (error, CmdExecCallBack)
}

type cmdMap map[string]cmdValue

var gCmdMap map[string]cmdValue = map[string]cmdValue{
	"echo": {narg: -1, exec: echo},
	"req":  {narg: 2, exec: req},
	"ret":  {narg: 1, exec: ret},
}

func CmdParser(line string) (*Cmd, error) {
	ss := strings.SplitN(line, " ", 2)
	if len(ss[0]) == 0 {
		return nil, errors.New("empty cmd")
	}
	if v, ok := gCmdMap[ss[0]]; ok {
		cmd := &Cmd{
			name: ss[0],
			narg: v.narg,
			exec: v.exec,
		}
		return parseCmdArgs(cmd, ss)
	}
	return nil, errors.New("cmd " + ss[0] + " not found")
}

func fillCmdArgs(cmd *Cmd, argStr string) error {
	index := 0
	remains := []rune(argStr)
	asm := NewASM()

	for index < len(remains) {
		arg, next, err := asm.GetString(remains[index:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		index += next
		cmd.args = append(cmd.args, arg)
		if cmd.narg > 0 && len(cmd.args) > cmd.narg {
			return errors.New("too many args for cmd: " + cmd.name +
				", only " + strconv.Itoa(cmd.narg) + " args expected")
		}
	}
	if cmd.narg < 0 {
		return nil
	}
	if cmd.narg != len(cmd.args) {
		return errors.New("unexpected " + strconv.Itoa(len(cmd.args)) +
			" args for cmd: " + cmd.name + ", only " +
			strconv.Itoa(cmd.narg) + " args expected")
	}
	return nil
}

func parseCmdArgs(cmd *Cmd, ss []string) (*Cmd, error) {
	switch {
	case cmd.narg == 0:
		if len(ss) != 1 {
			return nil, errors.New("too many args for cmd: " + cmd.name)
		}

	case cmd.narg < 0:
		if len(ss) == 1 {
			return cmd, nil
		}
		if err := fillCmdArgs(cmd, ss[1]); err != nil {
			return nil, err
		}

	case cmd.narg > 0:
		if len(ss) == 1 {
			return nil, errors.New("missing args for cmd: " + cmd.name)
		}
		if err := fillCmdArgs(cmd, ss[1]); err != nil {
			return nil, err
		}
	}
	return cmd, nil
}
