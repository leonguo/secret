package lib

import (
	"github.com/jinzhu/gorm"
	"math"
)

type Paginator struct {
	DB      *gorm.DB
	OrderBy []string
	Page    int
	Limit   int
}

type Data struct {
	TotalRecords int         `json:"total_records"`
	Records      interface{} `json:"records"`
	CurrentPage  int         `json:"current_page"`
	TotalPages   int         `json:"total_pages"`
}

func (p *Paginator) Paginate(dataSource interface{}) *Data {
	db := p.DB

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var output Data
	var count int
	var offset int

	go countRecords(db, dataSource, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(dataSource)
	<-done

	output.TotalRecords = count
	output.Records = dataSource
	output.CurrentPage = p.Page
	output.TotalPages = getTotalPages(p.Limit, count)

	return &output
}

func countRecords(db *gorm.DB, countDataSource interface{}, done chan bool, count *int) {
	db.Model(countDataSource).Count(count)
	done <- true
}

func getTotalPages(limit int, totalRecords int) int {
	totalPages := float64(totalRecords) / float64(limit)
	return int(math.Ceil(totalPages))
}
