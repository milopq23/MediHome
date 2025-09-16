<script>
	import { onMount, onDestroy } from 'svelte';

	let current = 0;
	let interval;

	const banners = [
		{
			img: 'https://production-cdn.pharmacity.io/digital/795x0/plain/e-com/images/banners/20250812062615-0-banner_Desktop1590x604px.webp?versionId=34xxuONDBYVZ9EQMw4Hi4oPUwBNKDGNZ',
			title: 'Quảng cáo 1',
			description: 'Mô tả quảng cáo 1'
		},
		// {
		// 	img: 'https://production-cdn.pharmacity.io/digital/795x0/plain/e-com/images/banners/20250910031822-0-qqq.png?versionId=cpd2omQWO8A41ZZZyf5RNq1PX1I3Q3fp',
		// 	title: 'Quảng cáo 2',
		// 	description: 'Mô tả quảng cáo 2'
		// },
		// {
		// 	img: 'https://production-cdn.pharmacity.io/digital/795x0/plain/e-com/images/banners/20250909075535-0-mweb.png?versionId=rvyVlA7FND_P_ZRzI2vb9zxzEQj_i2j_',
		// 	title: 'Quảng cáo 3',
		// 	description: 'Mô tả quảng cáo 3'
		// }
	];

	onMount(() => {
		interval = setInterval(() => {
			current = (current + 1) % banners.length;
		}, 4000); // 4 giây chuyển banner
	});

	onDestroy(() => {
		clearInterval(interval);
	});

	function prev() {
		current = (current - 1 + banners.length) % banners.length;
	}
	function next() {
		current = (current + 1) % banners.length;
	}
</script>

<!-- <div class=" flex items-center justify-center mx-auto my-4"> -->
<div
	class="flex-1 items-center justify-center grid-cols-3 grid-rows-2  gap-2 overflow-hidden md:grid"
>
	<div class="animate-scroll flex h-full w-full whitespace-nowrap md:col-span-2 md:row-span-2">
		{#each banners as banner, i}
			{#if i === current}
				<div class="relative flex h-full w-full items-center justify-center ">
					<!-- Container relative để chứa các nút và ảnh -->

					<!-- Ảnh -->
					<img
						src={banner.img}
						alt={banner.title}
						class="h-full w-full rounded-none md:rounded-xl object-contain"
					/>

					<!-- Nút trái -->
					<button
						class="absolute left-2 top-1/2 -translate-y-1/2 transform rounded-full bg-white p-2 shadow"
						aria-label="Previous Banner"
						on:click={prev}
					>
						<i class="fa-solid fa-arrow-left"></i>
					</button>

					<!-- Nút phải -->
					<button
						class="absolute right-2 top-1/2 -translate-y-1/2 transform rounded-full bg-white p-2 shadow"
						aria-label="Next Banner"
						on:click={next}
					>
						<i class="fa-solid fa-arrow-right"></i>
					</button>

					<div class="absolute bottom-4 left-1/2 flex -translate-x-1/2 transform gap-2 px-3">
						{#each banners as _, j}
							<button
								class="h-3 w-3 rounded-full transition-all duration-300
                            {current === j ? 'scale-125 bg-blue-600' : 'bg-gray-300'}"
								on:click={() => {
									current = j;
									resetInterval();
								}}
								aria-label="Go to banner {j + 1}"
							>
							</button>
						{/each}
					</div>
				</div>
			{/if}
		{/each}
	</div>
	<!-- Banner phụ trên  -->
	<div class="hidden md:col-start-3 md:row-start-1 md:flex">
		<img
			src="https://production-cdn.pharmacity.io/digital/795x0/plain/e-com/images/banners/20250910031822-0-qqq.png?versionId=cpd2omQWO8A41ZZZyf5RNq1PX1I3Q3fp"
			alt="Banner Phụ"
			class="h-full w-full rounded-xl object-contain shadow-md"
		/>
	</div>
	<!-- Banner phụ dưới -->
	<div class="hidden md:col-start-3 md:row-start-2 md:flex">
		<img
			src="https://production-cdn.pharmacity.io/digital/795x0/plain/e-com/images/banners/20250909075535-0-mweb.png?versionId=rvyVlA7FND_P_ZRzI2vb9zxzEQj_i2j_"
			alt="Banner Phụ"
			class="h-full w-full rounded-xl object-contain shadow-md"
		/>
	</div>
</div>
