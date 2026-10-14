package repository

import "go-rest-api/model"

// 「IUserRepository」の先頭の文字列が「I」から始まるのはインターフェイスを定義
//
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}