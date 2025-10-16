<script>
	import { Plus } from 'lucide-svelte';

	let previewUrl = '';
</script>

<!-- {#if showPopup} -->
<div class="flex items-center justify-center p-4">
	<div
		class="max-h-full w-full max-w-md overflow-y-auto rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800"
	>
		<form>
			<div class="mb-6 grid gap-6 md:grid-cols-2">
				<div>
					<label for="code" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Code</label
					>
					<input
						type="text"
						class="border-gray-300F block w-full rounded-lg border bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="MED001"
						required
					/>
				</div>
				<div>
					<label
						for="medicine_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Tên thuốc</label
					>
					<input
						type="text"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Paracetamol"
						required
					/>
				</div>
				<!-- này làm select -->
				<div>
					<label for="category" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Danh mục</label
					>
					<input
						type="text"
						id="company"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Thuốc dị ứng"
						required
					/>
				</div>
				<div>
					<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Thuốc kê đơn</label
					>
					<input
						type=""
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Có"
						pattern="[0-9]{3}-[0-9]{2}-[0-9]{3}"
						required
					/>
				</div>
				<div>
					<label for="website" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Đóng gói</label
					>
					<input
						type="url"
						id="website"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Có"
						required
					/>
				</div>
				<div>
					<label for="visitors" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Dạng bào chế</label
					>
					<input
						type="number"
						id="visitors"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder=""
						required
					/>
				</div>
			</div>

			<div>
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Hình chính:
				</label>
				<div class="flex justify-center">
					{#if previewUrl}
						<img
							src={previewUrl}
							alt="Ảnh preview"
							class="max-h-[200px] max-w-[200px] cursor-pointer rounded object-contain"
						/>
					{:else}
						<div
							class="flex h-[100px] w-[100px] items-center justify-center rounded-xl border-solid bg-gray-300"
						>
							<Plus class="h-5 w-5" />
						</div>
					{/if}
					<input class="hidden" type="file" accept="image/*" on:change={onFileChange} />
				</div>
			</div>
			<div>
				<label for="" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Hình chính:
				</label>
				<div class="flex justify-center">
					{#if previewUrl}
						<img
							src={previewUrl}
							alt="Ảnh preview"
							class="max-h-[200px] max-w-[200px] cursor-pointer rounded object-contain"
						/>
					{:else}
						<div
							class="flex h-[100px] w-[100px] items-center justify-center rounded-xl border-solid bg-gray-300"
						>
							<Plus class="h-5 w-5" />
						</div>
					{/if}
					<input class="hidden" type="file" multiple accept="image/*" on:change={onFileChange} />
				</div>
			</div>
			<div class="mb-6">
				<label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Email address</label
				>
				<input
					type="email"
					id="email"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="john.doe@company.com"
					required
				/>
			</div>
			<div class="mb-6">
				<label for="password" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Password</label
				>
				<input
					type="password"
					id="password"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="•••••••••"
					required
				/>
			</div>
			<div class="mb-6">
				<label
					for="confirm_password"
					class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
					>Confirm password</label
				>
				<input
					type="password"
					id="confirm_password"
					class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
					placeholder="•••••••••"
					required
				/>
			</div>
			<div class="mb-6 flex items-start">
				<div class="flex h-5 items-center">
					<input
						id="remember"
						type="checkbox"
						value=""
						class="h-4 w-4 rounded-sm border border-gray-300 bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600"
						required
					/>
				</div>
				<label for="remember" class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
					>I agree with the <a href="#" class="text-blue-600 hover:underline dark:text-blue-500"
						>terms and conditions</a
					>.</label
				>
			</div>
			<button
				type="submit"
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 focus:outline-none sm:w-auto dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Submit</button
			>
		</form>
	</div>
</div>
