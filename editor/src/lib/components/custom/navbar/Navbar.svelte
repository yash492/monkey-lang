<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import { Wasm } from '$lib/wasm';
	import { Github, Play } from '@lucide/svelte';
	import { codeState } from '../../../../routes/state.svelte';
	import { toast } from 'svelte-sonner';

	const wasm = new Wasm();
	function onInterpret() {
		const result = wasm.interpret(codeState.inputCode);
		codeState.result = result.result;
		codeState.isError = result.is_error;

		if (!codeState.isError) {
			toast.success('Success', {
				position: 'bottom-center',
				description: 'Check the output for more info.'
			});
		} else {
			toast.error('Error', {
				position: 'bottom-center',
				description: 'Check the output for more info.'
			});
		}
	}
</script>

<nav class="mx-3 flex h-16 items-center justify-between border-b">
	<div class="flex items-center">
		<h1 class="text-md mr-2 hidden text-xl font-bold md:block">The Monkey Playground</h1>
		<h1 class="text-md mr-2 text-xl font-bold md:hidden">Monkey</h1>
		<img src="/icons/monkey.svg" alt="monkey-logo" height="30" width="30" />
	</div>
	<div>
		<Button variant="outline" class="cursor-pointer md:font-bold" onclick={onInterpret}>
			<Play />
			Run
		</Button>
	</div>
</nav>
