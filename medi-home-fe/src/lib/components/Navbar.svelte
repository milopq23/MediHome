
<script>
  let isOpen = false;
  // let userName ;
  let userName = "Nguyễn Huỳnh Phú Quý aaaaaaaaaaa";
</script>

<nav class="bg-[#E3F2FD] shadow-lg px-4 py-3 rounded-b-lg ">
  <div class="max-w-7xl mx-auto flex items-center justify-between">
    
    <!-- Logo --> 
    <a href="/" class="text-2xl font-bold">
      <span class="text-[#2F80ED]">Medi</span><span class="text-[#09a94c]">Home</span>
    </a>

    <!-- Search Box (Center) -->
    <div class="hidden md:flex flex-1 justify-center px-4">
      <div class="w-full  max-w-md flex items-center gap-2 px-4 py-2 bg-blue-100  rounded-lg focus-within:ring-2 focus-within:ring-blue-400">
        <i class="fa-solid fa-magnifying-glass text-gray-500 "></i>
        <input
          type="text"
          placeholder="Tìm kiếm sản phẩm..."
          class="w-full bg-transparent focus:outline-none rounded-lg  border-gray-300"
        />
      </div>
    </div>

    <!-- Hamburger Icon (Mobile) -->
    <button class="md:hidden px-4" on:click={() => isOpen = !isOpen}  aria-label="Menu">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" 
           viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" 
              stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button>

    <!-- Menu Items -->
    <div class={`flex-col md:flex md:flex-row md:items-center gap-4 
                ${isOpen ? 'flex' : 'hidden'} md:flex`}>
      <!-- <a href="/" class="text-gray-700 hover:text-blue-600">Trang chủ</a> -->
      <!-- <a href="/about" class="text-gray-700 hover:text-blue-600">Giới thiệu</a>
      <a href="/contact" class="text-gray-700 hover:text-blue-600">Liên hệ</a> -->
      <a href="/cart" class="flex items-center gap-1 text-gray-700 hover:text-blue-600" aria-label="Giỏ hàng">
        <i class="fa-solid fa-cart-shopping text-xl text-gray-500 hover:text-blue-600 transition-colors duration-200 cursor-pointer"></i>
      </a>

      <!-- Login / User Info -->
       <div class="max-w-[180px]">
        {#if userName}
          <div class="flex items-center gap-2 text-gray-700 w-full hover:text-blue-600 transition-colors duration-200 cursor-pointer">
            <a href="/profile" class="truncate overflow-hidden whitespace-nowrap hover:underline">
            <!-- <i class="fa-solid fa-user"></i> -->
              {userName}
            </a>
          </div>
        {:else}
          <a href="/login" class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 w-full inline-block text-center">
            Đăng nhập/Đăng ký
          </a>
        {/if}
       <div>

      
    </div>
  </div>

  <!-- Sidebar Menu (Mobile) -->
  <div class={`fixed top-0 right-0 h-full w-60 bg-[#E3F2FD] shadow-lg transform 
              transition-transform duration-300 z-40 p-6 flex flex-col gap-4
              ${isOpen ? 'translate-0-x' : 'translate-x-full'} md:hidden`}>
    <button class="self-end mb-4" on:click={() => isOpen = false}>
      ❌
    </button>
    <a href="/" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Trang chủ</a>
    <a href="/about" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Giới thiệu</a>
    <a href="/contact" class="text-gray-800 hover:text-blue-600" on:click={() => isOpen = false}>Liên hệ</a>
    <a href="/login" class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700" on:click={() => isOpen = false}>Đăng nhập/Đăng ký</a>
  </div>

  <!-- Black overlay khi sidebar mở -->
  {#if isOpen}
    <button class="fixed inset-0 backdrop-blur-sm bg-black/20 z-30 md:hidden" aria-label="Menu" on:click={() => isOpen = false}></button>
  {/if}

</nav>
