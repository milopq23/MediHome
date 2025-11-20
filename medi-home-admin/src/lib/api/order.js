const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListOrder() {
	try {
		const res = await fetch(`${API_URL}/api/order/`);
		const order = await res.json();
		console.log(order);
		return order;
	} catch (error) {
		console.log('Lỗi Get Order' + error);
	}
}

export async function ListOrderStatus(status) {
	try {
		const res = await fetch(`${API_URL}/api/order/?status=${status}`, {
			headers: {
				'Content-Type': 'application/json'
			}
		});
		console.log(res);
		const order = await res.json();
		console.log(order);
		return order;
	} catch (error) {
		console.log('Lỗi Get Order Status' + error);
	}
}

export async function GetOrderDetail(order_id) {
	try {
		const res = await fetch(`${API_URL}/api/order/${order_id}`);
		const order = await res.json();
		console.log(order);
		return order;
	} catch (error) {
		console.log('Lỗi Get Order Detail' + error);
	}
}

export async function UpdateStatusOrder(order_id, newStatus) {
	try {
		const res = await fetch(`${API_URL}/api/order/approve/${order_id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ status: newStatus })
		});
		console.log(res);
		const order = await res.json();
		console.log(order);
		return order;
	} catch (error) {
		console.log('Lỗi tạo voucher' + error);
	}
}

// export async function UpdateVoucher(voucher_id, updatedData) {
// 	try {
// 		const res = await fetch(`${API_URL}/api/voucher/${voucher_id}`, {
// 			method: 'PATCH',
// 			headers: {
// 				'Content-Type': 'application/json'
// 			},
// 			body: JSON.stringify(updatedData)
// 		});
// 		const voucher = await res.json();
// 		console.log(voucher);
// 		return voucher;
// 	} catch (error) {
// 		console.log('Lỗi update voucher' + error);
// 	}
// }

// export async function DeleteVoucher(voucher_id) {
// 	try {
// 		const res = await fetch(`${API_URL}/api/voucher/${voucher_id}`, {
// 			method: 'DELETE',
// 			headers: {
// 				'Content-Type': 'application/json'
// 			}
// 		});
// 		const voucher = res.json();
// 		return voucher;
// 	} catch (error) {
// 		console.log('Lỗi xoá voucher' + error);
// 	}
// }
