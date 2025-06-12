// This file contains all the functions that can be involked from WASM

import type { InterpreterResult } from "./types"

export class Wasm {
    private _global = globalThis

    interpret(code: string): InterpreterResult {
        return this._global.interpret(code)
    }
}