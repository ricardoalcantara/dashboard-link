<script lang="ts">
	import Icon from '@iconify/svelte';

	const inventary = [
		{
			id: 1,
			name: 'Load balancer',
			ip: '10.0.0.1',
			status: 'online',
			selected: false
		},
		{
			id: 2,
			name: 'MySQL',
			ip: '10.0.0.2',
			status: 'offline',
			selected: true
		},
		{
			id: 3,
			name: 'Proxmox 1',
			ip: '10.0.0.3',
			status: 'online',
			selected: false
		},
		{
			id: 4,
			name: 'Mailcow',
			ip: '10.0.0.4',
			status: 'online',
			selected: false
		},
		{
			id: 5,
			name: 'Docker',
			ip: '10.0.0.5',
			status: 'online',
			selected: false
		}
	];

	const playbooks = [
		{
			id: 1,
			name: 'Ping',
			description: 'Does ping',
			selected: false
		},
		{
			id: 1,
			name: 'Update',
			description: 'Update all packages',
			selected: false
		},
		{
			id: 1,
			name: 'Basic Setup',
			description: 'Create User, Set Password, Set SSH Key',
			selected: false
		}
	];

	const logs: { timestamp: string; message: string }[] = [];
	for (let i = 0; i < 100; i++) {
		let log = {
			timestamp: new Date().toISOString(),
			message: `Log message ${i + 1}`
		};
		logs.push(log);
	}
</script>

<div class="flex h-dvh flex-col bg-zinc-900 text-sm text-zinc-100">
	<div class="mx-2 flex h-8 items-center justify-between">
		<div class="flex">
			<button class="rounded flex items-center px-1 sm:hidden active:bg-zinc-300">
				<Icon width="24" icon="mdi:menu"></Icon>
			</button>
			<h1 class="text-xl font-bold">Dashboard Link</h1>
		</div>
		<button class="rounded px-2 bg-zinc-500 active:bg-zinc-300 sm:hidden block">...</button>
		<div class="hidden gap-1 sm:flex">
			<button
				class="rounded border border-zinc-500 bg-zinc-500 px-1 hover:bg-zinc-400 focus:border-zinc-300 active:border-zinc-300"
				>Add Kubernetes</button
			>
			<button
				class="rounded border border-zinc-500 bg-zinc-500 px-1 hover:bg-zinc-400 focus:border-zinc-300 active:border-zinc-300"
				>Add Docker</button
			>
			<button
				class="rounded border border-zinc-500 bg-zinc-500 px-1 hover:bg-zinc-400 focus:border-zinc-300 active:border-zinc-300"
				>Add Ansible</button
			>
			<button
				class="rounded border border-zinc-500 bg-zinc-500 px-1 hover:bg-zinc-400 focus:border-zinc-300 active:border-zinc-300"
				>Add Server</button
			>
		</div>
	</div>
	<div class="flex flex-1 gap-1 px-1">
		<div class="hidden bg-zinc-800 sm:block">
			<section>
				<h2 class="my-2 pl-2">Servers</h2>
				<ul>
					{#each inventary as i}
						<li
							class="flex items-center gap-2 px-4 py-1 hover:bg-zinc-600 {i.selected
								? 'bg-zinc-600'
								: ''}"
						>
							<span
								class="h-2 w-2 rounded-full {i.status === 'online' ? 'bg-green-500' : 'bg-red-500'}"
							/>
							<div>
								<div>{i.name}</div>
								<!-- <div class="text-xs">{i.ip}</div> -->
							</div>
						</li>
					{/each}
				</ul>
			</section>
			<section>
				<h2 class="my-2 pl-2">Docker</h2>
			</section>
			<section>
				<h2 class="my-2 pl-2">Kubernetes</h2>
			</section>
			<section>
				<h2 class="my-2 pl-2">Ansible</h2>
				<ul>
					{#each playbooks as p}
						<li
							class="flex items-center gap-2 px-4 py-1 hover:bg-zinc-600 {p.selected
								? 'bg-zinc-600'
								: ''}"
						>
							<div>
								<div>{p.name}</div>
							</div>
						</li>
					{/each}
				</ul>
			</section>
		</div>
		<div class="flex-1 bg-zinc-800"><slot /></div>
		<!-- <div class="bg-zinc-800">Right</div> -->
	</div>
	<div class="m-1 h-40 overflow-auto bg-zinc-800">
		<table class="w-full table-auto">
			<thead>
				<tr class="h-6">
					<th class="border-b border-r border-zinc-700">Timestamp</th>
					<th class="border-b border-zinc-700">Message</th>
				</tr>
			</thead>
			<tbody>
				{#each logs as log}
					<tr class="h-6 even:bg-zinc-900">
						<td class="border-b border-r border-zinc-700 pl-1">{log.timestamp}</td>
						<td class="border-b border-zinc-700 pl-1">{log.message}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

<style lang="postcss">
	:global(html) {
		/* background-color: theme(colors.gray.800); */
	}
</style>
