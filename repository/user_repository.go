/* ----------------------------------------------------------------------------
Repository層の主な役割：DB内のデータの保存・取得・更新・削除。
考えること：DBにどのようにアクセスし、データを操作するか。

Repository層では、インターフェースと具体的な実装を分離する。
① インターフェース（IUserRepository）
・DB操作に必要なメソッドの仕様を定義する。
・メソッド名、引数、戻り値の型を定義する。
・具体的なDB操作の処理内容は記載しない（具体的な処理は「usrcase」などに任せる）

② Repositoryの具体的な実装
・IUserRepositoryで定義したメソッドを実装する。
・GORMなどを使用してDBへのアクセスを行う。
・ユーザー情報の検索や保存などを実行する。

③ Usecaseとの関係
・Usecaseはユーザー登録やログインなどの業務処理を担当する。
・DB操作が必要な場合、IUserRepositoryのメソッドを呼び出す。
・実際のDB操作はRepositoryの具体的な実装が担当する。
・これにより、UsecaseはDB操作の実装方法に直接依存しない。
---------------------------------------------------------------------------- */

package repository

import "go-rest-api/model"

// 「interface」はインターフェース型を定義するキーワード（あらかじめGoで決められた文字列）
// 「interface」とは、ある型が持つ必要のあるメソッドの仕様を定義したデータ型。
// 「IUserRepository」はインターフェースの型名。 先頭の「I」はインターフェースをだとわかるように命名（Go言語の必須ルールではない）
//
// 「GetUserByEmail」はメールアドレスからユーザー情報を取得するメソッド。
// 第1引数「user *model.User」は、User構造体へのポインタを受け取る。
// 第2引数「email string」は、検索するメールアドレスを文字列型で受け取る。
// 検索結果は第1引数のポインタを通じて格納する設計。
// 戻り値「error」は検索処理のエラー情報。
//
// 「CreateUser」はユーザー情報をDBに登録するメソッド。
// 引数「user *model.User」は登録するUser構造体へのポインタ。
// 戻り値「error」は登録処理のエラー情報。
//
// 「error」はGo標準で定義されているインターフェース型。処理成功時はnil、失敗時はエラー値を返す。
type IUserRepository interface {
    GetUserByEmail(user *model.User, email string) error
    CreateUser(user *model.User) error
}
