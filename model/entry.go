package model

import (
	"diary_api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	tx := database.Database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&entry).Error
	if err != nil {
		tx.Rollback()
		return &Entry{}, err
	}
	return entry, tx.Commit().Error
}
