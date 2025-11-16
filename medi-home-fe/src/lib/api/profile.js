const API_URL = import.meta.env.VITE_GO_PORT;

export async function GetProfile(user_id) {
    try {
        const res = await fetch(`${API_URL}/api/profile/${user_id}`);
        const user = await res.json();
        return user;
    } catch (error) {
        console.log('Failed GET Profile' + error);
    }
}