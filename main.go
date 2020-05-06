package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	usageMsg = `sleep sleeps proccess.

Usage:

	sleep <sec>

Aurguments are:

	sec	[0-36400] sleeping second time
`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usageMsg)
		os.Exit(1)
	}
	sec, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if sec < 0 || sec > 36400 {
		fmt.Fprintln(os.Stderr, "sec range over [0-36400]")
		os.Exit(1)
	}
	time.Sleep(time.Duration(sec) * time.Second)
}
