package qibo

import (
	"math"
)

// Pagination obj
type Pagination struct {
	CurrentPage int32 `json:"current_page"`
	PageSize    int32 `json:"page_size"`
	TotalPage   int32 `json:"total_page"`
	TotalResult int32 `json:"total_result"`
}

// SetTotalPage calculate total page by total count data
func (p *Pagination) SetTotalPage(total int32) *Pagination {
	totalPages := int32(1)
	if p.PageSize > 0 {
		d := float64(total) / float64(p.PageSize)
		totalPages = int32(math.Ceil(d))
	}
	p.TotalPage = totalPages
	p.TotalResult = total
	return p
}
