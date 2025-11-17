// admin.js
const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListAdmin(page = 1, pageSize = 10) {
	try {
		const response = await fetch(`${API_URL}/api/admin/?page=${page}&page_size=${pageSize}`);
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const result = await response.json();
		return {
			admins: result.data,
			page: result.page,
			pageSize: result.page_size,
			total: result.total
		};
	} catch (error) {
		console.error('Lỗi khi gọi danh sách admin:', error);
		throw error;
	}
}

export async function getTotalRole(role) {
	try {
		const response = await fetch(`${API_URL}/api/admin/total?role=${role}`);

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}

		const result = await response.json();
		return result;
	} catch {
		console.error('Lỗi khi gọi tổng số admin hoạt động:', error);
		throw error;
	}
}

export async function addAdmin(admin) {
	try {
		const res = await fetch(`${API_URL}/api/admin/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
				// 'Authorization': 'Bearer <token>' nếu có auth
			},
			body: JSON.stringify(admin)
		});
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		return result;
	} catch (error) {
		console.error('API POST admin lỗi:', err);
		throw err;
	}
}

export async function editAdmin(id, admin) {
	try {
		const res = await fetch(`${API_URL}/api/admin/${id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
				// 'Authorization': 'Bearer <token>' nếu cần
			},
			body: JSON.stringify(admin)
		});
		if (!res.ok) {
			const errorText = await res.text();
			throw new Error(`Lỗi khi cập nhật admin: ${res.status} - ${errorText}`);
		}
		const result = await res.json();
		return result;
	} catch (error) {
		console.error('API PATCH admin lỗi:', error);
		throw error;
	}
}

export async function deleteAdmin(id) {
	try {
		const res = await fetch(`${API_URL}/api/admin/${id}`, {
			method: 'DELETE'
		});
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const result = await res.json();
		return result;
	} catch (error) {
		console.error('API DELETE admin lỗi:', error);
		throw error;
	}
}
