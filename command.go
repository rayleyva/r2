package main

import (
	"errors"
	"io"
	"strings"
)

type Cmd struct {
	name string
	narg int
	args []string
	exec func(*Cmd) error
}

func (cmd *Cmd) Exec() error {
	return cmd.exec(cmd)
}

type cmdValue struct {
	narg int
	exec func(cmd *Cmd) error
}

type cmdMap map[string]cmdValue

var gCmdMap map[string]cmdValue = map[string]cmdValue{
	"echo": {narg: -1, exec: echo},
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

func parseCmdArgs(cmd *Cmd, ss []string) (*Cmd, error) {
	if cmd.narg == 0 {
		if len(ss) != 1 {
			return nil, errors.New("too many args for cmd: " + cmd.name)
		}
		return cmd, nil
	}
	if cmd.narg < 0 {
		if len(ss) == 1 {
			return cmd, nil
		}

		index := 0
		remains := []rune(ss[1])

		asm := NewASM()
		for index < len(remains) {
			arg, next, err := asm.GetString(remains[index:])
			if err != nil {
				if err == io.EOF {
					return cmd, nil
				}
				return nil, err
			}
			index += next
			cmd.args = append(cmd.args, arg)
		}
	}
	return cmd, nil
}
