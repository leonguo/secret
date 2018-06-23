package lib

import (
	"github.com/jinzhu/gorm"
	"math"
)

/**
    lib库-page用例
	var videos []Video
	db := pgorm.DBManager()
	if channelNumber > 0 {
		query := "channel_number = " + strconv.FormatInt(channelNumber, 10)
		db = db.Where(query)
	}
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	p := lib.Paginator{DB: db, Limit: limit, Page: page, OrderBy: []string{"video_id desc"}}
	data = p.Paginate(&videos)
 */
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
