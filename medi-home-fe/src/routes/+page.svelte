<script>
	import AdBanner from '$lib/components/AdBanner.svelte';
	import ProductCard from '$lib/components/ProductCard.svelte';
	import Navbar from '$lib/components/SideBar.svelte';

	// Mock danh sách sản phẩm
	let products = [
		{
			drugName: 'Thuốc Exopadin 60mg Trường Thọ điều trị viêm mũi dị ứng, mày đay (3 vỉ x 10 viên)',
			price: '60.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/768x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/IMG_0238_fc19904162.jpg',
			unit: '/hộp',
			packaging: '3 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Cetirizine Stada 10mg điều trị viêm mũi dị ứng, mày đay (10 vỉ x 10 viên)',
			price: '48.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/00033097_cetirizin_stada_10mg_10x10_6872_61c9_large_81d9335efe.jpg',
			unit: '/hộp',
			packaging: '10 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Vitamin C 500mg SPHARM phòng và trị thiếu vitamin C (10 vỉ x 10 viên)',
			price: '25.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/DSC_03587_e8e7e30119.jpg',
			unit: '/viên',
			packaging: '10 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Exopadin 60mg Trường Thọ điều trị viêm mũi dị ứng, mày đay (3 vỉ x 10 viên)',
			price: '60.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/768x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/IMG_0238_fc19904162.jpg',
			unit: '/hộp',
			packaging: '3 vỉ x 10 viên'
		},{
			drugName: 'Thuốc Exopadin 60mg Trường Thọ điều trị viêm mũi dị ứng, mày đay (3 vỉ x 10 viên)',
			price: '60.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/768x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/IMG_0238_fc19904162.jpg',
			unit: '/hộp',
			packaging: '3 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Cetirizine Stada 10mg điều trị viêm mũi dị ứng, mày đay (10 vỉ x 10 viên)',
			price: '48.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/00033097_cetirizin_stada_10mg_10x10_6872_61c9_large_81d9335efe.jpg',
			unit: '/hộp',
			packaging: '10 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Vitamin C 500mg SPHARM phòng và trị thiếu vitamin C (10 vỉ x 10 viên)',
			price: '25.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/DSC_03587_e8e7e30119.jpg',
			unit: '/viên',
			packaging: '10 vỉ x 10 viên'
		},
		{
			drugName: 'Thuốc Exopadin 60mg Trường Thọ điều trị viêm mũi dị ứng, mày đay (3 vỉ x 10 viên)',
			price: '60.000đ',
			image:
				'https://cdn.nhathuoclongchau.com.vn/unsafe/768x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/IMG_0238_fc19904162.jpg',
			unit: '/hộp',
			packaging: '3 vỉ x 10 viên'
		}
			
	];
</script>

<div class="flex-1 max-full bg-white py-10">

	<!-- Giới thiệu -->
	<section class="mb-10 text-center">
		<h1 class="text-3xl font-bold text-blue-800">Chào mừng đến với Nhà Thuốc MyPharma</h1>
		<p class="mt-2 text-gray-600">
			Chăm sóc sức khỏe gia đình bạn với sản phẩm chính hãng và giá tốt nhất.
		</p>
	</section>

	<!-- Banner quảng cáo -->
	<div class="mb-6 flex justify-center px-0 lg:px-5">
		<AdBanner />
	</div>


	<div class="flex-1 bg-deepglacier py-3	">
		<h1 class="px-4 text-red-500 font-bold text-3xl mb-2 items-content" >Săn sale</h1>
		<!-- <div class="px-4 text-white font-bold text-lg mb-2">
			Săn sale
		</div> -->
		<!-- Container cuộn ngang -->
		<div class="flex-1 overflow-x-auto scroll-hidden w-full px-2 py-2">
			<div class="flex gap-4 snap-x snap-mandatory px-4 flex-shrink-0 snap-start ">
				{#each products as product}
					<div
						class="flex-shrink-0 snap-start bg-white rounded-lg shadow-md p-3
							w-40 md:w-50 box-border"
					>
						<img
							src={product.image}
							alt={product.drugName}
							class="w-40 h-40 object-cover rounded-md mb-2 items-center"
						/>
						<div class="text-gray-800 text-base font-semibold truncate">{product.drugName}</div>
						<div class="text-red-500 font-bold text-base">{product.price}</div>
					</div>
				{/each}
			</div>
		</div>
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
</div>
