<script>
	// import ToolBar from '$lib/components/ToolBar.svelte';
	import { UserPlus, Pencil, Trash2, Eye } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { getAllAdmin, addAdmin, editAdmin, deleteAdmin, getTotalRole } from '$lib/api/admin.js';

	let showForm = false;
	let showConfirmDelete = false;

	let admins = [];
	let page = 1;
	let pageSize = 10;
	let total = 0;
	let totalPages = 1;

	let adminTotal = 1;
	let staffTotal = 1;

	let admin = {
		admin_id: null,
		name: '',
		email: '',
		password: '',
		phone: '',
		role: ''
	};

	let formMode = 'create' || 'view' || 'edit';

	$: formTitle =
		formMode === 'create'
			? 'Tạo quản trị viên'
			: formMode === 'edit'
				? 'Chỉnh sửa quản trị viên'
				: 'Chi tiết quản trị viên';

	async function reloadAction() {
		await loadAdmins();
		await totalByRole('Admin');
		await totalByRole('Staff');
	}
	async function totalByRole(role) {
		try {
			const count = await getTotalRole(role)
			if(role === 'Admin') adminTotal = count
			else staffTotal = count;
		} catch (err) {
			// error = 'Không thể tải.';
			console.error(err);
		}
	}
	// 

	async function loadAdmins(currentPage = 1) {
		try {
			const result = await getAllAdmin(currentPage, pageSize);
			admins = result.admins;
			page = result.page;
			pageSize = result.pageSize;
			total = result.total;
			totalPages = Math.ceil(total / pageSize);
		} catch (err) {
			error = 'Không thể tải danh sách.';
			console.error(err);
		}
	}

	async function submitForm(event) {
		event.preventDefault();
		try {
			if (formMode === 'create') {
				const result = await addAdmin(admin);
				console.log('Tạo thành công:', result);
			} else if (formMode === 'edit') {
				const result = await editAdmin(admin.admin_id, admin);
				console.log('Cập nhật thành công:', result);
			}
			toggleForm();
			await reloadAction();
		} catch (error) {
			console.log(error);
		}
	}

	async function confirmAdminDelete(id) {
		try {
			await deleteAdmin(id); // Gọi API xoá
			showConfirmDelete = false;

			await reloadAction();
		} catch (err) {
			console.error('Xoá thất bại:', err);
		}
	}

	onMount(() => {
		loadAdmins();
		totalByRole('Admin');
		totalByRole('Staff');
		// totalActice();
	});

	function confirmDelete(id) {
		showConfirmDelete = true;
		admin.admin_id = id;
	}

	function openForm(mode, adminData) {
		formMode = mode;
		if (adminData) {
			admin = { ...adminData };
		} else {
			admin = {
				admin_id: null,
				name: '',
				email: '',
				password: '',
				phone: '',
				role: ''
			};
		}
		showForm = true;
	}

	function goToPage(p) {
		if (p >= 1 && p <= totalPages) {
			loadAdmins(p);
		}
	}

	function toggleForm() {
		showForm = !showForm;
	}
