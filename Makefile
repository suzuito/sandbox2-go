OS=darwin
ARCH=amd64

#
# Tool chain
#
.bin/migrate:
	cd /tmp && rm migrate && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.${OS}-${ARCH}.tar.gz | tar xvz && cd - && mv /tmp/migrate .bin/migrate

PROTOC_VERSION = 27.2
PROTOC_OS_ARCH := osx-aarch_64
PROTOC_GO_VERSION = 1.34.2
.bin/protoc:
	rm -rf /tmp/protoc-download \
	&& mkdir /tmp/protoc-download \
	&& cd /tmp/protoc-download \
	&& curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS_ARCH}.zip \
	&& unzip protoc-${PROTOC_VERSION}-${PROTOC_OS_ARCH}.zip \
	&& cd - && mv /tmp/protoc-download/bin/protoc ./.bin/protoc
	&& go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GO_VERSION}

#
# common
#
common-test:
	sh test.sh ./common/...
common-mock:
	./mockgen common/cusecase/clog/logger.go
	./mockgen common/httpclientcache/client.go

#
# blog2
#
BLOG2_DB_NAME = blog2-dev
BLOG2_DB_NAME_FOR_UNIT_TEST = blog2-dev_for_unit_test
blog2-init:
	docker compose up -d blog2-mysql
	DB_HOST=127.0.0.1 DB_PORT=3307 sh wait-until-db-open.sh
blog2-init-rdb: .bin/migrate
	mysql -u root -h 127.0.0.1 -P 3307 -e 'create database if not exists `$(BLOG2_DB_NAME)`'
	mysql -u root -h 127.0.0.1 -P 3307 -e 'create database if not exists `$(BLOG2_DB_NAME_FOR_UNIT_TEST)`'
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME)" drop -f
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME)" up
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME_FOR_UNIT_TEST)" drop -f
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME_FOR_UNIT_TEST)" up
blog2-init-rdb-test-data:
	go run blog2/cmd/create-test-data/main.go
blog2-migrate-create: .bin/migrate
	# Example: make blog2-migrate-create NAME=create_article
	.bin/migrate create -ext sql -dir ./.service/blog2/.schema/ $(NAME)
blog2-build-server:
	go build -o blog2-server.out blog2/cmd/server/*.go
blog2-mock:
	./mockgen blog2/internal/markdown2html/markdown2html.go

#
# photodx/
#
photodx/protoc: ./.bin/protoc
	PATH=$(PATH):$(shell go env GOPATH)/bin ./.bin/protoc \
	--go_out=./ \
	./photodx/proto/entity/permission.proto

#
# photodx/db
#
PHOTODX_DB_NAME = photodx
photodx/db-init:
	docker compose up -d photodx-mysql
	DB_HOST=127.0.0.1 DB_PORT=3308 sh wait-until-db-open.sh
photodx/db-init-rdb: .bin/migrate
	mysql -u root -h 127.0.0.1 -P 3308 -e 'create database if not exists `$(PHOTODX_DB_NAME)`'
	.bin/migrate -source file://./.service/photodx/db/.schema/ -database "mysql://root:@tcp(127.0.0.1:3308)/$(PHOTODX_DB_NAME)" drop -f
	.bin/migrate -source file://./.service/photodx/db/.schema/ -database "mysql://root:@tcp(127.0.0.1:3308)/$(PHOTODX_DB_NAME)" up
	mysql -u root -h 127.0.0.1 -P 3308 -D $(PHOTODX_DB_NAME) < ./.service/photodx/db/localdata.sql

#
# photodx/bff
#
photodx/bff-init: photodx/db-init
	echo "TODO"
photodx/bff-init-rdb: photodx/db-init-rdb
	echo "TODO"
photodx/bff-build:
	go build -o photodx-bff.out photodx/cmd/bff/*.go
photodx/migrate-create: .bin/migrate
	# Example: make blog2-migrate-create NAME=create_article
	.bin/migrate create -ext sql -dir ./.service/photodx/db/.schema/ $(NAME)