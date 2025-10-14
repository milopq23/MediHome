<script>
	import { createEventDispatcher } from 'svelte';
	import { ChevronLeft, ChevronRight } from 'lucide-svelte';

	export let page = 1;
	export let totalPages = 1;

	const dispatch = createEventDispatcher();

	function goTo(p) {
		if (p >= 1 && p <= totalPages && p !== page) {
			dispatch('goToPage', p);
		}
	}

	// Mảng các trang hiển thị
	$: pages = [];

	$: {
		if (totalPages <= 9) {
			pages = Array.from({ length: totalPages }, (_, i) => i + 1);
		} else {
			let start, end;

			if (page <= 5) {
				start = 2;
				end = 7;
			} else if (page >= totalPages - 4) {
				start = totalPages - 6;
				end = totalPages - 1;
			} else {
				start = page - 2;
				end = page + 2;
			}

			pages = [1];

			if (start > 2) pages.push('...');
			for (let i = start; i <= end; i++) {
				pages.push(i);
			}
			if (end < totalPages - 1) pages.push('...');
			pages.push(totalPages);
		}
	}
</script>

<div class="mt-6 flex justify-center">
	<ul class="flex gap-1 text-gray-900 select-none items-center">
		<!-- Nút Previous -->
		<li>
			<button
				on:click={() => goTo(page - 1)}
				class="grid h-8 w-8 place-content-center rounded border border-gray-200 transition hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
				disabled={page === 1}
			>
				<ChevronLeft size="20" />
			</button>
		</li>

		<!-- Các trang -->
		{#each pages as p}
			{#if p === '...'}
				<li>
					<span class="h-8 w-8 grid place-content-center text-gray-400">…</span>
				</li>
			{:else}
				<li>
					<button
						on:click={() => goTo(p)}
						class="h-8 w-8 rounded border border-gray-200 text-center transition hover:bg-gray-100
							{page === p ? ' bg-blue-500 text-white font-semibold' : ''}"
						disabled={page === p}
					>
						{p}
					</button>
				</li>
			{/if}
		{/each}

		<!-- Nút Next -->
		<li>
			<button
				on:click={() => goTo(page + 1)}
				class="grid h-8 w-8 place-content-center rounded border border-gray-200 transition hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
				disabled={page === totalPages}
			>
				<ChevronRight size="20" />
			</button>
		</li>
	</ul>
</div>
