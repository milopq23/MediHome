<script>
	import ToolBar from '$lib/components/ToolBar.svelte';
	import { UserPlus, Pencil, Trash2, Eye } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { fetchUsers } from '$lib/api/user.js';



	let users = [];
	let showForm = false;
	// let user = {
	// 	name: '',
	// 	email: '',
	// 	password: '',
	// 	phone: '',
	// 	gender: '',
	// 	is_verified: false
	// };

	function toggleForm() {
		showForm = !showForm;
	}

	onMount(async () => {
		try {
			// const {users, page, pageSize, total} = await fetchUsers();

            const result = await fetchUsers();
            users = result.users;
            total = result.total;

		} catch (e) {
			error = e.message;
		}
	});
</script>

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách quản trị viên</h1>
		<div class="">
			<button class="btn-create flex flex-none gap-1" on:click={toggleForm}>
				<UserPlus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>
	<hr />

	<div class="flex">
		<div class="m-2 flex h-20 w-40 flex-col rounded-3xl bg-gray-200 p-4">
			<div>Total User</div>
			<h1 class="text-2xl font-bold">{total}</h1>
		</div>
		<div class="m-2 flex h-20 w-40 flex-col rounded-3xl bg-gray-200 p-4">
			<div class="">
				<div>Active User</div>
				<h1 class="text-2xl font-bold">24</h1>
			</div>
		</div>
	</div>
	{#if showForm}
		<!-- Background đen mờ -->
		<div
			class="bg-opacity-50 fixed inset-0 z-50 flex items-center justify-center bg-gray-100"
			on:click={toggleForm}
		></div>

		<!-- Popup form -->
		<div
			class="fixed top-1/2 left-1/2 z-60 mx-auto w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-lg"
			on:click|stopPropagation
		>
			<h2 class="mb-4 text-2xl font-bold">Tạo người dùng</h2>
			<form on:submit|preventDefault={createUser} class="space-y-4">
				<div>
					<label class="block font-medium">Họ tên:</label>
					<input
						type="text"
						bind:value={user.name}
						required
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block font-medium">Email:</label>
					<input
						type="email"
						bind:value={user.email}
						required
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block font-medium">Mật khẩu:</label>
					<input
						type="password"
						bind:value={user.password}
						required
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block font-medium">Số điện thoại:</label>
					<input
						type="text"
						bind:value={user.phone}
						required
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</div>

				<div>
					<label class="block font-medium">Giới tính:</label>
					<select
						bind:value={user.gender}
						required
						class="focus:ring-opacity-50 mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					>
						<option value="">-- Chọn --</option>
						<option value="male">Nam</option>
						<option value="female">Nữ</option>
					</select>
				</div>

				<div class="flex items-center space-x-2">
					<input
						type="checkbox"
						bind:checked={user.is_verified}
						id="verified"
						class="rounded border-gray-300 text-blue-600 shadow-sm focus:ring-blue-500"
					/>
					<label for="verified" class="font-medium select-none">Đã xác thực</label>
				</div>

				<div class="mt-6 flex justify-end space-x-3">
					<button
						type="button"
						on:click={toggleForm}
						class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
					>
						Hủy
					</button>
					<button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700">
						Tạo
					</button>
				</div>
			</form>
		</div>
	{/if}

	<table class=" border-seperate w-full table-auto border border-gray-300">
		<thead class="bg-gray-200">
			<tr class="">
				<th class="w-[50px] border p-2 text-left">#</th>
				<th class="items-content border p-2 text-left">Họ tên</th>
				<th class="border p-2 text-left">Email</th>
				<th class="border p-2 text-left">Số điện thoại</th>
				<th class="border p-2 text-center">Điểm</th>
				<th class="border p-2 text-center">Trạng thái</th>
				<th class="w-[180px] border p-2 text-center">Hành động</th>
			</tr>
		</thead>
		<tbody>
			{#each users as user}
				<tr>
					<td class="border p-2">{user.user_id}</td>
					<td class="border p-2">{user.name}</td>
					<td class="border p-2">{user.email}</td>
					<td class="border p-2">{user.phone}</td>
					<td class="border p-2">{user.point}</td>
					<td class="border p-2">{user.is_verified}</td>
					<td class="border p-2">
						<div class="flex gap-2">
							<button class="btn-create flex items-center gap-1">
								<Eye class="h-5  w-5" />
								Sửa</button
							>

							<button class="btn-edit flex items-center gap-1">
								<Pencil class="h-5  w-5" />
								Sửa</button
							>
							<button class="btn-delete flex items-center gap-1">
								<Trash2 class="h-5  w-5" />
								Xoá</button
							>
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
