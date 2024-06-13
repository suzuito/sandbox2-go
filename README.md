[![codecov](https://codecov.io/github/suzuito/sandbox2-go/branch/main/graph/badge.svg?token=Rj1wZ7rRgW)](https://codecov.io/github/suzuito/sandbox2-go)
[![check_in_main](https://github.com/suzuito/sandbox2-go/actions/workflows/check-in-main.yaml/badge.svg?branch=main)](https://github.com/suzuito/sandbox2-go/actions/workflows/check-in-main.yaml)


# sandbox2-go

趣味開発用のSandboxレポジトリ。

## ディレクトリ構造

ディレクトリ構造は下記の通り

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

```bash
# Install air
go install github.com/cosmtrek/air@latest
```

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
$(go env GOPATH)/bin/air -c ./.service/blog/.air.server.toml
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

#### Run

```bash
# Environment variables
cp ./.service/blog2/local.env.sh.sample ./.service/blog2/local.env.sh
## Modification of environment variables
vim ./.service/blog2/local.env.sh
source ./.service/blog2/local.env.sh

# create test data
make blog2-init-rdb-test-data

# server
$(go env GOPATH)/bin/air -c ./.service/blog2/.air.server.toml
curl http://localhost:8080/ping

# js
cd blog2/_ts && npm run serve
```

#### Migration

```bash
NAME=init make blog2-migrate-create
```

#### DB

Local

```bash
mysql -u root -h 127.0.0.1 -P 3307
```

## デプロイ

定義

- github.com/suzuito/sandbox2-go モジュールをインターネット上に公開すること
- blog CloudRun 用の Docker image を Google Container Resitory へアップロードすること

手順

1. Github action `create-release-draft` を実行する。
  a. Github action が release draft を作成する。
2. Release draft を公開する。
