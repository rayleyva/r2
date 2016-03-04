package main

import (
	"fmt"
)

func echo(cmd *Cmd) error {
	args := make([]interface{}, len(cmd.args))
	for i := 0; i < len(args); i++ {
		args[i] = interface{}(cmd.args[i])
	}
	_, err := fmt.Println(args...)
	return err
}
