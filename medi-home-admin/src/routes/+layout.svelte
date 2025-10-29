<script>
	import '../app.css';
	let sidebarOpen = false;
	import { pageTitle } from '$lib/store.js';
</script>

<div class="relative flex h-screen overflow-hidden">
	<!-- Sidebar -->
	<aside
		id="sidebar"
		class="fixed inset-y-0 left-0 z-20 w-64 transform bg-gray-800 text-white transition-transform duration-300 ease-in-out md:relative md:flex md:translate-x-0 md:flex-col"
		class:translate-x-0={sidebarOpen}
		class:-translate-x-full={!sidebarOpen}
	>
		<div class="flex items-center justify-between px-4 py-4 md:hidden">
			<span class="text-2xl font-extrabold">MediHome</span>
			<button on:click={() => (sidebarOpen = false)} class="text-2xl text-white">✕</button>
		</div>

		<a href="/" class="hidden items-center space-x-2 px-4 py-2 text-white md:flex">
			<span class="text-2xl font-extrabold">MediHome</span>
		</a>

		<nav class="mt-4 space-y-1 px-4">
			<a href="/" class="block rounded px-4 py-2.5 hover:bg-gray-700">Dashboard</a>
			<a href="/medicine" class="block rounded px-4 py-2.5 hover:bg-gray-700">Thuốc</a>
			<a href="/medicine-cate" class="block rounded px-4 py-2.5 hover:bg-gray-700">Loại thuốc</a>
			<a href="/order" class="block rounded px-4 py-2.5 hover:bg-gray-700">Đơn hàng</a>
			<a href="/voucher" class="block rounded px-4 py-2.5 hover:bg-gray-700">Khuyến mãi</a>
			<a href="/inventory" class="block rounded px-4 py-2.5 hover:bg-gray-700">Kho</a>
			<a href="/dosage-form" class="block rounded px-4 py-2.5 hover:bg-gray-700">Dạng bào chế</a>
			<a href="/user" class="block rounded px-4 py-2.5 hover:bg-gray-700">Người dùng</a>
			<a href="/admin" class="block rounded px-4 py-2.5 hover:bg-gray-700">Admin</a>
			<a href="/" class="block rounded px-4 py-2.5 hover:bg-gray-700">Đăng xuất</a>
		</nav>
	</aside>

	<div class="relative flex flex-1 flex-col">
		<header class="flex items-center justify-between bg-white px-4 py-2 shadow md:hidden">
			<button
				on:click={() => (sidebarOpen = true)}
				class="text-2xl text-gray-700 focus:outline-none"
			>
				☰
			</button>

			<h1 class="text-2xl font-extrabold">{$pageTitle}</h1>
		</header>
		<!-- <div class="hidden md:items-center md:justify-center bg-white px-6 py-4 md:flex">
			<h1 class="text-5xl font-extrabold">{$pageTitle}</h1>
		</div> -->

		<main class="relative flex-1 overflow-auto p-2 md:p-6">
			<slot />

			{#if sidebarOpen}
				<button
					class="absolute inset-0 z-10 bg-black/30 backdrop-blur-sm transition duration-300 md:hidden"
					on:click={() => (sidebarOpen = false)}
					type="button"
					aria-label="Overlay SideBar"
				>
				</button>
			{/if}
		</main>
	</div>
</div>
