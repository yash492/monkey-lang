<script>
	import { Heart } from '@lucide/svelte';
</script>

<footer class="mx-3 my-2 hidden md:block">
	<div class="flex justify-between">
		<div class="flex items-center">
			<p class="font-fira mr-3 flex items-center text-xs font-light">
				Thank you <span class="ml-2 font-bold hover:text-blue-400">
					<a href="https://interpreterbook.com" target="_blank">Thorsten Ball</a></span
				>
				<span class="mx-1 inline-block">
					<Heart color="#ed0a06" fill="#ed0a06" size="20" />
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
				using
				<img
					class="mr-2 ml-1 inline-block"
					src="/icons/go_icon.svg"
					alt="Go Icon"
					height="35"
					width="35"
				/>
				&
				<img
					class="mx-2 inline-block"
					src="/icons/svelte_icon.svg"
					alt="Svelte Icon"
					height="15"
					width="15"
				/>
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
				using
				<img
					class="mr-1 inline-block"
					src="/icons/go_icon.svg"
					alt="Go Icon"
					height="20"
					width="20"
				/>
				&
				<img
					class="mx-1 inline-block"
					src="/icons/svelte_icon.svg"
					alt="Svelte Icon"
					height="10"
					width="10"
				/>
			</p>
		</div>
	</div>
</footer>
