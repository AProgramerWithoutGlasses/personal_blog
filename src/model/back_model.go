package model

import "blog1/src/model/gorm_model"

type BackAllDataModel struct {
	Posts      []gorm_model.Post
	Base       gorm_model.Base
	Categories []gorm_model.Category
	Comments   []gorm_model.Comment
	Users      []gorm_model.User
}
