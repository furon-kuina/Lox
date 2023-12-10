package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	expr := &BinaryExpr{
		left: &UnaryExpr{
			right: &LiteralExpr{value: NewToken(Number, "123", 0, 0, 0)},
			op:    NewToken(Minus, "-", 0, 0, 0),
		},
		right: &GroupingExpr{&LiteralExpr{value: NewToken(Number, "12.34", 0, 0, 0)}},
		op:    NewToken(Star, "*", 0, 0, 0),
	}

	ue := &UnaryExpr{
		right: &LiteralExpr{value: NewToken(Number, "123", 0, 0, 0)},
		op:    NewToken(Minus, "-", 0, 0, 0),
	}

	ap := &AstPrinter{}
	fmt.Println(ap.print(expr))
	fmt.Println(ap.print(ue))

	// args := os.Args
	// if len(args) > 2 {
	// fmt.Println("Usage: glox [script]")
	// } else if len(args) == 2 {
	// runFile(args[1])
	// } else {
	// runPrompt()
	// }
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
