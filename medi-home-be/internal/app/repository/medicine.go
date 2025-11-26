package repository

import (
	"medi-home-be/config"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/model"
	"medi-home-be/pkg/helper"
)

type MedicineRepository interface {
	//Admin
	FindAll(page, pageSize int) (model.Pagination, error)
	FindByID(id int64) (model.Medicine, error)
	Create(medicine model.Medicine) (model.Medicine, error)
	Patch(medicine model.Medicine) (model.Medicine, error)
	Delete(id int64) error

	AddImages(medicine_id int64, urls []string) error
	GetMedicineImage(medicine_id int64) ([]model.Image, error)

	//User
	ListMedicineUser(page, pageSize int) (model.Pagination, error)
	DetailMedicineUser(medicine_id int64) (DetailMedicine, error)
	ListMedicineWithCate(page, pageSize int, medicinecate_id int64) (model.Pagination, error)
}

type medicineRepository struct{}

func NewMedicineRepository() MedicineRepository {
	return &medicineRepository{}
}

func (r *medicineRepository) FindAll(page, pageSize int) (model.Pagination, error) {
	var medicines []dto.AdminListMedicineDTO
	var total int64

	pagination := model.NewPagination(page, pageSize)

	// Đếm tổng bản ghi (phải dùng Table tương tự như truy vấn bên dưới)
	err := config.DB.
		Table("medicines AS m").
		Joins("LEFT JOIN medicinecates AS c ON c.medicinecate_id = m.medicinecate_id").
		Joins("LEFT JOIN dosageforms AS d ON d.dosageform_id = m.dosageform_id").
		Count(&total).Error

	if err != nil {
		return model.Pagination{}, err
	}
	pagination.Total = total

	// Truy vấn dữ liệu có phân trang
	err = config.DB.
		Table("medicines AS m").
		Select(`
			m.medicine_id,
			m.code,
			m.name,
			m.thumbnail,
			c.name AS category_name,
			d.name AS dosage_form_name
		`).
		Joins("LEFT JOIN medicinecates AS c ON c.medicinecate_id = m.medicinecate_id").
		Joins("LEFT JOIN dosageforms AS d ON d.dosageform_id = m.dosageform_id").
		Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Order("m.medicine_id ASC").
		Scan(&medicines).Error

	if err != nil {
		return model.Pagination{}, err
	}

	pagination.Data = medicines
	return *pagination, nil
}

func (r *medicineRepository) FindByID(id int64) (model.Medicine, error) {
	var medicine model.Medicine
	err := config.DB.Find(&medicine, id).Error
	return medicine, err
}

func (r *medicineRepository) Create(medicine model.Medicine) (model.Medicine, error) {
	err := config.DB.Create(&medicine).Error
	return medicine, err
}

