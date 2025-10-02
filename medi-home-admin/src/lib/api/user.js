const API_BASE = 'http://localhost:8000';

export async function fetchUsersPaginated(page = 1, pageSize = 10) {
    console.log("xxxxxxx")
    const res = await fetch(`${API_BASE}/user`, {
    method: 'GET',
    credentials: 'include', // Nếu backend dùng cookie
    headers: {
      'Content-Type': 'application/json',
      "Accept":'*'
    }
  });
  console.log('xxx: ',res)

//   if (!res.ok) {
//     const message = await res.text();
//     throw new Error(`API error (${res.status}): ${message}`);
//   }

//   return await res.json();
return {}
}
//   try {
//     const response = await fetch(`http://localhost:8000/user`);

//     console.log("a",response)

//     if (!response.ok) {
//       const message = await response.text();
//       throw new Error(`Lỗi API (${response.status}): ${message}`);
//     }

//     const result = await response.json();

//     // Kiểm tra có đầy đủ dữ liệu không
//     if (!result.data || !Array.isArray(result.data)) {
//       throw new Error('Dữ liệu trả về không đúng định dạng');
//     }

//     return result;
//   } catch (error) {
//     console.error('Lỗi khi fetch users:', error);
//     throw error; // để onMount bên Svelte bắt được
//   }

