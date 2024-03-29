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