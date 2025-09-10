<script>
  import {sidebarOpen} from '../index.js';
  
  // let userName ;
  let userName = "Nguyễn Huỳnh Phú Quý aaaaaaaaaaa";



</script>
<div>

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
</div>

