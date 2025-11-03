<script>
	import { toasts } from '$lib/stores/toastMessage.js';
	const typeColors = {
		success: 'bg-green-500',
		error: 'bg-red-500'
	};
</script>


<div class={`absolute top-0 right-0 w-70 p-3 `}>
	{#each $toasts as toast}
		<div
			class={`max-w-sm rounded-lg border border-gray-200 p-4 shadow-sm 
				dark:border-gray-700 dark:bg-gray-800 ${typeColors[toast.type] || ''}
				`}
		>
			<p class="text-lg font-bold text-gray-700 dark:text-gray-400">
				{toast.message}
			</p>
		</div>
	{/each}
</div>
