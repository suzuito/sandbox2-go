package viewmodel

import "github.com/suzuito/sandbox2-go/blog2/internal/entity"

type ComponentWasm struct {
	WasmBinaryURL string
}

type PageAdminArticle struct {
	ComponentCommonHead ComponentCommonHead
	ComponentWasm       ComponentWasm
	Article             *entity.Article
}
