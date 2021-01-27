package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/cli-tti/util"
	"log"
	"os/exec"
	"syscall"
	"time"
)

func timeCommandAverage(command *command, expectedOutput []byte, numTime int64) (duration time.Duration, err error) {
	total := time.Duration(0)

	for i := int64(0); i < numTime; i++ {
		duration, err := timeCommand(command, expectedOutput)
		if err != nil {
			return 0, err
		}

		total += duration
		fmt.Printf("Time %d: %d milliseconds\n", i, duration.Milliseconds())
	}

	return time.Duration(total.Nanoseconds() / numTime) * time.Nanosecond, nil
}

func timeCommand(command *command, expectedOutput []byte) (duration time.Duration, err error) {
	cmd := exec.Command(command.name, command.args...)
	cmd.Stderr = nil
	cmd.Stdin = nil

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("error opening stdout pipe: %w", err)
	}

	scanner := bufio.NewScanner(stdOut)
	scanner.Split(bufio.ScanBytes)

	go killAfterDuration(cmd, time.Duration(*timeToWait)*time.Millisecond)

	matcher := util.ByteMatcher(expectedOutput)

	err = cmd.Start()
	if err != nil {
		return 0, fmt.Errorf("error starting command: %w", err)
	}

	startTime := time.Now()
	for scanner.Scan() {
		err := matcher.Match(scanner.Bytes())
		if err != nil {
			return 0, fmt.Errorf("error matching bytes: %w", err)
		}

		if matcher.Complete() {
			break
		}
	}

	if !matcher.Complete() {
		return 0, errors.New("command exited before all bytes were matched")
	}

	return time.Now().Sub(startTime), nil
}

func killAfterDuration(cmd *exec.Cmd, duration time.Duration) {
	time.Sleep(duration)

	err := cmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		log.Fatal(err)
	}
}
