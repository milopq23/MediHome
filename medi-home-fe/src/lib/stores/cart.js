import { writable, derived } from 'svelte/store';

export const cartItems = writable([]); // lúc đầu để rỗng, sẽ set từ server
export const totalAmount = derived(cartItems, ($cartItems) =>
	$cartItems.reduce((sum, item) => sum + item.price * item.quantity, 0)
);
