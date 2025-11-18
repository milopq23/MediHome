const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListShipping() {
	try {
		const res = await fetch(`${API_URL}/api/shipping/`);
		const shipping = await res.json();
		return shipping;
	} catch (error) {
		console.log('Lỗi Get Voucher' + error);
	}
}

export async function DetailShipping(shipping_id) {
	try {
		const res = await fetch(`${API_URL}/api/shipping/${shipping_id}`);
		const shipping = await res.json();
		return shipping;
	} catch (error) {
		console.log('Lỗi Get Voucher' + error);
	}
}

export async function CreateShipping(shipping) {
	try {
		const res = await fetch(`${API_URL}/api/shipping/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(shipping)
		});
		const newShipping = res.json();
		return newShipping;
	} catch (error) {
		console.log('Lỗi tạo shipping' + error);
	}
}

export async function UpdateShipping(shipping_id, updatedData) {
	try {
		const res = await fetch(`${API_URL}/api/shipping/${shipping_id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(updatedData)
		});
		const shipping = res.json();
		return shipping;
	} catch (error) {
		console.log('Lỗi update shipping' + error);
	}
}

export async function DeleteShipping(shipping_id) {
	try {
		const res = await fetch(`${API_URL}/api/shipping/${shipping_id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const shipping = res.json();
		return shipping;
	} catch (error) {
		console.log('Lỗi xoá shipping' + error);
	}
}
