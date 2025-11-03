const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({ fetch, params }) {
	const { id } = params;
	console.log('Fetching user with ID:', id);
	try {
		const res = await fetch(`${API_URL}/api/admin/user/${id}`); 
        console.log('Response status:', res.status);
        console.log(res)

		if (!res.ok) throw new Error(`Failed to load user: ${res.status}`);
		const user = await res.json();
		console.log('Loaded user:', user);
		return user;
	} catch (error) {
		console.error('Lỗi load user:', error);
		return { user: null, error: error.message };
	}
}
