package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func findExpected(toRun *command) (expectedOutput []byte, err error) {
	cmd := exec.Command(toRun.name, toRun.args...)
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("error creating stdout pipe: %w", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting command: %w", err)
	}

	go killAfterDuration(cmd, time.Duration(*timeToWait)*time.Millisecond)

	output, err := ioutil.ReadAll(stdOut)
	if err != nil {
		return nil, fmt.Errorf("error reading stdout pipe: %w", err)
	}

	return output, nil
}