func (r *medicineRepository) AddImages(medicine_id int64, urls []string) error {

	for _, url := range urls {
		img := model.Image{
			MedicineID: medicine_id,
			ImageURL:   url,
		}
		if err := config.DB.Create(&img).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *medicineRepository) Patch(medicine model.Medicine) (model.Medicine, error) {
	err := config.DB.Save(&medicine).Error
	return medicine, err
}

func (r *medicineRepository) Delete(id int64) error {
	err := config.DB.Delete(&model.Medicine{}, id).Error
	return err
}

//User

func (r *medicineRepository) ListMedicineUser(page, pageSize int) (model.Pagination, error) {
	var medicines []dto.UserListMedicineDTO
	var total int64

	pagination := model.NewPagination(page, pageSize)

	err := config.DB.
		Table("medicines AS m").
		Joins("JOIN batchsellings bs ON bs.medicine_id = m.medicine_id").
		Joins("JOIN inventories i ON i.inventory_id = bs.inventory_id").
		Group("m.medicine_id").
		Count(&total).Error

	if err != nil {
		return model.Pagination{}, err
	}
	pagination.Total = total

	err = config.DB.
		Table("medicines AS m").
		Select(`
			m.medicine_id,
			m.name,
			m.thumbnail,
			m.package,
			ROUND(i.import_price * (1 + i.mark_up/100)) AS price
		`).
		Joins("JOIN batchsellings bs ON bs.medicine_id = m.medicine_id").
		Joins("JOIN inventories i ON i.inventory_id = bs.inventory_id").
		// Group("m.medicine_id").
		Order("m.medicine_id ASC").
		Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Scan(&medicines).Error

	if err != nil {
		return model.Pagination{}, err
	}

	pagination.Data = medicines
	return *pagination, nil
}

func (r *medicineRepository) ListMedicineWithCate(page, pageSize int, medicinecate_id int64) (model.Pagination, error) {
	var medicines []dto.UserListMedicineDTO
	var total int64

	pagination := model.NewPagination(page, pageSize)

	// 1️⃣ Lấy danh sách category (bao gồm cả subcategories)
	subQuery := config.DB.
		Table("medicinecates").
		Select("medicinecate_id").
		Where("parent_id = ? OR medicinecate_id = ?", medicinecate_id, medicinecate_id)

	// 2️⃣ Đếm tổng thuốc (unique theo medicine_id)
	err := config.DB.
		Table("medicines AS m").
		Joins("LEFT JOIN batchsellings bs ON bs.medicine_id = m.medicine_id").
		Joins("LEFT JOIN inventories i ON i.inventory_id = bs.inventory_id").
		Where("m.medicinecate_id IN (?)", subQuery).
		// Group("m.medicine_id").
		Count(&total).Error
	if err != nil {
		return model.Pagination{}, err
	}
	pagination.Total = total

	// 3️⃣ Lấy dữ liệu
	err = config.DB.
		Table("medicines AS m").
		Select(`
			m.medicine_id,
			m.name,
			m.thumbnail,
			m.package,
			ROUND(i.import_price * (1 + i.mark_up/100)) AS price
		`).
		Joins("LEFT JOIN batchsellings bs ON bs.medicine_id = m.medicine_id").
		Joins("LEFT JOIN inventories i ON i.inventory_id = bs.inventory_id").
		Where("m.medicinecate_id IN (?)", subQuery).
		Group("m.medicine_id").
		Order("m.medicine_id ASC").
		Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Scan(&medicines).Error

	if err != nil {
		return model.Pagination{}, err
	}

	pagination.Data = medicines
	return *pagination, nil
}

// func (r *medicineRepository) DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error) {
// 	var medicine dto.UserDetailMedicineDTO
// 	err := config.DB.Table("mv_detail_medicine").
// 		Where("medicine_id = ?", id).
// 		First(&medicine).Error // lấy lỗi từ .Error

// 	if err != nil {
// 		return dto.UserDetailMedicineDTO{}, err // trả về struct rỗng nếu lỗi
// 	}
// 	return medicine, nil
// }

type DetailMedicine struct {
	MedicineID uint   `json:"medicine_id" gorm:"column:medicine_id"`
	Code       string `json:"code" `
	Name       string `json:"name" `
	Thumbnail  string `json:"thumbnail" `
	// Images           []string `json:"images"`
	UnitPerStrip     int64   `json:"unit_per_strip"`
	UnitPerBox       int64   `json:"unit_per_box"`
	MedCategoryName  string  `json:"medcatename" gorm:"column:medcatename"`
	DosageFormName   string  `json:"dosagename" gorm:"column:dosagename"`
	PriceForStrip    float64 `json:"price_for_strip" gorm:"column:price_for_strip"`
	PriceForBox      float64 `json:"price_for_box" gorm:"column:price_for_box"`
	Prescription     bool    `json:"prescription" `
	Usage            string  `json:"usage"`
	Package          string  `json:"package"`
	Indication       string  `json:"indication"`
	Adverse          string  `json:"adverse"`
	Contraindication string  `json:"contraindication"`
	Precaution       string  `json:"precaution"`
	Ability          string  `json:"ability"`
	Pregnancy        string  `json:"pregnancy"`
	DrugInteraction  string  `json:"drug_interaction"`
	Storage          string  `json:"storage" `
	Manufacturer     string  `json:"manufacturer" `
	Note             string  `json:"note" `
}

func (r *medicineRepository) DetailMedicineUser(medicine_id int64) (DetailMedicine, error) {
	var medicine DetailMedicine
	query := `
		select m.medicine_id, m.code, m.name, m.thumbnail, m.unit_per_strip, m.unit_per_box,
			round((i.import_price * (1 + i.mark_up/100))/m.unit_per_strip) as price_for_strip,
			round((i.import_price * (1 + i.mark_up/100))) as price_for_box,
			mc.name as medcatename, df.name as dosagename, m.prescription, m.usage,
			m.package, m.indication, m.adverse, m.contraindication, m.precaution, m.ability, m.pregnancy,
			m.drug_interaction, m.storage, m.manufacturer, m.note
		from medicines m
		join medicinecates mc on mc.medicinecate_id = m.medicinecate_id
		join dosageforms df on df.dosageform_id = m.dosageform_id
		join batchsellings bs on bs.medicine_id = m.medicine_id
		join inventories i on i.inventory_id = bs.inventory_id
		where m.medicine_id = ?
		
	`
	err := config.DB.Raw(query, medicine_id).Scan(&medicine).Error
	return medicine, err
}

func (r *medicineRepository) GetMedicineImage(medicine_id int64) ([]model.Image, error) {
	var image []model.Image
	err := config.DB.Where("medicine_id = ?", medicine_id).Find(&image).Error
	return image, err
}

// func (r *medicineRepository) DetailMedicineWithPrice(id int64) (model.DetailMedicineVM, error) {
// 	var medicine model.DetailMedicineVM
// 	err := config.DB.Table("medicine_detail_view").
// 		Where("medicine_id = ?", id).
// 		First(&medicine).Error

// 	if err != nil {
// 		return model.DetailMedicineVM{}, err
// 	}
// 	log.Print(medicine)
// 	return medicine, nil
// }
