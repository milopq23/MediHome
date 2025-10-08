const API_URL = import.meta.env.VITE_GO_PORT;

export async function getAllCate(){
    try {
        const res = await fetch(`${API_URL}/api/medicine-cate/`)
        if(!res.ok){
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const result = await res.json();
        console.log(result)
        return result
    } catch (error) {
        console.error('Lỗi khi gọi danh sách loại thuốc:', error);
		throw error;
    }
}

export async function createCate(category){
    try {
        console.log("Dữ liệu POST", category)
        const res = await fetch(`${API_URL}/api/medicine-cate/children`,{
            method:'POST',
            body:JSON.stringify(category)
        })
        if(!res.ok){
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const result = await res.json()
        console.log("Dữ liệu POST END", category)
        return result
    } catch (error) {
        throw error
    }
}

export async function createCateChild(category){
     try {
        console.log("Dữ liệu POST", category)
        const res = await fetch(`${API_URL}/api/medicine-cate/children`,{
            method:'POST',
            body:JSON.stringify(category)
        })
        if(!res.ok){
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const result = await res.json()
        console.log("Dữ liệu POST END", category)
        return result
    } catch (error) {
        throw error
    }
}

export async function patchCate(id,category){
    try {
        console.log("Dữ liệu PATCH", category)
        const res = await fetch(`${API_URL}/api/medicine-cate/${id}`,{
            method:'PATCH',
            body: JSON.stringify(category)
        })
        if(!res.ok){
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const result = await res.json()
        console.log("Dữ liệu PATCH END", result)
        return result
    } catch (error) {
        throw error
    }
}

export async function deleteCate(id){
    try {
        const res = await fetch(`${API_URL}/api/medicine-cate/${id}`,{
            method:'DELETE',
        })
        if(!res.ok){
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const result = await res.json()
        console.log("Dữ liệu Delete END", result)
        return result
    } catch (error) {
        throw error
    }
}