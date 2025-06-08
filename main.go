//go:build js && wasm
// +build js,wasm

package main

import (
	"monkey/lexer"
	"monkey/parser"
	"strings"
	"syscall/js"
)

func main() {

	ch := make(chan bool)

	js.Global().Set("interpret", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return js.ValueOf("err: wrong data")
		}
		result := run(args[0].String())
		return js.ValueOf(result)
	}))

	<-ch

}

func run(code string) string {

	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return printParserErrors(p.Errors())
	}

	return program.String()
}

func printParserErrors(errors []string) string {

	out := strings.Builder{}
	out.WriteString("Woops! We ran into some monkey business here!\n")
	out.WriteString(" parser errors:\n")
	for _, msg := range errors {
		out.WriteString("\t" + msg + "\n")
	}
	return out.String()
}
