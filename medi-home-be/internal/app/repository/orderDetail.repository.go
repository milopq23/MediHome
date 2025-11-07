package repository

type OrderDetailRepository interface {
}

type orderDetailRepository struct {
}

func NewOrderDetailRepository() OrderDetailRepository {
	return &orderDetailRepository{}
}
