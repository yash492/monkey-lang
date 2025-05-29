<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/components/custom/navbar/Navbar.svelte';
	import Footer from '$lib/components/custom/footer/footer.svelte';

	let { children } = $props();
</script>

<div class="flex h-screen flex-col">
	<Navbar />
	{@render children()}
	<Footer />
</div>
