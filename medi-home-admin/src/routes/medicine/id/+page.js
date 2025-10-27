const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({params, fetch}){
    const {id} = params;
    try {
        const res = await fetch(`${API_URL}/api/admin/medicine/?${id}`)
        const medicine = await res.json();
        return medicine
    } catch (error) {
        
    }
}