package script

import (
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

func RunScript(filename string) {
	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	script := string(b) // convert content to a 'string'
	out := os.Stdout

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
	
	l := lexer.New(script)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluator.Eval(expanded, env)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
