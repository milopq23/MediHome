<script>
	import {
		DetailMedicine,
		GetDosage,
		GetMedicineCate,
		UpdateMedicine,
		UploadMedicine,
		UploadMultiMedicine
	} from '$lib/api/medicine.js';
	import { pageTitle } from '$lib/stores/store.js';
	import { toasts } from '$lib/stores/toastMessage.js';
	import { Plus } from 'lucide-svelte';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	let title = 'Chi tiết thuốc';

	let id = null;
	$: medicine_id = parseInt($page.params.id, 10);

	let previewUrl = '';
	let previewUrls = [];
	let medicine = {};
	let categories = {};
	let dosageForm = [];
	let selectedParent = null;
	let selectedChild = null;

	let file = null;
	let subFiles = [];

	$: childCategories = selectedParent
		? categories.find((c) => c.medicinecate_id === selectedParent)?.children || []
		: [];

	async function detailMedicine(medicine_id) {
		const data = await DetailMedicine(medicine_id);
		medicine = data || {};
		previewUrl = medicine.thumbnail || '';
		await loadCategorySelection();
	}

	async function getDosage() {
		const data = await GetDosage();
		dosageForm = data;
	}

	async function getMedicineCate() {
		const data = await GetMedicineCate();
		categories = data || [];
		await loadCategorySelection();
	}

	async function uploadFile(file) {
		if (!file) return '';
		const res = await UploadMedicine(file);
		console.log('upload url', res);
		return res;
	}

	export async function uploadMultiImages(medicine_id, files) {
		try {
			// Upload file lên server và lấy URLs
			const res = await UploadMultiMedicine(files);

			if (!res || !res.urls) {
				throw new Error('Upload thất bại hoặc không có urls trả về');
			}

			const urls = res.urls;

			// Gửi URLs lên backend gắn vào medicine
			const result = await AddImage(medicine_id, urls);

			return result;
		} catch (error) {
			console.error('Lỗi upload nhiều ảnh:', error);
			throw error;
		}
	}

	// async function uploadSubFiles(files) {
	// 	const urls = [];
	// 	for (const file of files) {
	// 		const url = await uploadFile(file);
	// 		if (url) urls.push(url);
	// 	}
	// 	return urls;
	// }

	function onMainFileChange(event) {
		const selectedFile = event.target.files[0];
		if (selectedFile) {
			file = selectedFile;
			previewUrl = URL.createObjectURL(selectedFile); // preview ảnh mới
		} else {
			file = null;
			previewUrl = medicine.thumbnail || '';
		}
	}

	function onSubFileChange(event) {
		const files = Array.from(event.target.files);
		subFiles = files;
		previewUrls = files.map((f) => URL.createObjectURL(f));
	}

	async function loadCategorySelection() {
		if (!medicine.medicinecate_id || categories.length === 0) return;

		for (const parent of categories) {
			const child = parent.children?.find((c) => c.medicinecate_id === medicine.medicinecate_id);
			console.log(child);
			if (child) {
				selectedParent = parent.medicinecate_id;
				console.log(selectedParent);
				selectedChild = child.medicinecate_id;
				break;
			}
		}
	}

	async function onSubmit() {
		try {
			let thumbnailUrl = medicine.thumbnail;
			if (file) {
				thumbnailUrl = await uploadFile(file);
				console.log(thumbnailUrl);
			}

			// 2. Upload ảnh phụ (nếu có)

			// let subImages = medicine.sub_images || [];
			// if (subFiles.length > 0) {
			// 	const uploadedUrls = await uploadSubFiles(subFiles);
			// 	subImages = [...subImages, ...uploadedUrls];
			// }

			// 3. Dữ liệu gửi đi
			const data = {
				...medicine,
				thumbnail: thumbnailUrl,
				// sub_images: subImages,
				medicinecate_id: selectedChild || selectedParent
			};

			const res = await UpdateMedicine(medicine_id, data);

			if (res?.status === 'success') {
				toasts.success('Cập nhật thành công!');

				medicine = { ...medicine, thumbnail: thumbnailUrl };
				previewUrl = thumbnailUrl;
				file = null;
			} else {
				toasts.error(res?.message || 'Cập nhật thất bại!');
			}
		} catch (err) {
			console.error(err);
			toasts.error('Lỗi khi cập nhật!');
		}
	}

	onMount(async () => {
		await getMedicineCate();
		await getDosage();
		if (medicine_id) {
			await detailMedicine(medicine_id);
		}
	});
</script>

<div class="flex items-center justify-center">
	<div class="max-h-full w-full overflow-y-auto rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
		<div class="hidden md:flex md:justify-center md:py-10">
			<h1 class="text-5xl font-extrabold">{title}</h1>
		</div>

		<form on:submit|preventDefault={onSubmit}>
			<div class="mb-6 grid gap-6 md:grid-cols-3">
				<div class="col-span-1">
					<label for="code" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Code</label
					>
					<input
						type="text"
						class="border-gray-300F block w-full rounded-lg border bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
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
				<div>
					<label for="category" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Danh mục cha</label
					>
					<select
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
						bind:value={selectedParent}
					>
						<option value="" disabled selected>-- Chọn danh mục --</option>
						{#each categories as parent}
							<option value={parent.medicinecate_id}>{parent.name}</option>
						{/each}
					</select>
				</div>

				<!-- Danh mục con -->
				<div>
					<label for="category" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Danh mục con</label
					>
					<select
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
						bind:value={selectedChild}
						disabled={!selectedParent}
					>
						<option value="" disabled selected>-- Chọn danh mục con --</option>
						{#each childCategories as child}
							<option value={child.medicinecate_id}>{child.name}</option>
						{/each}
					</select>
				</div>
				<div class="">
					<label
						for="dosageForm"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Dạng bào chế</label
					>
					<select
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
				<label>
					<div class="flex justify-center">
						{#if previewUrl}
							<div class="relative">
								<img
									src={previewUrl}
									alt="Preview"
									class="max-h-[200px] max-w-[200px] rounded object-contain"
								/>
							</div>
						{:else}
							<div
								class="flex h-[200px] w-[200px] items-center justify-center rounded-xl bg-gray-300"
							>
								<Plus class="h-5 w-5" />
							</div>
						{/if}
						<input class="hidden" type="file" accept="image/*" on:change={onMainFileChange} />
					</div>
				</label>
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
					<input class="hidden" type="file" multiple accept="image/*" on:change={onSubFileChange} />
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

			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
