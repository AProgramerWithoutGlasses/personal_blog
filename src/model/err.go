package model

import "errors"

var (
	ErrUserNotExist    = errors.New("user not exist")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserExist       = errors.New("user is exist")
)
