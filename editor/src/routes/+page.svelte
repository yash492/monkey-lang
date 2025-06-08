<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
	import * as Card from '$lib/components/ui/card';

	import { Wasm } from '$lib/wasm';
	import { Button } from '$lib/components/ui/button';
	import { Play } from '@lucide/svelte';

	let editor: Monaco.editor.IStandaloneCodeEditor;
	let monaco: typeof Monaco;
	let editorContainer: HTMLElement;

	const wasm = new Wasm();

	let inputCode = $state('blah blah');
	let output = $state("Press the 'run' button to see the result.");
	let currentModel = $state<Monaco.editor.ITextModel>();

	onMount(async () => {
		monaco = (await import('$lib/components/monaco/monaco')).default;
		editor = monaco.editor.create(editorContainer, {
			minimap: {
				enabled: false
			},
			automaticLayout: true,
			fontLigatures: true,
			fontSize: 14,
			fontFamily: 'Fira Code',
			scrollbar: {
				alwaysConsumeMouseWheel: false
			}
		});

		monaco.editor.setTheme('vs-dark');
		const model = monaco.editor.createModel(inputCode);
		currentModel = model;
		editor.setModel(model);
	});

	onDestroy(() => {
		monaco?.editor.getModels().forEach((model) => model.dispose());
		editor?.dispose();
	});

	$effect(() => {
		currentModel?.onDidChangeContent((_) => {
			inputCode = currentModel?.getValue() || '';
		});
	});
</script>

<main class="mx-3 flex flex-col">
	<div class="flex flex-col">
		<div>
			<div class="h-[55vh] w-full" bind:this={editorContainer}></div>
		</div>
		<Card.Root class="h-[32vh] rounded-none border-0">
			<Card.Content class="font-fira h-[calc(100%-0.5rem)]  overflow-y-scroll">
				<p class="font-fira mb-2 text-xs font-light">OUTPUT</p>
				<p class="text-sm">{output}</p>
			</Card.Content>
		</Card.Root>
	</div>
</main>
