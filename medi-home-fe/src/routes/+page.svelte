<!-- <script>
	import AdBanner from '$lib/components/AdBanner.svelte';
	import ProductCard from '$lib/components/ProductCard.svelte';
	import Navbar from '$lib/components/SideBar.svelte';
</script>


<main>
	<AdBanner />
	<h1 class="mb-4 text-3xl font-bold">Chào mừng đến với MyApp</h1>
	<p class="mb-2">Đây là phần nội dung chính. Kéo xuống để xem footer vẫn hiển thị.</p>
	<p>
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus
		tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor...
	</p>
	<div class="mt-6 h-[1500px] rounded bg-gray-200">Nội dung dài để test cuộn</div>
</main> -->
<script>
	import AdBanner from '$lib/components/AdBanner.svelte';
	import ProductCard from '$lib/components/ProductCard.svelte';
	import Navbar from '$lib/components/SideBar.svelte';

	// Mock danh sách sản phẩm
	let products = [
		{
			id: 1,
			drugName: 'Paracetamol 500mg',
			price: '25,000đ',
			image: 'https://via.placeholder.com/150x150?text=Paracetamol',
			unit: 'viên'
		},
		{
			id: 2,
			drugName: 'Vitamin C 1000mg',
			price: '45,000đ',
			image: 'https://via.placeholder.com/150x150?text=Vitamin+C',
			unit:'hộp'
		},
		{
			id: 3,
			drugName: 'Khẩu trang y tế',
			price: '30,000đ',
			image: 'https://via.placeholder.com/150x150?text=Khau+Trang'
		},
		{
			id: 4,
			drugName: 'Gel rửa tay khô',
			price: '35,000đ',
			image: 'https://via.placeholder.com/150x150?text=Gel+Rua+Tay'
		}
	];
</script>

<div class="flex min-h-screen bg-gray-50">
	<!-- Main content -->
	<main class="flex-1 p-4">
		<!-- Giới thiệu -->

		<section class="mb-10 text-center">
			<h1 class="text-3xl font-bold text-blue-800">Chào mừng đến với Nhà Thuốc MyPharma</h1>
			<p class="mt-2 text-gray-600">
				Chăm sóc sức khỏe gia đình bạn với sản phẩm chính hãng và giá tốt nhất.
			</p>
		</section>

		<!-- Banner quảng cáo -->
		<div class="mb-6 flex justify-center">
			<AdBanner />
		</div>


		<section class="mb-10 rounded-lg bg-white p-6 shadow-sm">
			<!-- Dán đoạn giới thiệu ở đây -->
			<h2 class="mb-4 text-2xl font-semibold text-blue-700 flex justify-center">Về chúng tôi</h2>
			<p class="leading-relaxed text-gray-700">
				MyPharma là hệ thống nhà thuốc trực tuyến uy tín, cam kết cung cấp các sản phẩm chăm sóc sức
				khỏe chính hãng với giá cả hợp lý. Với đội ngũ dược sĩ chuyên môn cao, chúng tôi luôn sẵn
				sàng tư vấn và đồng hành cùng bạn trong hành trình chăm sóc sức khỏe.
			</p>

			<ul class="mt-4 list-inside list-disc space-y-1 text-gray-700">
				<li>Sản phẩm được kiểm duyệt, có nguồn gốc rõ ràng</li>
				<li>Đội ngũ tư vấn viên, dược sĩ chuyên nghiệp</li>
				<li>Giao hàng toàn quốc, nhanh chóng, đúng hẹn</li>
				<li>Bảo mật thông tin khách hàng tuyệt đối</li>
			</ul>

			<div class="mt-6 flex flex-col items-center gap-6 md:flex-row">
				<img
					src="https://down-tx-vn.img.susercontent.com/5d12f5221cbad0685e95b80cbbe60a50_tn.webp"
					alt="Hình ảnh nhà thuốc"
					class="w-full rounded-lg object-cover shadow-md md:w-1/2"
				/>

				<div class="md:w-1/2">
					<h3 class="mb-2 text-lg font-medium text-green-700">Sứ mệnh của chúng tôi</h3>
					<p class="text-gray-700">
						Mang đến giải pháp y tế tiện lợi, an toàn và hiệu quả cho cộng đồng thông qua hệ thống
						nhà thuốc trực tuyến hiện đại, giúp mọi người tiếp cận thuốc và sản phẩm chăm sóc sức
						khỏe một cách dễ dàng và minh bạch.
					</p>
				</div>
			</div>
		</section>

		<!-- Sản phẩm nổi bật -->
		<section>
			<h2 class="mb-4 text-2xl font-semibold text-blue-700">Sản phẩm nổi bật</h2>
			<div class="grid grid-cols-2 gap-4 md:grid-cols-4">
				{#each products as product}
				<ProductCard
					drugName={product.drugName}
					price={product.price}
					image={product.image}
					unit={product.unit}
					packaging={product.packaging}
				/>
			{/each}
			</div>
		</section>

		<!-- Nội dung phụ hoặc ưu đãi -->
		<section class="mt-10">
			<h2 class="mb-3 text-xl font-semibold text-green-700">Ưu đãi đang diễn ra</h2>
			<ul class="list-inside list-disc text-gray-700">
				<li>Mua 1 tặng 1 với sản phẩm Vitamin C</li>
				<li>Miễn phí giao hàng cho đơn từ 300.000đ</li>
				<li>Giảm 10% cho khách hàng mới</li>
			</ul>
		</section>

		<!-- Nội dung dài để test cuộn -->
		<div
			class="mt-10 flex h-[1000px] items-center justify-center rounded bg-gray-200 text-gray-500"
		>
			Nội dung dài để test cuộn
		</div>
	</main>
</div>
