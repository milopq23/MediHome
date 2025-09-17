<script>
	let images = [
		'https://production-cdn.pharmacity.io/digital/640x640/plain/e-com/images/ecommerce/P01392_1.jpg',
		'https://production-cdn.pharmacity.io/digital/640x640/plain/e-com/images/ecommerce/P01392_3.jpg',
		'https://prod-cdn.pharmacity.io/e-com/images/ecommerce/P01392_6.jpg',
		'https://prod-cdn.pharmacity.io/e-com/images/ecommerce/P01392_7.jpg',
		'https://production-cdn.pharmacity.io/digital/640x640/plain/e-com/images/ecommerce/P01392_1.jpg',
		'https://production-cdn.pharmacity.io/digital/640x640/plain/e-com/images/ecommerce/P01392_3.jpg',
		'https://prod-cdn.pharmacity.io/e-com/images/ecommerce/P01392_6.jpg',
		'https://prod-cdn.pharmacity.io/e-com/images/ecommerce/P01392_7.jpg'
	];
	let cate = 'Giảm đau hạ sốt';
	let packageType = '12 vỉ/hộp';
	let dosage = 'Viên nén';
	let prescription = 'Không';
	let note =
		'Mọi thông tin trên đây chỉ mang tính chất tham khảo. Vui lòng đọc kĩ thông tin chi tiết ở tờ hướng dẫn sử dụng của sản phẩm.';
	// Ảnh lớn ban đầu
	let selectedImg = images[0];
	let quantity = 1;

	function decrease() {
		if (quantity > 1) quantity--;
	}
	function increase() {
		quantity++;
	}
</script>

<div class="bg-snowblue flex flex-col gap-6 lg:flex-row p-10">
	<!-- Left: ảnh -->
	<div class="flex flex-col gap-4 lg:w-1/2 ">
		<!-- Ảnh lớn -->
		<img
			src={selectedImg}
			alt="Ảnh sản phẩm"
			class="aspect-square w-full rounded-xl object-cover"
		/>

		<!-- List ảnh nhỏ -->
		<div class="scrollbar-hide flex flex-row gap-3 overflow-x-auto">
			{#each images as img}
				<img
					src={img}
					alt="thumbnail"
					class="w-30 h-30 flex-shrink-0 cursor-pointer rounded-md border-2 {selectedImg === img
						? 'border-blue-500'
						: 'border-transparent'} transition hover:scale-105"
					on:click={() => (selectedImg = img)}
				/>
			{/each}
		</div>
	</div>

	<!-- Right: thông tin -->
	<div class="flex flex-col gap-4 lg:w-1/2">
		<div>
			<h1 class="text-xl font-bold">
				Viên nén Panadol Extra Đỏ. Giảm triệu chứng đau, hạ sốt (15 vỉ x 12 viên)
			</h1>
		</div>
		<!-- <p class="text-gray-600">
			Đây là mô tả sản phẩm dài có thể xuống dòng. Nếu quá nhiều text thì vẫn responsive bình
			thường.
		</p> -->
		<h1 class="text-3xl font-bold text-red-500">19.000 vnd</h1>
		<p>Danh mục: {cate}</p>
		<p>Thuốc cần kê đơn: {prescription}</p>
		<p>Dạng bào chế: {dosage}</p>
        <p>Quy cách: {packageType}</p>
        <p>Lưu ý: {note}</p>
		<div>
			<label class="mb-2 block font-medium text-gray-700">Số lượng:</label>
			<div class="flex items-center space-x-3">
				<button
					type="button"
					class="flex h-8 w-8 items-center justify-center rounded-full border bg-gray-100 text-gray-600"
					on:click={decrease}
				>
					−
				</button>
				<span class="w-6 text-center">{quantity}</span>
				<button
					type="button"
					class="flex h-8 w-8 items-center justify-center rounded-full border bg-gray-100 text-gray-600"
					on:click={increase}
				>
					+
				</button>
			</div>
			<div class="space-y-2 p-2">
				<button
					class="w-full rounded-md bg-blue-600 py-2 font-semibold text-white transition hover:bg-blue-700"
				>
					Mua ngay
				</button>
				<button
					class="w-full rounded-md border border-blue-600 py-2 font-semibold text-blue-600 transition hover:bg-blue-50"
				>
					Thêm vào giỏ
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	/* Ẩn scrollbar cho đẹp */
	.scrollbar-hide::-webkit-scrollbar {
		display: none;
	}
	.scrollbar-hide {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
</style>
