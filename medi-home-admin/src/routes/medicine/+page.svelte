<script>
	import { Pencil, Trash2, MoreVertical, Eye, Plus } from 'lucide-svelte';
	import { goto } from '$app/navigation';
	import TitlePage from '$lib/components/TitlePage.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import { pageTitle } from '$lib/store.js';

	export let data;

	pageTitle.set('Trang quản lý thuốc');
	let search = '';
	let page = data.page;
	let pageSize = data.pageSize;
	let total = data.total;
	let totalPages = Math.ceil(total / pageSize);

	// let showPopup = true;
	let selectedActionId = null;
	let showDelete = false;
	let medicines = data.medicines;

	function viewMedicine(id){
		goto(`/medicine/${id}`);
	}

	function confirmDelete(id) {
		console.log('funciton', id);
		showDelete = false;
		deleteMedicine(id);
	}

	function openAction(id) {
		selectedActionId = selectedActionId === id ? null : id;
	}

	function goToPage(p) {
		if (p >= 1 && p <= totalPages) page = p;
	}
</script>

<div class="px-4 py-6 sm:px-6 lg:px-8">
	<div class="rounded-md bg-white p-4 shadow-sm">
		<TitlePage titleName="Danh sách thuốc" />

		<!-- Search bar -->
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
						<th class="p-2">Code</th>
						<th class="p-2">Tên thuốc</th>
						<th class="hidden p-2 lg:table-cell">Hình</th>
						<th class="hidden p-2 sm:table-cell">Danh mục</th>
						<th class="hidden p-2 lg:table-cell">Dạng bào chế</th>
						<th class="hidden p-2 lg:table-cell">Kê đơn</th>
						<th class="hidden p-2 md:table-cell">Hành động</th>
						<th class="w-10 p-2 md:hidden"></th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each medicines as medicine}
						<tr class="border-b hover:bg-gray-50">
							<td class="w-20 p-2">{medicine.code}</td>
							<td
								class="max-w-[200px] min-w-[200px] cursor-pointer truncate p-2"
								title={medicine.name}
							>
								{medicine.name}
							</td>
							<td class="hidden p-2 lg:table-cell">{medicine.thumbnail}</td>
							<td class="hidden max-w-[50px] truncate p-2 sm:table-cell"
								>{medicine.category_name}</td
							>
							<td class="hidden p-2 lg:table-cell">{medicine.dosageform_name}</td>
							<td class="hidden p-2 lg:table-cell">{medicine.prescription ? 'Có' : 'Không'}</td>
							<td class="p-2">
								<div class="flex items-center gap-2 text-gray-600">
									<button class="hidden cursor-pointer md:table-cell" title="Xem" on:click={()=>viewMedicine(medicine.medicine_id)}
										><Eye class="h-4 w-4" /></button
									>
									<button class="hidden cursor-pointer md:table-cell" title="Sửa"
										><Pencil class="h-4 w-4" /></button
									>
									<button
										class="hidden cursor-pointer md:table-cell"
										title="Xoá"
										on:click={() => (showDelete = true)}><Trash2 class="h-4 w-4" /></button
									>
									{#if showDelete}
										<!-- Overlay -->
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
													on:click={() => confirmDelete(medicine.medicine_id)}
												>
													Xác nhận
												</button>
											</div>
										</div>
									{/if}
									<div>
										<div class="relative inline-block w-2 md:hidden">
											<button on:click={() => openAction(medicine.medicine_id)}>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>

										{#if selectedActionId === medicine.medicine_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="flex w-30 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
														>
															<Eye class="h-4 w-4" /> Xem
														</button>
													</li>
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
														>
															<Pencil class="h-4 w-4" /> Sửa
														</button>
													</li>
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
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
		<Pagination {page} {pageSize} {totalPages} {total} {goToPage} />
	</div>
</div>
