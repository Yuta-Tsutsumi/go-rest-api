/*
repositoryパッケージ内のuser_repoditory.go内の「interface」の型を使用（interfaceに依存する関係）
user_repoditory.go内の「interface」の型の定義のみで処理を記載していないため、「usecase」パッケージで実装する必要がある
「usecase」では、DBで処理するためり「処理設計書」作成する役割。
その設計書は「repository」パッケージへ渡され、その処理設計書に従ってBDの編集を行う
type : 新しいデータ型を定義する
struct : 複数のフィールドを持つ構造体型を表す
func : 関数やメソッドを定義する
&userUsecase{ur} : 構造体の値を生成し、ポインタを取得する
*/

package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// 「controller」フォルダと連携をとるインターフェース（interface）を定義
// その中で「ユーザー登録」と「ログイン」の処理ルールを定義。
// 「SignUp」：ユーザーを新規登録するメソッド作成。「user」は引数名、「model.User」は引数（入力情報）、戻り値（出力値）は「model.UserResponse, error」の2つ
// 「Login」：ユーザーをログインさせるメソッドを作成
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// 「userUsecase」は新しく定義する型の名前。
// 「struct」は複数のフィールドをまとめる構造体。「ur」はその構造体が持つフィールド名。
// 「repository」はパッケージ名、「IUserRepository」は「repository」はパッケージ内の「user_repository.go」で定義された型の名前
type userUsecase struct {
	ur repository.IUserRepository
}

// 「NewUserUsecase」は、userUsecaseを生成するコンストラクター関数。
// 外部で用意したRepositoryの実装をUsecaseに渡す
// 依存性の注入（DI：Dependency Injection）を行う。
//
// 「ur」は引数名。
// 「repository.IUserRepository」は引数のデータ型（インターフェース型）。
// 「IUserUsecase」は戻り値のデータ型（インターフェース型）。
//
// 「&userUsecase{ur}」で構造体を生成する。
// 受け取ったRepositoryの値を、構造体のurフィールドに設定する。
// 「&」で構造体へのポインタを取得する。
// 生成した*userUsecaseをIUserUsecase型として呼び出し元へ返す。
//
// 構造体を生成する目的：
// ① Repositoryの実装を保持するため。
// ② SignUp・LoginからRepositoryのメソッドを呼び出すため。
// ③ 生成したUsecaseをControllerなどに渡すため。
	func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
		return &userUsecase{ur}
	}

//	func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
//		return &userUsecase{ur}
//	}
//上記で構造体（表データ）にした「IUserUsecase」を下で実装。
// return &userUsecase{ur}の返り値の「userUsecase」型が、「IUserUsecase」の「interface」を満たすには、
// 「interface」で定義されている全てのメソッドを実装する必要がある。（今回の場合は、「SignUp」と「Login」メソッド）
// 以下の処理で、「SignUp」と「Login」メソッドの処理を実装していく



func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse {
		ID: 		newUser.ID,
		Email:  newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
			return "", err
	}
	return tokenString, nil
}