package note

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/constant"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

func TestSubString(t *testing.T) {
	assert.Equal(t, "ab", subString("abc", 2))
	assert.Equal(t, "a", subString("a", 2))
}

func TestParse(t *testing.T) {
	testCases := []struct {
		desc        string
		inputR      io.Reader
		expected    NoteArticle
		expectedErr string
	}{
		{
			desc: "Success",
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expected: NoteArticle{
				Title:          "title1",
				Description:    "desc1",
				URL:            "https://www.example.com/v1",
				ArticleContent: "\n\t\t\t\t\t\tThis is content\n\t\t\t\t\t",
				ImageURL:       "https://www.example.com/v2",
				PublishedAt:    time.Date(2023, time.September, 29, 15, 0, 0, 0, constant.JST),
				Tags: []NoteArticleTag{
					{Name: "デザイン"},
					{Name: "デザイナー"},
					{Name: "ナレッジワーク"},
				},
			},
		},
		{
			desc: `Success
			Get content as a description if tag having 'meta[name=description]' is not found
			`,
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<!-- meta data-n-head="ssr" data-hid="description" name="description" content="desc1" -->
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expected: NoteArticle{
				Title:          "title1",
				Description:    "\n\t\t\t\t\t\tThis is content\n\t\t\t\t\t",
				URL:            "https://www.example.com/v1",
				ArticleContent: "\n\t\t\t\t\t\tThis is content\n\t\t\t\t\t",
				ImageURL:       "https://www.example.com/v2",
				PublishedAt:    time.Date(2023, time.September, 29, 15, 0, 0, 0, constant.JST),
				Tags: []NoteArticleTag{
					{Name: "デザイン"},
					{Name: "デザイナー"},
					{Name: "ナレッジワーク"},
				},
			},
		},
		{
			desc: `Error if title tag is not found in HTML`,
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<!-- <title>title1</title> -->
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `Cannot find title tag`,
		},
		{
			desc: `Error if meta tag for canonical URL is not found in HTML`,
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<!-- link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1" -->
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `Cannot find link\[rel=canonical\] tag`,
		},
		{
			desc: `Error if meta tag for canonical URL does not have href attribute in HTML`,
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `Cannot find href attr of link\[rel=canonical\] tag`,
		},
		{
			desc: "Error if tag having .p-article__content is not found",
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<!-- div class="p-article__content" data-v-c5502208>
						This is content
					</div -->
			</body>
			</html>
			`),
			expectedErr: `Cannot find html tag of \.p-article__content`,
		},
		{
			desc: "Error if tag having .p-article__content is not found",
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<!-- div class="o-noteContentHeader__name">
						<time datetime="2023-09-29T15:00:00.000+09:00">2023年9月29日 15:00</time>
					</div -->
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `Cannot find html tag of '\.o-noteContentHeader__info time'`,
		},
		{
			desc: "Error if tag having .p-article__content is not found",
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time>2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `Cannot find datetime attr in html tag of '\.o-noteContentHeader__info time'`,
		},
		{
			desc: "Error if tag having .p-article__content is not found",
			inputR: bytes.NewBufferString(`
			<!doctype html>
			<html data-n-head-ssr lang="ja" data-n-head="%7B%22lang%22:%7B%22ssr%22:%22ja%22%7D%7D">
			<head>
				<title>title1</title>
				<link data-n-head="ssr" rel="canonical" href="https://www.example.com/v1">
				<meta data-n-head="ssr" data-hid="description" name="description" content="desc1">
				<meta data-n-head="ssr" data-hid="og:image" property="og:image" content="https://www.example.com/v2">
			</head>
			<body>
				<div class="o-noteContentHeader__info">
					<div class="o-noteContentHeader__name">
						<time datetime="xxx">2023年9月29日 15:00</time>
					</div>
					<ul id="tagListBody" class="m-tagList__body">
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイン
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#デザイナー
								<!---->
							</div>
						</li>
						<li tabindex="-1" class="m-tagList__item" style="display:;">
							<div class="a-tag__label">
								#ナレッジワーク
							</div>
						</li>
					</ul>
					<div class="p-article__content" data-v-c5502208>
						This is content
					</div>
			</body>
			</html>
			`),
			expectedErr: `parsing time "xxx" as "2006-01-02T15:04:05Z07:00": cannot parse "xxx" as "2006"`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := Parser{}
			article, err := p.Parse(context.Background(), tC.inputR)
			test_helper.AssertError(t, tC.expectedErr, err)
			if err == nil {
				assert.Equal(t, *article, tC.expected)
			}
		})
	}
}
