<script>

</script>


<!-- {#if showPopup} -->
	<div class="fixed inset-0 z-40 bg-black/50"></div>
	<div
		class="fixed top-1/2 left-1/2 z-60 mx-auto w-full max-w-md -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white p-6 shadow-lg"
	>
		<form>
			<div class="mb-6 grid gap-6 md:grid-cols-2">
				<div>
					<label
						for="first_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Code</label
					>
					<input
						type="text"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="John"
						required
					/>
				</div>
				<div>
					<label
						for="last_name"
						class="mb-2 block text-sm font-medium text-gray-900 dark:text-white">Tên thuốc</label
					>
					<input
						type="text"
						id="last_name"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Doe"
						required
					/>
				</div>
				<div>
					<label for="company" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Company</label
					>
					<input
						type="text"
						id="company"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="Flowbite"
						required
					/>
				</div>
				<div>
					<label for="phone" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Phone number</label
					>
					<input
						type="tel"
						id="phone"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="123-45-678"
						pattern="[0-9]{3}-[0-9]{2}-[0-9]{3}"
						required
					/>
				</div>
				<div>
					<label for="website" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Website URL</label
					>
					<input
						type="url"
						id="website"
						class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
						placeholder="flowbite.com"
						required
					/>
				</div>
				<div>
					<label for="visitors" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
						>Unique visitors (per month)</label
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