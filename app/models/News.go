package models

import (
	"time"
)

const TableNameAppTrxNews = "app_trx_news"

// AppTrxNews mapped from table <app_trx_news>
type AppTrxNews struct {
	NewsID       string     `gorm:"column:news_id;primaryKey" json:"news_id"`
	NewsTitle    string     `gorm:"column:news_title" json:"news_title"`
	NewsDesc     string     `gorm:"column:news_desc" json:"news_desc"`
	NewsLink     *string    `gorm:"column:news_link" json:"news_link"`
	NewsPhoto    *string    `gorm:"column:news_photo" json:"news_photo"`
	NewsPublish  *time.Time `gorm:"column:news_publish" json:"news_publish"`
	NewsUnpublish *time.Time `gorm:"column:news_unpublish" json:"news_unpublish"`
	IsActive     *bool      `gorm:"column:is_active" json:"is_active"`
	SchoolID     *string    `gorm:"column:school_id" json:"school_id"`
}

// TableName AppTrxNews's table name
func (*AppTrxNews) TableName() string {
	return TableNameAppTrxNews
}
