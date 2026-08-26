package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/heige/a2s-issue-demo/textnorm"
)

func main() {
	input, err := readInput(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textnorm: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(textnorm.Normalize(input))
}

func readInput(arguments []string, stdin io.Reader) (string, error) {
	if len(arguments) != 0 {
		return strings.Join(arguments, " "), nil
	}

	input, err := io.ReadAll(stdin)
	if err != nil {
		return "", fmt.Errorf("read stdin: %w", err)
	}
	return string(input), nil
}
