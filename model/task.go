package model

import "time"

type Task struct{
	ID uint `json:"id" gprm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// user.goファイルのUserの定義
	//  constraint:OnDelete:CASCADEはUser情報を削除した時に、しれに紐ずくTask情報を削除する指定
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`

	// user.goファイルのUserのID情報を定義
	UserId uint `json:"user_id" gorm:"not unll"`
}

// クライアントに返す情報の定義
type TaskResponse struct {
	ID uint `json:"id" gprm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}