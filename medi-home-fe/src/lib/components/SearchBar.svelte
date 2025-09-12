<script>
    let showSearch = false;
</script>

<div class="md:flex justify-center px-4">
  <div class="flex items-center gap-2 px-4 py-2 bg-blue-100 rounded-lg focus-within:ring-2 focus-within:ring-blue-400 md:flex w-1/2">
    <i class="fa-solid fa-magnifying-glass text-gray-500"></i>
    <input
      type="text"
      placeholder="Tìm kiếm sản phẩm..."
      class="w-full bg-transparent focus:outline-none rounded-lg border-gray-300"
    />
  </div>

    <!-- Mobile search toggle button -->
  {#if !showSearch}
    <button
      class=" flex justify-end md:hidden text-gray-700 text-xl focus:outline-none"
      on:click={() => showSearch = !showSearch}
      aria-label="Toggle Search"
    >
      <i class="fa-solid fa-magnifying-glass"></i>
    </button>
  {/if}

    <!-- Mobile search input (show/hide) -->
  {#if showSearch}
    <div class="flex items-center gap-2 px-4 py-2 bg-blue-100 rounded-lg focus-within:ring-2 focus-within:ring-blue-400 md:flex w-1/2">
      <i class="fa-solid fa-magnifying-glass text-gray-500"></i>
      <input
        type="text"
        placeholder="Tìm kiếm sản phẩm..."
        class="w-full bg-transparent focus:outline-none rounded-lg border-gray-300"
      />
    </div>
  {/if}
</div>