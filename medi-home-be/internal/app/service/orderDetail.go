package service

import "medi-home-be/internal/app/repository"

type OrderDetailService interface{

}

type orderDetailService struct{
	repo repository.OrderDetailRepository
}

func NewOrderDetailService(repo repository.OrderDetailRepository) OrderDetailService{
	return &orderDetailService{repo}
}

