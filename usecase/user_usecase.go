/*
interfaceで要求されたメソッドを、
名前・引数・戻り値の型まで一致させて用意することが「interfaceを実装する」という意味
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