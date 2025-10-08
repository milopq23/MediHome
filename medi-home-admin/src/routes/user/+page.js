const API_URL = import.meta.env.VITE_GO_PORT;

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, url }) {
  const page = Number(url.searchParams.get('page') ?? 1);
  const pageSize = 10;

  const resUser = await fetch(`${API_URL}/api/user/?page=${page}&page_size=${pageSize}`);
  const userData = await resUser.json();

  const resTotal = await fetch(`${API_URL}/api/user/total`);
  const totalActiveUser = await resTotal.json();

  return {
    users: userData.data ?? [],
    page: userData.page ?? page,
    pageSize: userData.page_Size ?? pageSize,
    total: userData.total ?? 0,
    totalActiveUser
  };
}
