const API_URL = import.meta.env.VITE_GO_PORT;

export async function GetCartUser(user_id) {
	try {
		const cartRes = await fetch(`${API_URL}/api/cart/${user_id}`);
		const cart = await cartRes.json();
		return cart;
	} catch (error) {
		console.log('Failed GET Cart' + error);
	}
}

export async function UpdateQuantityOrTypeCart(cartitem_id, updateItem) {
	try {
		const cartItemRes = await fetch(`${API_URL}/api/cart/${cartitem_id}`, {
			method: 'PATCH',
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

export async function DeleteCartItem(cartitem_id) {
	try {
		const res = await fetch(`${API_URL}/api/cart/${cartitem_id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const cartitem = await res.json();
		console.log(cartitem);
		return cartitem;
	} catch (error) {
		console.log('Failed Delete Cart Item' + error);
	}
}

export async function GetShipping(shipping_id) {
	try {
		const res = await fetch(`${API_URL}/api/shipping/${shipping_id}`);
		const shipping = await res.json();
		console.log(shipping);
		return shipping;
	} catch (error) {
		console.log('Failed GET Shipping' + error);
	}
}

export async function GetVoucher(total) {
	try {
		const res = await fetch(`${API_URL}/api/voucher/?total=${total}`);
		const voucher = await res.json();
		console.log(voucher);
		return voucher;
	} catch (error) {
		console.log('Failed GET Shipping' + error);
	}
}
