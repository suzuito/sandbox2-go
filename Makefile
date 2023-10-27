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
	docker compose up -d blog-mysql
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

#
# crawler
#
crawler-init:
	docker compose up -d crawler-firebase-emulator
	while [ true ]; do curl http://localhost:8082 > /dev/null 2>&1 && echo 'Firebase emulator connection is OK' && break; echo 'Waiting until Firebase emulator connection is OK' && sleep 1; done
crawler-mock:
	./mockgen internal/crawler/crawler/internal/entity/crawler/crawler.go
	./mockgen internal/crawler/crawler/internal/entity/crawler/fetcher.go
	./mockgen internal/crawler/crawler/internal/entity/crawler/parser.go
	./mockgen internal/crawler/crawler/internal/entity/crawler/publisher.go
	./mockgen internal/crawler/crawler/internal/usecase/repository/repository.go
	./mockgen internal/crawler/crawler/internal/usecase/queue/queue.go
	./mockgen internal/crawler/crawler/internal/usecase/crawlerfactory/crawlerfactory.go
	./mockgen internal/crawler/notifier/internal/entity/notifier/notifier.go
	./mockgen internal/crawler/notifier/internal/usecase/repository/repository.go
	./mockgen internal/crawler/notifier/internal/usecase/notifierfactory/notifierfactory.go
	./mockgen internal/crawler/notifier/internal/usecase/discord/discordgo_session.go
crawler-test:
	sh test.sh ./internal/crawler/...