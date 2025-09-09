<script>
  let isOpen = false;
</script>

<nav class="bg-white shadow-md px-4 py-3">
  <div class="max-w-7xl mx-auto flex items-center justify-between">
    <!-- Logo -->
    <a href="/" class="text-2xl font-bold text-blue-600">MyApp</a>

    <!-- Hamburger Icon (Mobile) -->
    <button class="md:hidden px-4" on:click={() => isOpen = !isOpen}>
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" 
           viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" 
              stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button>

    <!-- Menu Items -->
    <div class={`flex-col md:flex md:flex-row md:items-center gap-4 
                ${isOpen ? 'flex' : 'hidden'} md:flex`}>
      <a href="/" class="text-gray-700 hover:text-blue-600">Trang chủ</a>
      <a href="/about" class="text-gray-700 hover:text-blue-600">Giới thiệu</a>
      <a href="/contact" class="text-gray-700 hover:text-blue-600">Liên hệ</a>
      <a href="/login" class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700">Đăng nhập</a>
    </div>
  </div>

  <!-- Sidebar Menu (Mobile) -->
  <div class={`fixed top-0 right-0 h-full w-60 bg-white shadow-lg transform 
              transition-transform duration-300 z-40 p-6 flex flex-col gap-4
              ${isOpen ? 'translate-x-0' : 'translate-x-full'} md:hidden`}>
    <button class="self-end mb-4" on:click={() => isOpen = false}>
      ❌
    </button>
    <a href="/" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Trang chủ</a>
    <a href="/about" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Giới thiệu</a>
    <a href="/contact" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Liên hệ</a>
    <a href="/login" class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700" on:click={() => isOpen = false}>Đăng nhập</a>
  </div>

  <!-- Black overlay khi sidebar mở -->
  {#if isOpen}
    <div class="fixed inset-0 backdrop-blur-sm bg-black/20 z-30 md:hidden" on:click={() => isOpen = false}></div>
  {/if}

</nav>
