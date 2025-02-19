package model

import "blog1/src/model/gorm_model"

type BackAllDataModel struct {
	Posts      []gorm_model.Post
	Base       gorm_model.Base
	Categories []gorm_model.Category
	Comments   []gorm_model.Comment
	Users      []gorm_model.User
}

type BackPostModel struct {
	Post gorm_model.Post
	Base gorm_model.Base
}

type BackEditPostModel struct {
	Post       gorm_model.Post
	Base       gorm_model.Base
	Categories []gorm_model.Category
}

type BackCategoriesModel struct {
	Categories []gorm_model.Category
	Base       gorm_model.Base
}
