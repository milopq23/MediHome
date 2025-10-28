const API_URL = import.meta.env.VITE_GO_PORT;

export async function apiGetAllCate() {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicinecate/`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		console.log('rés', res);
		const result = await res.json();
		console.log(result);
		return result;
	} catch (error) {
		console.error('Lỗi khi gọi danh sách loại thuốc:', error);
		throw error;
	}
}

export async function apiCreateCate(category) {
	try {
		console.log('Dữ liệu POST', category);
		const res = await fetch(`${API_URL}/api/medicine-cate/children`, {
			method: 'POST',
			body: JSON.stringify(category)
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		console.log('Dữ liệu POST END', category);
		return result;
	} catch (error) {
		throw error;
	}
}

export async function apiCreateCateChild(category) {
	try {
		console.log('Dữ liệu POST', category);
		const res = await fetch(`${API_URL}/api/medicine-cate/children`, {
			method: 'POST',
			body: JSON.stringify(category)
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		console.log('Dữ liệu POST END', category);
		return result;
	} catch (error) {
		throw error;
	}
}

export async function apiPatchCate(id, category) {
	try {
		console.log('Dữ liệu PATCH', category);
		const res = await fetch(`${API_URL}/api/medicine-cate/${id}`, {
			method: 'PATCH',
			body: JSON.stringify(category)
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		console.log('Dữ liệu PATCH END', result);
		return result;
	} catch (error) {
		throw error;
	}
}

export async function apiDeleteCate(id) {
	try {
		const res = await fetch(`${API_URL}/api/medicine-cate/${id}`, {
			method: 'DELETE'
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		console.log('Dữ liệu Delete END', result);
		return result;
	} catch (error) {
		throw error;
	}
}

export async function apiUploadIcon(fileName) {
	const formData = new FormData();
	formData.append('file', fileName);
	formData.append('folder', 'medicine-cate');
	try {
		const res = await fetch(`${API_URL}/api/upload`, {
			method: 'POST',
			body: formData
		});
		if (!res.ok) {
			console.log('Lỗi post ảnh');
		}
		console.log('Post ảnh thành công', res);
		return await res.json();
	} catch (error) {
		throw error;
	}
	F;
}
