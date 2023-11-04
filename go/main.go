package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
	} else if len(args) == 2 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}
	defer f.Close()
	src, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Failed to read file: %v", err)
	}
	return run(string(src))
}

func runPrompt() error {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		sc.Scan()
		input := sc.Text()
		if len(input) == 0 {
			return nil
		}
		err := run(input)
		if err != nil {
			return fmt.Errorf("Got error on input %s: %v", input, err)
		}
	}
}

func run(src string) error {
	s := NewScanner(src)
	s.scanTokens()
	if s.err != nil {
		return s.err
	}
	for _, token := range s.tokens {
		fmt.Println(token.lexeme)
	}
	return nil
}
