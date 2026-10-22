/* ---------------------------------------------------
dbパッケージの主な役割：
・アプリケーションとPostgreSQLの接続を管理する。
・開発環境では.envファイルから環境変数を読み込む。
・環境変数をもとにデータベースの接続URLを生成する。
・GORMを使用してPostgreSQLへの接続を初期化する。
・DB操作に使用するGORMのインスタンスを返す。
・データベース接続の終了処理を行う。

NewDB()：データベース接続を初期化し、GORMのインスタンスを返す関数。
CloseDB()：データベースの接続プールを閉じる関数。
godotenv.Load()：.envファイルから環境変数を読み込む関数。
os.Getenv()：OSの環境変数から指定した値を取得する関数。
fmt.Sprintf()：指定した形式に従って文字列を生成する関数。
gorm.Open()：GORMを初期化する関数。
--------------------------------------------------- */

package db

import (
	"fmt"
	"log"
	"os"

	// 「go mod tidy」コマンドで「gorm.DB」と「godotenv」をダウンロード
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDB() *gorm.DB {

	// 「GO_ENV」の値が「dev」の場合は、godotenv.Load」関数を呼び出して、「.env」ファイルの内容を読み込む
	//	本番環境の場合は以下のローカルの環境設定を実行されないように設定している
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()

		//上記で「err := godotenv.Load()」が使用されているため、「err」変数に入るの値のデータ型は「error型」となる
		//「nil」値（エラー情報）がない状態を表す識別子。（nilは特定の型を持たない）「!=」を使用しているため、今回は「値が場合」を見ている
		//つまり、「err != nil」はerr変数に格納してあるerror型の値があることの状態を確認する式
		// 「err != nil」の条件がtrue（エラー情報がある）場合、エラーメッセージをログに出力し、プログラムを終了させる
		if err != nil {
			log.Fatalln(err)
		}
	}

	// fmt.Sprintf()はGo言語のGORMを使用してPostgreSQLに接続するためのURLを生成する関数。
	//「.env」で定義した「POSTGRES_USER」などの環境変数の値を全て足した値がURLとなる
	// os.Getenv()は、OSの環境変数から指定した名前の値を文字列として取得する関数。
	url := fmt.Sprintf(
  	  "postgres://%s:%s@%s:%s/%s",
    	os.Getenv("POSTGRES_USER"),
  	  os.Getenv("POSTGRES_PW"),
    	os.Getenv("POSTGRES_HOST"),
  	  os.Getenv("POSTGRES_PORT"),
   	  os.Getenv("POSTGRES_DB"),
	)

	// 「db, err」は「db」と「err」の2つの変数を同時に宣言している。「db」はデータベース操作に使用するGORMのオブジェクトが格納され、「err」はエラー情報が格納される。
	//「:=」は、Go言語で変数の宣言と値の代入を同時に行う演算子。（例）「var number int = 10」→「number := 10」
	// gorm.Open()は、PostgreSQLへの接続設定を使用してGORMを初期化する関数。
	// postgres.Open(url)は、生成した接続URL（url変数）をもとにPostgreSQL用の接続設定を作成する関数。
	// &gorm.Config{}は、GORMの動作設定を指定するもので、今回はデフォルト設定を使用する。
	db, err := gorm.Open(
  	  postgres.Open(url),
   	 &gorm.Config{},
	)

	//上記コードでDB接続の際にエラーが発生した場合は、エラーメッセージをログに出力し、プログラムを終了させる
	if err != nil {
		log.Fatalln(err)
	}

	//エラーが発生しなかった場合は、「Connceted」という文字列を表示
	//「return db」の「db」変数は戻り値として「NewDB()」関数に返すようにする
	fmt.Println("Connceted")
	return db
}

// DBを閉じる処理
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}