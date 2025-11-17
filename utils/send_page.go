package utils

import "net/http"

type PaginatedData struct {
	Meta Pagination `json:"meta"`
	Data any        `json:"data"`
}

type Pagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
	TotalItems int64 `json:"totalItems"`
}

func SendPage(w http.ResponseWriter, data any, page,limit,count int64){
	paginatedData := PaginatedData{
		Meta:Pagination{
			Page:       page,
			Limit:      limit,
			TotalPages: count / limit,
			TotalItems: count,
		},
		Data: data,
	}
	SendData(w,http.StatusOK,paginatedData)
}
