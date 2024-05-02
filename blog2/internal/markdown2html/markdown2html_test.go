package markdown2html

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

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
			expectedHTML: `<h1 id="table-of-contents">Table of Contents</h1>
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
			expectedHTML: `<h1 id="table-of-contents">Table of Contents</h1>
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
			err := mdHTML.Generate(ctx, tC.inputSrc, &dst)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, tC.expectedHTML, dst)
			}
		})
	}
}
