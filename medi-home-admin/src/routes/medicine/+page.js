const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({ fetch, url }) {
	const page = Number(url.searchParams.get('page')) || 1;
	const pageSize = Number(url.searchParams.get('pageSize')) || 1;

	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/?page=${page}&page_size=${pageSize}`);

		if (!res.ok) {
			throw new Error(`Failed to fetch medicines: ${res.status}`);
		}

		const data = await res.json();

		return {
			medicines: data.data || [],
			page: data.page || page,
			pageSize: data.page_size || pageSize,
			total: data.total || 0
		};
	} catch (error) {
		console.error('Lỗi load medicines:', error);
		return {
			medicines: [],
			page,
			pageSize,
			total: 0
		};
	}
}
