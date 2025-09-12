<script>
	let showSearch = false;
</script>

<div class="justify-center px-4 md:flex">
	<div
		class="flex w-1/2 items-center gap-2 rounded-lg bg-blue-100 px-4 py-2 focus-within:ring-2 focus-within:ring-blue-400 md:flex"
	>
		<i class="fa-solid fa-magnifying-glass text-gray-500"></i>
		<input
			type="text"
			placeholder="Tìm kiếm sản phẩm..."
			class="w-full rounded-lg border-gray-300 bg-transparent focus:outline-none"
		/>
	</div>

	<!-- Mobile search toggle button -->
	{#if !showSearch}
		<button
			class=" flex justify-end text-xl text-gray-700 focus:outline-none md:hidden"
			on:click={() => (showSearch = !showSearch)}
			aria-label="Toggle Search"
		>
			<i class="fa-solid fa-magnifying-glass"></i>
		</button>
	{/if}

	<!-- Mobile search input (show/hide) -->
	{#if showSearch}
		<div
			class="flex w-1/2 items-center gap-2 rounded-lg bg-blue-100 px-4 py-2 focus-within:ring-2 focus-within:ring-blue-400 md:flex"
		>
			<i class="fa-solid fa-magnifying-glass text-gray-500"></i>
			<input
				type="text"
				placeholder="Tìm kiếm sản phẩm..."
				class="w-full rounded-lg border-gray-300 bg-transparent focus:outline-none"
			/>
		</div>
	{/if}
</div>
