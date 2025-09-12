<script>
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/SideBar.svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import SideBar from '$lib/components/SideBar.svelte';
	import {onMount} from 'svelte';

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

<div class="min-h-screen bg-gray-100">
	<Header />

	<!-- <SideBar /> -->

	<main class="mx-auto max-w-3xl p-4">
		<slot />
	</main>
	{#if show}
	<button
		class="hidden md:block md:fixed bottom-4 right-4 rounded-full  bg-blue-600 w-12 h-12 justify-center text-white shadow-lg hover:bg-blue-700 focus:outline-none"
		on:click={goToTop}
		aria-label="Go to top"
	>
		<i class="fa-solid fa-arrow-up"></i>
	</button>
	{/if}
	<Footer />
</div>
