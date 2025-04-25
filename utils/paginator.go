package utils

import (
	"math"
)

type Paginator struct {
	BackPage     int `json:"backPage"`
	CurrentPage  int `json:"currentPage"`
	LimitPerPage int `json:"limitPerPage"`
	NextPage     int `json:"nextPage"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

func NewPaginator(currentPage, limit, totalRecords int) Paginator {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	backPage := currentPage - 1
	if backPage < 1 {
		backPage = 1
	}

	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	return Paginator{
		BackPage:     backPage,
		CurrentPage:  currentPage,
		LimitPerPage: limit,
		NextPage:     nextPage,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}
}
