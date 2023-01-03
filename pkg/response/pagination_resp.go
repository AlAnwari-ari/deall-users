package response

import (
	"math"

	"github.com/deall-users/internal/model"
)

type PaginationResponse struct {
	Results  interface{} `json:"results"`
	Page     int         `json:"page,omitempty"`
	Limit    int         `json:"size,omitempty"`
	Total    int         `json:"total"`
	NextPage *bool       `json:"next_page,omitempty"`
	PrevPage *bool       `json:"prev_page,omitempty"`
}

func NewPaginationResponse(result interface{}, pagination model.Pagination, total int, needPagination bool) *PaginationResponse {
	paginationRes := &PaginationResponse{
		Results: result,
		Total:   total,
	}

	page := pagination.Page
	limit := pagination.Limit

	if needPagination {
		next := page < int(math.Ceil(float64(total)/float64(limit)))
		prev := page > 1
		paginationRes.Page = page
		paginationRes.Limit = limit
		paginationRes.NextPage = &next
		paginationRes.PrevPage = &prev
	}

	return paginationRes
}
