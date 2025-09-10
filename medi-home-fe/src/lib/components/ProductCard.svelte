<!-- src/lib/components/ProductCard.svelte -->
<script>
  export let drugName;
  export let image;
  export let price;
  export let unit; // đơn vị tính VD: /hộp, /viên
  export let packaging; // quy cách đóng gói VD: 10 vỉ x 10 viên
</script>

<div class="bg-white p-4 rounded shadow text-center min-h-[230px] min-w[180px]">
  <img src={image} alt={drugName} class="w-full aspect-square object-cover mx-auto rounded mb-2" />
  <p class="text-black line-clamp-2 font-bold">{drugName}</p>
  <!-- <p class="text-sm text-gray-700 mb-1">{dosage}</p>
  <p class="text-sm text-gray-500 mb-1 italic">{instructions}</p> -->
  <p class="text-blue-600 ">
    <span class="font-bold">{price}</span>
    <span>{unit}</span>
  </p>
  <p class="bg-gray-100 px-4 m-2 rounded-2xl text-gray-500 text-sm truncate">{packaging}</p>
  <button class="bg-blue-600 rounded-3xl py-1 px-4 w-full">Mua</button>  
</div>


