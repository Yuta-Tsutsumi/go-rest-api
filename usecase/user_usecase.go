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

/* ーーーーーーーーーーーーーーーーーーーーーーーーーー
以下は、具体的には、次の5つの処理を実行している。
1. パスワードをハッシュ化する。
2. エラーが発生していないか確認する。
3. 登録用のUser構造体を生成する。
4. Repositoryを呼び出してDBに保存する。
5. 登録したユーザーのID・Emailを返す。
ーーーーーーーーーーーーーーーーーーーーーーーーーー */
// 「SignUp」は、ユーザーの新規登録を行うUsecaseのメソッド。
// 「uu」はレシーバー名、「*userUsecase」はレシーバーのデータ型（ポインタ型）。
// 「user」は引数名、「model.User」は引数のデータ型。
// 戻り値は「model.UserResponse」と「error」の2つ。
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {

    // 「bcrypt.GenerateFromPassword」でパスワードをハッシュ化する。
    // 「user.Password」は入力されたパスワード。
    // 「[]byte」でstring型のパスワードを[]byte型に変換する。
		// 「10」はbcryptのコスト係数で、ハッシュ化処理の計算負荷を指定する。
		// コスト係数を1増やすと計算量は約2倍になり、総当たり攻撃への耐性が高まる。
		// ただし、コスト係数が高いほどユーザー登録・ログイン処理の時間も増える。
		// 「hash」は生成されたハッシュ値、「err」はエラーを格納する変数。
    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

    // ハッシュ化でエラーが発生した場合は、終了する処理。
    // 「model.UserResponse{}」は空のUserResponse構造体の値。
    // 「err」は発生したエラーを返す。（つまり、エラーの場合はエラーのみを返す）
    if err != nil {
        return model.UserResponse{}, err
    }

    // DBに登録するためのmodel.User型の構造体の値を生成する。
    // Emailには入力されたメールアドレスを設定する。
    // Passwordにはハッシュ化したパスワードを設定する。
    // 「string(hash)」で[]byte型をstring型に変換する。
    newUser := model.User{
        Email:    user.Email,
        Password: string(hash),
    }

    // RepositoryのCreateUserメソッドを呼び出してDB登録を依頼する。
    // 「uu.ur」はuserUsecaseが保持しているRepository。
    // 「&newUser」で構造体へのポインタを渡す。
    // DB登録に失敗した場合は空のUserResponseとエラーを返す。
    if err := uu.ur.CreateUser(&newUser); err != nil {
        return model.UserResponse{}, err
    }

    // Controllerに返すmodel.UserResponse型の構造体の値を生成する。
    // DB登録後のIDとEmailを設定する。
    // Passwordはレスポンスに含めない。
    resUser := model.UserResponse{
        ID:    newUser.ID,
        Email: newUser.Email,
    }

    // 登録したユーザー情報を呼び出し元（user_repository.go）へ返す。
    // 「nil」はエラーが発生していないことを表す。
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