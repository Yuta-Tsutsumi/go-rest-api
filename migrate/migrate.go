// 「package main」の特別で、このパケージを記載したファイルはマイグレーションを実行をするファイルとなる。
// 今回のmigrate.goは、データベースのマイグレーション（作成・更新）を実行するためのプログラムなので、package mainと記述
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

// 「db.NewDB()」は、先ほど作成した「db.go」ファイル内の「NewDB()」関数を呼び出す
// （つまり、dbConnにはデータベースを操作するためのGORMオブジェクトを格納）
// 「	defer fmt.Println()」関数はマイグレーション処理が成功した際に、文字列の「Successfully Migrated」を出力
// 「	defer db.CloseDB(dbConn)」はdb.goファイルの「CloseDB」関数を使用。DBを閉じる処理。
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	// 「model」フォルダ配下の「user.go」「task.go」で定義したカラムをDBのテーブルに反映
	// 「AutoMigrate()」関数は、Go言語の構造体に基づいてデータベースのテーブル構造を作成・更新する際に使用。
	// 「dbConn」は上記で定義している「db.go」ファイル内の「NewDB()」関数で、今回接続するDBを指定するために使用している
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}


/* ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
【プログラムとDBの接続を確認するコマンド】
・「Docker Desktop」を起動した後、以下のコマンドを実行。
docker compose up -d
・以下のコマンドで、PostgreSQLのコンテナが起動していることを確認
docker compose ps

・「migrate.go」の実行コマンド
go run migrate/migrate.go
（「db.go」ファイルの条件である「os.Getenv("GO_ENV") == "dev"」の部分が「GO_ENV=dev」の場合の実行コマンド）
（GO_ENV=dev go run migrate/migrate.go）

・「pdAdmin」を起動し、PostgreSQLに定義したテーブルを確認（ツリーは以下となる）
usersやtasksを右クリックし、「View/Edit Data」→「All Rows」の順にクリック

Servers（今回の場合はTodeDB）
 └── Databases
      └── todo
           └── Schemas
                └── public
                     └── Tables
                          ├── users
                          └── tasks
ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー */