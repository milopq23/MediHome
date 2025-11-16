const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListMedicines(page = 1, pageSize = 15) {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/?page=${page}&page_size=${pageSize}`);
		const medicine = await res.json();
		return {
			medicines: medicine.data,
			page: medicine.page,
			pageSize: medicine.page_size,
			total: medicine.total
		};
	} catch (error) {
		console.log('Lỗi gọi thuốc:', error);
	}
}

export async function DetailMedicine(medicine_id) {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/${medicine_id}`);
		const medicine = await res.json();
		console.log(medicine);
		return medicine;
	} catch (error) {
		console.log('Lỗi gọi thuốc:', error);
	}
}

export async function GetMedicineCate() {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicinecate/`);
		const medicineCate = await res.json();
		console.log(medicineCate);
		return medicineCate;
	} catch (error) {
		console.log('Lỗi gọi thuốc:', error);
	}
}

export async function GetDosage() {
	try {
		const res = await fetch(`${API_URL}/api/admin/dosage/`);
		const dosage = await res.json();
		console.log(dosage);
		return dosage;
	} catch (error) {
		console.log('Lỗi gọi thuốc:', error);
	}
}

export async function AddMedicine(medicine) {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine`, {
			medthod: 'POST',
			headers: {},
			body: JSON.stringify(medicine)
		});
		console.log('Log add medicine', res);
		const result = await res.json();
		console.log('Lỗi result add', result);
		return result;
	} catch (error) {
		console.log('Lỗi add user', error);
	}
}

export async function UpdateMedicine(medicine_id, medicine) {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/${medicine_id}`, {
			medthod: 'PATCH',
			body: JSON.stringify(medicine)
		});
		console.log('Log patch medicine', res);
		const result = await res.json();
		console.log('Lỗi result patch', result);
		return result;
	} catch (error) {
		console.log('Lỗi patch user', error);
	}
}

export async function DeleteMedicine(id) {
	try {
		const res = await fetch(`${API_URL}/api/admin/medicine/${id}`, {
			method: 'DELETE'
		});
		const result = await res.json();
		return result;
	} catch (error) {
		console.log('Lỗi delete medicine', error);
	}
}

export async function apiUploadMedicine(file) {
	const formData = new FormData();
	formData.append('file', file);
	formData.append('folder', 'medicine');
	try {
		const res = await fetch(`${API_URL}/api/upload`, {
			method: 'POST',
			body: formData
		});
		const result = await res.json();
		return result;
	} catch (error) {
		console.log('Upload ảnh thuốc', error);
	}
}
