<script>
	import { ListDosage, DeleteDosage } from '$lib/api/dosageform.js';
	import { Pencil, Trash2, MoreVertical, Eye, Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let search = '';
	let dosageforms = [];
	let showDelete = false;
	let selectOption = null;

	onMount(() => {
		listDosages();
	});

	async function listDosages() {
		const data = await ListDosage();
		dosageforms = data;
	}

	async function deleteDosage(id) {
		await DeleteDosage(id);
	}

	async function confirmDelete(id) {
		try {
			await deleteDosage(id);
			// toasts.add({ message: 'Xóa thành công', type: 'success' });
		} catch (error) {
			// toasts.add({ message: 'Xóa thất bại', type: 'error' });
		}
		showDelete = false;
		await listDosages();
	}

	function viewDetailDosage(id) {
		// toasts.add({ message: 'Lưu thành công', type: 'success' });
		goto(`/dosage-form/${id}`);
	}

	function create() {
		goto(`/dosage-form/create`);
	}

	function openMoreOption(id) {
		selectOption = selectOption === id ? null : id;
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
						<th class="table-cell max-w-50 min-w-15 p-2">Mã</th>
						<th class="table-cell w-30 p-2">Dạng bào chế</th>
						<th class="table-cell p-2">Mô tả</th>
						<th class=""></th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each dosageforms as dosage}
						<tr class="border-b hover:bg-gray-50">
							<td class="table-cell p-2">{dosage.dosageform_id}</td>
							<td class="table-cell p-2">
								{dosage.name}
							</td>
							<td class="table-cell p-2">
								<div class="line-clamp-3">{dosage.description}</div>
							</td>
							<td class="p-2">
								<div class="flex items-center gap-2 text-gray-600">
									<button
										class="hidden cursor-pointer md:table-cell"
										title="Sửa"
										on:click={() => viewDetailDosage(dosage.dosageform_id)}
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
													on:click={() => confirmDelete(dosage.dosageform_id)}
												>
													Xác nhận
												</button>
											</div>
										</div>
									{/if}
									<div>
										<div class="relative inline-block w-2 md:hidden">
											<button on:click={() => openMoreOption(dosage.dosageform_id)}>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>

										{#if selectOption === dosage.dosageform_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
															on:click={() => viewDetailDosage(dosage.dosageform_id)}
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
									</div>
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
