const API_URL = import.meta.env.VITE_GO_PORT;

export async function listMedicine(page = 1, pageSize = 10) {
    try {
        const medicineRes =  await fetch(`${API_URL}/api/medicine/?page=${page}&page_size=${pageSize}`)
        const medicine = await medicineRes.json();
        return {
            medicines: medicine.data,
            page: medicine.page,    
			pageSize: medicine.page_size,
			total: medicine.total,  
        }
    } catch (error) {
        console.log("Lỗi gọi thuốc:", error)
    }
}