</script>

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách quản trị viên</h1>
		<div class="">
			<button class="btn-create flex flex-none gap-1" on:click={openForm('create')}>
				<UserPlus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>
	<hr />

	<div class="flex px-10 py-2">
		<div class="m-2 mx-5 flex h-23 w-40 flex-col rounded-3xl bg-gray-200 p-4">
			<h1>Tổng nhân viên</h1>
			<h1 class="flex items-center justify-center text-4xl font-bold">{total}</h1>
		</div>
		<div class="m-2 mx-5 flex h-23 w-40 flex-col items-center rounded-3xl bg-gray-200 p-4">
			<h1 class="">Admin</h1>
			<div class="flex items-center justify-center">
				<h1 class="text-4xl font-bold">
					{adminTotal}
				</h1>
			</div>
		</div>
		<div class="m-2 mx-5 flex h-23 w-40 flex-col items-center rounded-3xl bg-gray-200 p-4">
			<h1 class="">Staff</h1>
			<div class="flex items-center justify-center">
				<h1 class="text-4xl font-bold">
					{staffTotal}
				</h1>
			</div>
		</div>
	</div>

	<table class="max-w-full table-auto border-collapse border border-gray-300">
		<thead class="bg-gray-200">
			<tr class="">
				<th class="w-[50px] border p-2 text-left">ID</th>
				<th class="items-content border p-2 text-left">Họ tên</th>
				<th class="border p-2 text-left">Email</th>
				<th class="border p-2 text-left">Số điện thoại</th>
				<th class="border p-2 text-center">Chức vụ</th>
				<th class="border p-2 text-center">Hành động</th>
			</tr>
		</thead>
		<tbody>
			{#each admins as admin}
				<tr>
					<td class="w-[50px] border p-2">{admin.admin_id}</td>
					<td class="w-[300px] border p-2">{admin.name}</td>
					<td class="w-[400px] border p-2">{admin.email}</td>
					<td class="w-[150px] border p-2">{admin.phone}</td>
					<td class="w-[100px] border p-2">{admin.role}</td>
					<td class="w-[350px] border p-2">
						<div class="flex gap-2">
							<button
								on:click={() => openForm('view', admin)}
								class="btn-create flex items-center gap-1"
							>
								<Eye class="h-5  w-5" />
								Chi tiết</button
							>

							<button
								on:click={() => openForm('edit', admin)}
								class="btn-edit flex items-center gap-1"
							>
								<Pencil class="h-5  w-5" />
								Sửa</button
							>
							<button
								on:click={() => confirmDelete(admin.admin_id)}
								class="btn-delete flex items-center gap-1"
							>
								<Trash2 class="h-5  w-5" />
								Xoá</button
							>
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
	<div class="mt-4 flex items-center justify-center space-x-2">
		<button
			class="rounded bg-gray-300 px-3 py-1 disabled:opacity-50"
			on:click={() => goToPage(page - 1)}
			disabled={page === 1}
		>
			Trước
		</button>

		{#each Array(totalPages) as _, i}
			<button
				class="rounded px-3 py-1 {page === i + 1 ? 'bg-blue-600 text-white' : 'bg-gray-200'}"
				on:click={() => goToPage(i + 1)}
			>
				{i + 1}
			</button>
		{/each}

		<button
			class="rounded bg-gray-300 px-3 py-1 disabled:opacity-50"
			on:click={() => goToPage(page + 1)}
			disabled={page === totalPages}
		>
			Sau
		</button>
	</div>
</div>

{#if showForm}
	<!-- Background đen mờ -->
	<div class="fixed inset-0 z-40 bg-black/50"></div>

	<!-- Popup form -->
	<div
		class="fixed top-1/2 left-1/2 z-60 mx-auto w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-lg"
	>
		<h2 class="mb-4 text-2xl font-bold">{formTitle}</h2>
		<form on:submit={submitForm} class="space-y-4">
			<div>
				<label class="block font-medium"
					>Họ tên:
					<input
						type="text"
						bind:value={admin.name}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</label>
			</div>

			<div>
				<label class="block font-medium"
					>Email:
					<input
						type="email"
						bind:value={admin.email}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</label>
			</div>
			{#if formMode === 'create'}
				<div>
					<label class="block font-medium"
						>Mật khẩu:
						<input
							type="password"
							bind:value={admin.password}
							required
							disabled={formMode === 'view'}
							class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
						/>
					</label>
				</div>
			{/if}

			<div>
				<label class="block font-medium"
					>Số điện thoại:
					<input
						type="text"
						bind:value={admin.phone}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</label>
			</div>

			<div>
				<label class="block font-medium"
					>Chức vụ:
					<select
						bind:value={admin.role}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					>
						<option value="">-- Chọn --</option>
						<option value="Admin">Admin</option>
						<option value="Staff">Staff</option>
					</select>
				</label>
			</div>

			<div class="mt-6 flex justify-end space-x-3">
				{#if formMode === 'view'}
					<button
						type="button"
						on:click={toggleForm}
						class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
					>
						Đóng
					</button>
				{:else}
					<button
						type="button"
						on:click={toggleForm}
						class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
					>
						Huỷ
					</button>
					<button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700">
						{formMode === 'edit' ? 'Cập nhật' : 'Tạo'}
					</button>
				{/if}
			</div>
		</form>
	</div>
{/if}

{#if showConfirmDelete}
	<!-- Overlay -->
	<div class="fixed inset-0 z-40 bg-black/50"></div>

	<!-- Modal content -->
	<div
		class="fixed top-1/2 left-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-xl"
	>
		<h2 class="mb-4 text-xl font-bold">Xác nhận xoá</h2>
		<p>Bạn có chắc chắn muốn xoá người dùng <strong> {admin.name}</strong> không?</p>

		<div class="mt-6 flex justify-end gap-3">
			<button
				class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
				on:click={() => (showConfirmDelete = false)}
			>
				Hủy
			</button>
			<button
				class="rounded bg-red-600 px-4 py-2 text-white hover:bg-red-700"
				on:click={() => confirmAdminDelete(admin.admin_id)}
			>
				Xác nhận
			</button>
		</div>
	</div>
{/if}
