OS=darwin
ARCH=amd64

#
# Tool chain
#
.bin/migrate:
	cd /tmp && rm migrate && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.${OS}-${ARCH}.tar.gz | tar xvz && cd - && mv /tmp/migrate .bin/migrate

#
# common
#
common-test:
	sh test.sh ./common/...
common-mock:
	./mockgen common/cusecase/clog/logger.go
	./mockgen common/httpclientcache/client.go

#
# blog
#
DB_NAME = blog
DB_NAME_FOR_UNIT_TEST = blog_for_unit_test
blog-init:
	docker compose up -d blog-mysql
	while [ true ]; do mysql -u root -h 127.0.0.1 -e 'show databases' > /dev/null 2>&1 && echo 'DB connection is OK' && break; echo 'Waiting until DB connection is OK' && sleep 1; done
	mysql -u root -h 127.0.0.1 -e "create database if not exists $(DB_NAME)"
	mysql -u root -h 127.0.0.1 -e "create database if not exists $(DB_NAME_FOR_UNIT_TEST)"
blog-build:
	go build -o blog-server.exe blog/cmd/server/*.go
blog-test:
	sh test.sh ./blog/...
blog-init-rdb: .bin/migrate
	.bin/migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" drop -f
	.bin/migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" up
	.bin/migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" drop -f
	.bin/migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" up
blog-mock:
	./mockgen blog/usecase/repository_article_html.go
	./mockgen blog/usecase/repository_article_source.go
	./mockgen blog/usecase/repository_article.go
	./mockgen blog/usecase/markdown2html/markdown2html.go
	./mockgen blog/usecase/usecase.go
	./mockgen blog/web/presenter.go
	./mockgen blog/web/presenters.go

#
# blog2
#
BLOG2_DB_NAME = blog2-dev
BLOG2_DB_NAME_FOR_UNIT_TEST = blog2-dev_for_unit_test
blog2-init:
	docker compose up -d blog2-mysql
	DB_HOST=127.0.0.1 DB_PORT=3307 sh wait-until-db-open.sh
blog2-init-rdb:
	mysql -u root -h 127.0.0.1 -P 3307 -e "create database if not exists $(BLOG2_DB_NAME)"
	mysql -u root -h 127.0.0.1 -P 3307 -e "create database if not exists $(BLOG2_DB_NAME_FOR_UNIT_TEST)"
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME)" drop -f
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME)" up
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME_FOR_UNIT_TEST)" drop -f
	.bin/migrate -source file://./.service/blog2/.schema/ -database "mysql://root:@tcp(127.0.0.1:3307)/$(BLOG2_DB_NAME_FOR_UNIT_TEST)" up
blog2-init-rdb-test-data:
	go run blog2/cmd/create-test-data/main.go
blog2-migrate-create:
	# Example: make blog2-migrate-create NAME=create_article
	.bin/migrate create -ext sql -dir ./.service/blog2/.schema/ $(NAME)
blog2-build-server:
	go build -o blog2-server.out blog2/cmd/server/*.go
	GOOS=js GOARCH=wasm go build -o blog2/internal/web/_wasm/page_admin_article.wasm blog2/wasm/page_admin_article/*.go
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js blog2/internal/web/_wasm/wasm_exec.js

#
# crawler
#
crawler-init:
	docker compose up -d crawler-firebase-emulator
	while [ true ]; do curl http://localhost:8082 > /dev/null 2>&1 && echo 'Firebase emulator connection is OK' && break; echo 'Waiting until Firebase emulator connection is OK' && sleep 1; done
crawler-mock:
	./mockgen crawler/pkg/entity/crawler/fetcher.go
	./mockgen crawler/pkg/entity/crawler/parser.go
	./mockgen crawler/pkg/entity/crawler/publisher.go
	./mockgen crawler/pkg/entity/notifier/notifier.go
	./mockgen crawler/internal/usecase/factory/crawler.go
	./mockgen crawler/internal/usecase/factory/notifier.go
	./mockgen crawler/internal/usecase/repository/crawler.go
	./mockgen crawler/internal/usecase/repository/time_series_data.go
	./mockgen crawler/internal/usecase/repository/notifier.go
	./mockgen crawler/internal/usecase/repository/crawler_configuration.go
	./mockgen crawler/internal/usecase/queue/trigger_crawler.go
	./mockgen crawler/internal/usecase/discord/discord_go_session.go
	./mockgen crawler/internal/infra/fetcher/httpclientwrapper/http_client_wrapper.go
crawler-test:
	sh test.sh ./crawler/...