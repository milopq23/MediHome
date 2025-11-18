<script>
	import { DetailVoucher, UpdateVoucher } from '$lib/api/voucher.js';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import HeaderDetail from '$lib/components/HeaderDetail.svelte';

	let start_date = '';
	let end_date = '';
	let voucher = {};

	let id = null;
	$: voucher_id = parseInt($page.params.id, 10);

	onMount(() => {
		loadDetailVoucher(voucher_id);
	});

	async function loadDetailVoucher(voucher_id) {
		const data = await DetailVoucher(voucher_id);
		voucher = data;

		start_date = voucher.start_date ? voucher.start_date.split('T')[0] : '';
		end_date = voucher.end_date ? voucher.end_date.split('T')[0] : '';
	}

	
	async function updateVoucher() {
		const data = { ...voucher };
		await UpdateVoucher(voucher_id, data);
	}

</script>

<div class="flex items-center justify-center py-10">
	<div class="w-full max-w-3xl rounded-2xl bg-white p-8 shadow-xl dark:bg-gray-800">
		<HeaderDetail title="Trang chi tiết voucher" route="voucher" />

		<form class="space-y-6" on:submit|preventDefault={updateVoucher}>
			<!-- Code -->
			<div>
				<label for="code" class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200">
					Code:
				</label>
				<input
					type="text"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={voucher.code}
				/>
			</div>

			<!-- Name -->
			<div>
				<label for="name" class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200">
					Tên voucher:
				</label>
				<input
					type="text"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={voucher.name}
				/>
			</div>

			<!-- Description -->
			<div>
				<label
					for="description"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Mô tả:
				</label>
				<textarea
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					rows="3"
					bind:value={voucher.description}
				></textarea>
			</div>

			<!-- Discount type -->
			<div>
				<label
					for="discount_type"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Loại:
				</label>
				<select
					bind:value={voucher.discount_type}
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
				>
					<option value="" disabled selected>Chọn loại voucher</option>
					<option value="Phần trăm">Phần trăm</option>
					<option value="Cố định">Cố định</option>
				</select>
			</div>
			{#if voucher.discount_type === 'Phần trăm'}
				<div>
					<label
						for="discount_value"
						class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
					>
						Giá trị giảm (%):
					</label>
					<input
						type="number"
						class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
						bind:value={voucher.discount_value}
					/>
				</div>
			{/if}
			<div>
				<label
					for="min_order_value"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Ngày bắt đầu:
				</label>
				<input
					type="date"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={start_date}
				/>
			</div>

			<div>
				<label
					for="min_order_value"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Ngày kết thúc:
				</label>
				<input
					type="date"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={end_date}
				/>
			</div>
			<!-- Min order value -->
			<div>
				<label
					for="min_order_value"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Đơn hàng tối thiểu:
				</label>
				<input
					type="number"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={voucher.min_order_value}
				/>
			</div>

			<!-- Max discount -->
			<div>
				<label
					for="max_discount_value"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Giảm tối đa:
				</label>
				<input
					type="number"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={voucher.max_discount_value}
				/>
			</div>

			<!-- Usage limit -->
			<div>
				<label
					for="usage_limit"
					class="mb-2 block text-sm font-semibold text-gray-700 dark:text-gray-200"
				>
					Giới hạn lượt sử dụng
				</label>
				<input
					type="number"
					class="w-full rounded-lg border border-gray-300 bg-gray-50 p-3 text-sm shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-400 dark:border-gray-600 dark:bg-gray-700 dark:text-white"
					bind:value={voucher.usage_limit}
				/>
			</div>
			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
