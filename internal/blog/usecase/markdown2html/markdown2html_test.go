package markdown2html

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/common/test_helper"
)

func TestMarkdown2HTMLImpl_Generate_ParseMeta(t *testing.T) {
	testCases := []struct {
		desc            string
		inputSrc        string
		expectedArticle entity.Article
		expectedError   string
	}{
		{
			desc: `Success`,
			inputSrc: `---
id: "001"
tags: [タグ１,タグ２,タグ３]
version: 1
description: |-
  あいうえお。
  かきくけこ。
date: 2020-11-24
---

# これはテスト

こんにちは！世界！
`,
			expectedArticle: entity.Article{
				Title:       "これはテスト",
				ID:          entity.ArticleID("001"),
				Tags:        []entity.Tag{{ID: "タグ１"}, {ID: "タグ２"}, {ID: "タグ３"}},
				Version:     1,
				Description: "あいうえお。\nかきくけこ。",
				Date:        time.Date(2020, time.November, 24, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			desc: `Success. empty tag is OK`,
			inputSrc: `---
id: "001"
tags: []
version: 1
description: |-
  あいうえお。
  かきくけこ。
date: 2020-11-24
---

# これはテスト

こんにちは！世界！
`,
			expectedArticle: entity.Article{
				Title:       "これはテスト",
				ID:          entity.ArticleID("001"),
				Tags:        []entity.Tag{},
				Version:     1,
				Description: "あいうえお。\nかきくけこ。",
				Date:        time.Date(2020, time.November, 24, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			desc:     `Success. Validation of Article is not executed`,
			inputSrc: ``,
			expectedArticle: entity.Article{
				Tags: []entity.Tag{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mdHTML := Markdown2HTMLImpl{}
			dst := ""
			article := entity.Article{}
			err := mdHTML.Generate(ctx, tC.inputSrc, &dst, &article)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedArticle, article)
			}
		})
	}
}

// HTML生成部のテスト
// ほとんどの処理は、goldmarkライブラリにより実行されるが
// 一部、HTML生成処理を追加している。
// 追加している部分のテスト。
func TestMarkdown2HTMLImpl_Generate_HTML(t *testing.T) {
	testCases := []struct {
		desc          string
		inputSrc      string
		expectedError string
		expectedHTML  string
	}{
		{
			desc: `Success
			headingにclass="md-heading"が付与されること
			`,
			inputSrc: `
# タイトル1
テキスト1
## タイトル1-1
テキスト1-1
### タイトル1-1-1
テキスト1-1-1
`,
			expectedHTML: `<h1>Table of Contents</h1>
<ul>
<li>
<a href="#1">タイトル1</a><ul>
<li>
<a href="#1-1">タイトル1-1</a><ul>
<li>
<a href="#1-1-1">タイトル1-1-1</a></li>
</ul>
</li>
</ul>
</li>
</ul>
<h1 id="1" class="md-heading">タイトル1</h1>
<p>テキスト1</p>
<h2 id="1-1" class="md-heading">タイトル1-1</h2>
<p>テキスト1-1</p>
<h3 id="1-1-1" class="md-heading">タイトル1-1-1</h3>
<p>テキスト1-1-1</p>
`,
		},
		{
			desc: `Success aタグにtargetが付与されること（ただし、TOCのaタグにはtargetは付与されない）`,
			inputSrc: `
# タイトル

https://www.example.com
`,
			expectedHTML: `<h1>Table of Contents</h1>
<ul>
<li>
<a href="#heading">タイトル</a></li>
</ul>
<h1 id="heading" class="md-heading">タイトル</h1>
<p><a href="https://www.example.com" target="blank">https://www.example.com</a></p>
`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mdHTML := Markdown2HTMLImpl{}
			dst := ""
			article := entity.Article{}
			err := mdHTML.Generate(ctx, tC.inputSrc, &dst, &article)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedHTML, dst)
			}
		})
	}
}
