<script>
	import { Pencil, Trash2, MoreVertical } from 'lucide-svelte';
	import { ListInventory } from '$lib/api/inventory.js';
	import { onMount } from 'svelte';

	let search;
	let showDelete = false;
	let inventories = [];

	onMount(() => {
		listInventory();
	});

	async function listInventory() {
		const data = await ListInventory();
		inventories = data;
		console.log(inventories);
	}

	
	const formatVND = (value) => {
		return new Intl.NumberFormat('vi-VN').format(value) + 'đ';
	};

	function formatDate(isoString) {
		const date = new Date(isoString);
		const dd = String(date.getDate()).padStart(2, '0');
		const mm = String(date.getMonth() + 1).padStart(2, '0');
		const yyyy = date.getFullYear();
		return `${dd}/${mm}/${yyyy}`;
	}
</script>

<div class="px-4 py-6 sm:px-6 lg:px-8">
	<div class="rounded-md bg-white p-4 shadow-sm">
		<div class="mb-4 flex flex-col gap-2 sm:flex-row sm:items-center">
			<input
				type="text"
				class="w-full rounded-md border px-3 py-2 text-sm sm:max-w-xs"
				placeholder="Tìm kiếm thuốc..."
				bind:value={search}
			/>
		</div>
		<div class="overflow-x-auto">
			<table class="min-w-full divide-y divide-gray-200 text-sm">
				<thead class="bg-gray-50 text-left text-gray-500">
					<tr>
						<th class="table-cell max-w-50 min-w-15 p-2">Mã thuốc</th>
						<th class="table-cell p-2">Tên thuốc</th>
						<th class="table-cell p-2">Lô</th>
						<th class="table-cell p-2">Giá nhập</th>
						<th class="table-cell p-2">Lợi nhuận</th>
						<th class="table-cell p-2">Giá bán</th>
						<th class="table-cell p-2">Ngày hết hạn</th>
						<th class="table-cell p-2">Số lượng tồn</th>
						<th class="table-cell p-2">Trạng thái</th>
						<th class=""></th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each inventories as inventory}
						<tr class="border-b hover:bg-gray-50">
							<td class="table-cell p-2">{inventory.code}</td>
							<td class="table-cell p-2">
								{inventory.name}
							</td>
							<td class="table-cell p-2">{inventory.batch_number}</td>
							<td class="table-cell p-2">{formatVND(inventory.import_price)}</td>
							<td class="table-cell p-2">{inventory.mark_up}%</td>
							<td class="table-cell p-2">{formatVND(inventory.selling_price)}</td>
							<td class="table-cell p-2">{formatDate(inventory.expiration_date)}</td>
							<td class="table-cell p-2">{inventory.quantity}</td>
							<td class="table-cell p-2">{inventory.status}</td>
							<td class="p-2">
								<div class="flex items-center gap-2 text-gray-600">
									<button
										class="hidden cursor-pointer md:table-cell"
										title="Sửa"
										on:click={() => viewDetailinventory(inventory.inventory_id)}
										><Pencil class="h-4 w-4" /></button
									>
									<button
										class="hidden cursor-pointer md:table-cell"
										title="Xoá"
										on:click={() => (showDelete = true)}><Trash2 class="h-4 w-4" /></button
									>
									{#if showDelete}
										<div class="fixed inset-0 z-40 bg-black/50"></div>

										<div
											class="fixed top-1/2 left-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-xl"
										>
											<h2 class="mb-4 text-xl font-bold">Xác nhận xoá</h2>
											<p>Bạn có chắc chắn muốn xoá không?</p>

											<div class="mt-6 flex justify-end gap-3">
												<button
													class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
													on:click={() => (showDelete = false)}
												>
													Hủy
												</button>
												<button
													class="rounded bg-red-600 px-4 py-2 text-white hover:bg-red-700"
													on:click={() => confirmDelete(inventory.inventory_id)}
												>
													Xác nhận
												</button>
											</div>
										</div>
									{/if}
									<!-- <div>
										<div class="relative inline-block w-2 md:hidden">
											<button on:click={() => openMoreOption(inventory.inventory_id)}>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>

										{#if selectOption === inventory.inventory_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
															on:click={() => viewDetailinventory(inventory.inventory_id)}
														>
															<Pencil class="h-4 w-4" /> Sửa
														</button>
													</li>
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
															on:click={() => (showDelete = true)}
														>
															<Trash2 class="h-4 w-4" /> Xoá
														</button>
													</li>
												</ul>
											</div>
										{/if}
									</div> -->
								</div>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
		<!-- <Pagination {page} {pageSize} {totalPages} {total} {goToPage} /> -->
	</div>
</div>
