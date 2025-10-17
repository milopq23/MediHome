<script>
	import { apiDetailMedicine } from '$lib/api/medicine';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

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
	let medicine = [];
	let cate = 'Giảm đau hạ sốt';
	let packageType = '12 vỉ/hộp';
	let dosage = 'Viên nén';
	let prescription = 'Không';
	let note =
		'Mọi thông tin trên đây chỉ mang tính chất tham khảo. Vui lòng đọc kĩ thông tin chi tiết ở tờ hướng dẫn sử dụng của sản phẩm.';

	let selectedImg = images[0];
	let quantity = 1;

	let id = null;
	$: id = $page.params.id;

	async function detailMedicine(id) {
		try {
			const medicine = await apiDetailMedicine(id);
			console.log(medicine);
		} catch (error) {}
	}

	onMount(() => {
		// detailMedicine();
	});

	function decrease() {
		if (quantity > 1) quantity--;
	}
	function increase() {
		quantity++;
	}
</script>

<div class="[scroll-padding-top:4rem] scroll-smooth">
	<div class="flex flex-col gap-6 bg-white p-10 lg:flex-row">
		<div class="flex flex-col gap-4 lg:w-1/2">
			<img
				src={selectedImg}
				alt="Ảnh sản phẩm"
				class="aspect-square w-full rounded-xl object-cover"
			/>

			<div class="scrollbar-hide flex flex-row gap-3 overflow-x-auto">
				{#each images as img}
					<button
						type="button"
						class="h-30 w-30 flex-shrink-0 cursor-pointer overflow-hidden rounded-md border-2 p-0 {selectedImg ===
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
		<div class="flex flex-col gap-4 lg:w-1/2">
			<div>
				<h1 class="text-xl font-semibold text-slate-900">
					Viên nén Panadol Extra Đỏ. Giảm triệu chứng đau, hạ sốt (15 vỉ x 12 viên)
				</h1>
				<div class="mt-8 flex flex-wrap gap-4">
					<p class="text-4xl font-semibold text-red-500">1200 đ</p>
					<!-- <p class="text-base text-slate-500">
					<strike>15000</strike> <span class="ml-1 text-sm">Sale</span>
				</p> -->
				</div>
			</div>
			<div class="flex flex-col">
				<ul>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Danh mục:</span>
						<span>Thuốc cảm</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Dạng bào chế:</span>
						<span>Viên nén</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Quy cách:</span>
						<span>Hộp 10 vỉ x 10 viên</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Thành phần:</span>
						<span>Paracetamol 500mg</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Chỉ định:</span>
						<span>Giảm đau, hạ sốt</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Thuốc kê đơn:</span>
						<span>Có</span>
					</li>
					<li class="mb-2 flex">
						<span class="w-40 font-semibold">Nhà sản xuất:</span>
						<span>CTCP Dược phẩm ABC</span>
					</li>
				</ul>
				<span class="w-40 font-semibold">Lưu ý:</span>
				<span class="p-2">{note}</span>
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

	<div class="flex [scroll-padding-top:4rem] flex-col gap-6 scroll-smooth bg-white p-4 md:flex-row">
		<nav
			class="sticky top-0 flex gap-4 overflow-x-auto
           bg-gray-200 p-4 font-bold
           md:max-h-screen md:w-64 md:flex-col md:overflow-y-auto"
		>
			<a href="#mota" class="whitespace-nowrap hover:underline">Mô tả</a>
			<a href="#chidinh" class="whitespace-nowrap hover:underline">Chỉ định</a>
			<a href="#lieudung" class="whitespace-nowrap hover:underline">Liều dùng</a>
			<a href="#chongchidinh" class="whitespace-nowrap hover:underline">Chống chỉ định</a>
			<a href="#tacdungphu" class="whitespace-nowrap hover:underline">Tác dụng phụ</a>
			<a href="#tuongtacthuoc" class="whitespace-nowrap hover:underline">Tương tác thuốc</a>
			<a href="#baoquan" class="whitespace-nowrap hover:underline">Bảo quản</a>
		</nav>

		<!-- Nội dung bên phải -->
		<div class="flex flex-col gap-10 md:flex-[2]">
			<section id="mota" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Mô tả</h2>
				<p>
					Thuốc ABC là sản phẩm điều trị các bệnh lý về XYZ. Được bào chế dạng viên nén dễ sử dụng
					và hấp thu nhanh.
				</p>
			</section>

			<section id="lieudung" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Liều dùng</h2>
				<p>Người lớn: 1 viên/lần, 2 lần/ngày sau ăn. Trẻ em: Theo chỉ định của bác sĩ.</p>
			</section>

			<section id="chidinh" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Chỉ định</h2>
				<p>
					Thuốc được chỉ định trong các trường hợp nhiễm khuẩn nhẹ đến trung bình như viêm họng,
					viêm phế quản, nhiễm trùng da và mô mềm.
				</p>
			</section>

			<section id="chongchidinh" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Chống chỉ định</h2>
				<p>
					Không dùng thuốc cho người mẫn cảm với bất kỳ thành phần nào của thuốc hoặc có tiền sử dị
					ứng với nhóm thuốc tương tự.
				</p>
			</section>

			<section id="tacdungphu" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Tác dụng phụ</h2>
				<p>
					Các tác dụng không mong muốn có thể gặp bao gồm buồn nôn, tiêu chảy, phát ban, chóng mặt.
					Thông báo cho bác sĩ nếu gặp phản ứng nghiêm trọng.
				</p>
			</section>

			<section id="mangthai" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Phụ nữ mang thai</h2>
				<p>
					Chưa có đủ nghiên cứu để xác định tính an toàn. Chỉ dùng thuốc khi thật sự cần thiết và
					theo chỉ định của bác sĩ.
				</p>
			</section>

			<section id="laixe" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Khả năng lái xe và vận hành máy móc</h2>
				<p>
					Thuốc có thể gây chóng mặt ở một số người. Nên thận trọng khi lái xe hoặc vận hành máy
					móc.
				</p>
			</section>

			<section id="tuongtacthuoc" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Tương tác thuốc</h2>
				<p>
					Thuốc có thể tương tác với một số loại thuốc kháng sinh, thuốc chống đông máu. Tham khảo ý
					kiến bác sĩ khi dùng chung với thuốc khác.
				</p>
			</section>

			<section id="baoquan" class="scroll-mt-16">
				<h2 class="mb-2 text-xl font-bold">Bảo quản</h2>
				<p>
					Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánBảo quản nơi khô ráo, nhiệt độ dưới
					30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ em. h sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em.Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em. Bảo quản nơi khô ráo, nhiệt độ dưới 30°C, tránh ánh sáng trực tiếp. Để xa tầm tay trẻ
					em.
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

<style>
	.scrollbar-hide {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}

	.scrollbar-hide::-webkit-scrollbar {
		display: none; /* Chrome, Safari, Opera */
	}
</style>
