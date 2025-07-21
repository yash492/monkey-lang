<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { Wasm } from '$lib/wasm';
	import { Play, TreePineIcon, Copy, Network } from '@lucide/svelte';
	import { codeState } from '../../../../routes/state.svelte';
	import { toast } from 'svelte-sonner';
	import * as Dialog from '$lib/components/ui/dialog';
	import Highlight, { LineNumbers } from 'svelte-highlight';
	import json from 'svelte-highlight/languages/json';
	import { compress } from 'compress-json';
	import base16IrBlack from 'svelte-highlight/styles/base16-ir-black';

	const BASE_VISUALIZE_LINK = 'https://omute.net/editor';
	let openDialog = $state(false);
	let astContent = $state('');
	let astContentParsed = $state({});
	let astError = $state(false);
	let visualizeLink = $state(BASE_VISUALIZE_LINK);

	$effect(() => {
		const compressedJSON = compress(astContentParsed);
		const encodedJSON = encodeURIComponent(JSON.stringify(compressedJSON));
		visualizeLink = `${BASE_VISUALIZE_LINK}?json=${encodedJSON}`;
	});

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
		astContentParsed = JSON.parse(result.result);
	}

	async function onCopy(text: string) {
		try {
			await navigator.clipboard.writeText(text);
			toast.success('Copied!', {
				position: 'top-right'
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
			<TreePineIcon />
			View AST
		</Button>
	</div>
</nav>

<Dialog.Root open={openDialog} onOpenChange={(isOpen) => (openDialog = isOpen)}>
	<Dialog.Content class="hidden w-[600px] bg-black sm:max-h-[600px] md:flex">
		<Dialog.Header>
			<Dialog.Title>
				<div class="flex items-center gap-6">
					<div>
						<p>AST Explorer</p>
					</div>
					<div class="flex gap-2">
						<Button variant="outline" class="cursor-pointer">
							<Network />
							<a href={visualizeLink} target="_blank">Visualize</a>
						</Button>
						<Button
							variant="outline"
							class="w-fit cursor-pointer"
							onclick={() => onCopy(astContent)}
						>
							<Copy />
						</Button>
					</div>
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
							code={JSON.stringify(astContentParsed, null, 2)}
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
