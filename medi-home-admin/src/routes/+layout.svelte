<script>
	import '../app.css';
</script>

<div class="flex min-h-screen">
  <!-- Sidebar -->
  <nav class="w-64 bg-gray-100 border-r border-gray-300 p-6">
    <h2 class="text-lg font-semibold mb-4">Admin</h2>
    <ul class="space-y-2">
      <li><a href="/medicine" class="block hover:underline">Thuốc</a></li>
      <li><a href="/medicine-cate" class="block hover:underline">Loại thuốc</a></li>
      <li><a href="/voucher" class="block hover:underline">Voucher</a></li>
      <li><a href="/dosage-form" class="block hover:underline">Dạng bào chế</a></li>
      <li><a href="/admin" class="block hover:underline">Quản trị viên</a></li>
      <li><a href="/user" class="block hover:underline">Người dùng</a></li>
      <li><a href="/order" class="block hover:underline">Đơn hàng</a></li>
      <li><a href="/logout" class="block hover:text-red-500">Đăng xuất</a></li>
    </ul>
  </nav>

  <!-- Main content -->
  <main class="flex-1 px-20 py-10">
    <slot />
  </main>
</div>