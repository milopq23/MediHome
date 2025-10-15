<script>
    import {Plus} from 'lucide-svelte';
    export let titleName = ''
</script>

<div>
    <div class="flex flex-col justify-between sm:flex-row sm:items-center">
		<!-- Title -->
		<h1 class="text-2xl font-bold sm:text-4xl">{titleName}</h1>

		<!-- Button -->
		<div class="flex justify-center p-3">
			<button
				class="btn-create flex items-center rounded gap-2 bg-blue-600 transition hover:bg-blue-700"
				on:click={openForm('create')}
			>
				<Plus class="h-5 w-5" />
				<span class="text-lg sm:text-xl">Thêm</span>
			</button>
		</div>
	</div>
</div>