<script>
	import { onMount } from 'svelte';
	import {
		DeleteCartItem,
		GetCartUser,
		GetShipping,
		UpdateQuantityOrTypeCart
	} from '$lib/api/cart';
	import { updated } from '$app/state';
	import { Trash2 } from 'lucide-svelte';

	let cart = {
		cartitems: [],
		totalamount: 0
	};
	let shipping = {};
	let shipping_id;

	$: if (shipping_id) {
		getShipping(shipping_id);
	}
	$: total = (shipping.price || 0) + cart.totalamount;

	// $: total = shipping.price + calculateTotal()

	function getPrice(item) {
		return item.select_type === 'Box' ? item.price_box : item.price_strip;
	}

	function calculateTotal() {
		cart.totalamount = cart.cartitems.reduce(
			(acc, item) => acc + getPrice(item) * item.quantity,
			0
		);
		console.log(cart.totalamount);
		cart = { ...cart };
	}

	const formatVND = (value) => {
		return new Intl.NumberFormat('vi-VN').format(value) + 'đ';
	};

	async function getCart(user_id) {
		const data = await GetCartUser(user_id);
		if (data) {
			cart = data;
		}
	}

	async function getShipping(shipping_id) {
		shipping = await GetShipping(shipping_id);
	}

	async function updateQuantityOrType(cartitem_id, newQuantity, newType) {
		const updatedData = {
			quantity: Number(newQuantity),
			select_type: newType
		};

		const updatedItem = await UpdateQuantityOrTypeCart(cartitem_id, updatedData);
		if (updatedItem) {
			cart.cartitems = cart.cartitems.map((item) =>
				item.cartitem_id === cartitem_id ? { ...item, ...updatedItem } : item
			);
		}
		calculateTotal();
	}

	async function deleteCartItem(cartitem_id) {
		// const data = await DeleteCartItem(cartitem_id);
		await DeleteCartItem(cartitem_id);

		cart.cartitems = cart.cartitems.filter((item) => item.cartitem_id !== cartitem_id);

		calculateTotal();
	}

	onMount(() => {
		const user_id = 31;
		getCart(user_id);
	});
</script>

