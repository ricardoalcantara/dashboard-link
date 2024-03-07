<script lang="ts">
	import Icon from '@iconify/svelte';

	const menu = [
		{
			name: 'Summary',
			selected: true
		},
		{
			name: 'Terminal',
			selected: false
		}
	];
</script>

<div class="m-2 flex h-full flex-col pb-4">
	<div class="h-8 border-b border-zinc-700">
		<h2>Server 'MySQL'</h2>
	</div>
	<div class="flex flex-1">
		<div class="border-r border-zinc-700">
			<ul class="">
				{#each menu as m}
					<li
						class="flex items-center gap-2 px-4 py-1 hover:bg-zinc-600 {m.selected
							? 'bg-zinc-600'
							: ''}"
					>
						<div class="flex items-center gap-1">
							<Icon icon="mdi:home"></Icon>
							<div class="hidden sm:block">{m.name}</div>
						</div>
					</li>
				{/each}
			</ul>
		</div>
		<div>
			<div class="flex h-8 items-center border-b border-zinc-700">
				<button
					class="ml-2 rounded border border-slate-500 bg-slate-500 px-1 hover:bg-slate-400 focus:border-slate-300 active:border-slate-300"
					>Package verions</button
				>
			</div>
			<div class="m-2">
				<h3 class="text-lg">Summary</h3>
				<div class="mt-2">
					Lorem, ipsum dolor sit amet consectetur adipisicing elit. Vero, delectus, perspiciatis
					consequatur laborum dignissimos consectetur praesentium fugit corrupti amet voluptas
					officia, dolore quibusdam rem voluptates et facilis animi commodi! Fugit.
				</div>
			</div>
		</div>
	</div>
</div>
