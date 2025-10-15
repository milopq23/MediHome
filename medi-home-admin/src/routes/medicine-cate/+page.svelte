<script>
	import { onMount } from 'svelte';
	import { Plus, ChevronUp, Pencil, Trash2 } from 'lucide-svelte';
	import { getAllCate, createCate, patchCate, deleteCate, upload } from '$lib/api/medicine-cate.js';

	let categories = [];
	let expandedIds = [];

	let category = {
		medicinecate_id: null,
		name: '',
		icon: '',
		parent_id: null
	};

	let file = null;
	let previewUrl = '';
	let formMode = 'create';
	let showForm = false;
	let showConfirmDelete = false;

	$: formTitle =
		formMode === 'create'
			? 'Tạo người dùng'
			: formMode === 'edit'
				? 'Chỉnh sửa người dùng'
				: 'Chi tiết người dùng';

	onMount(() => {
		loadCates();
	});

	async function loadCates() {
		try {
			const res = await getAllCate();
			categories = res;
			// categories = res;
		} catch (error) {
			console.error('Lỗi load danh mục:', error);
		}
	}

	async function onFileChange(event) {
		file = event.target.files[0];
		if (file) {
			// Tạo URL tạm thời để preview
			previewUrl = URL.createObjectURL(file);
		} else {
			previewUrl = '';
		}
	}

	async function submitForm(event) {
		event.preventDefault();
		try {
			if (formMode === 'create') {
				await upload(file);
				await createCate(category);
			} else if (formMode === 'edit') {
				await upload(file);
				await patchCate(category.medicinecate_id, category);
			}
			await loadCates();
			toggleForm();
		} catch (error) {
			console.error('Lỗi submit form:', error);
		}
	}

	async function confirmCateDelete(id) {
		try {
			await deleteCate(id);
			showConfirmDelete = false;
			await loadCates();
		} catch (err) {
			console.error('Xoá thất bại:', err);
		}
	}

	function confirmDelete(id) {
		showConfirmDelete = true;
		category.medicinecate_id = id;
	}

	function toggleForm() {
		showForm = !showForm;
	}

	function openForm(mode, data = null) {
		formMode = mode;
		if (data) {
			category = { ...data };
		} else {
			category = {
				medicinecate_id: null,
				name: '',
				icon: '',
				parent_id: null
			};
		}
		showForm = true;
	}

	function toggleList(id) {
		if (expandedIds.includes(id)) {
			expandedIds = expandedIds.filter((i) => i !== id);
		} else {
			expandedIds = [...expandedIds, id];
		}
	}
</script>

<div>
	<div class="flex items-center">
		<h1 class="flex-1 p-5 text-4xl font-bold">Danh sách danh mục thuốc</h1>
		<div>
			<button class="btn-create flex items-center gap-1" on:click={() => openForm('create')}>
				<Plus class="h-5 w-5" />
				Thêm
			</button>
		</div>
	</div>

	<hr class="m-4" />

	<ul class="space-y-2">
		{#each categories as parent}
			<li class="rounded border">
				<button
					type="button"
					class="flex w-full items-center justify-between bg-gray-100 p-3 hover:bg-gray-200"
					on:click={() => toggleList(parent.medicinecate_id)}
				>
					<span>{parent.name}</span>
					<span
						class="transform transition-transform"
						class:rotate-180={expandedIds.includes(parent.medicinecate_id)}
					>
						<ChevronUp class="h-5 w-5" />
					</span>
				</button>

				{#if expandedIds.includes(parent.medicinecate_id)}
					<ul class="bg-white py-2 pl-6">
						{#each parent.children as child}
							<li class="flex items-center justify-between py-1 hover:bg-gray-50">
								{child.name}
								<div class="flex gap-1">
									<button
										class="btn-edit flex items-center gap-1"
										on:click|stopPropagation={() => openForm('edit', child)}
									>
										<Pencil class="h-4 w-4" />
										Sửa
									</button>
									<button
										on:click={() => confirmDelete(child.medicinecate_id)}
										class="btn-delete flex items-center gap-1"
									>
										<Trash2 class="h-5 w-5" />
										Xoá
									</button>
								</div>
							</li>
						{/each}
					</ul>
				{/if}
			</li>
		{/each}
	</ul>
</div>

{#if showForm}
	<!-- Overlay mờ -->
	<div class="fixed inset-0 z-40 bg-black/50"></div>

	<!-- Popup form -->
	<div
		class="fixed top-1/2 left-1/2 z-50 w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-lg"
	>
		<h2 class="mb-4 text-2xl font-bold">
			{formTitle}
		</h2>
		<form on:submit={submitForm} class="space-y-4">
			<!-- Tên danh mục -->
			<div>
				<label class="block font-medium"
					>Tên danh mục:
					<input
						type="text"
						bind:value={category.name}
						required
						disabled={formMode === 'view'}
						class="mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					/>
				</label>
			</div>

			<!-- Chọn danh mục cha -->
			<div>
				<label class="block font-medium"
					>Danh mục cha:
					<select
						bind:value={category.parent_id}
						disabled={formMode === 'view'}
						class="mt-1 block w-full rounded border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
					>
						<option value={null}>-- Không có (Danh mục cha) --</option>
						{#each categories as cate}
							<option value={cate.medicinecate_id}>{cate.name}</option>
						{/each}
					</select>
				</label>
			</div>

			<div>
				<label class="block font-medium text-gray-900 dark:text-white"
					>Icon:
					<div class="flex justify-center">
						{#if previewUrl}
							<img
								src={previewUrl}
								alt="Ảnh preview"
								class="max-h-[200px] max-w-[200px] cursor-pointer rounded object-contain"
							/>
						{:else}
							<div
								class="flex h-[200px] w-[200px] items-center justify-center rounded-xl border-solid bg-gray-300"
							>
								<Plus class="h-5 w-5" />
							</div>
						{/if}
					</div>
					<input class="hidden" type="file" accept="image/*" on:change={onFileChange} />
				</label>
			</div>

			<!-- Nút hành động -->
			<div class="mt-6 flex justify-end gap-2">
				<button
					type="button"
					on:click={toggleForm}
					class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
				>
					{formMode === 'view' ? 'Đóng' : 'Huỷ'}
				</button>
				{#if formMode !== 'view'}
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
		<p>Bạn có chắc chắn muốn xoá danh mục không?</p>
		<div class="mt-6 flex justify-end gap-3">
			<button
				class="rounded bg-gray-300 px-4 py-2 hover:bg-gray-400"
				on:click={() => (showConfirmDelete = false)}
			>
				Hủy
			</button>
			<button
				class="rounded bg-red-600 px-4 py-2 text-white hover:bg-red-700"
				on:click={() => confirmCateDelete(category.medicinecate_id)}
			>
				Xác nhận
			</button>
		</div>
	</div>
{/if}
