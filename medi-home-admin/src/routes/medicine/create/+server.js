const API_URL = import.meta.env.VITE_GO_PORT;
import { json } from '@sveltejs/kit';

export async function POST({ request, fetch }) {
	try {
		const medicine = await request.json();
		const res = await fetch(`${API_URL}/api/admin/medicine`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(medicine)
		});
		console.log(medicine);

		if (!res.ok) throw new Error(`Failed to add medicine: ${res.status}`);

		const result = await res.json();
		return json(result, { status: 201 });
	} catch (error) {
		console.error('Lỗi add medicine:', error);
		return json({ error: error.message }, { status: 500 });
	}
}
