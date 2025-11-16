<script>
	import ProductCard from '$lib/components/ProductCard.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import { ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight } from 'lucide-svelte';
	import { ListMedicine } from '$lib/api/medicine.js';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let totalPages = 1;
	let pageSize = 10;
	let total = 0;
	let page = 1;
	let medicines = [];

	async function listMedicine(currentPage = 1) {
		const result = await ListMedicine(currentPage, pageSize);
		medicines = result.medicines;
		page = result.page;
		pageSize = result.pageSize;
		total = result.total;
		totalPages = Math.ceil(total / pageSize);
	}


	onMount(() => {
		listMedicine();
	});

	function goToDetail(id) {
		console.log(id);
		goto(`/medicines/${id}`);
	}
	function handleGoToPage(event) {
		const selectedPage = event.detail;
		listMedicinePage(selectedPage);
	}
</script>

<h1 class="mb-4 text-2xl font-bold">Sản phẩm</h1>

<div class="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-4">
	{#each medicines as medicine}
		<button on:click={() => goToDetail(medicine.medicine_id)}>
			<ProductCard
				name={medicine.name}
				thumbnail={medicine.thumbnail}
				price={medicine.price}
				packaging={medicine.package}
			/>
		</button>
	{/each}
</div>

<Pagination {page} {totalPages} on:goToPage={handleGoToPage} />
<!-- <Pagination {page} {totalPages} on:goToPage={handleGoToPage} /> -->
