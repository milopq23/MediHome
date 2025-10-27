const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({ params, fetch }) {
	const { id } = params;
	console.log(id)
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/${id}`);
		const medicine = await res.json();
		console.log(medicine)
		return medicine;
		
	} catch (error) {}
}
