package handler

import "medi-home-be/internal/app/service"

type OrderDetailHandler struct {
	service service.OrderDetailService
}

func NewOrderDetailHandler(service service.OrderDetailService) *OrderDetailHandler {
	return &OrderDetailHandler{service}
}

