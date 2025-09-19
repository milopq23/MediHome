<script>
	let showPassword = false;
</script>

<div class="bg-snowblue flex min-h-screen items-center justify-center">
	<!-- Container -->
	<div
		class="w-full mx-3 md:mx-10 lg:mx-50 flex-1 rounded-2xl bg-white md:w-4/5 lg:w-2/3 shadow-md"
	>
		<h1 class="text-charcoal p-4 text-center text-2xl font-bold">Đăng nhập</h1>

		<div class="flex flex-col lg:flex-row items-stretch p-5">
			<!-- Form -->
			<div class="w-full lg:w-2/4  min-w-">
				<form class="space-y-4 p-10">
					<!-- Email -->
					<div>
						<label for="email" class="text-md text-charcoal mb-1 block px-2 font-medium">Email</label>
						<input
							type="email"
							placeholder="youremail@gmail.com"
							class="w-full rounded-md border border-gray-300 py-2 pl-4 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-400"
						/>
					</div>

					<!-- Password -->
					<div class="relative">
						<label for="password" class="text-md text-black-300 mb-1 block px-2 font-medium">Mật khẩu</label>
						<input
							type={showPassword ? 'text' : 'password'}
							placeholder="••••••••"
							class="w-full rounded-md border border-gray-300 py-2 pl-4 pr-8 focus:outline-none focus:ring-2 focus:ring-blue-400"
						/>
						<button
							type="button"
							class="absolute right-2 top-10 text-gray-500 focus:outline-none"
							on:click={() => (showPassword = !showPassword)}
							aria-label={showPassword ? 'Ẩn mật khẩu' : 'Hiện mật khẩu'}
						>
							<i
								class="fa-solid"
								class:fa-eye={!showPassword}
								class:fa-eye-slash={showPassword}
							></i>
						</button>
					</div>

					<!-- Submit button -->
					<button
						type="submit"
						class="w-full rounded-md bg-blue-500 py-2 text-white transition duration-300 hover:bg-blue-600"
					>
						Đăng nhập
					</button>
				</form>

				<!-- Link -->
				<p class="mt-4 pb-2 text-center text-sm text-gray-500">
					Chưa có tài khoản?
					<a href="/register" class="text-blue-500 hover:underline">Đăng ký</a>
				</p>

				<div class="m-4 grid grid-cols-3 items-center">
					<hr class="border-gray-400" />
					<h2 class="text-center">OR</h2>
					<hr class="border-gray-400" />
				</div>
			</div>

			<!-- Image -->
			<div class="hidden md:flex-1/4 lg:flex w-2/4 	 items-center justify-center p-4">
				<img
					src="https://img.pikbest.com/origin/09/29/90/98rpIkbEsTvK8.png!sw800"
					alt="Img medicine"
					class="rounded-2xl object-contain max-h-[400px] w-full"
				/>
			</div>
		</div>
	</div>
</div>
