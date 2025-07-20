// See https://svelte.dev/docs/kit/types#app.d.ts

import type { InterpreterResult } from "$lib/wasm/types";

// for information about these interfaces
declare global {
	function interpret(code: string): InterpreterResult;
	function getAST(code: string): InterpreterResult;
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}

	}
}

export { };
