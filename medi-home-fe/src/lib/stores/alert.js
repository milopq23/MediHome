import { writable } from "svelte/store";

export const alerts = writable([]);

export function addAlert(message, type = "success", timeout = 3000) {
  const id = Date.now();
  alerts.update(all => [...all, { id, message, type }]);

  // tự động xóa sau timeout
  setTimeout(() => {
    alerts.update(all => all.filter(a => a.id !== id));
  }, timeout);
}