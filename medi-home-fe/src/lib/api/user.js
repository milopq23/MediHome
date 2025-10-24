import { json } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;

export async function apiLogin(email, password) {
	try {
		console.log(email);
		const res = await fetch(`${API_URL}/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				email, // phải là chuỗi
				password
			}) // phải là chuỗi
		});

		if (!res.ok) {
			return 'Không nhận req';
		}

		const data = await res.json(); // Lấy token + user
		return data;
	} catch (error) {
		console.log('Lỗi token', error);
	}
}

// export async function apiProfile() {
// 	try {
// 		const res = await fetch(`${API_URL}/api/user/profile`, {
// 			headers: {
// 				'Content-Type': 'application/json'
// 			},
// 			credentials: 'include'
// 		});
// 		if (!res.ok) {
// 			return 'Không nhận req';
// 		}
// 		console.log('res', res);
// 		const data = await res.json();
// 		console.log('data', data);
// 		return data;
// 	} catch (error) {
// 		console.log('Lỗi fetch profile', data);
// 	}
// }
