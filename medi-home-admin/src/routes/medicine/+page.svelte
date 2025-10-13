<script>
	import { Pencil, Trash2, MoreVertical, Plus } from 'lucide-svelte';
	import { loadMedicines, addMedicine, patchMedicine } from '$lib/api/medicine.js';
	import { onMount } from 'svelte';

	let search = '';
	let page = 1;
	let pageSize = 10;
	let total = 0;
	let totalPages = 1;

	let medicines = []
	let medicine = {
		code: '',
		name: '',
		thumbnail: '',
		category_name: '',
		dosageform_name: '',	
		prescription: ''
	}

	// $: filtered = products.filter((p) => p.name.toLowerCase().includes(search.toLowerCase()));
	// $: total = filtered.length;
	// $: totalPages = Math.ceil(total / pageSize);
	// $: paginated = filtered.slice((page - 1) * pageSize, page * pageSize);

	async function loadMedicinePage(currentPage = 1) {
		try {
			const result = await loadMedicines(currentPage, pageSize);
			console.log(result);
			medicines = result.medicines;
			page = result.page;
			pageSize = result.pageSize;
			total = result.total;
			totalPages = Math.ceil(total / pageSize);
		} catch (error) {
			console.log('load medicinee', error);
		}
	}

	onMount(() => {
		loadMedicinePage();
	});

	function goToPage(p) {
		if (p >= 1 && p <= totalPages) page = p;
	}
</script>

<div class="rounded-md bg-white p-4 shadow-sm">
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách người dùng</h1>
		<div class="">
			<button class="btn-create flex flex-none gap-1" on:click={openForm('create')}>
				<Plus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>

	<!-- Search bar -->
	<div class="mb-4 flex items-center">
		<input
			type="text"
			class="w-full max-w-xs rounded-md border px-3 py-2 text-sm"
			placeholder="Search for items"
			bind:value={search}
		/>
	</div>

	<!-- Table -->
	<table class="min-w-full table-auto text-sm">
		<thead class="border-b text-left text-gray-500">
			<tr>
				<th class="py-2"><input type="checkbox" /></th>
				<th class="py-2">Code</th>
				<th class="py-2">Tên thuốc</th>
				<th class="py-2">Hình</th>
				<th class="py-2">Danh mục</th>
				<th class="py-2">Dạng bào chế</th>
				<th class="py-2">Kê đơn</th>
				<th class="py-2">Hành động</th>
			</tr>
		</thead>
		<tbody class="text-gray-800">
			{#each medicines as medicine}
				<tr class="border-b hover:bg-gray-50">
					<td class="py-3"><input type="checkbox" /></td>
					<td class="py-3">{medicine.code}</td>
					<td class="py-3">{medicine.name}</td>
					<td class="py-3">{medicine.thumbnail}</td>
					<td class="py-3">{medicine.category_name}</td>
					<td class="py-3">{medicine.dosageform_name}</td>
					<td class="py-3">{medicine.prescription ? 'Có':'Không'}</td>

					<!-- <td class="py-3">
						<span class={`px-2 py-1 rounded text-xs font-medium ${badgeColor(item.availability)}`}>
							{item.availability}
						</span>
					</td> -->
					<td class="py-3">
						<div class="flex items-center gap-2 text-gray-600">
							<button><Pencil class="h-4 w-4" /></button>
							<button><Trash2 class="h-4 w-4" /></button>
							<button><MoreVertical class="h-4 w-4" /></button>
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>

	<!-- Pagination -->
	<div class="mt-4 flex items-center justify-between text-sm text-gray-500">
		<div>
			Showing {(page - 1) * pageSize + 1} to {Math.min(page * pageSize, total)} of {total} products
		</div>
		<div class="flex items-center gap-1">
			<button
				class="rounded px-2 py-1 disabled:opacity-50"
				on:click={() => goToPage(page - 1)}
				disabled={page === 1}
			>
				«
			</button>
			{#each Array(totalPages) as _, i}
				<button
					class={`rounded px-3 py-1 ${page === i + 1 ? 'bg-indigo-100 font-semibold text-indigo-600' : 'hover:bg-gray-100'}`}
					on:click={() => goToPage(i + 1)}
				>
					{i + 1}
				</button>
			{/each}
			<button
				class="rounded px-2 py-1 disabled:opacity-50"
				on:click={() => goToPage(page + 1)}
				disabled={page === totalPages}
			>
				»
			</button>
		</div>
	</div>
</div>
