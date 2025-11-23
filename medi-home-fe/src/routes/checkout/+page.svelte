<script>
</script>

<div>
	<!-- order -->
	<div></div>
	<!-- order summary -->
	<div></div>
</div>
<div class="bg-white px-20 py-10 sm:px-8">
	<div class="mx-auto max-md:max-w-xl">
		<div class="mb-16 flex items-start">
			<div class="w-full">
				<div class="flex w-full items-center">
					<div
						class=" flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-600 p-1.5"
					>
						<span class="text-sm font-semibold text-white">1</span>
					</div>
					<div class="mx-4 h-[3px] w-full rounded-lg bg-blue-600"></div>
				</div>
				<div class="mt-2 mr-4">
					<h6 class="text-sm font-semibold text-slate-900">Cart</h6>
				</div>
			</div>
			<div class="w-full">
				<div class="flex w-full items-center">
					<div
						class="mx-[-1px] flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-600 p-1.5"
					>
						<span class="text-sm font-semibold text-white">2</span>
					</div>
					<div class="mx-4 h-[3px] w-full rounded-lg bg-slate-300"></div>
				</div>
				<div class="mt-2 mr-4">
					<h6 class="text-sm font-semibold text-slate-900">Checkout</h6>
				</div>
			</div>
			<div>
				<div class="flex items-center">
					<div
						class="mx-[-1px] flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-slate-300 p-1.5"
					>
						<span class="text-sm font-semibold text-white">3</span>
					</div>
				</div>
				<div class="mt-2">
					<h6 class="text-sm font-semibold text-slate-400">Order</h6>
				</div>
			</div>
		</div>

		<div class="grid gap-x-8 gap-y-12 md:grid-cols-2 lg:grid-cols-3 lg:gap-x-12">
			<div class="lg:col-span-2">
				<form>
					<div>
						<h2 class="mb-6 text-xl font-semibold text-slate-900">Delivery Details</h2>
						<div class="grid gap-x-4 gap-y-6 lg:grid-cols-2">
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">First Name</label>
								<input
									type="text"
									placeholder="Enter First Name"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">Last Name</label>
								<input
									type="text"
									placeholder="Enter Last Name"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">Email</label>
								<input
									type="email"
									placeholder="Enter Email"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">Phone No.</label>
								<input
									type="number"
									placeholder="Enter Phone No."
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">Address Line</label>
								<input
									type="text"
									placeholder="Enter Address Line"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">City</label>
								<input
									type="text"
									placeholder="Enter City"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">State</label>
								<input
									type="text"
									placeholder="Enter State"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
							<div>
								<label class="mb-2 block text-sm font-medium text-slate-900">Zip Code</label>
								<input
									type="text"
									placeholder="Enter Zip Code"
									class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
								/>
							</div>
						</div>
					</div>

					<div class="mt-12">
						<h2 class="mb-6 text-xl font-semibold text-slate-900">Thanh toán</h2>
						<div class="grid gap-4 lg:grid-cols-2">
							<div class="max-w-sm rounded-md border border-gray-300 bg-gray-100 p-4">
								<div>
									<div class="flex items-center">
										<input
											type="radio"
											name="method"
											class="h-5 w-5 cursor-pointer"
											id="card"
											checked
										/>
										<label for="card" class="ml-4 flex cursor-pointer gap-2">
											<img
												src="https://cdn-icons-png.flaticon.com/512/2897/2897832.png"
												class="w-20 bg-gray-100"
												alt="COD"
											/>
										</label>
									</div>
								</div>
								<p class="mt-4 text-sm font-medium text-slate-500">Tiền mặt</p>
							</div>
							<div class="max-w-sm rounded-md border border-gray-300 bg-gray-100 p-4">
								<div>
									<div class="flex items-center">
										<input type="radio" name="method" class="h-5 w-5 cursor-pointer" id="paypal" />
										<label for="paypal" class="ml-4 flex cursor-pointer gap-2">
											<img
												src="https://cdn.haitrieu.com/wp-content/uploads/2022/10/Icon-VNPAY-QR.png"
												class="w-20 bg-gray-100"
												alt="VNPay"
											/>
										</label>
									</div>
								</div>
								<p class="mt-4 text-sm font-medium text-slate-500">VNPAY</p>
							</div>
						</div>
					</div>

					<div class="mt-12 max-w-md">
						<p class="mb-2 text-sm font-medium text-slate-900">Do you have a promo code?</p>
						<div class="flex gap-4">
							<input
								type="email"
								placeholder="Promo code"
								class="w-full rounded-md border border-gray-400 bg-white px-4 py-2.5 text-sm text-slate-900 focus:outline-blue-600"
							/>
							<button
								type="button"
								class="flex cursor-pointer items-center justify-center rounded-md bg-blue-600 px-4 py-2.5 text-sm font-medium tracking-wide text-white hover:bg-blue-700"
							>
								Apply
							</button>
						</div>
					</div>
				</form>
			</div>

			<div class="relative">
				<h2 class="mb-6 text-xl font-semibold text-slate-900">Order Summary</h2>
				<ul class="space-y-4 font-medium text-slate-500">
					<li class="flex flex-wrap gap-4 text-sm">
						Subtotal <span class="ml-auto font-semibold text-slate-900">$72.00</span>
					</li>
					<li class="flex flex-wrap gap-4 text-sm">
						Discount <span class="ml-auto font-semibold text-slate-900">$0.00</span>
					</li>
					<li class="flex flex-wrap gap-4 text-sm">
						Shipping <span class="ml-auto font-semibold text-slate-900">$6.00</span>
					</li>
					<li class="flex flex-wrap gap-4 text-sm">
						Tax <span class="ml-auto font-semibold text-slate-900">$5.00</span>
					</li>
					<hr class="border-slate-300" />
					<li class="flex flex-wrap gap-4 text-[15px] font-semibold text-slate-900">
						Total <span class="ml-auto">$83.00</span>
					</li>
				</ul>
				<div class="mt-8 space-y-4">
					<button
						type="button"
						class="w-full cursor-pointer rounded-md bg-blue-600 px-4 py-2.5 text-sm font-medium tracking-wide text-white hover:bg-blue-700"
						>Complete Purchase</button
					>
					<button
						type="button"
						class="w-full cursor-pointer rounded-md border border-gray-300 bg-gray-100 px-4 py-2.5 text-sm font-medium tracking-wide text-slate-900 hover:bg-gray-200"
						>Continue Shopping</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
