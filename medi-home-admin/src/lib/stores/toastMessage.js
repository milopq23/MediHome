// src/lib/stores/toastMessage.js
import { writable } from 'svelte/store';

function createToastStore() {
	const { subscribe, update } = writable([]);

	function add({ message, type = 'success', timeout = 3000 }) {
		const id = crypto.randomUUID?.() || Date.now();
		const toast = { id, message, type };

		update((toasts) => [...toasts, toast]);

		// Xoá sau khi hết timeout
		setTimeout(() => {
			update((toasts) => toasts.filter((t) => t.id !== id));
		}, timeout);
	}

	function remove(id) {
		update((toasts) => toasts.filter((t) => t.id !== id));
	}

	return { subscribe, add, remove };
}

export const toasts = createToastStore();
