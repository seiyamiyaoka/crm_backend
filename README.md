## CRM backend project

## version
go 1.19

## how to setup
- setup docker and docker-compose

## how to run

```sh
make build
make start
```

## if you want to mysql server
- access from docker
- mysql crm_gackend_database -u crm_gackend_user -p  => then (input password)


## db schema

```
mysql> desc `customers`;
+------------+-----------------+------+-----+---------+----------------+
| Field      | Type            | Null | Key | Default | Extra          |
+------------+-----------------+------+-----+---------+----------------+
| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)     | YES  |     | NULL    |                |
| updated_at | datetime(3)     | YES  |     | NULL    |                |
| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
| name       | longtext        | YES  |     | NULL    |                |
| age        | bigint          | YES  |     | NULL    |                |
| role       | longtext        | YES  |     | NULL    |                |
| email      | longtext        | YES  |     | NULL    |                |
| phone      | longtext        | YES  |     | NULL    |                |
| contacted  | tinyint(1)      | YES  |     | NULL    |                |
+------------+-----------------+------+-----+---------+----------------+
```

## my todo
- [x] dockerで環境構築
- [x] ルートエンドポイントで固定データを返すようにする
- [x] db containerを作成する
- [x] db => mysqlと接続
- [x] appとdbを接続
- [x] applicationのディレクトリ構成調査
- [x] ディレクトリ構築
- [x] readmeに使い方を書く
- [x] seedデータを投入できるようにする(3人以上)
- [x] localhostを通じてアクセス
- [x] データは全てjsonで返す
- [x] /ルートは簡単な説明文を記述
- [x] gorila/muxを使う
- [x] routing list
  - [x]Getting a single customer through a /customers/{id} path
  - [x] Getting all customers through a the /customers path
  - [x] Creating a customer through a /customers path
  - [x] Updating a customer through a /customers/{id} path
  - [x] Deleting a customer through a /customers/{id} path
- [x]存在しないユーザの更新や削除は404とからのjsonを返す
- [x]適切なコンテンツheaderを返す
- [x]io/ioutilを使って使ってeー他を取り出す
- option batch処理で顧客をupdateするendpoint作成(複数parameterを作成)
- option apiをwebにdeploy
- option unit test入れる

```go
getCustomers()
getCustomer()
addCustomer()
updateCustomer()
deleteCustomer()
```

## reference
- packageをまとめる go mod init github.com/seiyamiyaoka/crm_backend
- http://psychedelicnekopunch.com/archives/1308
- https://gorm.io/ja_JP/docs/models.html
- https://zenn.dev/hrs/articles/go-gin-air-docker
- https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
- https://www.bogotobogo.com/GoLang/GoLang_Modules_2_Adding_and_Updating_Dependencies.php
- https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh
- https://qiita.com/fetaro/items/31b02b940ce9ec579baf
- https://eaceto.dev/2020/02/15/listing-all-endpoints-with-gorilla-mux-router/