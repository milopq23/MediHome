<script>
    import {sidebarOpen} from '../index.js';

    let userName;

    function toggleSidebar() {
        sidebarOpen.update(n => !n);
    }
</script>

<nav class="bg-[#E3F2FD] shadow-lg px-4 py-3 rounded-b-lg ">
  <div class="max-w-7xl mx-auto flex items-center justify-between">
    
    <!-- Logo --> 
    <a href="/" class="text-2xl font-bold">
      <span class="text-[#2F80ED]">Medi</span><span class="text-[#09a94c]">Home</span>
    </a>

    <!-- Search Box (Center) -->
    <div class="flex md:flex flex-1 justify-center px-4">
      <div class="w-full  max-w-md flex items-center gap-2 px-4 py-2 bg-blue-100  rounded-lg focus-within:ring-2 focus-within:ring-blue-400">
        <i class="fa-solid fa-magnifying-glass text-gray-500 "></i>
        <input
          type="text"
          placeholder="Tìm kiếm sản phẩm..."
          class="w-full bg-transparent focus:outline-none rounded-lg  border-gray-300"
        />
      </div>
      

    </div>

    
    <div class="flex md:flex md:flex-row md:items-center gap-4 ">
      <a href="/cart" class="flex items-center gap-1 text-gray-700 hover:text-blue-600" aria-label="Giỏ hàng">
        <i class="fa-solid fa-cart-shopping text-xl text-gray-500 hover:text-blue-600 transition-colors duration-200 cursor-pointer"></i>
      </a>

      <!-- <div>
        <a href="/login" class="flex items-center md:hidden text-gray-700 hover:text-blue-600 transition-colors duration-200 cursor-pointer" aria-label="Đăng nhập">
            <i class="fa-solid fa-user text-xl text-gray-500 hover:text-blue-600 transition-colors duration-200 cursor-pointer"></i>
          </a>
      </div> -->
      <!-- Login / User Info -->
      <div class="max-w-[190px] ">
        {#if userName}
          <div class="hidden md:flex items-center gap-2 text-gray-700 w-full hover:text-blue-600 transition-colors duration-200 cursor-pointer">
            <a href="/profile" class="truncate overflow-hidden whitespace-nowrap hover:underline">
            
              {userName}
            </a>
          </div>
          <a href="/profile" class="md:hidden items-center text-gray-700 hover:text-blue-600 transition-colors duration-200 cursor-pointer text-xl">
            <i class="fa-solid fa-user"></i>
          </a>
        {:else}
          <a href="/login" class="hidden bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 w-full md:inline-block text-center">
            <i class="fa-solid fa-user"></i>  
            Đăng nhập/Đăng ký
          </a>
          <div>
            <a href="/login" class="flex items-center md:hidden text-gray-700 hover:text-blue-600 transition-colors duration-200 cursor-pointer" aria-label="Đăng nhập">
                <i class="fa-solid fa-user text-xl text-gray-500 hover:text-blue-600 transition-colors duration-200 cursor-pointer"></i>
              </a>
          </div> 
          
        {/if}
      </div>
    </div>
    <!-- <button class="md:hidden px-4" on:click={toggleSidebar}  aria-label="Menu">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" 
            viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" 
              stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button> -->
</nav>