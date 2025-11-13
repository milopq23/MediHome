const API_URL = import.meta.env.VITE_GO_PORT;

export async function GetCartUser(user_id) {
	try {
		const cartRes = await fetch(`${API_URL}/api/cart/${user_id}`);
		const cart = await cartRes.json();
		console.log('res', cartRes);
		console.log('cart', cart);
		return {
			cart
		};
	} catch (error) {
		console.log('Failed GET Cart' + error);
	}
}

export async function UpdateQuantityOrTypeCart(cartitem_id, updateItem) {
	try {
		const cartItemRes = await fetch(`${API_URL}/api/cart/${cartitem_id}`, {
			method: 'PATCH', // Hoặc PATCH tùy API của bạn
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(updateItem)
		});
		const cartItem = await cartItemRes.json();
		console.log(cartItem);
		return cartItem;
	} catch (error) {
		console.log('Failed Update Cart' + error);
	}
}
