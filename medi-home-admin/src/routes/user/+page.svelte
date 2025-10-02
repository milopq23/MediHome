<script>
	import ToolBar from '$lib/components/ToolBar.svelte';
	import { UserPlus, Pencil, Trash2, Eye } from 'lucide-svelte';
	import { onMount } from 'svelte';
	// import { fetchUser,getUser } from '$lib/api/user.js';

	import { fetchUsersPaginated } from '$lib/api/user.js';

  let users = [];
  let errorMsg = '';

  onMount(async () => {
    try {
      const res = await fetchUsersPaginated(1, 10);
      users = res.data;
    } catch (err) {
      errorMsg = err.message;
      console.error('Lỗi khi gọi API:', err);
    }
  });


	// let users = [
	// 	{
	// 		id: 1,
	// 		name: 'Nguyễn Văn A',
	// 		email: 'a@gmail.com',
	// 		phone: '08125610580',
	// 		point: '84',
	// 		status: 'active'
	// 	},
	// 	{
	// 		id: 2,
	// 		name: 'Trần Thị B',
	// 		email: 'b@gmail.com',
	// 		phone: '08125610580',
	// 		point: '84',
	// 		status: 'active'
	// 	},
	// 	{
	// 		id: 3,
	// 		name: 'Lê Văn C',
	// 		email: 'c@gmail.com',
	// 		phone: '08125610580',
	// 		point: '84',
	// 		status: 'active'
	// 	},
	// 	{
	// 		id: 4,
	// 		name: 'Nguyễn Huỳnh Phú Quý',
	// 		email: 'nguyenhuynhphuquytv2012@gmail.com',
	// 		phone: '08125610580',
	// 		point: '84',
	// 		status: 'active'
	// 	}
	// ];
</script>

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách người dùng</h1>
		<div class="">
			<button class="btn-create flex flex-none gap-1">
				<UserPlus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>
	<hr />

	<div class="flex">
		<div class="m-2 flex h-20 w-40 flex-col rounded-3xl bg-gray-200 p-4">
			<div>Total User</div>
			<h1 class="text-2xl font-bold">24</h1>
		</div>
		<div class="m-2 flex h-20 w-40 flex-col rounded-3xl bg-gray-200 p-4">
			<div class="">
				<div>Active User</div>
				<h1 class="text-2xl font-bold">24</h1>
			</div>
		</div>
	</div>


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
