<script>
	import { onMount } from 'svelte';
	import { Plus, ChevronUp, Pencil, Trash2 } from 'lucide-svelte';
	import {getAllCate} from '$lib/api/medicine-cate.js'

	let categories = [];

	let expandedIds = [];

	// Trạng thái mở rộng từng danh mục cha (theo ID)
	let expanded = new Set();


	async function loadCates() {
		try {
			const res = await getAllCate();
			categories = res;
			
		} catch (error) {
			console.error(error)
		}
		
	}
	onMount( () => {
		loadCates();
		// const res = await fetch('http://localhost:8000/api/medicine-cate/'); // URL thực tế của bạn
		// categories = await res.json();
	});

	function toggle(id) {
		if (expandedIds.includes(id)) {
			expandedIds = expandedIds.filter((i) => i !== id);
		} else {
			expandedIds = [...expandedIds, id];
		}
	}
</script>

<!-- <style>
	.chevron {
		transition: transform 0.3s ease;
	}
	.rotate-90 {
		transform: rotate(90deg);
	}
</style> -->

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách người dùng</h1>
		<div class="">
			<button class="btn-create flex flex-none gap-1" on:click={openForm('create')}>
				<Plus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>
	<hr class="m-4" />

	<ul class="space-y-2">
		{#each categories as parent}
			<li class="rounded border">
				<!-- Toggle cha -->
				<!-- <div
					class="flex cursor-pointer items-center justify-between bg-gray-100 p-3 hover:bg-gray-200"
					on:click={() => toggle(parent.MedicineCateID)}
				>
					<span>{parent.name}</span>
					<span class="chevron {expanded.has(parent.MedicineCateID) ? 'rotate-180' : ''}">
						<ChevronUp />
					</span>
				</div> -->

				<button
					type="button"
					class="flex w-full cursor-pointer items-center justify-between bg-gray-100 p-3 text-left hover:bg-gray-200"
					on:click={() => toggle(parent.MedicineCateID)}
				>
					<span>{parent.name}</span>
					<span class="chevron {expanded.has(parent.MedicineCateID) ? 'rotate-180' : ''}">
						<ChevronUp />
					</span>
				</button>

				<!-- Children sổ ra -->
				{#if expandedIds.includes(parent.MedicineCateID)}
					<ul class="bg-white py-2 pl-6">
						{#each parent.children as child}
							<li class="flex items-center justify-between border-b p-10 py-1 hover:bg-gray-50">
								<div>
									{child.name}
									<button class="btn-edit flex items-center gap-1" on:click|stopPropagation>
										<Pencil class="h-4 w-4" />
										Sửa
									</button>
									<button
										on:click={() => confirmDelete(user.user_id)}
										class="btn-delete flex items-center gap-1"
									>
										<Trash2 class="h-5  w-5" />
										Xoá</button
									>
								</div>
							</li>
						{/each}
					</ul>
				{/if}
			</li>
		{/each}
	</ul>
</div>
