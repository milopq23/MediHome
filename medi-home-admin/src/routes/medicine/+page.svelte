<script>
	import { Pencil, Trash2, MoreVertical, Eye, Plus } from 'lucide-svelte';
	import { loadMedicines, addMedicine, patchMedicine } from '$lib/api/medicine.js';
	import { onMount, onDestroy } from 'svelte';
	import TitlePage from '$lib/components/TitlePage.svelte';
	import Pagination from '$lib/components/Pagination.svelte';

	let search = '';
	let page = 1;
	let pageSize = 2;
	let total = 0;
	let totalPages = 1;

	let showPopup = true;
	let isOpenAction = false;
	let popupRef;
	let selectedActionId = null;

	let medicines = [];

	// let medicine = {
	// 	code: '',
	// 	name: '',
	// 	thumbnail: '',
	// 	category_name: '',
	// 	dosageform_name: '',
	// 	prescription: ''
	// };

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

	onDestroy(() => {});

	// function handleClickOutside(event) {
	// 	if (showPopup && popupRef && !popupRef.contains(event.target)) {
	// 		showPopup = false;
	// 	}
	// }

	function openAction(id) {
		selectedActionId = selectedActionId === id ? null : id;
	}

	function goToPage(p) {
		if (p >= 1 && p <= totalPages) page = p;
		loadMedicinePage(page);
	}
</script>

<div class="px-4 py-6 sm:px-6 lg:px-8">
	<div class="rounded-md bg-white p-4 shadow-sm">
		<TitlePage titleName="Danh sách thuốc" />

		<!-- Search bar -->
		<div class="mb-4 flex flex-col gap-2 sm:flex-row sm:items-center">
			<input
				type="text"
				class="w-full rounded-md border px-3 py-2 text-sm sm:max-w-xs"
				placeholder="Tìm kiếm thuốc..."
				bind:value={search}
			/>
		</div>
		<div class="overflow-x-auto">
			<table class="min-w-full divide-y divide-gray-200 text-sm">
				<thead class="bg-gray-50 text-left text-gray-500">
					<tr>
						<th class="p-2">Code</th>
						<th class="p-2">Tên thuốc</th>
						<th class="hidden p-2 lg:table-cell">Hình</th>
						<th class="hidden p-2 sm:table-cell">Danh mục</th>
						<th class="hidden p-2 lg:table-cell">Dạng bào chế</th>
						<th class="hidden p-2 lg:table-cell">Kê đơn</th>
						<th class="hidden p-2 md:table-cell">Hành động</th>
						<th class="w-10 p-2 md:hidden"></th>
					</tr>
				</thead>
				<tbody class="text-gray-800">
					{#each medicines as medicine}
						<tr class="border-b hover:bg-gray-50">
							<td class="w-20 p-2">{medicine.code}</td>
							<td
								class="max-w-[200px] min-w-[200px] cursor-pointer truncate p-2"
								title={medicine.name}
							>
								{medicine.name}
							</td>
							<td class="hidden p-2 lg:table-cell">{medicine.thumbnail}</td>
							<td class="hidden max-w-[50px] truncate p-2 sm:table-cell"
								>{medicine.category_name}</td
							>
							<td class="hidden p-2 lg:table-cell">{medicine.dosageform_name}</td>
							<td class="hidden p-2 lg:table-cell">{medicine.prescription ? 'Có' : 'Không'}</td>
							<td class="p-2">
								<div class="flex items-center gap-2 text-gray-600">
									<button class="hidden cursor-pointer md:table-cell" title="Xem"
										><Eye class="h-4 w-4" /></button
									>
									<button class="hidden cursor-pointer md:table-cell" title="Sửa"
										><Pencil class="h-4 w-4" /></button
									>
									<button class="hidden cursor-pointer md:table-cell" title="Xoá"
										><Trash2 class="h-4 w-4" /></button
									>
									<div>
										<div class="relative inline-block w-2 md:hidden">
											<button on:click={() => openAction(medicine.medicine_id)}>
												<MoreVertical class="h-4 w-4" />
											</button>
										</div>

										{#if selectedActionId === medicine.medicine_id}
											<div
												class="absolute right-5 z-10 mt-1 w-30 divide-y divide-gray-100 rounded-lg bg-white shadow-sm dark:bg-gray-700"
											>
												<ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
													<li>
														<button
															class="flex w-30 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
														>
															<Eye class="h-4 w-4" /> Xem
														</button>
													</li>
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
														>
															<Pencil class="h-4 w-4" /> Sửa
														</button>
													</li>
													<li>
														<button
															class="flex w-20 items-center gap-2 px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600"
														>
															<Trash2 class="h-4 w-4" /> Xoá
														</button>
													</li>
												</ul>
											</div>
										{/if}
									</div>
								</div>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
		<Pagination {page} {pageSize} {totalPages} {total} {goToPage} />
	</div>
</div>


{#if showPopup}
	<div class="fixed inset-0 z-40 bg-black/50"></div>
	<div
		class="fixed top-1/2 left-1/2 z-60 mx-auto w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-lg"
	>
		<form>
			<div class="mb-6 grid gap-6 md:grid-cols-2">
				<div>
					<label
						for="first_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Code</label
					>
					<input
						type="text"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="John"
						required
					/>
				</div>
				<div>
					<label
						for="last_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Tên thuốc</label
					>
					<input
						type="text"
						id="last_name"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Doe"
						required
					/>
				</div>
				<div>
					<label for="company" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Company</label
					>
					<input
						type="text"
						id="company"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Flowbite"
						required
					/>
				</div>
				<div>
					<label for="phone" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Phone number</label
					>
					<input
						type="tel"
						id="phone"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="123-45-678"
						pattern="[0-9]{3}-[0-9]{2}-[0-9]{3}"
						required
					/>
				</div>
				<div>
					<label for="website" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Website URL</label
					>
					<input
						type="url"
						id="website"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="flowbite.com"
						required
					/>
				</div>
				<div>
					<label for="visitors" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Unique visitors (per month)</label
					>
					<input
						type="number"
						id="visitors"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder=""
						required
					/>
				</div>
			</div>
			<div class="mb-6">
				<label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Email address</label
				>
				<input
					type="email"
					id="email"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="john.doe@company.com"
					required
				/>
			</div>
			<div class="mb-6">
				<label for="password" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Password</label
				>
				<input
					type="password"
					id="password"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="•••••••••"
					required
				/>
			</div>
			<div class="mb-6">
				<label
					for="confirm_password"
					class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Confirm password</label
				>
				<input
					type="password"
					id="confirm_password"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="•••••••••"
					required
				/>
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
					>I agree with the <a href="#" class="text-blue-600 hover:underline dark:text-blue-500"
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
{/if}
