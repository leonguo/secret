package models

import (
	pgorm "../../db/gorm"
	"errors"
)

// 共用部分方法
func Create(v interface{}) (err error) {
	if !pgorm.DBManager().NewRecord(v) {
		return errors.New("create error,  not new record")
	}
	if err = pgorm.DBManager().Create(v).Error; err != nil {
		return err
	}
	return
}

