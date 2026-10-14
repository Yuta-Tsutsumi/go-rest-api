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
