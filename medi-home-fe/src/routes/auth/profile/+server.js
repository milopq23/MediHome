const API_URL = import.meta.env.VITE_GO_PORT;

export async function GET({ locals }) {
	try {
	} catch (error) {}
}

export async function PATCH({ locals, request }) {
	try {
		const data = await request.json();
		const user_id = locals.user?.user_id;

		if (!user_id) {
			return new Response('Unauthorized', { status: 401 });
		}

		const res = await fetch(`${API_URL}/api/user/profile/`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				Cookie: request.headers.get('cookie') || ''
			},
			body: JSON.stringify(data)
		});

		if (!res.ok) {
			const errorText = await res.text();
			return new Response('Lỗi cập nhật', { status: res.status });
		}

		const updated = await res.json();

		return new Response(
			JSON.stringify({
				success: true,
				message: 'Cập nhật thành công!',
				user: updated.user
			}),
			{ headers: { 'Content-Type': 'application/json' } }
		);
	} catch (err) {
		console.error('PATCH error:', err);
		return new Response('Lỗi server', { status: 500 });
	}
}

export async function POST({ locals, request }) {
	try {
		const res = await fetch(`${API_URL}/api/logout/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Cookie: request.headers.get('cookie') || ''
			},
			body: JSON.stringify(data)
		});
	} catch (error) {
		console.log('Lỗi đăng xuất');
	}
}
