<script>
	export let page = 1;
	export let pageSize = 10;
	export let total = 1;
	export let totalPages = 1;
	export let goToPage;
	// const dispatch = createEventDispatcher();

	// function goToPage(p) {
	// 	dispatch('pageChange', p);
	// }
</script>

<div class="mt-4 flex items-center justify-between text-sm text-gray-500">
	<div>
		Showing {(page - 1) * pageSize + 1} to {Math.min(page * pageSize, total)} of {total} products
	</div>
	<div class="flex items-center gap-1">
		<button
			class="rounded px-2 py-1 disabled:opacity-50"
			on:click={() => goToPage(page - 1)}
			disabled={page === 1}
		>
			«
		</button>
		{#each Array(totalPages) as _, i}
			{#if totalPages <= 5}
				<!-- Nếu tổng số trang nhỏ hơn hoặc bằng 5, hiển thị tất cả -->
				<button
					class={`rounded px-3 py-1 ${page === i + 1 ? 'bg-indigo-100 font-semibold text-indigo-600' : 'hover:bg-gray-100'}`}
					on:click={() => goToPage(i + 1)}
				>
					{i + 1}
				</button>
			{:else}
				<!-- Tổng số trang > 5 -->
				{#if i === 0 || i === totalPages - 1 || (i + 1 >= page - 1 && i + 1 <= page + 1)}
					<!-- Hiển thị: Trang đầu, trang cuối, và các trang gần current page -->
					<button
						class={`rounded px-3 py-1 ${page === i + 1 ? 'bg-indigo-100 font-semibold text-indigo-600' : 'hover:bg-gray-100'}`}
						on:click={() => goToPage(i + 1)}
					>
						{i + 1}
					</button>
				{:else if (i === 1 && page > 3) || (i === totalPages - 2 && page < totalPages - 2)}
					<!-- Hiển thị dấu ... sau trang đầu và trước trang cuối -->
					<span class="px-2 py-1 text-gray-400">...</span>
				{/if}
			{/if}
		{/each}
		<button
			class="rounded px-2 py-1 disabled:opacity-50"
			on:click={() => goToPage(page + 1)}
			disabled={page === totalPages}
		>
			»
		</button>
	</div>
</div>
