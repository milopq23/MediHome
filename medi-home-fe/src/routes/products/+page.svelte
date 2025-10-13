<!-- src/routes/products/+page.svelte -->
<script>
	import ProductCard from '$lib/components/ProductCard.svelte';
	import { listMedicine } from '$lib/api/medicine.js';
	import { onMount } from 'svelte';

	let totalPages = 10;
	let pageSize = 10;
	let total = 0;
	let page = 1;
	let medicines = [];

	async function listMedicinePage(currentPage = 1) {
		try {
			const result = await listMedicine(currentPage, pageSize);
			medicines = result.medicines;
			page = result.page;
			pageSize = result.pageSize;
			total = result.total;
			totalPages = Math.ceil(total / pageSize);
		} catch (error) {
			console.log('medicine', error);
		}
	}
	onMount(() => {
		listMedicinePage();
	});

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
	{#each medicines as medicine}
		<ProductCard
			name={medicine.name}
			thumbnail={medicine.thumbnail}
			price={medicine.price}
			packaging={medicine.package}
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
