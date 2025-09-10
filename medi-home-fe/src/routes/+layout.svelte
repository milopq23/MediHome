<script>
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/SideBar.svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import SideBar from '$lib/components/SideBar.svelte';

	let y = 0;
	let innerWidth = 0;
	let innerHeight = 0;

	function goToTop() {
		document.body.scrollIntoView();
	}
</script>

<div class="min-h-screen bg-gray-100">

	<Header />

  	<!-- <SideBar /> -->

	<main class="max-w-3xl mx-auto p-4">
		<slot />
	</main>

	<Footer />
</div>




