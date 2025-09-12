<!-- src/routes/products/+page.svelte -->
<script>
	import ProductCard from '$lib/components/ProductCard.svelte';

	let totalPages = 10;
	let currentPage = 1;

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
		}
		// {
		//   drugName: "Vitamin C",
		//   price: "15.000đ",
		//   image: "https://via.placeholder.com/100"
		// }
	];

	function goToPage(page) {
		if (page >= 1 && page <= totalPages) {
			currentPage = page;
			// Logic tải dữ liệu sản phẩm cho trang hiện tại
		}
	}

	function nextPage() {
		if (currentPage < totalPages) {
			goToPage(currentPage + 1);
		}
	}
	function prevPage() {
		if (currentPage > 1) {
			goToPage(currentPage - 1);
		}
	}
</script>

<h1 class="mb-4 text-2xl font-bold">Sản phẩm</h1>

<div class="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-4">
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

<div class="mt-6 flex justify-center space-x-2">
	<button
		on:click={prevPage}
		class="rounded border border-blue-600 bg-white px-4 py-2 text-blue-600 hover:bg-blue-100"
	>
		Prev
	</button>
	<button class="rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700"> 1 </button>
	<button class="rounded border border-blue-600 bg-white px-4 py-2 text-blue-600 hover:bg-blue-100">
		2
	</button>
	<p class="border bg-white px-4 py-2 text-blue-600">...</p>
	<button class="rounded border border-blue-600 bg-white px-4 py-2 text-blue-600 hover:bg-blue-100">
		{totalPages}
	</button>

	<button
		on:click={nextPage}
		class="rounded border border-blue-600 bg-white px-4 py-2 text-blue-600 hover:bg-blue-100"
	>
		Next
	</button>
</div>

<!-- dosage={product.dosage}
      instructions={product.instructions} -->
