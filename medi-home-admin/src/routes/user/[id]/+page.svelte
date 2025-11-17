<script>
	import { pageTitle } from '$lib/stores/store.js';
	import { addToast } from '$lib/stores/toastMessage.js';
	import { page } from '$app/stores';

	let id = null;
	$: user_id = parseInt($page.params.id, 10);
	let title = 'Chi tiết người dùng';
	pageTitle.set(title);
</script>

<div class="flex items-center justify-center">
	<div class="max-h-full w-full overflow-y-auto rounded-lg bg-white p-6 shadow-lg dark:bg-gray-800">
		<div class="hidden md:flex md:justify-center md:py-10">
			<h1 class="text-5xl font-extrabold">{title}</h1>
		</div>

		<div>
			<form>
				<div>
					<div class="">
						<label for="id" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
							>ID</label
						>
						<input
							type="text"
							class="border-gray-300F block w-full rounded-lg border bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
							placeholder="MED001"
							bind:value={user.user_id}
							required
						/>
					</div>
				</div>
			</form>
		</div>
	</div>
</div>
