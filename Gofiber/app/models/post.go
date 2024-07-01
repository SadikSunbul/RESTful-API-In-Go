package models

type Post struct {
	Base
	Title   string
	Content string `gorm:"type:text"`
}
