// const API_URL = import.meta.env.VITE_GO_PORT; // chắc chắn đây là full URL, ví dụ: http://localhost:3000

// export async function handle({ event, resolve }) {
// 	try {
// 		// Lấy token từ cookie
// 		const token = event.cookies.get('token');

// 		if (token) {

// 			const res = await fetch(`${API_URL}/api/profile/`, {
// 				headers: {
// 					'Content-Type': 'application/json',
// 					'Cookie': `token=${token}`
// 				}
// 			});

// 			console.log('Profile fetch status:', res.status);

// 			if (res.ok) {
// 				const user = await res.json();
// 				event.locals.user = user;
// 			} else {
// 				// const text = await res.text();
// 				event.locals.user = null;
// 			}
// 		} else {
// 			console.log('No token found in cookies.');
// 			event.locals.user = null;
// 		}
// 	} catch (error) {
// 		console.error('Error fetching profile:', error);
// 		event.locals.user = null;
// 	}


// 	return resolve(event);
// }
