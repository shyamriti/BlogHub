package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Blogs    []Blog
	Comments []Comment
}

type Blog struct {
	gorm.Model
	Title    string
	Caption  string
	UserID   uint
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:BlogID"`
}

type Comment struct {
	gorm.Model
	Content string
	BlogID  uint
	Blog    Blog `gorm:"foreignKey:BlogID"`
	UserID  uint
	User    User `gorm:"foreignKey:BlogID"`
}
