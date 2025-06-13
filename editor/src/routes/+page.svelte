<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';
	import * as Card from '$lib/components/ui/card';
	import { codeState } from './state.svelte';
	import {
		MonacoLanguageConfiguration,
		MonarchConfig,
		MONKEY_LANGUAGE
	} from '$lib/components/monaco/language-config';

	let editor: Monaco.editor.IStandaloneCodeEditor;
	let monaco: typeof Monaco;
	let editorContainer: HTMLElement;

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
		
		monaco.languages.register({
			id: MONKEY_LANGUAGE
		});

		monaco.languages.setMonarchTokensProvider(MONKEY_LANGUAGE, MonarchConfig);
		monaco.languages.setLanguageConfiguration(MONKEY_LANGUAGE, MonacoLanguageConfiguration);
		monaco.editor.defineTheme('monkey-theme', {
			base: 'vs-dark',
			inherit: true,
			rules: [],
			colors: {}
		});

		monaco.editor.setTheme('monkey-theme');

		const model = monaco.editor.createModel(codeState.inputCode, MONKEY_LANGUAGE);
		currentModel = model;
		editor.setModel(model);
	});

	onDestroy(() => {
		monaco?.editor.getModels().forEach((model) => model.dispose());
		editor?.dispose();
	});

	$effect(() => {
		currentModel?.onDidChangeContent((_) => {
			codeState.inputCode = currentModel?.getValue() || '';
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
				<p class="text-sm">{codeState.result}</p>
			</Card.Content>
		</Card.Root>
	</div>
</main>
