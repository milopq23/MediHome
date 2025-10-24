<script>
	// import { apiProfile } from '$lib/api/user';
	import { onMount } from 'svelte';

	const props = $props();

	const { data } = props;
	const { user } = data;

	async function saveChanges() {
		const res = await fetch('/auth/profile/', {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json',  },
			body: JSON.stringify({ name: user.name, phone: user.phone })
		});

		console.log(res);

		if (!res.ok) {
			alert('Cập nhật thất bại');
			return;
		}

		const result = await res.json();
		console.log('Response from /profile:', result);
		alert(result.message);
	}

	let isEditing = $state(false);
	let showPassword = $state(false);
	let successMessage = $state('');

	function toggleEdit() {
		isEditing = !isEditing;
		console.log(user);
		if (!isEditing) {
			successMessage = '';
		}
	}

	function handleSubmit(e) {
		e.preventDefault();
		// Simulate save
		successMessage = 'Thông tin đã được cập nhật thành công!';
		isEditing = false;
		setTimeout(() => {
			successMessage = '';
		}, 3000);
	}

	let orders = $state([
		{
			id: 'ORD-001',
			date: '2024-10-20',
			total: '1,250,000 ₫',
			status: 'completed',
			items: 3
		},
		{
			id: 'ORD-002',
			date: '2024-10-18',
			total: '890,000 ₫',
			status: 'completed',
			items: 2
		},
		{
			id: 'ORD-003',
			date: '2024-10-15',
			total: '2,100,000 ₫',
			status: 'shipped',
			items: 5
		},
		{
			id: 'ORD-004',
			date: '2024-10-10',
			total: '450,000 ₫',
			status: 'processing',
			items: 1
		},
		{
			id: 'ORD-005',
			date: '2024-10-05',
			total: '550,000 ₫',
			status: 'completed',
			items: 2
		},
		{
			id: 'ORD-006',
			date: '2024-09-28',
			total: '1,800,000 ₫',
			status: 'completed',
			items: 4
		}
	]);

	let selectedTab = $state('all');
	const filteredOrders = $derived(
		selectedTab === 'all' ? orders : orders.filter((order) => order.status === selectedTab)
	);

	function getStatusBadge(status) {
		const statusMap = {
			completed: { label: 'Hoàn Thành', color: 'bg-green-100 text-[var(--color-success)]' },
			shipped: { label: 'Đang Giao', color: 'bg-blue-100 text-[var(--color-primary)]' },
			processing: { label: 'Đang Xử Lý', color: 'bg-yellow-100 text-[var(--color-warning)]' },
			cancelled: { label: 'Đã Hủy', color: 'bg-red-100 text-[var(--color-error)]' }
		};
		return statusMap[status] || statusMap.processing;
	}

	function formatDate(dateString) {
		const date = new Date(dateString);
		return date.toLocaleDateString('vi-VN', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}
</script>

<div class="mx-auto max-w-6xl px-4 py-8">
	<!-- Header -->
	<div class="mb-8">
		<h1 class="mb-2 text-3xl font-bold text-[var(--color-text)]">Hồ Sơ Cá Nhân</h1>
		<p class="text-[var(--color-text-muted)]">Quản lý thông tin tài khoản và đơn hàng của bạn</p>
	</div>

	<!-- Profile Section -->
	<div class="mb-8 grid grid-cols-1 gap-8 lg:grid-cols-3">
		<!-- Personal Info Form -->
		<div class="lg:col-span-2">
			<div class="rounded-lg border border-[var(--color-border)] bg-white p-6 shadow-sm">
				<div class="mb-6 flex items-center justify-between">
					<h2 class="text-xl font-semibold text-[var(--color-text)]">Thông Tin Cá Nhân</h2>
					<button
						onclick={toggleEdit}
						class="rounded-lg px-4 py-2 font-medium transition-all duration-200 {isEditing
							? 'bg-[var(--color-surface)] text-[var(--color-text)] hover:bg-gray-200'
							: 'bg-[var(--color-primary)] text-black hover:bg-[var(--color-primary-dark)]'}"
					>
						{isEditing ? 'Hủy' : 'Chỉnh Sửa'}
					</button>
				</div>

				{#if successMessage}
					<div class="mb-4 rounded-lg border border-[var(--color-success)] bg-green-50 p-4">
						<p class="font-medium text-[var(--color-success)]">{successMessage}</p>
					</div>
				{/if}

				<form onsubmit={saveChanges} class="space-y-6">
					<!-- Full Name -->
					<div>
						<label for="fullName" class="mb-2 block text-sm font-medium text-[var(--color-text)]">
							Họ và Tên
						</label>
						<input
							type="text"
							id="fullName"
							bind:value={user.name}
							disabled={!isEditing}
							class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none disabled:cursor-not-allowed disabled:bg-[var(--color-surface)]"
						/>
					</div>

					<!-- Email -->
					<div>
						<label for="email" class="mb-2 block text-sm font-medium text-[var(--color-text)]">
							Email
						</label>
						<input
							type="email"
							id="email"
							bind:value={user.email}
							disabled={!isEditing}
							class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none disabled:cursor-not-allowed disabled:bg-[var(--color-surface)]"
						/>
					</div>

					<!-- Phone -->
					<div>
						<label for="phone" class="mb-2 block text-sm font-medium text-[var(--color-text)]">
							Số Điện Thoại
						</label>
						<input
							type="tel"
							id="phone"
							bind:value={user.phone}
							disabled={!isEditing}
							class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none disabled:cursor-not-allowed disabled:bg-[var(--color-surface)]"
						/>
					</div>

					<!-- Gender -->
					<div>
						<label for="gender" class="mb-2 block text-sm font-medium text-[var(--color-text)]">
							Giới Tính
						</label>
						<select
							id="gender"
							bind:value={user.gender}
							disabled={!isEditing}
							class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none disabled:cursor-not-allowed disabled:bg-[var(--color-surface)]"
						>
							<option value="Nam">Nam</option>
							<option value="Nữ">Nữ</option>
							<!-- <option value="other">Khác</option> -->
						</select>
					</div>

					<!-- Password Section -->
					{#if isEditing}
						<div class="border-t border-[var(--color-border)] pt-4">
							<h3 class="mb-4 text-sm font-semibold text-[var(--color-text)]">Đổi Mật Khẩu</h3>

							<!-- Password -->
							<div class="mb-4">
								<label
									for="password"
									class="mb-2 block text-sm font-medium text-[var(--color-text)]"
								>
									Mật Khẩu Mới
								</label>
								<div class="relative">
									<input
										type={showPassword ? 'text' : 'password'}
										id="password"
										bind:value={user.password}
										class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none"
									/>
									<button
										type="button"
										onclick={() => (showPassword = !showPassword)}
										class="absolute top-1/2 right-3 -translate-y-1/2 text-[var(--color-text-muted)] hover:text-[var(--color-text)]"
									>
										{showPassword ? '👁️' : '👁️‍🗨️'}
									</button>
								</div>
							</div>

							<!-- Confirm Password -->
							<div>
								<label
									for="confirmPassword"
									class="mb-2 block text-sm font-medium text-[var(--color-text)]"
								>
									Xác Nhận Mật Khẩu
								</label>
								<input
									type={showPassword ? 'text' : 'password'}
									id="confirmPassword"
									bind:value={user.confirmPassword}
									class="w-full rounded-lg border border-[var(--color-border)] px-4 py-2 focus:ring-2 focus:ring-[var(--color-primary)] focus:outline-none"
								/>
							</div>
						</div>
					{/if}

					<!-- Submit Button -->
					{#if isEditing}
						<button
							type="submit"
							class="w-full rounded-lg bg-[var(--color-primary)] px-4 py-2 font-medium text-black transition-colors duration-200 hover:bg-[var(--color-primary-dark)]"
						>
							Lưu Thay Đổi
						</button>
					{/if}
				</form>
			</div>
		</div>

		<!-- Stats Card -->
		<div class="rounded-lg border border-[var(--color-border)] bg-white p-6 shadow-sm">
			<h2 class="mb-6 text-lg font-semibold text-[var(--color-text)]">Thống Kê</h2>
			<div class="space-y-4">
				<div class="rounded-lg bg-[var(--color-surface)] p-4">
					<p class="mb-1 text-sm text-[var(--color-text-muted)]">Tích Điểm</p>
					<p class="text-2xl font-bold text-[var(--color-primary)]">{user.point}</p>
				</div>
				<div class="rounded-lg bg-[var(--color-surface)] p-4">
					<p class="mb-1 text-sm text-[var(--color-text-muted)]">Đơn Hàng Hoàn Thành</p>
					<p class="text-2xl font-bold text-[var(--color-success)]">12</p>
				</div>
				<div class="rounded-lg bg-[var(--color-surface)] p-4">
					<p class="mb-1 text-sm text-[var(--color-text-muted)]">Tổng Chi Tiêu</p>
					<p class="text-2xl font-bold text-[var(--color-text)]">5,240,000 ₫</p>
				</div>
			</div>
		</div>
	</div>

	<!-- Orders Section -->
	<div class="rounded-lg border border-[var(--color-border)] bg-white shadow-sm">
		<!-- Header -->
		<div class="border-b border-[var(--color-border)] px-6 py-4">
			<h2 class="text-xl font-semibold text-[var(--color-text)]">Lịch Sử Đơn Hàng</h2>
		</div>

		<!-- Tabs -->
		<div class="flex gap-4 border-b border-[var(--color-border)] px-6 pt-4">
			{#each ['all', 'completed', 'shipped', 'processing'] as tab}
				<button
					onclick={() => (selectedTab = tab)}
					class="border-b-2 px-4 py-2 text-sm font-medium transition-colors duration-200 {selectedTab ===
					tab
						? 'border-[var(--color-primary)] text-[var(--color-primary)]'
						: 'border-transparent text-[var(--color-text-muted)] hover:text-[var(--color-text)]'}"
				>
					{tab === 'all'
						? 'Tất Cả'
						: tab === 'completed'
							? 'Hoàn Thành'
							: tab === 'shipped'
								? 'Đang Giao'
								: 'Đang Xử Lý'}
				</button>
			{/each}
		</div>

		<!-- Table -->
		<div class="overflow-x-auto">
			<table class="w-full">
				<thead>
					<tr class="border-b border-[var(--color-border)] bg-[var(--color-surface)]">
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Mã Đơn Hàng
						</th>
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Ngày Đặt
						</th>
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Số Lượng
						</th>
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Tổng Tiền
						</th>
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Trạng Thái
						</th>
						<th class="px-6 py-4 text-left text-sm font-semibold text-[var(--color-text)]">
							Hành Động
						</th>
					</tr>
				</thead>
				<tbody>
					{#each filteredOrders as order (order.id)}
						<tr
							class="border-b border-[var(--color-border)] transition-colors hover:bg-[var(--color-surface)]"
						>
							<td class="px-6 py-4 text-sm font-medium text-[var(--color-primary)]">
								{order.id}
							</td>
							<td class="px-6 py-4 text-sm text-[var(--color-text)]">
								{formatDate(order.date)}
							</td>
							<td class="px-6 py-4 text-sm text-[var(--color-text)]">
								{order.items} sản phẩm
							</td>
							<td class="px-6 py-4 text-sm font-semibold text-[var(--color-text)]">
								{order.total}
							</td>
							<td class="px-6 py-4 text-sm">
								<span
									class="rounded-full px-3 py-1 text-xs font-medium {getStatusBadge(order.status)
										.color}"
								>
									{getStatusBadge(order.status).label}
								</span>
							</td>
							<td class="px-6 py-4 text-sm">
								<button
									class="font-medium text-[var(--color-primary)] transition-colors hover:text-[var(--color-primary-dark)]"
								>
									Chi Tiết
								</button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		<!-- Empty State -->
		<!-- {#if filteredOrders.length === 0}
			<div class="px-6 py-12 text-center">
				<p class="text-[var(--color-text-muted)]">Không có đơn hàng nào</p>
			</div>
		{/if} -->

		<!-- Footer -->
		<!-- <div class="flex items-center justify-between border-t border-[var(--color-border)] px-6 py-4">
			<p class="text-sm text-[var(--color-text-muted)]">
				Hiển thị {filteredOrders.length} trong {orders.length} đơn hàng
			</p>
			<button
				class="rounded-lg px-4 py-2 text-sm font-medium text-[var(--color-primary)] transition-colors hover:bg-[var(--color-surface)]"
			>
				Xem Tất Cả
			</button>
		</div> -->
	</div>
</div>
