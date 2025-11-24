import { error } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListInventory() {
	try {
		const res = await fetch(`${API_URL}/api/inventory/`);
		const inventory = await res.json();
		return inventory;
	} catch (error) {
		console.log('Lỗi Get Inventory' + error);
	}
}

export async function FindMedicine(medicine_id) {
	try {
		const res = await fetch(`${API_URL}/api/inventory/${medicine_id}`);
		const inventory = await res.json();
		return inventory;
	} catch (error) {
		console.log('Lỗi Get Medicine Inveentory' + error);
	}
}

export async function UpdateBatchSelling(medicine_id, inventory_id) {
	try {
		const res = await fetch(`${API_URL}/api/inventory/${medicine_id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ inventory_id })
		});
		console.log(medicine_id);
		console.log(inventory_id);
		const batchSelling = res.json();
		return batchSelling;
	} catch {
		console.log('Lỗi Update BatchSeling' + error);
	}
}
