<script>
	import { pageTitle } from '$lib/stores/store.js';
	import { toasts } from '$lib/stores/toastMessage.js';
	import { Plus } from 'lucide-svelte';
	let title = 'Chi tiết thuốc';

	pageTitle.set(title);

	export let data;
	let { medicine, medicineCate, selectedParent, selectedChild, dosageForm } = data;

	let previewUrl = '';
	let previewUrls = [];

	let categories = medicineCate;

	$: childCategories = selectedParent
		? medicineCate.find((c) => c.medicinecate_id == selectedParent)?.children || []
		: [];

	async function addMedicine() {
		try {
			const res = await fetch('/admin/medicine', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(medicine)
			});
		} catch (error) {
			console.error('Lỗi khi thêm thuốc:', error);
			responseMessage.set('Lỗi khi thêm thuốc!');
		}
	}

	// async function onFileChange(event) {
	// 	file = event.target.files[0];
	// 	if (file) {
	// 		// Tạo URL tạm thời để preview
	// 		previewUrl = URL.createObjectURL(file);
	// 	} else {
	// 		previewUrl = '';
	// 	}
	// }

	function onFileChange(event) {
		const files = event.target.files;
		for (const file of files) {
			const reader = new FileReader();
			reader.onload = (e) => {
				previewUrls = [...previewUrls, e.target.result]; // thêm URL mới vào mảng
			};
			reader.readAsDataURL(file);
		}
	}
</script>

<div class="flex items-center justify-center">
	<div class="max-h-full w-full overflow-y-auto rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
		<div class="hidden md:flex md:justify-center md:py-10">
			<h1 class="text-5xl font-extrabold">{title}</h1>
		</div>

		<form>
			<div class="mb-6 grid gap-6 md:grid-cols-3">
				<div class="col-span-1">
					<label for="code" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Code</label
					>
					<input
						type="text"
						class="border-gray-300F block w-full rounded-lg border bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="MED001"
						bind:value={medicine.code}
						required
					/>
				</div>
				<div class="md:col-span-2">
					<label
						for="medicine_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Tên thuốc</label
					>
					<input
						type="text"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Tên thuốc"
						bind:value={medicine.name}
						required
					/>
				</div>
				<div class="">
					<label for="category" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Danh mục cha</label
					>
					<select
						id="category"
						bind:value={selectedParent}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						required
					>
						<option value="" disabled selected>Chọn danh mục</option>
						{#each categories as parent}
							<option value={parent.medicinecate_id}>{parent.name}</option>
						{/each}
					</select>
				</div>
				<div class="">
					<label for="category" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Danh mục con</label
					>
					<select
						id="childCategory"
						bind:value={selectedChild}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900"
						required
						disabled={!selectedParent}
					>
						<option value="" disabled selected>Chọn danh mục</option>
						{#if childCategories}
							{#each childCategories as child}
								<option value={child.medicinecate_id}>{child.name}</option>
							{/each}
						{/if}
					</select>
					{#if childCategories.length > 0}
						<ul class="mt-2">
							{#each selectedParent.children as child}
								<li>{child.name}</li>
							{/each}
						</ul>
					{/if}
				</div>
				<div class="">
					<label
						for="dosageForm"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Dạng bào chế</label
					>
					<select
						id="dosage"
						bind:value={medicine.dosageform_id}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900"
						required
					>
						<option class="" value="" disabled selected>Chọn danh mục</option>

						{#each dosageForm as dosage}
							<option class="" value={dosage.dosageform_id}>{dosage.name}</option>
						{/each}
					</select>
				</div>

				<div class="">
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Đơn vị(viên/vỉ)</label
					>
					<input
						type="number"
						bind:value={medicine.unitstrip}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder=""
						required
					/>
				</div>
				<div class="">
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Đơn vị(vỉ/hộp)</label
					>
					<input
						type="number"
						bind:value={medicine.unitbox}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder=""
						required
					/>
				</div>

				<div class="">
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Đóng gói</label
					>
					<input
						type="url"
						bind:value={medicine.package}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Hộp 10 vỉ x 10 viên"
						required
					/>
				</div>

				<div class="">
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Thuốc kê đơn</label
					>
					<select
						bind:value={medicine.prescription}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900"
						required
					>
						<option value="" disabled>Chọn danh mục</option>
						<option value={true} selected>Kê đơn </option>
						<option value={false} selected>Không kê đơn</option>
					</select>
				</div>

				<div class="">
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Nhà sản xuất</label
					>
					<input
						bind:value={medicine.manufacturer}
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Công ty ABC"
						required
					/>
				</div>
			</div>

			<div>
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Ảnh chính:
				</label>
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
					<input class="hidden" type="file" accept="image/*" on:change={onFileChange} />
				</div>
			</div>
			<div>
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Ảnh phụ:
				</label>
				<label class="flex cursor-pointer flex-wrap justify-center gap-2">
					{#if previewUrls.length > 0}
						{#each previewUrls as url, index}
							<img
								src={url}
								alt="Ảnh preview"
								class="max-h-[200px] max-w-[200px] rounded object-contain"
							/>
						{/each}
					{:else}
						<div
							class="flex h-[150px] w-[150px] items-center justify-center rounded-xl border-solid bg-gray-300"
						>
							<Plus class="h-5 w-5" />
						</div>
					{/if}
					<input class="hidden" type="file" multiple accept="image/*" on:change={onFileChange} />
				</label>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Cách dùng</label
				>
				<textarea
					type="text"
					bind:value={medicine.usage}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Chỉ định</label
				>
				<textarea
					type="text"
					bind:value={medicine.indication}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Tác dụng phụ</label
				>
				<textarea
					type="text"
					bind:value={medicine.adverse}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Chống chỉ định</label
				>
				<textarea
					type="text"
					bind:value={medicine.contraindication}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Thận trọng</label
				>
				<textarea
					type="text"
					bind:value={medicine.precaution}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Ảnh hưởng của thuốc lên khả năng lái xe và vận hành máy móc</label
				>
				<textarea
					type="text"
					bind:value={medicine.ability}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Phụ nữ có thai hoặc đang cho con bú</label
				>
				<textarea
					type="text"
					bind:value={medicine.pregnancy}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Tương tác thuốc
				</label>
				<textarea
					type="text"
					bind:value={medicine.drug_interaction}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Bảo quản</label
				>
				<textarea
					type="text"
					bind:value={medicine.storage}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6">
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Lưu ý</label
				>
				<textarea
					type="text"
					bind:value={medicine.note}
					class="block h-30 w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="Uống mỗi lần 2 viên, ngày 3 lần"
					required
				></textarea>
			</div>
			<div class="mb-6 flex items-start">
				<div class="flex h-5 items-center">
					<input
						id="remember"
						type="checkbox"
						value=""
						class="h-4 w-4 rounded-sm border border-gray-300 bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600"
						required
					/>
				</div>
				<label for="remember" class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
					>I agree with the <a href="" class="text-blue-600 hover:underline dark:text-blue-500"
						>terms and conditions</a
					>.</label
				>
			</div>
			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
