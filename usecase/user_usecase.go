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

	// 「repository」パッケージで定義されたインターフェースを通じて、Repositoryの実装を「usecase」パッケージに渡す依存性の注入（DI：Dependency Injection）を行う。
	// 「NewUserUsecase」は、外部からRepositoryの実装を受け取り、userUsecase構造体を生成するコンストラクター関数。
	// 「ur」は引数名、「repository.IUserRepository」は引数。引数のデータ型はインターフェース型。
	// 「IUserUsecase」は戻り値のデータ型（インターフェース型）。
	// 「&userUsecase{ur}」で構造体を生成し、受け取ったRepositoryをurフィールドに格納する。
	// 「&」で構造体のポインタを取得し、「IUserUsecase」型として呼び出し元（user_repository.go）に返す。

	/* 下記はurというフィールドを持つ構造体型を定義
	type userUsecase struct {
	ur repository.IUserRepository
	}
	return &userUsecase{ur}はその構造体を生成する処理
ここで構造体を生成する目的は、主に次の3つです。
1. Repositoryの実装を保持するため
2. SignUp・LoginメソッドからRepositoryを利用できるようにするため
3. 生成したUsecaseをControllerなどに渡すため
	*/
	func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
		return &userUsecase{ur}
	}

	
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