/* ---------------------------------------------------
repositoryパッケージの主な役割：
・データベース内のデータの保存・取得・更新・削除を担当する。
・DB操作に必要なメソッドをインターフェースで定義する。
・インターフェースと具体的なDB操作の実装を分離する。
・GORMを使用してPostgreSQLへのデータ操作を実行する。
・メールアドレスを条件にユーザー情報を検索する。
・新しいユーザー情報をデータベースに登録する。
・Usecase層にDB操作の機能を提供する。
・Usecase層が具体的なDB操作の実装に直接依存しないようにする。

IUserRepository：DB操作に必要なメソッドの仕様を定義するインターフェース型。
userRepository：GORMのDB接続情報を保持し、DB操作を実装する構造体。
NewUserRepository()：userRepositoryを生成し、IUserRepository型として返す関数。
GetUserByEmail()：メールアドレスを条件にユーザー情報を取得するメソッド。
CreateUser()：新しいユーザー情報をデータベースに登録するメソッド。

interface：メソッドの集合などを定義するキーワード。
struct：複数のフィールドをまとめた構造体型を表すキーワード。
「*gorm.DB」：GORMのDB操作用インスタンスを参照するポインタ型。
error：処理中に発生したエラー情報を表すインターフェース型。
nil：エラーがない場合などに使用される、値が存在しないことを表す識別子。
Where()：データベースの検索条件を指定するGORMのメソッド。
First()：条件に一致する最初のレコードを取得するGORMのメソッド。
Create()：データベースに新しいレコードを登録するGORMのメソッド。
.Error：GORMの処理結果からエラー情報を取得するフィールド。
--------------------------------------------------- */

package repository

import (
	"go-rest-api/model"

	// GORM：Go言語からDBを操作するためのORMライブラリ。
	"gorm.io/gorm"
)

// 「interface」はインターフェース型を定義するキーワード。
// 「interface」とは、ある型が持つ必要のあるメソッドの仕様を定義したデータ型。
// 「IUserRepository」はインターフェースの型名。
// 先頭の「I」はインターフェースだと分かるように命名（Go言語の必須ルールではない）。
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
// 「error」はGo標準で定義されているインターフェース型。
// 処理成功時はnil、失敗時はエラー値を返す。
type IUserRepository interface {
    GetUserByEmail(user *model.User, email string) error
    CreateUser(user *model.User) error
}

// 「userRepository」は、DB接続情報を保持する構造体型。
// 「db」はフィールド名、「*gorm.DB」はGORMのDB接続を扱うポインタ型。
// この構造体を通じて、ユーザー情報の検索・登録などのDB操作を実行する。
type userRepository struct {
    db *gorm.DB
}

// 「NewUserRepository」は、userRepositoryを生成するコンストラクター関数。
// 「db」は引数名、「*gorm.DB」は引数のデータ型（ポインタ型）。
// 「IUserRepository」は戻り値のデータ型（インターフェース型）。
// 「&userRepository{db}」で構造体を生成し、受け取ったDB接続情報をdbフィールドに設定する。
// 「&」で構造体へのポインタを取得する。
// 生成した*userRepositoryをIUserRepository型として呼び出し元へ返す。
func NewUserRepository(db *gorm.DB) IUserRepository {
    return &userRepository{db}
}

// 「GetUserByEmail」は、メールアドレスを条件にDBからユーザー情報を取得するメソッド。
// 「ur」はレシーバー名、「*userRepository」はレシーバーのデータ型（ポインタ型）。
// 「user」は検索結果の格納先、「email」は検索対象のメールアドレス。
// 「ur.db」はuserRepository構造体が保持するGORMのDB接続情報。
//
// 「Where("email=?", email)」は、emailカラムが指定した値と一致する条件を設定する。
// 「?」には引数emailの値が安全にバインドされる。
// 「First(user)」は条件に一致する最初のレコードを取得し、userに格納する。
// 「.Error」はGORMの処理結果からエラー情報を取得する。
// 「err != nil」の場合はエラーを返し、正常終了時はnilを返す。
// 対象ユーザーが存在しない場合も、通常はエラーが返される。
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
    if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
        return err
    }
    return nil
}

// 「CreateUser」は、新しいユーザー情報をDBに登録するメソッド。
// 「ur」はレシーバー名、「*userRepository」はレシーバーのデータ型（ポインタ型）。
// 「user *model.User」は登録するユーザー情報へのポインタ。
// 「ur.db」はuserRepository構造体が保持するGORMのDB接続情報。
//
// 「Create(user)」はGORMを使用してDBに新しいレコードを登録する。
// 登録が成功すると、自動採番されたIDや作成日時などがuserに反映される。
// 「.Error」はDB登録処理で発生したエラー情報を取得する。
// 「err != nil」の場合はエラーを返し、正常終了時はnilを返す。
func (ur *userRepository) CreateUser(user *model.User) error {
    if err := ur.db.Create(user).Error; err != nil {
        return err
    }
    return nil
}
