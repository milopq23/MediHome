const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({ fetch, url }) {
	const page = Number(url.searchParams.get('page') ?? 1);
	// const page = Number(url.searchParams.get('page') ?? 1);
	const pageSize = 15;

	const fetchUser = async () => {
		try {
			const res = await fetch(`${API_URL}/api/admin/user/?page=${page}&page_size=${pageSize}`);
			const user = await res.json();
			return user;
		} catch (error) {
			console.log('Lỗi load list user:', error);
		}
	};
	// const resUser = await fetch(`${API_URL}/api/admin/user/?page=${page}&page_size=${pageSize}`);
	// const userData = await resUser.json();

	const fetchTotalUserActive = async () => {
		try {
			const res = await fetch(`${API_URL}/api/admin/user/total`);
			const total = await res.json();
			return total;
		} catch (error) {
			console.log('Lỗi load user active:', error);
		}
	};

	// const resTotal = await fetch(`${API_URL}/api/admin/user/total`);
	// const totalActiveUser = await resTotal.json();

	const [user, total] = await Promise.all([fetchUser(), fetchTotalUserActive()]);

	return {
		users: user.data ?? [],
		selectedPage: user.page ?? page, // thêm dòng này
		pageSize: user.page_Size ?? pageSize,
		total: user.total ?? 0,
		totalActive: total
	};
}
