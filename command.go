package main

import (
	"errors"
	"strings"
)

type command struct {
	name string
	args []string
}

func parseCommand(s string) (result *command, err error) {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return nil, errors.New("empty command string")
	}

	result = new(command)
	result.name = fields[0]
	result.args = fields[1:]
	return
}
