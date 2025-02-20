package model

import "blog1/src/model/gorm_model"

type BackAllDataModel struct {
	Posts       []gorm_model.Post
	Base        gorm_model.Base
	Categories  []gorm_model.Category
	Comments    []gorm_model.Comment
	Users       []gorm_model.User
	CurrentPage string
}

type BackPostModel struct {
	Post        gorm_model.Post
	Base        gorm_model.Base
	CurrentPage string
}

type BackEditPostModel struct {
	Post        gorm_model.Post
	Base        gorm_model.Base
	Categories  []gorm_model.Category
	CurrentPage string
}

type BackCategoriesModel struct {
	Categories  []gorm_model.Category
	Base        gorm_model.Base
	CurrentPage string
}

type BackEditCategoriesModel struct {
	Id          string `form:"id"`
	Name        string `form:"name"`
	CurrentPage string
}

type BackNewCategoriesModel struct {
	Name string `form:"name"`
}

type BackUsersModel struct {
	Users       []gorm_model.User
	Base        gorm_model.Base
	CurrentPage string
}
