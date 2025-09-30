package model

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int64 `json:"total"`
	Data       interface{} `json:"data"`
}

func 	NewPagination(page, pageSize int) *Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}	
	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}	