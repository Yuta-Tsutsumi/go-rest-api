/*
modelパッケージの主な役割：
・アプリケーションで扱うデータの構造を定義する。
・ユーザー情報などの構造体型を定義する。
・GORMがDBのテーブル構造を判断するための情報を定義する。
・クライアントに返すレスポンス用のデータ構造を定義する。

type：新しいデータ型を定義するキーワード。
struct：複数のフィールドをまとめた構造体型を表すキーワード。
*/

package model

import "time"

// ユーザー属性の定義
type User struct{
	ID				uint			`json:"id" gorm:"primaryKey"`
	Email			string		`json:"email" gorm:"unique"`
	Password	string		`json:"password"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAd time.Time	`json:"updated_at"`
}

// クライアントに返す新しく作成される値
type UserResponse struct{
	ID 				uint 	 `json:"id" gorm:"primaryKey"`
	Email 		string `json:"email" gorm:"unique"`

}
