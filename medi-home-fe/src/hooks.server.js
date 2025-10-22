
const API_URL = import.meta.env.VITE_GO_PORT;

export async function handle({ event, resolve }) {
	const token = event.cookies.get('token');

	if (token) {
		const res = await fetch(`${API_URL}/api/login`, {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});

		if (res.ok) {
			const user = await res.json();
			event.locals.user = user;
		} else {
			event.locals.user = null;
		}
	} else {
		event.locals.user = null;
	}

	return resolve(event);
}
