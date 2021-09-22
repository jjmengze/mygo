package model

import (
	"gorm.io/gorm"
)

// Post ...
type Post struct {
	gorm.Model
	Author     []User
	Title      string
	Content    string
	LikeGivers []User
}
