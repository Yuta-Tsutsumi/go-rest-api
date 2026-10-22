/* ---------------------------------------------------
modelパッケージの主な役割：
・アプリケーションで扱うデータの構造を定義する。
・ユーザー情報などの構造体型を定義する。
・GORMがDBのテーブル構造を判断するための情報を定義する。
・クライアントに返すレスポンス用のデータ構造を定義する。

type：新しいデータ型を定義するキーワード。
struct：複数のフィールドをまとめた構造体型を表すキーワード。
--------------------------------------------------- */

package model

import "time"

// Userは、DBに保存するユーザー情報を定義する構造体。
// GORMタグで主キーや一意制約を指定し、作成日時・更新日時を管理する。
type User struct{
	ID				uint			`json:"id" gorm:"primaryKey"`
	Email			string		`json:"email" gorm:"unique"`
	Password	string		`json:"password"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAd time.Time	`json:"updated_at"`
}
/*
「User」（型名）は、DBに保存するユーザー情報を定義する構造体。
「struct（キーワード）」は複数のフィールドをまとめた構造体型を表す。
uint：0以上の整数を扱うデータ型。
json：JSON変換時のキー名を指定するタグ。これを指定することで、小文字で値が保存される
id：ユーザーを識別するID。
password：パスワードを表すJSONのキー名。
gorm：DBのテーブルやカラムの設定を指定するタグ。
primaryKey：主キーを指定し、レコードを一意に識別する。
unique：一意制約を指定し、値の重複を禁止する。
created_at：作成日時を表すJSONのキー名。
updated_at：最終更新日時を表すJSONのキー名。
*/


// UserResponseは、クライアントに返す、新しく作成されるユーザー情報を定義する構造体。
// IDとEmailのみを返し、パスワードなどの機密情報を除外する。（機密情報はユーザーにレスポンスしない）
type UserResponse struct{
	ID 				uint 	 `json:"id" gorm:"primaryKey"`
	Email 		string `json:"email" gorm:"unique"`

}
