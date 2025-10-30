package repository

import (
	"log"
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

	//User
	ListMedicineUser(page, pageSize int) (model.Pagination, error)
	DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error)
	DetailMedicineWithPrice(id int64) (model.DetailMedicineVM, error)
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

	// Đếm tổng bản ghi (phải dùng Table tương tự như truy vấn bên dưới)
	err := config.DB.
		Table("medicines AS m").
		// Joins("LEFT JOIN medicinecates AS c ON c.medicinecate_id = m.medicinecate_id").
		// Joins("LEFT JOIN dosageforms AS d ON d.dosageform_id = m.dosageform_id").
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
			m.name,
			m.thumbnail,
			m.package
		`).
		Scopes(helper.Paginate(pagination.Page, pagination.PageSize)).
		Order("m.medicine_id ASC").
		Scan(&medicines).Error

	if err != nil {
		return model.Pagination{}, err
	}

	pagination.Data = medicines
	return *pagination, nil
}

func (r *medicineRepository) DetailMedicine(id int64) (dto.UserDetailMedicineDTO, error) {
	var medicine dto.UserDetailMedicineDTO
	err := config.DB.Table("mv_detail_medicine").
		Where("medicine_id = ?", id).
		First(&medicine).Error // lấy lỗi từ .Error

	if err != nil {
		return dto.UserDetailMedicineDTO{}, err // trả về struct rỗng nếu lỗi
	}
	return medicine, nil
}

// func (r *medicineRepository) DetailMedicinePrice(id int64) (model.Medicine, error) {
// 	var medicine model.Medicine
// 	err := config.DB.Where("medicine_id = ?", id).Find(&medicine).Error
// 	return medicine, err
// }

func (r *medicineRepository) priceMedicineInventory(id int64) (dto.UserDetailMedicineDTO, error) {
	var medicine dto.UserDetailMedicineDTO
	err := config.DB.Table("mv_detail_medicine").
		Where("medicine_id = ?", id).
		First(&medicine).Error // lấy lỗi từ .Error

	if err != nil {
		return dto.UserDetailMedicineDTO{}, err // trả về struct rỗng nếu lỗi
	}
	return medicine, nil
}

func (r *medicineRepository) DetailMedicineWithPrice(id int64) (model.DetailMedicineVM, error) {
	var medicine model.DetailMedicineVM
	err := config.DB.Table("medicine_detail_view").
		Where("medicine_id = ?", id).
		First(&medicine).Error 
	
	if err != nil {
		return model.DetailMedicineVM{}, err 
	}	
	log.Print(medicine)
	return medicine, nil
}

