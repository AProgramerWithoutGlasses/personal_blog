package model

import "staging/src/model/gorm_model"

type IndexModel struct {
	Posts      []gorm_model.Post
	Base       gorm_model.Base
	Categories []gorm_model.Category
}

type PostModel struct {
	Post       gorm_model.Post
	PrePost    gorm_model.Post
	NextPost   gorm_model.Post
	Base       gorm_model.Base
	Categories []gorm_model.Category
}
