package servers

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
}

func (p *Pagination) GetPageSize() int {
	var limits = []int{10, 30, 50, 100}

	for limit := range limits {
		if p.PageSize == limits[limit] {
			return p.PageSize
		}
	}

	return 10
}

func (p *Pagination) GetCurrentPage() int {
	var currentPage int
	PageSize := p.GetPageSize()
	currentPage = PageSize * (p.CurrentPage - 1)

	if currentPage < 0 {
		return 0
	}
	return currentPage
}
