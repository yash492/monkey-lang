# Monkey Language Interpreter

This repository implements a Monkey programming language interpreter in Go. It includes components for lexical analysis, parsing, evaluation, and a REPL for interactive use.

## Features

- **Lexer**: Tokenizes the input source code.
- **Parser**: Builds an Abstract Syntax Tree (AST).
- **Evaluator**: Executes the AST.
- **REPL**: Interactive shell for testing Monkey code.

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
