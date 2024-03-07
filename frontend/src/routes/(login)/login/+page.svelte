<div class="rounded-md bg-zinc-800 shadow-xl shadow-black">
	<h1 class="m-2 text-xl font-bold">Dashboard Link Login</h1>
	<div class="mt-2 flex gap-2 border border-zinc-700 p-2">
		<div class="flex flex-col justify-center gap-2">
			<label class="text-right" for="email"> Email: </label>
			<label class="text-right" for="password"> Password: </label>
		</div>
		<div class="flex w-full flex-col gap-2">
			<input class="h-6 w-full text-zinc-950 px-2" type="email" id="email" placeholder="Email" />
			<input class="h-6 w-full text-zinc-950 px-2" type="password" id="password" placeholder="password" />
		</div>
	</div>
	<div class="m-2 flex justify-end gap-2">
        <label for="remember">
            Remember-me <input type="checkbox" id="remember">
        </label>
		<button
			class="rounded border border-zinc-500 bg-zinc-500 px-2 hover:bg-zinc-400 focus:border-zinc-300 active:border-zinc-300"
			>Login</button
		>
	</div>
</div>
