<script>
	import { Heart } from '@lucide/svelte';
</script>

<footer class="mx-3 my-2 hidden md:block">
	<div class="flex justify-between">
		<div class="flex items-center">
			<p class="font-fira mr-3 flex items-center text-xs font-light">
				Thank you <span class="ml-2 font-bold hover:text-blue-400">
					<a href="https://interpreterbook.com" target="_blank">Thorsten Ball</a> 🙌
				</span>
			</p>
			<!-- <Github size="20" /> -->
		</div>
		<div>
			<p class="font-fira flex items-center justify-center text-center text-xs">
				Made with
				<span class="mx-1 inline-block">
					<Heart color="#ed0a06" fill="#ed0a06" size="20" />
				</span>
				by
				<span class="ml-2 font-bold hover:text-blue-400">
					<a href="https://github.com/yash492/monkey-lang" target="_blank">Yash Jain</a></span
				>
			</p>
		</div>
	</div>
</footer>

<footer class="mx-3 my-2 md:hidden">
	<div class="flex justify-between">
		<div class="font-fira text-xs">
			<span class="ml-2 font-bold hover:text-blue-400">
				<a href="https://interpreterbook.com" target="_blank"
					>Thorsten Ball <span class="text-sm">🙌</span></a
				></span
			>
		</div>
		<div>
			<p class="font-fira flex items-center justify-center text-center text-xs">
				Made with
				<span class="mx-1 inline-block">
					<Heart color="#ed0a06" fill="#ed0a06" size="15" />
				</span>
				by
				<span class="ml-2 font-bold hover:text-blue-400">
					<a href="https://github.com/yash492/monkey-lang" target="_blank">Yash Jain</a></span
				>
			</p>
		</div>
	</div>
</footer>
