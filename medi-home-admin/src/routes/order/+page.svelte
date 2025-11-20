<script>
	import { ListOrder, GetOrderDetail, ListOrderStatus, UpdateStatusOrder } from '$lib/api/order.js';
	import { Eye, Trash2, MoreVertical, Check, X } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let search = '';
	let orders = [];
	let showCancel = false;
	let selectOption = null;
	let selectedOrder = {};
	let showModal = false;
	let selectedOrderId = null;

	const nextStatusMap = {
		'Chờ xác nhận': 'Đã xác nhận',
		'Đã xác nhận': 'Đang giao',
		'Đang giao': 'Hoàn thành',
		'Hoàn thành': null,
		'Đã hủy': null
	};

	const cancelStatusMap = {
		'Chờ xác nhận': 'Đã hủy',
		'Đã xác nhận': 'Đã hủy',
		'Đang giao': null,
		'Hoàn thành': null,
		'Đã hủy': null
	};

	onMount(() => {
		listOrders();
	});

	async function openOrderDetail(order_id) {
		try {
			const result = await GetOrderDetail(order_id);
			selectedOrder = result.data;
			showModal = true;
		} catch (error) {
			console.error('Lỗi khi lấy chi tiết đơn hàng:', error);
			alert('Không thể tải chi tiết đơn hàng');
		}
	}

	async function listOrders() {
		const data = await ListOrder();
		orders = data.data;
	}

	async function listOrderStatus(status) {
		const data = await ListOrderStatus(status);
		orders = data.data;
	}

	async function approveOrder(order) {
		console.log(order.order_id);
		const next = nextStatusMap[order.order_status];
		if (!next) {
			alert('Trạng thái hiện tại không thể duyệt tiếp');
			return;
		}
		await UpdateStatusOrder(order.order_id, next);
		await listOrders();
	}

	async function cancelOrder(order) {
		console.log(order.order_id);
		const cancle = cancelStatusMap[order.order_status];
		if (!cancle) {
			alert('Trạng thái hiện tại không thể duyệt tiếp');
			return;
		}
		await UpdateStatusOrder(order.order_id, cancle);
		await listOrders();
	}

	function closeModal() {
		showModal = false;
		selectedOrder = null;
	}

	function openCancel(order_id) {
		selectedOrderId = order_id;
		showCancel = true;
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

<div class="px-2 py-6 sm:px-6 lg:px-8">
	<div class="rounded-md bg-white p-4 shadow-sm">
		<!-- <TitlePage titleName="Danh sách thuốc" {create} /> -->
		<div class="mb-4 flex flex-col gap-2 sm:flex-row sm:items-center">
			<input
				type="text"
				class="w-full rounded-md border px-3 py-2 text-sm sm:max-w-xs"
				placeholder="Tìm kiếm thuốc..."
				bind:value={search}
			/>
		</div>
		<div class="block w-full max-w-screen overflow-x-auto py-2 whitespace-nowrap">
			<div class="flex flex-wrap gap-2 overflow-x-auto py-2">
				<button class="rounded bg-gray-300 px-4 py-2" on:click={() => listOrders()}>Tất cả</button>
				<button
					class="rounded bg-yellow-500 px-4 py-2 text-white"
					on:click={() => listOrderStatus('Chờ xác nhận')}>Chờ xác nhận</button
				>
				<button
					class="rounded bg-blue-500 px-4 py-2 text-white"
					on:click={() => listOrderStatus('Đang giao')}>Đang giao</button
				>
				<button
					class="rounded bg-green-500 px-4 py-2 text-white"
					on:click={() => listOrderStatus('Hoàn thành')}>Hoàn thành</button
				>
				<button
					class="rounded bg-red-500 px-4 py-2 text-white"
					on:click={() => listOrderStatus('Đã hủy')}>Hủy</button
				>
			</div>
		</div>
		<div class="">
			<table class="min-w-full divide-y divide-gray-200 text-sm">
				<thead class="bg-gray-50 text-left text-gray-500">
					<tr>
						<th class="p-2 text-center md:table-cell">Mã đơn hàng</th>
						<th class="hidden p-2 text-center lg:table-cell">Ngày tạo đơn</th>
						<th class="hidden p-2 text-center lg:table-cell">Tên khách hàng</th>
						<th class="hidden p-2 text-center lg:table-cell">Sản phẩm</th>
						<th class="table-cell p-2 text-center">Trạng thái đơn</th>
						<th class="table-cell p-2 text-center">Tổng tiền</th>
						<th class="hidden p-2 text-center md:table-cell">Hình thức thanh toán</th>
						<th class="hidden p-2 text-center md:table-cell">Trạng thái thanh toán</th>
						<th class="w-10 p-2 md:hidden"></th>
						<th class="hidden w-50 p-2 md:table-cell">Hành động</th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each orders as order}
						<tr class="border-b hover:bg-gray-50">
							<td class="table-cell truncate p-2 text-center">#{order.order_id} </td>
							<td class="hidden p-2 text-center lg:table-cell">{formatDate(order.date)}</td>
							<td class="hidden p-2 text-center lg:table-cell">
								<div class="flex flex-1 flex-col">
									<span>{order.full_name}</span>
									<span class="font-semibold">08869100</span>
								</div>
							</td>
							<td class="hidden p-2 text-center lg:table-cell">{order.order_item}</td>
							<td class="table-cell p-2 text-center">{order.order_status}</td>
							<td class="table-cell p-2 text-center">{formatVND(order.final_amount)}</td>
							<td class="hidden p-2 text-center md:table-cell">{order.payment_method}</td>
							<td class="hidden p-2 text-center md:table-cell">{order.payment_status}</td>

							<td class="p-2">
								<div class="flex items-center gap-4 text-gray-500">
									<button
										class="hidden cursor-pointer p-2 hover:text-blue-500 md:table-cell"
										title="Xem"
										on:click={() => openOrderDetail(order.order_id)}><Eye class="h-4 w-4" /></button
									>
									<button
										class="hidden cursor-pointer p-2 hover:text-green-500 md:table-cell"
										title="Duyệt"
										on:click={() => approveOrder(order)}><Check class="h-4 w-4" /></button
									>
									<button
										class="hidden cursor-pointer p-2 hover:text-red-500 md:table-cell"
										title="Hủy"
										on:change={() => console.log(order.order_id)}
										on:click={() => cancelOrder(order)}><X class="h-4 w-4" /></button
									>
									{#if showCancel}
										<div class="fixed inset-0 z-40 bg-black/50"></div>

										<div
											class="fixed top-1/2 left-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-xl"
										>
											<h2 class="mb-4 text-xl font-bold">Hủy đơn hàng</h2>
											<p>Bạn có chắc chắn muốn hủy đơn hàng #{selectedOrderId} không?</p>

											<div class="mt-6 flex justify-end gap-3">
												<button
													class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
													on:click={() => (showCancel = false)}
												>
													Hủy
												</button>
												<button
													class="rounded bg-red-600 px-4 py-2 text-white hover:bg-red-700"
													on:click={() => cancelOrder(order)}
												>
													Xác nhận
												</button>
											</div>
										</div>
									{/if}
									<div>
										<div class="relative inline-block w-2 md:hidden">
											<button
												on:change={console.log(order.order_id)}
												on:click={() => openMoreOption(order.order_id)}
											>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>
										{#if selectOption === order.order_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="hidden cursor-pointer p-2 hover:text-blue-500 md:table-cell"
															title="Sửa"
															on:click={() => openOrderDetail(order.order_id)}
															><Eye class="h-4 w-4" /></button
														>
													</li>
													<li>
														<button
															class="hidden cursor-pointer p-2 hover:text-green-500 md:table-cell"
															title="Duyệt"
															on:click={() => openOrderDetail(order.order_id)}
															><Check class="h-4 w-4" /></button
														>
													</li>
													<li>
														<button
															class="hidden cursor-pointer p-2 hover:text-red-500 md:table-cell"
															title="Hủy"
															on:click={() => (showCancel = true)}><X class="h-4 w-4" /></button
														>
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

{#if showModal && selectedOrder}
	<div class="bg-opacity-50 fixed inset-0 z-50 flex items-center justify-center bg-black p-4">
		<div class="max-h-full w-full max-w-4xl overflow-y-auto rounded-lg bg-white shadow-xl">
			<div class="sticky top-0 border-b bg-white p-6">
				<div class="flex items-center justify-between">
					<h3 class="text-2xl font-bold">Chi Tiết Đơn Hàng #{selectedOrder.order_id}</h3>
					<button on:click={closeModal} class="text-gray-400 hover:text-gray-600">
						<X class="h-5 w-5" />
					</button>
				</div>
			</div>

			<div class="space-y-8 p-6">
				<!-- Thông tin khách hàng -->
				<div>
					<h4 class="mb-3 text-lg font-semibold text-gray-700">Thông Tin Nhận Hàng</h4>
					<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
						<div>
							<p class="text-sm text-gray-500">Họ tên</p>
							<p class="font-medium">{selectedOrder.full_name}</p>
						</div>
						<div>
							<p class="text-sm text-gray-500">Số điện thoại</p>
							<p class="font-medium">{selectedOrder.phone}</p>
						</div>
						<div class="md:col-span-2">
							<p class="text-sm text-gray-500">Địa chỉ giao hàng</p>
							<p class="font-medium">{selectedOrder.address}</p>
						</div>
					</div>
				</div>

				<!-- Thông tin đơn hàng -->
				<div>
					<h4 class="mb-3 text-lg font-semibold text-gray-700">Thông Tin Đơn Hàng</h4>
					<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
						<div>
							<p class="text-sm text-gray-500">Phương thức thanh toán</p>
							<p class="font-medium">{selectedOrder.payment_method}</p>
						</div>
						<div>
							<p class="text-sm text-gray-500">Trạng thái thanh toán</p>
							<p class="font-medium">
								<span
									class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium
									{selectedOrder.payment_status === 'Chưa thanh toán'
										? 'bg-red-100 text-red-800'
										: 'bg-green-100 text-green-800'}"
								>
									{selectedOrder.payment_status}
								</span>
							</p>
						</div>
						<div>
							<p class="text-sm text-gray-500">Đơn vị vận chuyển</p>
							<p class="font-medium">{selectedOrder.shipping_name}</p>
						</div>
						<div>
							<p class="text-sm text-gray-500">Mã giảm giá</p>
							<p class="font-medium">{selectedOrder.voucher_code || 'Không sử dụng'}</p>
						</div>
					</div>
				</div>

				<!-- Danh sách sản phẩm -->
				<div>
					<h4 class="mb-3 text-lg font-semibold text-gray-700">Sản Phẩm</h4>
					<table class="w-full">
						<thead>
							<tr class="border-b text-left text-sm font-medium text-gray-700">
								<th class="pb-2">Tên thuốc</th>
								<th class="pb-2 text-center">Loại</th>
								<th class="pb-2 text-center">Số lượng</th>
								<th class="pb-2 text-right">Đơn giá</th>
								<th class="pb-2 text-right">Thành tiền</th>
							</tr>
						</thead>
						<tbody>
							{#each selectedOrder.order_detail as item}
								<tr class="border-b">
									<td class="py-3 font-medium">{item.medicine_name}</td>
									<td class="py-3 text-center text-sm text-gray-600">{item.type}</td>
									<td class="py-3 text-center">{item.quantity}</td>
									<td class="py-3 text-right">{formatVND(item.unit_price)}</td>
									<td class="py-3 text-right font-medium"
										>{formatVND(item.unit_price * item.quantity)}</td
									>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>

				<!-- Tổng tiền -->
				<div class="border-t pt-4">
					<div class="flex justify-between text-lg font-semibold">
						<span>Tổng cộng:</span>
						<span class="text-blue-600">{formatVND(selectedOrder.final_amount)}</span>
					</div>
					{#if selectedOrder.voucher_code}
						<div class="mt-1 flex justify-between text-sm text-gray-600">
							<span>Giảm giá ({selectedOrder.voucher_code}):</span>
							<span>-{formatVND(selectedOrder.total_amount - selectedOrder.final_amount)}</span>
						</div>
					{/if}
				</div>
			</div>

			<div class="border-t bg-gray-50 p-6">
				<div class="flex justify-end gap-3">
					<button
						on:click={closeModal}
						class="rounded-lg border border-gray-300 px-5 py-2.5 text-sm font-medium text-gray-700 hover:bg-gray-100"
					>
						Đóng
					</button>
				</div>
			</div>
		</div>
	</div>
{/if}
