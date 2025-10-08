// user.js
const API_URL = import.meta.env.VITE_GO_PORT;

export async function loadUsers(page = 1, pageSize = 10) {
	try {
		const userRes = await fetch(`${API_URL}/api/user/?page=${page}&page_size=${pageSize}`);
		const user = await userRes.json();
		// if(!userRes.ok){

		// }

		const userActiveRes = await fetch(`${API_URL}/api/user/total`);
		const activeUser = await userActiveRes.json();

		console.log(user)
		console.log(activeUser)

	

		return {
			users: user.data,
			page: user.page,
			pageSize: user.page_size,
			total: user.total,
			activeUser
		};
	} catch (error) {
		console.error('Lỗi khi gọi danh sách user:', error);
		throw error;
	}
}

export async function addUser(user) {
	try {
		console.log(user);
		const res = await fetch(`${API_URL}/api/user/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
				// 'Authorization': 'Bearer <token>' nếu có auth
			},
			body: JSON.stringify(user)
		});
		console.log(res);
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}

		if (!res.ok) {
			const errorText = await res.text();
			throw new Error(`Lỗi khi tạo user: ${res.status} - ${errorText}`);
		}

		const result = await res.json();
		return result;
	} catch (error) {
		console.error('API POST user lỗi:', err);
		throw err;
	}
}

export async function editUser(id, user) {
	try {
		console.log('Dữ liệu gửi đi:', user);

		const res = await fetch(`${API_URL}/api/user/${id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
				// 'Authorization': 'Bearer <token>' nếu cần
			},
			body: JSON.stringify(user)
		});

		console.log('Phản hồi từ server:', res);

		if (!res.ok) {
			const errorText = await res.text();
			throw new Error(`Lỗi khi cập nhật user: ${res.status} - ${errorText}`);
		}

		const result = await res.json();
		return result;
	} catch (error) {
		console.error('API PUT user lỗi:', error);
		throw error;
	}
}

export async function deleteUser(id) {
	try {
		const res = await fetch(`${API_URL}/api/user/${id}`,{
			method:'DELETE'

		})
		if (!res.ok) {
			throw new Error(`HTTP error! status: ${res.status}`);
		}
		const result = await res.json();
		return result
	} catch (error) {
		console.error('API DELETE user lỗi:', error);
		throw error;
	}
}