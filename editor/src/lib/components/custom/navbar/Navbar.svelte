<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import { Github, Moon, Play, Sun } from '@lucide/svelte';
</script>

<nav class="mx-3 flex h-16 items-center justify-between border-b">
	<div class="flex items-center">
		<h1 class="text-md mr-2 hidden text-xl font-bold md:block">The Monkey Playground</h1>
		<h1 class="text-md mr-2 text-xl font-bold md:hidden">Monkey</h1>
		<img src="/icons/monkey.svg" alt="monkey-logo" height="30" width="30" />
	</div>
	<div>
		<Button variant="outline" class="cursor-pointer md:font-bold">
			<Play />
			Run
		</Button>
		<Button variant="outline" class="cursor-pointer md:font-bold">
			<Github />
		</Button>
	</div>
</nav>
