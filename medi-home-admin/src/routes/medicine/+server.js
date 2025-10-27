import { json } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;



// DELETE: xoá thuốc
export async function DELETE({ url }) {
	try {
		const id = url.searchParams.get('id');
		if (!id) throw new Error('id is required');

		const res = await fetch(`${API_URL}/api/admin/medicine/${id}`, {
			method: 'DELETE'
		});

		if (!res.ok) throw new Error(`Failed to delete medicine: ${res.status}`);

		const result = await res.json();
		return json(result);
	} catch (error) {
		console.error('Lỗi delete medicine:', error);
		return json({ error: error.message }, { status: 500 });
	}
}
