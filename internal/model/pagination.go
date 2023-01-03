package model

type Pagination struct {
	Page  int `form:"page" json:"page" binding:"omitempty,max=20"`
	Limit int `form:"limit" json:"limit" binding:"omitempty,max=1000"`
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
