package liner

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strconv"
	"strings"
)

func Liner() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: liner <num>")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid argument:", args[0])
		return
	}

	hashes := strings.Repeat("#", num)
	fmt.Println(hashes)

	if err := clipboard.WriteAll(hashes); err != nil {
		fmt.Println("Failed to copy to clipboard:", err)
	}
}
