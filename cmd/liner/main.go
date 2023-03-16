package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/isksss/go-liner"
)

var (
	length int
	moji   string
	cp     bool
)

// copy clipboard
func copy(line string) {
	if err := clipboard.WriteAll(line); err != nil {
		fmt.Println("Failed to copy to clipboard:", err)
	}
	fmt.Printf("Copy to clipboard:%s\n", line)
}

// flag register
func flagRegister() {
	flag.IntVar(&length, "l", 10, "line length") // length
	flag.StringVar(&moji, "m", "#", "char")      // moji
	flag.BoolVar(&cp, "no", false, "no copy")    // copy
}

func run() int {
	// flag
	flagRegister()
	flag.Parse()

	// main
	line := liner.MakeLine(moji, length)

	// copy
	if !cp {
		copy(line)
	}

	return 0
}

func main() {
	os.Exit(run())
}
