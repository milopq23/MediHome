<script>
	// import { apiLogin } from '$lib/api/user';
	import { goto } from '$app/navigation';
	import { Login } from '$lib/api/user';
	import { addAlert } from '$lib/stores/alert';

	let showPassword = false;
	let email = '';
	let password = '';
	// let error = '';
	// let success = '';

	async function handleSubmit() {
		try {
			const res = await Login(email, password);
			addAlert('Đăng nhập thành công', 'success');
			setTimeout(() => goto('/'), 3000);
		} catch (err) {
			addAlert('Đăng nhập thất bại', 'error');
			console.error(err);
		}
	}

	// async function handleSubmit(event) {
	// 	event.preventDefault();
	// 	// error = '';
	// 	// success = '';

	// 	// const res = await fetch('/auth/login', {
	// 	// 	method: 'POST',
	// 	// 	headers: { 'Content-Type': 'application/json' },
	// 	// 	body: JSON.stringify({ email, password })
	// 	// });
	// 	const data = await Login(email, password);

	// 	goto('/');
	// 	// try {
	// 	// 	data = await res.json();
	// 	// } catch (err) {
	// 	// 	console.error('Invalid JSON response', err);
	// 	// }

	// 	// if (res.ok) {
	// 	// 	success = `Đăng nhập thành công! Chào mừng`;
	// 	// 	addAlert(success, 'success');

	// 	// } else {
	// 	// 	error = `Lỗi: ${data.error || 'Đăng nhập thất bại'}`;
	// 	// 	addAlert(error, 'error');
	// 	// }
	// }
</script>

<div class="flex min-h-screen items-center justify-center bg-snowblue">
	<!-- Container -->
	<div
		class="mx-3 w-full flex-1 rounded-2xl bg-white shadow-md md:mx-10 md:w-4/5 lg:mx-50 lg:w-2/3"
	>
		<h1 class="p-4 text-center text-2xl font-bold text-charcoal">Đăng nhập</h1>

		<div class="flex flex-col items-stretch p-5 lg:flex-row">
			<!-- Form -->
			<div class="min-w- w-full lg:w-2/4">
				<form class="space-y-4 p-10" on:submit|preventDefault={handleSubmit}>
					<!-- Email -->
					<div>
						<label for="email" class="text-md mb-1 block px-2 font-medium text-charcoal"
							>Email</label
						>
						<input
							type="email"
							placeholder="youremail@gmail.com"
							class="w-full rounded-md border border-gray-300 py-2 pr-8 pl-4 focus:ring-2 focus:ring-blue-400 focus:outline-none"
							bind:value={email}
						/>
					</div>

					<!-- Password -->
					<div class="relative">
						<label for="password" class="text-md text-black-300 mb-1 block px-2 font-medium"
							>Mật khẩu</label
						>
						<input
							type={showPassword ? 'text' : 'password'}
							placeholder="••••••••"
							class="w-full rounded-md border border-gray-300 py-2 pr-8 pl-4 focus:ring-2 focus:ring-blue-400 focus:outline-none"
							bind:value={password}
						/>
						<button
							type="button"
							class="absolute top-10 right-2 text-gray-500 focus:outline-none"
							on:click={() => (showPassword = !showPassword)}
							aria-label={showPassword ? 'Ẩn mật khẩu' : 'Hiện mật khẩu'}
						>
							<i class="fa-solid" class:fa-eye={!showPassword} class:fa-eye-slash={showPassword}
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
			<div class="hidden w-2/4 items-center justify-center p-4 md:flex-1/4 lg:flex">
				<img
					src="https://img.pikbest.com/origin/09/29/90/98rpIkbEsTvK8.png!sw800"
					alt="Img medicine"
					class="max-h-[400px] w-full rounded-2xl object-contain"
				/>
			</div>
		</div>
	</div>
</div>
