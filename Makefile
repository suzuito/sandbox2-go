test:
	go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...
	go test -timeout 30s -coverprofile=coverage.pre.txt -covermode=atomic ./...
	grep -v "_mock.go" coverage.pre.txt > coverage.txt
	go tool cover -html=coverage.txt -o coverage.html
	go tool cover -func coverage.txt
mock:
	./mockgen internal/common/cusecase/clog/logger.go
	./mockgen internal/blog/usecase/repository_article_html.go
	./mockgen internal/blog/usecase/repository_article_source.go
	./mockgen internal/blog/usecase/repository_article.go
	./mockgen internal/blog/usecase/markdown2html/markdown2html.go
	./mockgen internal/blog/usecase/usecase.go
	./mockgen internal/blog/web/presenter.go
	./mockgen internal/blog/web/presenters.go