<!-- src/lib/components/ProductCard.svelte -->
<script>
	export let name;
	export let thumbnail;
	export let price;
	export let packaging; // quy cách đóng gói VD: 10 vỉ x 10 viên
</script>

<div class="min-w[180px] min-h-[230px] rounded-2xl bg-white p-8 text-center shadow">
	<img src={thumbnail} alt={name} class="mx-auto mb-2 aspect-square w-full rounded object-cover" />
	<p class="line-clamp-2 font-bold text-black">{name}</p>
	<!-- <p class="text-sm text-gray-700 mb-1">{dosage}</p>
  <p class="text-sm text-gray-500 mb-1 italic">{instructions}</p> -->
	<p class="text-blue-600">
		<span class="font-bold">{price}</span>
		<!-- <span>{unit}</span> -->
	</p>
	<p class="m-2 truncate rounded-2xl bg-gray-100 px-4 text-sm text-gray-500">{packaging}</p>
	<div class="flex items-center justify-center">
		<button class="h-8 w-2/3 rounded-3xl bg-blue-500 py-1 hover:bg-blue-600">Mua</button>
	</div>
</div>
