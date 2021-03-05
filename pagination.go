package main

type Paginated struct {
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func DefaultPagination() Paginated {
	return Paginated{
		PageSize: 100,
	}
}
