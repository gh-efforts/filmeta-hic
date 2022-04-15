package paginator

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

const (
	PageDefaultSize    = 20
	PageDefaultNum     = 1
	PageDefaultMaxSize = 100
)

type (
	Page struct {
		Page        int64 `json:"paginator" form:"paginator" binding:"required"`
		PageSize    int64 `json:"pageSize" form:"pageSize"`
		PageMaxSize int64 `json:"-" form:"-"`
	}

	WithOption func(p *Page)
)

func (p *Page) With(options ...WithOption) {
	for _, opt := range options {
		opt(p)
	}
}

func WithPageMaxSize(size int64) WithOption {
	return func(p *Page) {
		if size == 0 {
			size = PageDefaultMaxSize
		}

		p.PageMaxSize = size
	}
}

func (p *Page) Offset() int64 {
	if p.PageSize == 0 {
		p.PageSize = PageDefaultSize
	}
	if p.Page < PageDefaultNum {
		p.Page = PageDefaultNum
	}
	return p.PageSize * (p.Page - 1)
}

func (p *Page) Limit() int64 {
	if p.PageSize == 0 {
		p.PageSize = PageDefaultSize
	} else if p.PageSize > p.PageMaxSize && p.PageMaxSize != 0 {
		p.PageSize = p.PageMaxSize
	}
	return p.PageSize
}

func (p *Page) Calculate(total int64) int64 {
	return int64(math.Ceil(float64(total) / float64(p.Limit())))
}

func (p *Page) ToListPage(total int64) ListPage {
	return ListPage{
		TotalPage: p.Calculate(total),
		Page:      p.Page,
		PageSize:  p.PageSize,
		Total:     total,
	}
}

func (p *Page) Paginate() *options.FindOptions {
	return options.Find().SetSkip(p.Offset()).SetLimit(p.Limit())
}

type ListPage struct {
	TotalPage int64 `json:"totalPage"`
	Page      int64 `json:"paginator"`
	PageSize  int64 `json:"pageSize"`
	Total     int64 `json:"total"`
}
