[![codecov](https://codecov.io/gh/suzuito/sandbox2-go/graph/badge.svg?token=Rj1wZ7rRgW)](https://codecov.io/gh/suzuito/sandbox2-go)

# sandbox2-go

趣味開発用のSandboxレポジトリ。

## ディレクトリ構造

ディレクトリ構造は下記の通り。

|||
|---|---|
|`cmd`|サブプロジェクト（後述）に依存しないコマンドラインユーティリティ|
|`.service`|サブプロジェクト毎の開発環境を構築するためのリソース群|
|`$subProject/`|サブプロジェクト（後述）のルートディレクトリ|

### サブプロジェクト

本レポジトリのソースコードはサブプロジェクト毎に分類されている。
あるサブプロジェクトと別のサブプロジェクトは、基本的には、全く関係がなく独立している。
例 `blog`はブログサイト用のソースコードがあり、`crawler`は自前クローラー用のソースコードがある。
しかしながら、`blog`配下のソースコードから`crawler`配下のソースコードを呼び出すことについて、何も問題はない。
趣味開発なので、その辺りは緩めで運用する。
とはいえ、いうまでもなく、公開する必要のないソースコードは、なるべくサブプロジェクト外には公開しないようにすることが鉄則である。

### サブプロジェクト構造制約

サブプロジェクト配下のディレクトリに対して、[package-check-list.yaml](./package-check-list.yaml)制約を課す。
なるべく、ソースコードを綺麗に保っておくためである。

### パッケージ依存制約

サブプロジェクト配下のディレクトリには、[import-check-list.yaml](./import-check-list.yaml)制約を課す。

## Development

### common

#### Test

```bash
make common-test
```

### blog

```bash
make blog-init
make blog-init-rdb
```

#### Run

```bash
# env
cp ./.service/blog/local.env.sh.sample ./.service/blog/local.env.sh
## Add GH Token
vim ./.service/blog/local.env.sh
source ./.service/blog/local.env.sh

# server
air -c ./.service/blog/.air.server.toml
curl http://localhost:8080/ping

# check rdb
docker compose exec blog-mysql mysql
```

Insert test articles into local db

- Login as admin
  - Access /admin/login
  - Input password
  - Click Login
- Import markdown on repository
  - Access /admin/
  - Click `*`

Migration

```bash
# Create new migration
.bin/migrate create -dir .schema -ext sql init
```

#### Test

```bash
make blog-test
```

### blog2

```bash
make blog2-init
make blog2-init-rdb
```

#### Migration

```bash
NAME=init make blog2-migrate-create
```

### crawler

Make develop environment

```bash
make crawler-init
```

#### Run

Open filebase UI in local.
http://localhost:8082

Load environment variables.

```bash
cp ./.service/crawler/local.env.sh.sample ./.service/crawler/local.env.sh
vi ./.service/crawler/local.env.sh
source ./.service/crawler/local.env.sh
```

Run crawler app

```bash
./crawler-crawl.exe
```

Run notifier app

```bash
go run ./crawler/cmd/crawl/main.go -crawler-id knowledgeworkblog -crawler-input-data '{"URL":"https://note.com/knowledgework/n/n4d7b97ff802c"}'
go run ./crawler/cmd/notify/main.go -full-path Crawler/TimeSeriesData/goblog/goblog-2023-08-14
```

## デプロイ

定義

- github.com/suzuito/sandbox2-go モジュールをインターネット上に公開すること
- blog CloudRun 用の Docker image を Google Container Resitory へアップロードすること

手順

1. Github action `create-release-draft` を実行する。
  a. Github action が release draft を作成する。
2. Release draft を公開する。
