const API_URL = import.meta.env.VITE_GO_PORT;

// export async function GET({  fetch }) {
// 	try {
// 		const res = await fetch(`${API_URL}/api/admin/medicinecate/parent`);

// 		if (!res.ok) throw new Error(`Load cate parent lỗi: ${res.status}`);

// 		const result = await res.json();
// 		console.log(result);
// 		return json(result, { status: 200 });
// 	} catch (error) {
// 		console.error('Lỗi add medicine:', error);
// 		return json({ error: error.message }, { status: 500 });
// 	}
// }

// POST: tạo thuốc mới
export async function POST({ request }) {
	try {
		

		const medicine = await request.json();
		const res = await fetch(`${API_URL}/api/admin/medicine`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(medicine)
		});

		if (!res.ok) throw new Error(`Failed to add medicine: ${res.status}`);

		const result = await res.json();
		return json(result, { status: 201 });
	} catch (error) {
		console.error('Lỗi add medicine:', error);
		return json({ error: error.message }, { status: 500 });
	}
}

// PATCH: cập nhật thuốc
export async function PATCH({ request }) {
	try {
		const medicine = await request.json();
		if (!medicine.medicine_id) throw new Error('medicine_id is required');

		const res = await fetch(`${API_URL}/api/admin/medicine/${medicine.medicine_id}`, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(medicine)
		});

		if (!res.ok) throw new Error(`Failed to patch medicine: ${res.status}`);

		const result = await res.json();
		return json(result);
	} catch (error) {
		console.error('Lỗi patch medicine:', error);
		return json({ error: error.message }, { status: 500 });
	}
}
