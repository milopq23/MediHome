<script>
	import { DetailShipping, UpdateShipping } from '$lib/api/shipping.js';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import HeaderDetail from '$lib/components/HeaderDetail.svelte';

	let shipping = {};

	let id = null;
	$: shipping_id = parseInt($page.params.id, 10);

	onMount(() => {
		loadDetailShipping(shipping_id);
	});

	async function loadDetailShipping(shipping_id) {
		const data = await DetailShipping(shipping_id);
		shipping = data;
	}

	async function updateShipping() {
		const data = { ...shipping };
		await UpdateShipping(shipping_id, data);
	}
</script>

<div class="flex items-center justify-center py-10">
	<div class="w-full max-w-3xl rounded-2xl bg-white p-8 shadow-xl dark:bg-gray-800">
		<HeaderDetail title="Trang chi tiết shipping" route="shipping" />

		<form class="space-y-6" on:submit|preventDefault={updateShipping}>
			<div>
				<label for="name" class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200">
					Tên shipping:
				</label>
				<input
					type="text"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={shipping.name}
				/>
			</div>

			<div>
				<label
					for="price"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Giá:
				</label>
				<input
					type="number"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={shipping.price}
				/>
			</div>

			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
