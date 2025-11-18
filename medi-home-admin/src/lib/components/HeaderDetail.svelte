<script>
	import { goto } from '$app/navigation';
	import { ArrowLeft } from 'lucide-svelte';

	export let title = '';
	export let route = '';

	function goBack() {
		goto(`/${route}`);
	}
</script>

<div>
	<button
		class="text-grey rounded-lg bg-white px-5 py-2.5 text-center text-sm font-medium hover:bg-gray-100 dark:hover:bg-gray-600"
		on:click={() => goBack()}
	>
		<ArrowLeft class="h-5 w-10" />Quay về
	</button>
	<h1 class="mb-10 text-center text-4xl font-extrabold text-gray-900 dark:text-white">
		{title}
	</h1>
</div>
