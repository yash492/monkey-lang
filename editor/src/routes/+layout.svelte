<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/components/custom/navbar/Navbar.svelte';
	import Footer from '$lib/components/custom/footer/footer.svelte';
	import { Toaster } from '$lib/components/ui/sonner';

	let { children } = $props();
</script>

<Toaster richColors />
<div class="flex h-screen flex-col">
	<Navbar />
	{@render children()}
	<Footer />
</div>
