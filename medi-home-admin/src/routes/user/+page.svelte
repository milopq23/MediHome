<script>
	import ToolBar from '$lib/components/ToolBar.svelte';
	import { UserPlus, Pencil, Trash2, Eye } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { loadUsers, addUser, editUser, deleteUser } from '$lib/api/user.js';
	let users = [];
	let total = 0;
	let page = 1;
	let pageSize = 10;
	let userActive = 1;
	let totalPages = 1;

	let showForm = false;
	let showConfirmDelete = false;

	let user = {
		user_id: null,
		name: '',
		email: '',
		password: '',
		phone: '',
		gender: '',
		is_verified: 'false'
	};

	let formMode = 'create';

	$: formTitle =
		formMode === 'create'
			? 'Tạo người dùng'
			: formMode === 'edit'
				? 'Chỉnh sửa người dùng'
				: 'Chi tiết người dùng';

	async function loadUserPage(currentPage = 1) {
		try {
			const result = await loadUsers(currentPage, pageSize);
			users = result.users;
			page = result.page;
			pageSize = result.pageSize;
			total = result.total;
			totalPages = Math.ceil(total / pageSize);

			userActive = result.activeUser;
		} catch (err) {
			console.error(err);
		}
	}

	async function submitForm(event) {
		event.preventDefault();
		try {
			if (formMode === 'create') {
				const result = await addUser(user);
				console.log('Tạo thành công:', result);
			} else if (formMode === 'edit') {
				const result = await editUser(user.user_id, user);
				console.log('Cập nhật thành công:', result);
			}
			toggleForm();
		} catch (error) {
			console.error(error);
		}
	}

	async function confirmUserDelete(id) {
		try {
			await deleteUser(id);
			showConfirmDelete = false;
			await loadUserPage();
		} catch (err) {
			console.error('Xoá thất bại:', err);
		}
	}

	onMount(() => {
		loadUserPage();
	});

	function confirmDelete(id) {
		showConfirmDelete = true;
		user.user_id = id;
	}

	function openForm(mode, userData) {
		formMode = mode;
		if (userData) {
			user = { ...userData };
			console.log(user);
		} else {
			user = {
				user_id: null,
				name: '',
				email: '',
				password: '',
				phone: '',
				gender: '',
				is_verified: 'Chưa xác thực'
			};
		}
		showForm = true;
	}

	function goToPage(p) {
		if (p >= 1 && p <= totalPages) {
			loadUserPage(p);
		}
	}

	function toggleForm() {
		showForm = !showForm;
	}
</script>

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách người dùng</h1>
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
			<h1>Tổng người dùng</h1>
			<h1 class="flex items-center justify-center text-4xl font-bold">{total}</h1>
		</div>
		<div class="m-2 mx-5 flex h-23 w-40 flex-col items-center rounded-3xl bg-gray-200 p-4">
			<h1 class="">Đã kích hoạt</h1>
			<div class="flex items-center justify-center">
				<h1 class="text-4xl font-bold">
					{userActive}<span class="text-xl font-normal">/{total}</span>
				</h1>
			</div>
		</div>
	</div>

	<table class="	 max-w-full table-auto border-collapse border border-gray-300">
		<thead class="bg-gray-200">
			<tr class="">
				<th class="w-[50px] border p-2 text-left">ID</th>
				<th class="items-content border p-2 text-left">Họ tên</th>
				<th class="border p-2 text-left">Email</th>
				<th class="border p-2 text-left">Số điện thoại</th>
				<th class="border p-2 text-center">Điểm</th>
				<th class="border p-2 text-center">Trạng thái</th>
				<th class="border p-2 text-center">Hành động</th>
			</tr>
		</thead>
		<tbody>
			{#each users as user}
				<tr>
					<td class="w-[50px] border p-2">{user.user_id}</td>
					<td class="w-[300px] border p-2">{user.name}</td>
					<td class="w-[400px] border p-2">{user.email}</td>
					<td class="w-[150px] border p-2">{user.phone}</td>
					<td class="w-[100px] border p-2">{user.gender}</td>
					<td class="w-[100px] border p-2">{user.is_verified ? 'Đã xác thực' : 'Chưa xác thực'}</td>
					<td class="w-[350px] border p-2">
						<div class="flex gap-2">
							<button
								on:click={() => openForm('view', user)}
								class="btn-create flex items-center gap-1"
							>
								<Eye class="h-5  w-5" />
								Chi tiết</button
							>

							<button
								on:click={() => openForm('edit', user)}
								class="btn-edit flex items-center gap-1"
							>
								<Pencil class="h-5  w-5" />
								Sửa</button
							>
							<button
								on:click={() => confirmDelete(user.user_id)}
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
						bind:value={user.name}
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
						bind:value={user.email}
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
							bind:value={user.password}
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
						bind:value={user.phone}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</label>
			</div>

			<div>
				<label class="block font-medium"
					>Giới tính:
					<select
						bind:value={user.gender}
						required
						disabled={formMode === 'view'}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					>
						<option value="">-- Chọn --</option>
						<option value="Nam">Nam</option>
						<option value="Nữ">Nữ</option>
					</select>
				</label>
			</div>

			<div class="flex items-center space-x-2">
				<label for="verified" class="font-medium select-none">
					Kích hoạt:
					<select
						bind:value={user.is_verified}
						required
						disabled={formMode === 'view'}
						on:change={() => console.log('user.is_verified:', user.is_verified)}
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					>
						<!-- <option value="">-- Chọn --</option> -->
						<option value="true">Đã xác thực</option>
						<option value="false">Chưa xác thực</option>
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
		<p>Bạn có chắc chắn muốn xoá người dùng <strong> {user.name}</strong> không?</p>

		<div class="mt-6 flex justify-end gap-3">
			<button
				class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
				on:click={() => (showConfirmDelete = false)}
			>
				Hủy
			</button>
			<button
				class="rounded bg-red-600 px-4 py-2 text-white hover:bg-red-700"
				on:click={() => confirmUserDelete(user.user_id)}
			>
				Xác nhận
			</button>
		</div>
	</div>
{/if}
