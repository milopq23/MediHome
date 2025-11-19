<script>
	import { goto } from '$app/navigation';
	import { CreateDosage } from '$lib/api/dosageform.js';
	import HeaderDetail from '$lib/components/HeaderDetail.svelte';

	let dosage = {
		name: '',
		description: ''
	};

	async function createdosage() {
		try {
			const data = { ...dosage };
			await CreateDosage(data);
			goto('/dosage-form');
		} catch (err) {
			console.error('Error:', err);
		}
	}
</script>

<div class="flex items-center justify-center py-10">
	<div class="w-full max-w-3xl rounded-2xl bg-white p-8 shadow-xl dark:bg-gray-800">
		<HeaderDetail title="Tạo dosage" route="dosage" />

		<form class="space-y-6" on:submit|preventDefault={createdosage}>
			<div>
				<label for="name" class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200">
					Tên dosage:
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
					Mô tả:
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
