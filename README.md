# Monkey Language Interpreter

This repository implements a Monkey programming language interpreter in Go. It includes components for lexical analysis, parsing, evaluation, and a REPL for interactive use.

## Features

- **Lexer**: Tokenizes the input source code.
- **Parser**: Builds an Abstract Syntax Tree (AST).
- **Evaluator**: Executes the AST.
- **REPL**: Interactive shell for testing Monkey code.

## WebAssembly Integration

The project leverages WebAssembly (Wasm) to enable high-performance execution of the Monkey language in web environments. This integration allows the interpreter to run directly in browsers, providing a seamless experience for users without requiring native installations.

### Purpose
WebAssembly enhances the portability and performance of the Monkey language interpreter, making it accessible to a broader audience.

### Functionality
- **Compilation**: The Go-based interpreter is compiled to WebAssembly using `wasm_exec.js`.
- **Frontend Integration**: The `editor/src/lib/wasm/index.ts` file handles communication between the WebAssembly module and the frontend editor.
- **Execution**: Monkey code can be executed in the browser with near-native performance.

### Relevant Files
- [`editor/src/lib/wasm/index.ts`](editor/src/lib/wasm/index.ts)
- [`editor/static/wasm/wasm_exec.js`](editor/static/wasm/wasm_exec.js)

## Examples

### Let Statement

```monkey
let x = 10;
let y = x + 5;
```

### If Conditions

```monkey
if (x > 5) {
  let result = 40;
} else {
  let result = 50;
}
```

## Directory Structure

- **lexer/**: Handles tokenization.
- **parser/**: Builds AST.
- **evaluator/**: Evaluates AST.
- **object/**: Defines runtime objects.
- **repl/**: Interactive shell.
- **editor/**: Frontend editor for Monkey code.

## Acknowledgments

This project is based on the book [Writing An Interpreter In Go](https://interpreterbook.com/). Special thanks to [Thorsten Ball](https://thorstenball.com/) for providing an excellent resource for learning and implementing interpreters.

## License

This project is licensed under the MIT License.
