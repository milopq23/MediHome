const API_URL = import.meta.env.VITE_GO_PORT;


export async function load({ fetch }) {
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
		} catch (error) {
			console.log('Lỗi load dosage form:', error);
		}
	};

	const [medicineCate, dosageForm] = await Promise.all([fetchMedicineCate(), fetchDosageForm()]);

	return {
		medicineCate,
		dosageForm
	};
}
