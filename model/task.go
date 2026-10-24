/* ---------------------------------------------------
task.goファイルの主な役割：
・Todoタスクのデータ構造を定義する。
・タスクのID・タイトル・作成日時・更新日時を定義する。
・GORMがDBのテーブル構造を判断するための情報を定義する。
・UserとTaskの1対多のリレーションシップを定義する。
・UserIDを外部キーとして、タスクとユーザーを関連付ける。
・ユーザー削除時に関連するタスクも削除するCASCADE制約を定義する。
・クライアントに返すレスポンス用のデータ構造を定義する。

Task：DBに保存するTodoタスクの情報を定義する構造体。
TaskResponse：クライアントに返すタスク情報を定義する構造体。
primaryKey：DBの主キーを指定するGORMタグ。
not null：DBのカラムにNULL値を許可しない制約。
foreignKey：関連する構造体との外部キーを指定するGORMタグ。
OnDelete:CASCADE：親レコード削除時に関連する子レコードも削除する外部キー制約。
time.Time：日時を表すGoのデータ型。
--------------------------------------------------- */

package model

import "time"

// Taskは、DBに保存するTodoタスクの情報を定義する構造体。
// Userとは1対多の関係で、1人のユーザーが複数のTaskを所有できる。
//
// ID：タスクを一意に識別するID。
// uint：0以上の整数型。primaryKey：DBの主キーを指定。
// Title：タスクのタイトル。
// not null：DBでNULL値を禁止する制約。
// CreatedAt：タスクの作成日時。GORMが自動設定する。
// UpdatedAt：タスクの最終更新日時。GORMが自動管理する。
// User：user.goで定義したUser構造体との関連付け。
// foreignKey:UserID：TaskのUserIDフィールドを外部キーに指定する。
// constraint:OnDelete:CASCADE：関連するUserがDBから削除された場合、
// 外部キー制約によって、そのUserに紐づくTaskも自動削除する。
// UserID：このタスクを所有するユーザーのID。
// User構造体のID（usersテーブルの主キー）を参照する外部キー。
// not null：DBでNULL値を禁止する制約。
type Task struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Title string `json:"title" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    User User `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    UserID uint `json:"user_id" gorm:"not null"`
}

// TaskResponseは、クライアントに返すタスク情報を定義する構造体。
// タスクのID・タイトル・作成日時・更新日時のみを返す。
type TaskResponse struct {
    ID        uint      `json:"id"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}