<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
	import * as Card from '$lib/components/ui/card';

	let editor: Monaco.editor.IStandaloneCodeEditor;
	let monaco: typeof Monaco;
	let editorContainer: HTMLElement;

	let value = $state('');
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
		const model = monaco.editor.createModel('blah blah');
		currentModel = model;
		editor.setModel(model);
	});

	onDestroy(() => {
		monaco?.editor.getModels().forEach((model) => model.dispose());
		editor?.dispose();
	});

	$effect(() => {
		currentModel?.onDidChangeContent((_) => {
			value = currentModel?.getValue() || '';
		});
	});
</script>

<main class="mx-3 flex h-full flex-col">
	<div class="flex grow flex-col">
		<div>
			<div class="h-[55vh] w-full" bind:this={editorContainer}></div>
		</div>
		<Card.Root class="h-full grow rounded-none border-0">
			<Card.Content class="font-fira">
				<p class="font-fira mb-2 text-xs font-light">OUTPUT</p>
				<p class="text-sm">Press "Run" to see the result.</p>
			</Card.Content>
		</Card.Root>
	</div>
</main>
