package gorm_model

import "gorm.io/gorm"

// User 结构体映射 users 表
type User struct {
	gorm.Model
	Username   string `gorm:"not null" form:"Username"`
	Password   string `gorm:"not null" form:"Password"`
	Name       string `gorm:"not null" form:"Name"`
	Age        string `gorm:"default:0" form:"Age"`
	Gender     string `gorm:"default:男" form:"Gender"`
	Avatar     string `gorm:"default:null" form:"Avatar"`
	Permission string `gorm:"default:0" form:"Permission"`

	//
	Comments []Comment `gorm:"foreignKey:UserID"`
}

// Post 结构体映射 posts 表
type Post struct {
	gorm.Model
	Title    string `gorm:"not null; foreignKey" form:"title"`
	Content  string `gorm:"not null" form:"content"`
	CoverImg string `gorm:"default:null" form:"coverImg"` // 封面
	Summary  string `gorm:"default: " form:"summary"`
	Comment  string `gorm:"type:text" form:"comment"`
	Views    int    `gorm:"default:0" form:"views"` // 浏览次数
	Slug     string `gorm:"not null; unique" form:"slug"`

	// 外键关系
	CategoryID int       `form:"category_id"`
	Comments   []Comment `gorm:"foreignKey:PostID"`
	Category   Category
}

// Category 结构体映射 categories 表
type Category struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Posts []Post `gorm:"foreignKey:CategoryID"` // 字段可以去除不会有问题
}

type Comment struct {
	gorm.Model
	// 内容
	Content string `gorm:"type:text"`

	// 外键关系
	UserID int `gorm:"not null"`
	PostID int `gorm:"not null"`
	User   User
	Post   Post
}

type Base struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
}
