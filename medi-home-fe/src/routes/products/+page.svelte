<!-- src/routes/products/+page.svelte -->
<script>
  import ProductCard from '$lib/components/ProductCard.svelte';

  let products = [
    {
      drugName: "Thuốc Exopadin 60mg Trường Thọ điều trị viêm mũi dị ứng, mày đay (3 vỉ x 10 viên)",
      price: "60.000đ",
      image: "https://cdn.nhathuoclongchau.com.vn/unsafe/768x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/IMG_0238_fc19904162.jpg",
      unit: "/hộp",
      packaging: "3 vỉ x 10 viên"
    },
    {
      drugName: "Thuốc Cetirizine Stada 10mg điều trị viêm mũi dị ứng, mày đay (10 vỉ x 10 viên)",
      price: "48.000đ",
      image: "https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/00033097_cetirizin_stada_10mg_10x10_6872_61c9_large_81d9335efe.jpg",
      unit: "/hộp",
      packaging: "10 vỉ x 10 viên"
    },
    {
      drugName: "Thuốc Vitamin C 500mg SPHARM phòng và trị thiếu vitamin C (10 vỉ x 10 viên)",
      price: "25.000đ",
      image: "https://cdn.nhathuoclongchau.com.vn/unsafe/375x0/filters:quality(90)/https://cms-prod.s3-sgn09.fptcloud.com/DSC_03587_e8e7e30119.jpg",
      unit: "/viên",
      packaging: "10 vỉ x 10 viên"
    }
    // {
    //   drugName: "Vitamin C",
    //   price: "15.000đ",
    //   image: "https://via.placeholder.com/100"
    // }
  ];
</script>

<h1 class="text-2xl font-bold mb-4">Sản phẩm</h1>

<div class="grid grid-cols-2 sm:grid-cols-3  lg:grid-cols-4 gap-4">
  {#each products as product}
    <ProductCard
      drugName={product.drugName}
      price={product.price}
      image={product.image}
      unit={product.unit}
      packaging={product.packaging}
    />
  {/each}
</div>

      <!-- dosage={product.dosage}
      instructions={product.instructions} -->
