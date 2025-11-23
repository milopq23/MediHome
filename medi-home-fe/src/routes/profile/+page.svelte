<script>
	import { GetAllOrder, GetProfile, GetDetailOrder } from '$lib/api/user';
	import { onMount } from 'svelte';

	let info = {};
	let orders = [];
	let selectedOrder = null; // Lưu chi tiết đơn hàng khi mở popup
	let showModal = false; // Điều khiển hiển thị modal

	const statusMap = {
		all: null,
		completed: 'Hoàn thành',
		shipped: 'Đang giao',
		processing: 'Chờ xác nhận',
		cancelled: 'Đã hủy'
	};

	const countCompleted = () => orders.filter((o) => o.order_status === 'Hoàn thành').length;

	async function getProfile() {
		info = await GetProfile();
	}

	async function getOrderUser(status) {
		const data = await GetAllOrder(status);
		console.log('daat', data);
		orders = data.data;
	}

	async function openOrderDetail(order_id) {
		try {
			const result = await GetDetailOrder(order_id);
			selectedOrder = result.data;
			showModal = true;
		} catch (error) {
			console.error('Lỗi khi lấy chi tiết đơn hàng:', error);
			alert('Không thể tải chi tiết đơn hàng');
		}
	}

	function closeModal() {
		showModal = false;
		selectedOrder = null;
	}

	onMount(() => {
		getProfile();
		getOrderUser(null); // mặc định load tất cả
	});

	let selectedTab = 'all';

	function formatDate(dateString) {
		const date = new Date(dateString);
		return date.toLocaleDateString('vi-VN');
	}

	function formatCurrency(amount) {
		return new Intl.NumberFormat('vi-VN', {
			style: 'currency',
			currency: 'VND'
		}).format(amount);
	}
</script>

<div class="mx-auto max-w-6xl px-4 py-8">
	<div class="mb-8">
		<h1 class="mb-2 text-center text-5xl font-bold">Thông tin cá nhân</h1>
		<p class="text-center text-xl">Quản lý thông tin tài khoản và đơn hàng của bạn</p>
	</div>

	<div class="mb-8 grid grid-cols-1 gap-8 lg:grid-cols-3">
		<!-- PROFILE -->
		<div class="lg:col-span-2">
			<div class="rounded-lg border bg-white p-6 shadow-sm">
				<h2 class="mb-6 text-xl font-semibold">Thông Tin Cá Nhân</h2>

				<div class="space-y-6">
					<div>
						<label for="fullname" class="mb-2 block text-sm font-medium">Họ và Tên</label>
						<input
							type="text"
							bind:value={info.name}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label for="email" class="mb-2 block text-sm font-medium">Email</label>
						<input
							type="email"
							bind:value={info.email}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label for="phone" class="mb-2 block text-sm font-medium">Số Điện Thoại</label>
						<input
							type="tel"
							bind:value={info.phone}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label for="gender" class="mb-2 block text-sm font-medium">Giới Tính</label>
						<input
							type="text"
							bind:value={info.gender}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>
				</div>
			</div>
		</div>

		<!-- Stats -->
		<div class="rounded-lg border bg-white p-6 shadow-sm">
			<h2 class="mb-6 text-lg font-semibold">Thống Kê</h2>
			<div class="space-y-4">
				<div class="rounded-lg bg-gray-50 p-4">
					<p class="text-sm text-gray-500">Tích Điểm</p>
					<p class="text-2xl font-bold text-blue-600">{info.point}</p>
				</div>

				<!-- <div class="rounded-lg bg-gray-50 p-4">
					<p class="text-sm text-gray-500">Đơn Hoàn Thành</p>
					<p class="text-2xl font-bold text-green-600">{countCompleted()}</p>
				</div> -->

				<!-- <div class="rounded-lg bg-gray-50 p-4">
					<p class="text-sm text-gray-500">Tổng Chi Tiêu</p>
					<p class="text-2xl font-bold">5,240,000 ₫</p>
				</div> -->
			</div>
		</div>
	</div>

	<!-- ORDERS -->
	<div class="rounded-lg border bg-white shadow-sm">
		<div class="border-b px-6 py-4">
			<h2 class="text-xl font-semibold">Lịch Sử Đơn Hàng</h2>
		</div>

		<div class="flex gap-4 border-b px-6 pt-4">
			{#each ['all', 'completed', 'shipped', 'processing', 'cancelled'] as tab}
				<button
					on:click={() => {
						selectedTab = tab;
						getOrderUser(statusMap[tab]); // ← GỌI API
					}}
					class="border-b-2 px-4 py-2 text-sm font-medium
        {selectedTab === tab
						? 'border-blue-500 text-blue-600'
						: 'border-transparent text-gray-500 hover:text-gray-700'}"
				>
					{tab === 'all'
						? 'Tất Cả'
						: tab === 'completed'
							? 'Hoàn Thành'
							: tab === 'shipped'
								? 'Đang Giao'
								: tab === 'processing'
									? 'Chờ xác nhận'
									: 'Đã huỷ'}
				</button>
			{/each}
		</div>

		<div class="overflow-x-auto">
			<table class="w-full">
				<thead>
					<tr class="border-b bg-gray-50">
						<th class="px-6 py-4 text-left text-sm font-semibold">Mã Đơn</th>
						<th class="px-6 py-4 text-left text-sm font-semibold">Ngày</th>
						<th class="px-6 py-4 text-left text-sm font-semibold">SL</th>
						<th class="px-6 py-4 text-left text-sm font-semibold">Tổng</th>
						<th class="px-6 py-4 text-left text-sm font-semibold">Trạng Thái</th>
						<th class="px-6 py-4 text-left text-sm font-semibold">Hành Động</th>
					</tr>
				</thead>

				<tbody>
					{#each orders as order}
						<tr class="border-b hover:bg-gray-50">
							<td class="px-6 py-4 font-medium text-blue-600">{order.order_id}</td>
							<td class="px-6 py-4">{formatDate(order.date)}</td>
							<td class="px-6 py-4">{order.order_item}</td>
							<td class="px-6 py-4 font-semibold">{formatCurrency(order.final_amount)}</td>
							<td class="px-6 py-4">
								<span class="rounded-full px-3 py-1 text-xs font-medium">
									{order.order_status}
								</span>
							</td>
							<td class="px-6 py-4">
								<button
									on:click={() => openOrderDetail(order.order_id)}
									class="font-medium text-blue-600 hover:underline"
								>
									Chi Tiết
								</button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</div>
<!-- Modal Chi Tiết Đơn Hàng -->

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
									<td class="py-3 text-right">{formatCurrency(item.unit_price)}</td>
									<td class="py-3 text-right font-medium"
										>{formatCurrency(item.unit_price * item.quantity)}</td
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
						<span class="text-blue-600">{formatCurrency(selectedOrder.final_amount)}</span>
					</div>
					{#if selectedOrder.voucher_code}
						<div class="mt-1 flex justify-between text-sm text-gray-600">
							<span>Giảm giá ({selectedOrder.voucher_code}):</span>
							<span>-{formatCurrency(selectedOrder.total_amount - selectedOrder.final_amount)}</span
							>
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
