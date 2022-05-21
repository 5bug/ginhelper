package rest

import "math"

// PageOption ...
type PageOption struct {
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
	Sort      string `form:"sort"` //asc or desc
}

// TotalOption ...
type TotalOption struct {
	PageIndex  int   `json:"pageIndex"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	TotalPage  int   `json:"totalPage"`
}

// ListData ...
type ListData struct {
	TotalOption
	List interface{} `json:"list"`
}

// NewListData fill data
func NewListData(opt PageOption, total int64, list interface{}) *ListData {
	data := &ListData{}
	data.PageIndex = opt.PageIndex
	data.PageSize = opt.PageSize
	data.TotalCount = total
	data.TotalPage = int(math.Ceil(float64(total) / float64(opt.PageSize)))
	data.List = list
	return data
}

// Check ...
func (opt *PageOption) Check() {
	if opt.PageIndex <= 0 {
		opt.PageIndex = 0
	}
	if opt.PageSize <= 0 {
		opt.PageSize = 10
	} else if opt.PageSize < 5 {
		opt.PageSize = 5
	}
}
