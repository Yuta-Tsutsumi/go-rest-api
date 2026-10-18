/*
主な役割 : DB内のデータの保存・取得・更新・削除
考えること : DBにどうアクセスするか
DBへの直接アクセスを行う
ユーザー情報をDBに保存
repositoryパッケージはメソッドの定義は行うが、「interface」では具体的な処理は書かない・
具体的な処理は「usrcase」に任せる
*/

package repository

import "go-rest-api/model"

// 「IUserRepository」の先頭の文字列が「I」から始まるのはインターフェイスを定義
// 「GetUserByEmail」、「CreateUser」の2つのメソットを定義
// 「user *model.User」（第一引数）でユーザーオブジェクトのポインタを受け取れるようにしている
// 「email string」（第二引数）で検索したいeメールを文字列で受け取れるようにする
// 返り値は「error」インターフェース型でを設定
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}