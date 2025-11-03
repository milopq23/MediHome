import { json } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;

// DELETE: xoá thuốc
export async function DELETE({ request }) {
	try {
		const { id } = await request.json();
		console.log(id);
		if (!id) throw new Error('id is required');

		const res = await fetch(`${API_URL}/api/admin/medicine/${id}`, {
			method: 'DELETE',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ id })
		});

		if (!res.ok) throw new Error(`Failed to delete medicine: ${res.status}`);

		const result = await res.json();
		return json(result, { status: 200 });
	} catch (error) {
		console.error('Lỗi delete medicine:', error);
		return json({ error: error.message }, { status: 500 });
	}
}
