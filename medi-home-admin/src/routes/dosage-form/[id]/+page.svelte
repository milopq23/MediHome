<script>
	import { DetailDosage, UpdateDosage } from '$lib/api/dosageform.js';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import HeaderDetail from '$lib/components/HeaderDetail.svelte';

	let dosage = {};

	let id = null;
	$: dosage_id = parseInt($page.params.id, 10);

	onMount(() => {
		loadDetailDosage(dosage_id);
	});

	async function loadDetailDosage(dosage_id) {
		const data = await DetailDosage(dosage_id);
		dosage = data;
	}

	async function updateDosage() {
		const data = { ...dosage };
		await UpdateDosage(dosage_id, data);
	}
</script>

<div class="flex items-center justify-center py-10">
	<div class="w-full max-w-3xl rounded-2xl bg-white p-8 shadow-xl dark:bg-gray-800">
		<HeaderDetail title="Trang chi tiết shipping" route="shipping" />

		<form class="space-y-6" on:submit|preventDefault={updateDosage}>
			<div>
				<label for="name" class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200">
					Tên shipping:
				</label>
				<input
					type="text"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={dosage.name}
				/>
			</div>

			<div>
				<label
					for="price"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Giá:
				</label>
				<textarea
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={dosage.description}
					rows="4"
				></textarea>
			</div>

			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
