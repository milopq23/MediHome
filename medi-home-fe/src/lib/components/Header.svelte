<script>
	import { sidebarOpen } from '../index.js';
	import SearchBar from './SearchBar.svelte';

	let showSearch = false;

	export let user;

	function toggleSidebar() {
		sidebarOpen.update((n) => !n);
	}
</script>

<nav class="sticky top-0 z-50 w-full rounded-b-lg bg-ice px-4 py-3 shadow-lg">
	<div class="mx-auto flex w-full max-w-7xl items-center justify-between">
		<!-- Logo -->
		<a href="/" class="text-2xl font-bold">
			<span class="text-[#2F80ED]">Medi</span><span class="text-[#09a94c]">Home</span>
		</a>

		<!-- Search Bar (Desktop) -->
		<div class="hidden flex-1 justify-center px-4 md:flex">
			<div
				class="flex w-full max-w-[500px] items-center gap-2 rounded-lg bg-blue-100 px-4 py-2 focus-within:ring-2 focus-within:ring-blue-400 md:flex"
			>
				<i class="fa-solid fa-magnifying-glass text-gray-500"></i>
				<input
					type="text"
					placeholder="Tìm kiếm sản phẩm..."
					class="w-full rounded-lg border-gray-300 bg-transparent focus:outline-none"
				/>
			</div>
		</div>

		<!-- <SearchBar (Mobile)/> -->
		<div class="flex flex-row items-center gap-4 md:flex">
			{#if !showSearch}
				<button
					class=" flex justify-end text-xl text-gray-700 focus:outline-none md:hidden"
					on:click={() => (showSearch = !showSearch)}
					aria-label="Toggle Search"
				>
					<i class="fa-solid fa-magnifying-glass"></i>
				</button>
			{/if}
			<!-- Mobile search input (show/hide) -->
			{#if showSearch}
				<div
					class="flex w-1/2 items-center gap-2 rounded-lg bg-blue-100 px-4 py-2 focus-within:ring-2 focus-within:ring-blue-400 md:flex"
				>
					<i class="fa-solid fa-magnifying-glass text-gray-500"></i>
					<input
						type="text"
						placeholder="Tìm kiếm sản phẩm..."
						class="w-full rounded-lg border-gray-300 bg-transparent focus:outline-none"
					/>
				</div>
			{/if}

			<a
				href="/cart"
				class="flex items-center gap-1 text-gray-700 hover:text-blue-600"
				aria-label="Giỏ hàng"
			>
				<i
					class="fa-solid fa-cart-shopping cursor-pointer text-xl text-gray-500 transition-colors duration-200 hover:text-blue-600"
				></i>
			</a>

			<!-- Login / User Info -->
			<div class="flex max-w-[190px] items-center">
				{#if user}
					<div
						class="text-g ray-700 hidden w-full cursor-pointer items-center gap-2 transition-colors duration-200 hover:text-blue-600 md:flex"
					>
						<a href="/profile" class="truncate overflow-hidden whitespace-nowrap hover:underline">
							Xin chào {user.name}
						</a>
					</div>
					<a
						href="/profile"
						class="cursor-pointer items-center text-xl text-gray-700 transition-colors duration-200 hover:text-blue-600 md:hidden"
						aria-label="Tài khoản"
					>
						<i class="fa-solid fa-user"></i>
					</a>
				{:else}
					<a
						href="/login"
						class="hidden w-full rounded bg-blue-600 px-3 py-1 text-center text-white hover:bg-blue-700 md:inline-block"
					>
						<i class="fa-solid fa-user"></i>
						Đăng nhập/Đăng ký
					</a>
					<div>
						<a href="/login" class="flex items-center gap-1 md:hidden" aria-label="Đăng nhập">
							<i
								class="fa-solid fa-user cursor-pointer text-xl text-gray-500 transition-colors duration-200 hover:text-blue-600"
							></i>
						</a>
					</div>
				{/if}
			</div>
		</div>
	</div>
</nav>
