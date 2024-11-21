package dto

import (
	"time"
)

type Post struct {
	ID        int64 `gorm:"primaryKey"`
	Title     string
	Image     string
	CreatorID int64
	CreatedAt time.Time
}

type Comment struct {
	ID        int64 `gorm:"primaryKey"`
	PostID    int64
	Comment   string
	CreatorID int64
	CreatedAt time.Time
}
