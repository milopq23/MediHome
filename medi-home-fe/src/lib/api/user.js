// import { json } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;

export async function GetProfile() {
	try {
		const res = await fetch(`${API_URL}/api/profile/`, {
			credentials: 'include'
		});
		const user = await res.json();
		return user;
	} catch (error) {
		console.log('Failed GET Profile' + error);
	}
}

export async function GetAllOrder(status) {
	try {
		let url = `${API_URL}/api/order/user/`;

		if (status != null && String(status).trim() !== '') {
			url += `?status=${encodeURIComponent(String(status).trim())}`;
		}

		const res = await fetch(url, {
			credentials: 'include'
		});
		console.log(url);
		const order = await res.json();
		return order;
	} catch (error) {
		console.log('Failed GET Order' + error);
	}
}

export async function GetDetailOrder(order_id) {
	try {
		const res = await fetch(`${API_URL}/api/order/${order_id}`);
		const order = await res.json();
		return order;
	} catch (error) {
		console.log('Failed GET OrderDetail' + error);
	}
}

export async function Login(email, password) {
	try {
		const res = await fetch(`${API_URL}/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify({
				email,
				password
			})
		});
		console.log(res);
		const data = await res.json();
		console.log(data);
		return data;
	} catch (error) {
		console.log('Lỗi token', error);
	}
}

export async function Logout() {
	try {
		const res = await fetch(`${API_URL}/api/logout`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include'
		});
		const data = await res.json();
		return data;
	} catch (error) {
		console.log('Lỗi LogOut', +error);
	}
}

// export aysnc function GetLogin(){
// 	try {
// 		return
// 	} catch (error) {

// 	}
// }

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
