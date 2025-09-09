<script>
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';

	let y = 0;
	let innerWidth = 0;
	let innerHeight = 0;

	function goToTop() {
		document.body.scrollIntoView();
	}
</script>

<!-- <svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head> -->


<!-- <Navbar />

<div class="relative flex flex-col max-w-[1400px] min-h-screen mx-auto text-sm sm:text-base min-h-screen">
	<div class={"fixed bottom-0 w-full duration-200 flex p-10 z-[10]"+ (y > 0
			?'opactity-full pointer-events-auto'
			:'opacity-0 pointer-events-none')}>
		<button on:click={goToTop} class="ml-auto bg-blue-600 hover:bg-blue-700 text-white px-3 sm:px-4 curso-:pointer rounded-full shadow-lg">
			<i class="fa-solid fa-arrow-up grid place-items-center aspect-squre" />
		</button>
	</div>
	<Header {y} />
	<slot />
	<Foooter />

	
</div> -->
<!-- <svelte:window bind:scrollY={y} bind:innerWidth bind:innerHeight />	 -->


<div class="min-h-screen bg-gray-100">

  <!-- <Header /> -->

  <Navbar />
  <main class="max-w-3xl mx-auto p-4">
    <h2 class="text-2xl font-semibold mb-4">Trang chủ</h2>
    <p class="mb-2">
      Đây là phần nội dung chính. Kéo xuống để xem footer vẫn hiển thị.
    </p>
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse
      lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor...
    </p>
    <div class="h-[1500px] bg-gray-200 mt-6 rounded">Nội dung dài để test cuộn</div>
  </main>

  <Footer />
</div>




