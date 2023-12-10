package tool

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: ./generate_ast <output directory>")
		return
	}
	outputDir := args[1]
	defineAst(outputDir, "Expr", []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : Object value",
		"Unary    : Token operator, Expr right",
	})
}

func defineAst(outputDir, baseName string, types []string) error {
	path := outputDir + "/" + baseName + ".go"
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create file: %v", err)
	}
	defer f.Close()

	f.Write()
}
