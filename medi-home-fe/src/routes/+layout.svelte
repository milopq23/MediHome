<script>
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/SideBar.svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import SideBar from '$lib/components/SideBar.svelte';
	import { onMount } from 'svelte';
	import AdBanner from '$lib/components/AdBanner.svelte';

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

<div class="min-h-screen max-h-screen ">
	<Header />
	<main class="flex-1 justify-center items-center px-0  md:px-20 lg:px-60 bg-white w-full">
		<slot />
	</main>
	{#if show}
		<button
			class="bottom-4 right-4 hidden h-12 w-12 justify-center rounded-full bg-blue-600 text-white shadow-lg hover:bg-blue-700 focus:outline-none md:fixed md:block"
			on:click={goToTop}
			aria-label="Go to top"
		>
			<i class="fa-solid fa-arrow-up"></i>
		</button>
	{/if}
	<Footer />
</div>
