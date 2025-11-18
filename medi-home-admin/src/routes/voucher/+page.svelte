<script>
	import { ListVoucher, DeleteVoucher } from '$lib/api/voucher.js';
	import { Pencil, Trash2, MoreVertical, Eye, Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let search = '';
	let vouchers = [];
	let showDelete = false;
	let selectOption = null;

	onMount(() => {
		listVouchers();
	});

	async function listVouchers() {
		const data = await ListVoucher();
		vouchers = data;
	}

	async function deleteVoucher(id) {
		await DeleteVoucher(id);
	}

	async function confirmDelete(id) {
		try {
			await deleteVoucher(id);
			toasts.add({ message: 'Xóa thành công', type: 'success' });

			await listVouchers(page);
		} catch (error) {
			toasts.add({ message: 'Xóa thất bại', type: 'error' });
		}
		showDelete = false;
	}

	function viewDetailVoucher(id) {
		toasts.add({ message: 'Lưu thành công', type: 'success' });
		goto(`/voucher/${id}`);
	}

	function create() {
		goto(`/voucher/create`);
	}

	function openMoreOption(id) {
		selectOption = selectOption === id ? null : id;
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
		<!-- <TitlePage titleName="Danh sách thuốc" {create} /> -->

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
						<th class="hidden p-2 md:table-cell">Code</th>
						<th class="p-2">Voucher</th>
						<th class="hidden w-30 p-2 sm:table-cell">Loại</th>
						<th class="hidden w-30 p-2 lg:table-cell">Giá trị</th>
						<th class="hidden w-30 p-2 lg:table-cell">Số lượng mã</th>
						<th class="w-30 p-2">Trạng thái</th>
						<th class="hidden p-2 lg:table-cell">Thời gian</th>
						<th class="w-10 p-2 md:hidden"></th>
						<th class="hidden w-50 p-2 md:table-cell">Hành động</th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each vouchers as voucher}
						<tr class="border-b hover:bg-gray-50">
							<td class="hidden w-10 cursor-pointer truncate p-2 md:table-cell">{voucher.code}</td>
							<td class="max-w-[100px] cursor-pointer truncate p-2">
								{voucher.name}
							</td>
							<td class="hidden w-30 p-2 sm:table-cell">{voucher.discount_type}</td>
							{#if voucher.discount_type === 'Phần trăm'}
								<td class="hidden w-30 p-2 lg:table-cell">{voucher.discount_value}%</td>
							{:else}
								<td class="hidden w-30 p-2 lg:table-cell">{formatVND(voucher.max_discountvalue)}</td
								>
							{/if}
							<td class="hidden w-30 p-2 lg:table-cell"
								>{voucher.used_count}/{voucher.usage_limit}</td
							>
							<td>{voucher.is_active === true ? 'Hết hạn' : 'Sẵn sàng'}</td>
							<td class="hidden max-w-[50px] truncate p-2 lg:table-cell"
								>{formatDate(voucher.start_date)} - {formatDate(voucher.end_date)}
							</td>
							<td class="p-2">
								<div class="flex items-center gap-2 text-gray-600">
									<button
										class="hidden cursor-pointer md:table-cell"
										title="Sửa"
										on:click={() => viewDetailVoucher(voucher.voucher_id)}
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
													on:click={() => confirmDelete(voucher.voucher_id)}
												>
													Xác nhận
												</button>
											</div>
										</div>
									{/if}
									<div>
										<div class="relative inline-block w-2 md:hidden">
											<button on:click={() => openMoreOption(voucher.voucher_id)}>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>

										{#if selectOption === voucher.voucher_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
															on:click={() => viewDetailVoucher(voucher.voucher_id)}
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
