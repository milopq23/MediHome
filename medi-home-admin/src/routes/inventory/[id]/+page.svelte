<script>
	import { page } from '$app/stores';
	import { FindMedicine, UpdateBatchSelling } from '$lib/api/inventory.js';
	import { onMount } from 'svelte';

	let inventories = [];
	let selectedBatch = null;
	let loading = false;

	$: medicine_id = parseInt($page.params.id, 10);

	onMount(() => {
		getInventoryMedicine(medicine_id);
	});

	function selectBatch(id) {
		selectedBatch = id;
	}

	async function saveSelectedBatch() {
		if (!selectedBatch) {
			alert('Vui lòng chọn batch trước khi lưu!');
			return;
		}
		loading = true;
		try {
			await UpdateBatchSelling(medicine_id, selectedBatch);

			alert('✔️ Lưu batch bán thành công!');
		} catch (error) {
			console.error(error);
			alert('❌ Lỗi khi lưu batch!');
		} finally {
			loading = false;
		}
	}

	function daysLeft(expirationdate) {
		const today = new Date();
		const exp = new Date(expirationdate);
		return Math.ceil((exp - today) / (1000 * 60 * 60 * 24));
	}

	async function getInventoryMedicine(medicine_id) {
		const data = await FindMedicine(medicine_id);
		inventories = data.inventories;
		selectedBatch = data.inventory_id;
	}
</script>

<div class="px-4 py-6">
	<div class="rounded-lg bg-white p-6 shadow">
		{#if inventories.length > 0}
			<table class="w-full rounded-lg border border-gray-200">
				<thead class="bg-gray-100 text-gray-700">
					<tr>
						<th class="px-3 py-2">Batch Number</th>
						<th class="px-3 py-2">Số lượng</th>
						<th class="px-3 py-2">Giá nhập</th>
						<th class="px-3 py-2">Lợi nhuận</th>
						<th class="px-3 py-2">Ngày hết hạn</th>
						<th class="px-3 py-2">Chọn</th>
					</tr>
				</thead>
				<tbody>
					{#each inventories as inventory}
						<tr
							class="
								border-t
								hover:bg-gray-50
								{selectedBatch === inventory.inventory_id ? 'bg-green-100' : ''}
								{daysLeft(inventory.expiration_date) < 0 ? 'bg-red-100' : ''}
								{daysLeft(inventory.expiration_date) >= 0 && daysLeft(inventory.expiration_date) < 30
								? 'bg-yellow-100'
								: ''}
							"
						>
							<td class="px-3 py-2 text-center">{inventory.batchnumber}</td>
							<td class="px-3 py-2 text-center">{inventory.quantity}</td>
							<td class="px-3 py-2 text-center">{inventory.importprice}</td>
							<td class="px-3 py-2 text-center">{inventory.markup}</td>
							<td class="px-3 py-2 text-center">{inventory.expiration_date}</td>
							<td class="px-3 py-2 text-center">
								<input
									type="radio"
									name="batch"
									class="h-5 w-5 cursor-pointer"
									bind:group={selectedBatch}
									disabled={daysLeft(inventory.expiration_date) < 0}
									on:change={() => selectBatch(inventory.inventory_id)}
								/>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>

			<div class="mt-4 flex items-center gap-4">
				<button
					on:click={saveSelectedBatch}
					class="rounded bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 disabled:bg-gray-400"
					disabled={!selectedBatch || loading}
				>
					Lưu
				</button>
			</div>
		{:else}
			<p>Không có batch nào.</p>
		{/if}
	</div>
</div>
