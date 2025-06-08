// This file contains all the functions that can be involked from WASM

export class Wasm {
    private _global = globalThis

    interpret(code: string): string {
        return this._global.interpret(code)
    }
}