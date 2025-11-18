const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListVoucher() {
	try {
		const res = await fetch(`${API_URL}/api/voucher/`);
		const voucher = await res.json();
		console.log(voucher);
		return voucher;
	} catch (error) {
		console.log('Lỗi Get Voucher' + error);
	}
}

export async function DetailVoucher(voucher_id) {
	try {
		const res = await fetch(`${API_URL}/api/voucher/${voucher_id}`);
		const voucher = await res.json();
		console.log(voucher);
		return voucher;
	} catch (error) {
		console.log('Lỗi Get Detail Voucher' + error);
	}
}

export async function CreateVoucher(voucher) {
	try {
		const res = await fetch(`${API_URL}/api/voucher/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(voucher)
		});
		const newVoucher = await res.json();
		return newVoucher;
	} catch (error) {
		console.log('Lỗi tạo voucher' + error);
	}
}

export async function UpdateVoucher(voucher_id, updatedData) {
	try {
		const res = await fetch(`${API_URL}/api/voucher/${voucher_id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(updatedData)
		});
		const voucher = await res.json();
		console.log(voucher);
		return voucher;
	} catch (error) {
		console.log('Lỗi update voucher' + error);
	}
}

export async function DeleteVoucher(voucher_id) {
	try {
		const res = await fetch(`${API_URL}/api/voucher/${voucher_id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const voucher = res.json();
		return voucher;
	} catch (error) {
		console.log('Lỗi xoá voucher' + error);
	}
}
