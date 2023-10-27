#
# common
#
common-test:
	sh test.sh ./internal/common/...
common-mock:
	./mockgen internal/common/cusecase/clog/logger.go

#
# blog
#
DB_NAME = blog
DB_NAME_FOR_UNIT_TEST = blog_for_unit_test
blog-init:
	docker compose up blog-mysql -d
	while [ true ]; do mysql -u root -h 127.0.0.1 -e 'show databases' > /dev/null 2>&1 && echo 'DB connection is OK' && break; echo 'Waiting until DB connection is OK' && sleep 1; done
	mysql -u root -h 127.0.0.1 -e "create database if not exists $(DB_NAME)"
	mysql -u root -h 127.0.0.1 -e "create database if not exists $(DB_NAME_FOR_UNIT_TEST)"
blog-build:
	go build -o blog-server.exe internal/blog/cmd/server/*.go
blog-test:
	sh test.sh ./internal/blog/...
blog-init-rdb:
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" drop -f
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME)" up
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" drop -f
	migrate -source file://./.service/blog/.schema/ -database "mysql://root:@tcp(127.0.0.1:3306)/$(DB_NAME_FOR_UNIT_TEST)" up
blog-mock:
	./mockgen internal/blog/usecase/repository_article_html.go
	./mockgen internal/blog/usecase/repository_article_source.go
	./mockgen internal/blog/usecase/repository_article.go
	./mockgen internal/blog/usecase/markdown2html/markdown2html.go
	./mockgen internal/blog/usecase/usecase.go
	./mockgen internal/blog/web/presenter.go
	./mockgen internal/blog/web/presenters.go
