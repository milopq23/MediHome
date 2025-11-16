<script>
	import { GetAllOrder, GetProfile } from '$lib/api/user';
	import { onMount } from 'svelte';

	let info = {};
	let orders = [];

	const statusMap = {
		all: null,
		completed: 'Hoàn thành',
		shipped: 'Đang giao',
		processing: 'Chờ xác nhận',
		cancelled: 'Đã hủy'
	};

	const countCompleted = () => orders.filter((o) => o.order_status === 'Hoàn thành').length;

	async function getProfile(user_id) {
		info = await GetProfile(user_id);
	}

	async function getOrderUser(user_id, status) {
		const data = await GetAllOrder(user_id, status);
		orders = data.data;
	}

	onMount(() => {
		getProfile(31);
		getOrderUser(31, null); // mặc định load tất cả
	});

	let selectedTab = 'all';

	function formatDate(dateString) {
		const date = new Date(dateString);
		return date.toLocaleDateString('vi-VN');
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
						<label class="mb-2 block text-sm font-medium">Họ và Tên</label>
						<input
							type="text"
							bind:value={info.name}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label class="mb-2 block text-sm font-medium">Email</label>
						<input
							type="email"
							bind:value={info.email}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label class="mb-2 block text-sm font-medium">Số Điện Thoại</label>
						<input
							type="tel"
							bind:value={info.phone}
							readonly
							class="w-full rounded-lg border bg-gray-50 px-4 py-2"
						/>
					</div>

					<div>
						<label class="mb-2 block text-sm font-medium">Giới Tính</label>
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

				<div class="rounded-lg bg-gray-50 p-4">
					<p class="text-sm text-gray-500">Đơn Hoàn Thành</p>
					<p class="text-2xl font-bold text-green-600">{countCompleted()}</p>
				</div>

				<div class="rounded-lg bg-gray-50 p-4">
					<p class="text-sm text-gray-500">Tổng Chi Tiêu</p>
					<p class="text-2xl font-bold">5,240,000 ₫</p>
				</div>
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
						getOrderUser(31, statusMap[tab]); // ← GỌI API
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
							<td class="px-6 py-4 font-semibold">{order.final_amount}</td>
							<td class="px-6 py-4">
								<span class="rounded-full px-3 py-1 text-xs font-medium">
									{order.order_status}
								</span>
							</td>
							<td class="px-6 py-4">
								<button class="font-medium text-blue-600 hover:underline">Chi Tiết</button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</div>
