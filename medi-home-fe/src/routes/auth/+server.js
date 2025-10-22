// src/routes/auth/login/+server.ts
import { json } from '@sveltejs/kit';

const API_URL = import.meta.env.VITE_GO_PORT;

export async function POST({ request, cookies, fetch }) {
	const { email, password } = await request.json();

	// Gửi thông tin đến Go backend
	const res = await fetch(`${API_URL}/api/login`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ email, password })
	});

	if (!res.ok) {
		return json({ error: 'Sai email hoặc mật khẩu' }, { status: 401 });
	}

	const { token, user } = await res.json();

	// Set cookie HttpOnly
	cookies.set('token', token, {
		httpOnly: true,
		sameSite: 'strict',
		secure: false, // ⚠️ true nếu dùng HTTPS
		path: '/',
		maxAge: 60 * 60 * 24
	});

	return json({ message: 'Đăng nhập thành công', user });
}
