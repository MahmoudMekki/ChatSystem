package models

type Paginator struct {
	Limit   int
	Page    int
	KeyWord string
}

func (p *Paginator) GetLimit() int {
	return p.Limit
}
func (p *Paginator) GetOffset() int {
	return (p.Page - 1) * p.Limit
}
func (p *Paginator) GetKeyWord() string {
	return p.KeyWord
}
