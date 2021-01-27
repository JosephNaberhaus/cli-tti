package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var timeToWait = flag.Int("wait-time", 3000, "the number of milliseconds to wait until the program is assumed to be interactive")
var numTimes = flag.Int64("n", 5, "the number of runs to time a command before computing the average")
var printExpected = flag.Bool("print-expected", false, "whether to print what the time will consider the interactive state of the program")

func main() {
	flag.Parse()

	cmd, err := parseCommand(strings.Join(flag.Args(), " "))
	if err != nil {
		fmt.Println("Invalid command")
		log.Fatal(err)
	}

	fmt.Println("Finding expected output...")
	fmt.Printf("Program is assumed to be interactive after %d milliseconds\n", *timeToWait)
	fmt.Println()

	expectedOutput, err := findExpected(cmd)
	if err != nil {
		log.Fatal(err)
	}

	if *printExpected {
		fmt.Println("Expected output:")
		fmt.Println(string(expectedOutput))
		fmt.Println()
	}

	fmt.Println("Timing command...")
	duration, err := timeCommandAverage(cmd, expectedOutput, *numTimes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time to interactive: %d milliseconds\n", duration.Milliseconds())
}
