const API_URL = import.meta.env.VITE_GO_PORT;

export async function ListDosage() {
	try {
		const res = await fetch(`${API_URL}/api/dosage/`);
		const dosage = await res.json();
		return dosage;
	} catch (error) {
		console.log('Lỗi Get Dosage' + error);
	}
}

export async function DetailDosage(dosage_id) {
	try {
		const res = await fetch(`${API_URL}/api/dosage/${dosage_id}`);
		const dosage = await res.json();
		return dosage;
	} catch (error) {
		console.log('Lỗi Get Dosage' + error);
	}
}

export async function CreateDosage(dosage) {
	try {
		const res = await fetch(`${API_URL}/api/dosage/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(dosage)
		});
		const newdosage = res.json();
		return newdosage;
	} catch (error) {
		console.log('Lỗi tạo dosage' + error);
	}
}

export async function UpdateDosage(dosage_id, updatedData) {
	try {
		const res = await fetch(`${API_URL}/api/dosage/${dosage_id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(updatedData)
		});
		const dosage = res.json();
		return dosage;
	} catch (error) {
		console.log('Lỗi update dosage' + error);
	}
}

export async function DeleteDosage(dosage_id) {
	try {
		const res = await fetch(`${API_URL}/api/dosage/${dosage_id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		const dosage = res.json();
		return dosage;
	} catch (error) {
		console.log('Lỗi xoá dosage' + error);
	}
}
