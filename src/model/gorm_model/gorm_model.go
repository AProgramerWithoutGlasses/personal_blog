package gorm_model

import "gorm.io/gorm"

// User 结构体映射 users 表
type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Avatar       string `gorm:"default:null"`
	Permission   string `gorm:"default:0"`
	Posts        []Post `gorm:"foreignKey:UserID"` // 外键关联
}

// Post 结构体映射 posts 表
type Post struct {
	gorm.Model
	UserID     int      `gorm:"not null"`
	CategoryID *int     // 字段去除了会报错
	Title      string   `gorm:"not null; foreignKey"`
	Content    string   `gorm:"not null"`
	coverimg   string   `gorm:"default:null"` // 封面
	Status     string   `gorm:"default:draft"`
	Views      int      `gorm:"default:0"` // 浏览次数
	User       User     `gorm:"foreignKey:UserID"`
	Category   Category `gorm:"foreignKey:CategoryID"` // 字段去了会报错  标签可以去除不会有问题
}

// Category 结构体映射 categories 表
type Category struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Posts []Post `gorm:"foreignKey:CategoryID"` // 字段可以去除不会有问题
}