<div class="mx-auto px-4 py-8">
	<h1 class="mb-6 text-2xl font-semibold text-gray-800">Giỏ hàng</h1>
	<div class="flex px-30">
		<div class="w-3/4 px-10">
			<div class="pb-10">
				<h1 class="mb-4 text-xl font-semibold">Danh sách giỏ hàng</h1>
				<table class="min-w-full pb-10">
					<thead class="bg-gray-50 text-gray-700">
						<tr>
							<th class="px-4 py-3 text-left text-sm font-medium">Sản phẩm</th>
							<th class="hidden px-4 py-3 text-sm font-medium md:block">Giá</th>
							<th class="px-4 py-3 text-sm font-medium">Loại</th>
							<th class="px-4 py-3 text-sm font-medium">Số lượng</th>
							<th class="px-4 py-3 text-sm font-medium">Tổng</th>
							<th class="px-4 py-3"></th>
						</tr>
					</thead>

					<tbody class="divide-y divide-gray-100">
						{#each cart.cartitems as item}
							<tr class="transition hover:bg-gray-50">
								<td class="flex items-center gap-4 px-4 py-3">
									<img
										src="https://images.unsplash.com/photo-1618354691373-d851c5c3a990?auto=format&fit=crop&w=100&q=80"
										alt={item.name}
										class="size-16 rounded-md object-cover"
									/>
									<span class="text-xs text-gray-800 md:font-medium">{item.name}</span>
								</td>

								<td class="px-4 py-3 text-center text-gray-700">{formatVND(getPrice(item))}</td>

								<td class="px-4 py-3 text-center text-xs">
									<select
										bind:value={item.select_type}
										on:change={(e) =>
											updateQuantityOrType(item.cartitem_id, item.quantity, e.target.value)}
										class="rounded border border-gray-300 bg-gray-50 px-2 py-1 text-xs focus:ring-1 focus:ring-blue-400 focus:outline-none md:text-sm"
									>
										<option value="Box">Hộp</option>
										<option value="Strip">Vỉ</option>
									</select>
								</td>

								<td class="px-4 py-3 text-center">
									<input
										type="number"
										min="1"
										bind:value={item.quantity}
										on:input={(e) =>
											updateQuantityOrType(item.cartitem_id, e.target.value, item.select_type)}
										class="w-16 rounded border border-gray-300 bg-gray-50 px-2 py-1 text-center text-sm focus:ring-1 focus:ring-blue-400 focus:outline-none"
									/>
								</td>

								<td class="px-4 py-3 text-center text-gray-700"
									>{formatVND(getPrice(item) * item.quantity)}</td
								>

								<td class="px-4 py-3 text-center">
									<button
										class="rounded p-1 text-gray-500 transition hover:bg-red-100 hover:text-red-600"
										title="Xóa sản phẩm"
										on:click={() => deleteCartItem(item.cartitem_id)}
									>
										<Trash2 />
									</button>
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>

			<div>
				<h1 class="mb-4 text-xl font-semibold">Chọn phương thức giao hàng</h1>
				<div class="flex flex-1 gap-6 p-4">
					<!-- Option 1 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="shipping"
							bind:group={shipping_id}
							value="1"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">Giao hàng tiết kiệm</span>
					</label>

					<!-- Option 2 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="shipping"
							bind:group={shipping_id}
							value="2"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">Giao hàng nhanh</span>
					</label>

					<!-- Option 3 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="shipping"
							bind:group={shipping_id}
							value="3"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">Giao hàng cấp tốc</span>
					</label>
				</div>
			</div>
			<div>
				<h1 class="mb-4 text-xl font-semibold">Chọn phương thức thanh toán</h1>
				<div class="flex flex-1 gap-4 p-4">
					<!-- Option 1 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="payment_method"
							value="COD"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">Tiền mặt</span>
					</label>

					<!-- Option 2 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="payment_method"
							value="MOMO"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">Momo</span>
					</label>

					<!-- Option 3 -->
					<label
						class="bg-neutral-primary-soft hover:border-brand flex w-full cursor-pointer items-center gap-2 rounded-xl border px-4 py-3 transition"
					>
						<input
							type="radio"
							name="payment_method"
							value="VNPAY"
							class="text-brand border-default-medium focus:ring-brand-subtle h-4 w-4 rounded-full"
						/>
						<span class="text-sm font-medium select-none">VNPAY</span>
					</label>
				</div>
			</div>
		</div>
		<div class="w-1/4 px-5">
			<h1 class="mb-4 text-xl font-semibold">Tổng tiền</h1>
			<button
				class="mb-3 flex w-full items-center justify-between rounded-lg bg-blue-50 px-3 py-2 text-blue-600"
			>
				<span>Áp dụng ưu đãi để được giảm giá</span>
				<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M19 9l-7 7-7-7"
					/>
				</svg>
			</button>

			<!-- Bảng giá -->
			<div class="space-y-2">
				<div class="flex justify-between">
					<span class="font-semibold">Tổng tiền</span>
					<span class="font-medium">{formatVND(cart.totalamount)}</span>
				</div>

				<div class="flex justify-between">
					<div class="flex items-center gap-1">
						<span>Giảm giá voucher</span>
						<span
							class="flex h-4 w-4 cursor-help items-center justify-center rounded-full bg-gray-200 text-[10px] text-gray-500"
							>i</span
						>
					</div>
					<span>0đ</span>
				</div>

				<div class="flex justify-between text-green-600">
					<span>Tiết kiệm được</span>
					<span>29.800đ</span>
				</div>

				<div class="flex justify-between">
					<span>Phí vận chuyển</span>
					<span class="font-medium text-blue-600"
						>{shipping?.price ? formatVND(shipping.price) : '0đ'}</span
					>
				</div>
			</div>

			<!-- Thành tiền -->
			<div class="mt-4 border-t pt-3">
				<div class="flex justify-between">
					<span class="text-xl font-bold">Thành tiền</span>
					<div class="flex justify-between">
						<div class="text-xl font-bold text-blue-500">{formatVND(total)}</div>
					</div>
				</div>

				<!-- Nút -->
				<button
					class="mt-3 w-full rounded-full bg-blue-600 py-3 text-center text-sm font-medium text-white transition hover:bg-blue-700"
				>
					Hoàn tất
				</button>
			</div>
		</div>
	</div>
</div>
