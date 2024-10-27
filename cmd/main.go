package main

import (
	"fmt"
	"os"

	"shabashov.com/wc/internal/counter"
)

const DEBUG = true

func getPath() string {
	return os.Args[1]
}

func RecoverPanic(msg string) {
	if r := recover(); r != nil {
		if DEBUG {
			fmt.Println(msg, "-", r)
		} else {
			fmt.Println(msg)
		}
	}
}

func main() {
	defer RecoverPanic("Empty path")
	path := getPath()
	defer RecoverPanic("Something went wrong:")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	chars, lines, words := counter.NewTextFileCounter(file).Count()
	fmt.Println(
		"Lines:", lines,
		"Chars:", chars,
		"Words:", words,
	)
}
