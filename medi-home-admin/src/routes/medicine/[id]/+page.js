const API_URL = import.meta.env.VITE_GO_PORT;

export async function load({ params, fetch }) {
	const { id } = params;
	const fetchMedicine = async () => {
		try {
			const res = await fetch(`${API_URL}/api/admin/medicine/${id}`);
			const medicine = await res.json();
			return medicine;
		} catch (error) {
			console.log('Lỗi load medicine:', error);
		}
	};
	const fetchMedicineCate = async () => {
		try {
			const res = await fetch(`${API_URL}/api/admin/medicinecate/`);
			const medicineCate = await res.json();
			return medicineCate;
		} catch (error) {
			console.log('Lỗi load medicine cate parent:', error);
		}
	};
	const fetchDosageForm = async () => {
		try {
			const res = await fetch(`${API_URL}/api/admin/dosage/`);
			const dosageForm = await res.json();
			return dosageForm;
		}
		catch (error) {
			console.log('Lỗi load dosage form:', error);
		}
	};

	const [medicine, medicineCate, dosageForm] = await Promise.all([
		fetchMedicine(),
		fetchMedicineCate(),
		fetchDosageForm()
	]);

	let selectedParent = null;
	let selectedChild = medicine.medicinecate_id;

	for (const parent of medicineCate) {
		const child = parent.children?.find((c) => c.medicinecate_id === medicine.medicinecate_id);
		if (child) {
			selectedParent = parent.medicinecate_id;
			selectedChild = child.medicinecate_id;
			break;
		}
	}

	return {
		medicine,
		medicineCate,
		dosageForm,
		selectedParent,
		selectedChild
	};
}
