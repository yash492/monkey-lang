//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"strings"
	"syscall/js"
)

const defaultOuput = "No Result. Code executed successfully."

func main() {
	ch := make(chan bool)
	js.Global().Set("interpret", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return js.ValueOf("err: wrong data")
		}
		result, isError := run(args[0].String())

		response := map[string]any{
			"result":   result,
			"is_error": isError,
		}
		return js.ValueOf(response)
	}))

	js.Global().Set("getAST", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return js.ValueOf("err: wrong data")
		}
		result, isError := getAST(args[0].String())

		response := map[string]any{
			"result":   result,
			"is_error": isError,
		}
		return js.ValueOf(response)
	}))

	<-ch

}

// run returns result and whether error occurred after
// lexing -> parsing -> evaluation
func run(code string) (string, bool) {
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	if len(p.Errors()) != 0 {
		return printParserErrors(p.Errors()), true
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated == nil {
		return defaultOuput, false
	}

	if evaluated.Type() == object.ErrorObj {
		return evaluated.Inspect(), true
	}

	return evaluated.Inspect(), false
}

func getAST(code string) (string, bool) {
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return printParserErrors(p.Errors()), true
	}

	bytes, err := json.Marshal(program)
	if err != nil {
		return err.Error(), true
	}

	return string(bytes), false

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
