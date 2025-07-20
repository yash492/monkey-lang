<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { Wasm } from '$lib/wasm';
	import { Play, BadgeInfo, Copy } from '@lucide/svelte';
	import { codeState } from '../../../../routes/state.svelte';
	import { toast } from 'svelte-sonner';
	import * as Dialog from '$lib/components/ui/dialog';
	import Highlight, { LineNumbers } from 'svelte-highlight';
	import json from 'svelte-highlight/languages/json';
	import base16IrBlack from 'svelte-highlight/styles/base16-ir-black';

	let openDialog = $state(false);
	let astContent = $state('');
	let astError = $state(false);

	const wasm = new Wasm();
	function onInterpret() {
		if (codeState.inputCode.trim() === '') {
			codeState.result = "Press the 'Run' button to see the result.";
			return;
		}

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

	function onOpenASTExplorer() {
		openDialog = true;
		const result = wasm.getAST(codeState.inputCode);
		if (result.is_error) {
			const errMsg = `Error while showing AST, ${result.result}`;
			astError = true;
			astContent = errMsg;
			return;
		}

		astContent = result.result;
	}

	async function onCopy(text: string) {
		try {
			await navigator.clipboard.writeText(text);
			toast.success('Success', {
				position: 'top-right',
				description: 'Copied the AST successfully'
			});
		} catch (e) {
			toast.error('Success', {
				position: 'top-right',
				description: `Error while copying the AST ${JSON.stringify(e)}`
			});
		}
	}
</script>

<svelte:head>
	{@html base16IrBlack}
</svelte:head>

<nav class="mx-3 flex h-16 items-center justify-between border-b">
	<div class="flex items-center">
		<h1 class="text-md mr-2 hidden text-xl font-bold md:block">The Monkey Playground</h1>
		<h1 class="text-md mr-2 text-xl font-bold md:hidden">Monkey</h1>
		<img src="/icons/monkey.svg" alt="monkey-logo" height="30" width="30" />
	</div>
	<div class="flex gap-3">
		<Button variant="outline" class="cursor-pointer md:font-bold" onclick={onInterpret}>
			<Play />
			Run
		</Button>
		<Button
			variant="outline"
			class="hidden cursor-pointer md:flex md:font-bold"
			onclick={onOpenASTExplorer}
		>
			<BadgeInfo />
			View AST
		</Button>
	</div>
</nav>

<Dialog.Root open={openDialog} onOpenChange={(isOpen) => (openDialog = isOpen)}>
	<Dialog.Content class="hidden w-[600px] bg-black sm:max-h-[600px] md:flex">
		<Dialog.Header>
			<Dialog.Title>
				<div class="flex items-center gap-6">
					<p>AST Explorer</p>
					<Button variant="outline" class="w-fit cursor-pointer" onclick={() => onCopy(astContent)}>
						<Copy />
						Copy</Button
					>
				</div>
			</Dialog.Title>
			<Dialog.Description
				class="max mt-2 w-[450px] overflow-x-auto overflow-y-auto rounded-sm border-2 py-2 text-white"
			>
				<div class="h-[410px] w-full p-2 font-mono">
					{#if astError}
						{astContent}
					{:else}
						<Highlight
							language={json}
							code={JSON.stringify(JSON.parse(astContent), null, 2)}
							let:highlighted
						>
							<LineNumbers {highlighted} />
						</Highlight>
					{/if}
				</div>
			</Dialog.Description>
		</Dialog.Header>

		<Dialog.Footer></Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
