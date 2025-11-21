<script>
	import { AddCart, DetailMedicine } from '$lib/api/medicine';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let medicine = {};
	let images = [];

	let selectedImg;
	let quantity = 1;
	let select_type = 'Strip';

	let id = null;
	$: medicine_id = parseInt($page.params.id, 10);

	async function detailMedicine(id) {
		medicine = await DetailMedicine(id);
		images = medicine.images;
		selectedImg = images[0];
		console.log(medicine);
	}

	async function addCart() {
		const user_id = 31;
		const res = await AddCart(user_id, medicine_id, select_type, quantity);
	}

	onMount(() => {
		detailMedicine(medicine_id);
	});

	function decrease() {
		if (quantity > 1) quantity--;
	}
	function increase() {
		quantity++;
	}

	function formatCurrency(amount) {
		return new Intl.NumberFormat('vi-VN', {
			style: 'currency',
			currency: 'VND'
		}).format(amount);
	}
</script>

<div class="scroll-pt-16 scroll-smooth px-20 py-10">
	<div class="flex flex-col gap-20 rounded-2xl bg-white p-10 lg:flex-row">
		<div></div>
		<div class="flex flex-col gap-4 lg:w-1/2">
			<img
				src={selectedImg}
				alt="Ảnh sản phẩm"
				class="w-150px aspect-square h-150 rounded-xl object-cover"
			/>

			<div class="scrollbar-hide flex flex-row gap-3 overflow-x-auto">
				{#each images as img}
					<button
						type="button"
						class="h-30 w-30 shrink-0 cursor-pointer overflow-hidden rounded-md border-2 p-0 {selectedImg ===
						img
							? 'border-blue-500'
							: 'border-transparent'} transition"
						on:click={() => (selectedImg = img)}
						aria-label="Select image"
					>
						<img src={img} alt="thumbnail" class="h-full w-full object-cover" />
					</button>
				{/each}
			</div>
		</div>

		<div class="flex flex-col gap-4 p-4 lg:w-1/2">
			<div>
				<h1 class="text-2xl font-semibold text-slate-900">
					{medicine.name}
				</h1>
				<span class="p-4">Code : {medicine.code}</span>
				<div class="mt-4 flex flex-wrap gap-4">
					<p class="text-4xl font-semibold text-blue-500">
						{formatCurrency(
							select_type === 'Strip' ? medicine.price_for_strip : medicine.price_for_box
						)}
					</p>
				</div>
			</div>
			<div class="mt-4 flex items-center gap-4">
				<span class="px-4">Chọn đơn vị tính: </span>
				<button
					class="rounded-md border px-4 py-2 font-medium {select_type === 'Strip'
						? 'border-blue-500 text-blue-500'
						: ''}"
					on:click={() => (select_type = 'Strip')}
				>
					Vỉ
				</button>

				<button
					class="rounded-md border px-4 py-2 font-medium {select_type === 'Box'
						? 'border-blue-500 text-blue-500'
						: ''}"
					on:click={() => (select_type = 'Box')}
				>
					Hộp
				</button>
			</div>
			<div class="flex flex-col">
				<ul>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Danh mục:</span>
						<span>{medicine.medcatename}</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Dạng bào chế:</span>
						<span>{medicine.dosagename}</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Quy cách:</span>
						<span>{medicine.package}</span>
					</li>
					<!-- <li class="mb-2 flex">
						<span class="w-40 font-semibold">Thành phần:</span>
						<span>Paracetamol 500mg</span>
					</li> -->
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Chỉ định:</span>
						<span>{medicine.indication}</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Thuốc kê đơn:</span>
						<span>{medicine.prescription === true ? 'Kê đơn' : 'Không kê đơn'}</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Nhà sản xuất:</span>
						<span>{medicine.manufacturer}</span>
					</li>
				</ul>
				<span class="w-40 font-semibold">Lưu ý:</span>
				<span class="p-2">{medicine.note}</span>
			</div>
			<div>
				<div class="flex gap-10">
					<p class="mb-2 block font-medium text-gray-700">Số lượng:</p>
					<div class="flex items-center gap-2">
						<button
							type="button"
							class="flex h-8 w-8 items-center justify-center rounded-lg border bg-gray-100 text-gray-600"
							on:click={decrease}
						>
							−
						</button>
						<span class="flex h-8 w-8 items-center justify-center text-gray-600">{quantity}</span>
						<button
							type="button"
							class="flex h-8 w-8 items-center justify-center rounded-lg border bg-gray-100 text-gray-600"
							on:click={increase}
						>
							+
						</button>
					</div>
				</div>

				<div class="space-y-2 p-2">
					<!-- <button
						class="w-full rounded-md bg-blue-600 py-2 font-semibold text-white transition hover:bg-blue-700"
					>
						Mua ngay
					</button> -->
					<button
						class="w-full rounded-md border border-blue-600 py-2 font-semibold text-blue-600 transition hover:bg-blue-50"
						on:click={() => addCart()}
					>
						Thêm vào giỏ
					</button>
				</div>
			</div>
		</div>
	</div>

	<div class="h-10 w-full"></div>

	<div class="flex scroll-pt-16 flex-col gap-6 scroll-smooth rounded-2xl bg-white p-10 md:flex-row">
		<nav
			class="sticky top-0 flex gap-8 overflow-x-auto
            font-bold
           md:max-h-screen md:w-64 md:flex-col md:overflow-y-auto"
		>
			<!-- <a href="#mota" class="whitespace-nowrap hover:underline">Mô tả</a> -->
			<a href="#chidinh" class="whitespace-nowrap hover:underline">Chỉ định</a>
			<a href="#lieudung" class="whitespace-nowrap hover:underline">Liều dùng</a>
			<a href="#chongchidinh" class="whitespace-nowrap hover:underline">Chống chỉ định</a>
			<a href="#tacdungphu" class="whitespace-nowrap hover:underline">Tác dụng phụ</a>
			<a href="#tuongtacthuoc" class="whitespace-nowrap hover:underline">Tương tác thuốc</a>
			<a href="#baoquan" class="whitespace-nowrap hover:underline">Bảo quản</a>
		</nav>

		<!-- Nội dung bên phải -->
		<div class="flex flex-col gap-10 md:flex-2">
			<!-- <section id="mota" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Mô tả</h2>
				<p>{medicine.description}
				</p>
			</section> -->

			<section id="lieudung" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Liều dùng</h2>
				<p>{medicine.usage}</p>
			</section>

			<section id="chidinh" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Chỉ định</h2>
				<p>
					{medicine.indication}
				</p>
			</section>

			<section id="chongchidinh" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Chống chỉ định</h2>
				<p>
					{medicine.contraindication}
				</p>
			</section>

			<section id="tacdungphu" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Tác dụng phụ</h2>
				<p>
					{medicine.adverse}
				</p>
			</section>

			<section id="mangthai" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Phụ nữ mang thai</h2>
				<p>
					{medicine.pregnancy}
				</p>
			</section>

			<section id="laixe" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Khả năng lái xe và vận hành máy móc</h2>
				<p>
					{medicine.ability}
				</p>
			</section>

			<section id="tuongtacthuoc" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Tương tác thuốc</h2>
				<p>
					{medicine.drug_interaction}
				</p>
			</section>

			<section id="baoquan" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-semibold">Bảo quản</h2>
				<p>
					{medicine.storage}
				</p>
			</section>
		</div>
	</div>

	<!-- <div class="scrollbar-hide w-full flex-1 overflow-x-auto px-2 py-2">
		<div class="flex flex-shrink-0 snap-x snap-mandatory snap-start gap-4 px-4">
			{#each products as product}
				<div
					class="box-border w-40 flex-shrink-0 snap-start rounded-lg bg-white
							p-3 shadow-md md:w-50"
				>
					<img
						src={product.image}
						alt={product.name}
						class="mb-2 h-40 w-40 items-center rounded-md object-cover"
					/>
					<div class="truncate text-base font-semibold text-gray-800">{product.drugName}</div>
					<div class="text-base font-bold text-red-500">{product.price}</div>
				</div>
			{/each}
		</div>
	</div> -->
</div>

<!-- <style>
	.scrollbar-hide {
		-ms-overflow-style: none;
		scrollbar-width: none; /* Firefox */
	}

	.scrollbar-hide::-webkit-scrollbar {
		display: none; /* Chrome, Safari, Opera */
	}
</style> -->
