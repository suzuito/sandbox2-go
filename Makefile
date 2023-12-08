#
# common
#
common-test:
	sh test.sh ./common/...
common-mock:
	./mockgen common/cusecase/clog/logger.go

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
blog-init-rdb:
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" drop -f
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" up
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" drop -f
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" up
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
crawler-test:
	sh test.sh ./crawler/...