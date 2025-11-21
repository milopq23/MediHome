<script>
	import '../app.css';
	import Navbar from '$lib/components/SideBar.svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import SideBar from '$lib/components/SideBar.svelte';
	import { onMount } from 'svelte';
	import AdBanner from '$lib/components/AdBanner.svelte';

	export let data;
	const { user } = data;
	// console.log(data);

	let show = false;

	function goToTop() {
		window.scrollTo({ top: 0, behavior: 'smooth' });
	}

	onMount(() => {
		const handleScroll = () => {
			show = window.scrollY > 200;
		};

		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	});
</script>

<div class="max-h-screen min-h-screen">
	<Header {user} />
	<main class="w-full flex-1 items-center justify-center bg-gray-100 px-0 md:px-20 lg:px-30">
		<slot />
	</main>
	{#if show}
		<button
			class="right-4 bottom-4 hidden h-12 w-12 justify-center rounded-full bg-blue-600 text-white shadow-lg hover:bg-blue-700 focus:outline-none md:fixed md:block"
			on:click={goToTop}
			aria-label="Go to top"
		>
			<i class="fa-solid fa-arrow-up"></i>
		</button>
	{/if}
	<Footer />
</div>
