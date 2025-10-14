<!-- src/routes/products/+page.svelte -->
<script>
	import ProductCard from '$lib/components/ProductCard.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import { ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight } from 'lucide-svelte';
	import { listMedicine } from '$lib/api/medicine.js';
	import { onMount } from 'svelte';

	let totalPages = 1;
	let pageSize = 1;
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

	function handleGoToPage(event) {
		const selectedPage = event.detail;
		listMedicinePage(selectedPage);
	}

	// function goToPage(page) {
	// 	if (page >= 1 && page <= totalPages) {
	// 		listMedicinePage(page); // Gọi lại API để load dữ liệu trang mới
	// 	}
	// }

	// function nextPage() {
	// 	if (currentPage < totalPages) {
	// 		goToPage(page + 1);	
	// 	}
	// }
	// function prevPage() {
	// 	if (currentPage > 1) {
	// 		goToPage(page - 1);
	// 	}
	// }
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

<Pagination {page} {totalPages} on:goToPage={handleGoToPage} />
<!-- <Pagination {page} {totalPages} on:goToPage={handleGoToPage} /> -->