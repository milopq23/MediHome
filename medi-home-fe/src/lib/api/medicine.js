const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListMedicine(page = 1, pageSize = 10) {
	try {
		const medicineRes = await fetch(`${API_URL}/api/medicine/?page=${page}&page_size=${pageSize}`);
		const medicine = await medicineRes.json();
		return {
			medicines: medicine.data,
			page: medicine.page,
			pageSize: medicine.page_size,
			total: medicine.total
		};
	} catch (error) {
		console.log('Lỗi gọi thuốc:', error);
	}
}

export async function DetailMedicine(id) {
	try {
		const medicine = await fetch(`${API_URL}/api/medicine/${id}`);
		return await medicine.json();
	} catch (error) {
		console.log('Lỗi detail', error);
	}
}

export async function AddCart(medicine_id, select_type, quantity) {
	try {
		const res = await fetch(`${API_URL}/api/cart/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify({
				medicine_id,
				select_type,
				quantity
			})
		});
		console.log('res cart:', res);
		const cart = await res.json();
		return cart;
	} catch (error) {
		console.log('Lỗi ADD CART', error);
	}
}
