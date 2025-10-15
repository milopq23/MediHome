const API_URL = import.meta.env.VITE_GO_PORT

export async function loadMedicines(page = 1, pageSize = 10){
    try {
        const medicineRes =  await fetch(`${API_URL}/api/admin/medicine/?page=${page}&page_size=${pageSize}`)
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

export async function addMedicine(medicine) {
    try {
        const res =  await fetch(`${API_URL}/api/admin/medicine`,{
            medthod:'POST',
            headers:{

            },
            body: JSON.stringify(medicine)
        })
        console.log('Log add medicine',res)
        const result =  await res.json();
        console.log('Lỗi result add',result)
        return result
    } catch (error) {
        console.log("Lỗi add user", error)
    }
}

export async function patchMedicine(medicine){
    try {
        const res =  await fetch(`${API_URL}/api/admin/medicine`,{
            medthod:'PATCH',
            headers:{

            },
            body: JSON.stringify(medicine)
        })
        console.log('Log patch medicine',res)
        const result =  await res.json();
        console.log('Lỗi result patch',result)
        return result
    } catch (error) {
        console.log("Lỗi patch user", error)
    }
} 

export async function  apiDeleteMedicne(id) {
    try {
        const res =  await fetch(`${API_URL}/api/admin/medicine/${id}`,{
            method:'DELETE'
        })
        const result =  await res.json();
        // console.log('Lỗi result delete',result)
        return result
    } catch (error) {
        console.log("Lỗi delete medicine", error)
    }
}