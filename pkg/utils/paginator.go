package utils

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (p *Pagination) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.Size
}

func (p *Pagination) Limit() int {
	if p.Size <= 0 {
		p.Size = 10
	}
	return p.Size
}
