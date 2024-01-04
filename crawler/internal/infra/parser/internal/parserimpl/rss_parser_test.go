package parserimpl

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/constant"
	"github.com/suzuito/sandbox2-go/common/test_helper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

func TestRSSParser(t *testing.T) {
	testCases := []struct {
		desc          string
		inputR        string
		expected      []*timeseriesdata.TimeSeriesDataBlogFeed
		expectedError string
	}{
		{
			desc: "Success",
			inputR: `
			<?xml version="1.0" encoding="UTF-8"?>
			<rss xmlns:webfeeds="http://webfeeds.org/rss/1.0" xmlns:note="https://note.com" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" version="2.0">
			  <channel>
				<title>注目</title>
				<description>noteの注目記事</description>
				<link>https://note.com/recommend</link>
				<atom:link rel="self" type="application/rss+xml" href="https://note.com/recommend/rss/"/>
				<copyright>(C) note inc.</copyright>
				<webfeeds:icon>https://d2l930y2yx77uc.cloudfront.net/assets/default/default_note_logo_202212-f2394a9e5b60c49f48650eee13f6e75987c8c4f1cfa7555629a9697dc6015cd9.png</webfeeds:icon>
				<webfeeds:logo>https://d2l930y2yx77uc.cloudfront.net/assets/default/default_note_logo_202212-f2394a9e5b60c49f48650eee13f6e75987c8c4f1cfa7555629a9697dc6015cd9.png</webfeeds:logo>
				<webfeeds:accentColor>249F80</webfeeds:accentColor>
				<webfeeds:related layout="card" target="browser"/>
				<webfeeds:analytics id="UA-48687000-1" engine="GoogleAnalytics"/>
				<language>ja</language>
				<lastBuildDate>Tue, 07 Nov 2023 18:21:30 +0900</lastBuildDate>
				<item>
				  <title>もしも猫展</title>
				  <media:thumbnail>https://assets.st-note.com/production/uploads/images/119934613/rectangle_large_type_2_b0777e430e38b6df6e62ac9dac71bc1c.jpeg?width=800</media:thumbnail>
				  <description><![CDATA[<p name="317a7ac4-f384-4b2c-85b3-7849d2892724" id="317a7ac4-f384-4b2c-85b3-7849d2892724">京都文化博物館で「もしも猫展」を見てきました。文博へは烏丸で下りて北の方へ歩いて行きました。これが真夏だったら10分間程度歩くのも死にそうなところですが、今日は秋晴れの爽やかな良い天気でしたので歩くのも全く問題なしでした。</p><figure name="297012e2-34af-4028-a433-415659b2f1ba" id="297012e2-34af-4028-a433-415659b2f1ba" data-src="https://www.bunpaku.or.jp/exhi_special_post/20230923-1112/" data-identifier="null" embedded-service="external-article" embedded-content-key="emb6a929ac55bbe"> <a href="https://www.bunpaku.or.jp/exhi_special_post/20230923-1112/" rel="nofollow noopener" target="_blank"><strong>もしも猫展 - 京都府京都文化博物館</strong><em>基本情報 展示構成と主な出品作品 関連イベント 音声ガイド 展覧会図録 グッズ 基本情報 会期 2023年9月</em><em>www.bunpaku.or.jp</em></a><a href="https://www.bunpaku.or.jp/exhi_special_post/20230923-1112/" rel="nofollow noopener" target="_blank"></a> </figure><br/><a href='https://note.com/swingmammav2/n/n6e38203a479c'>続きをみる</a>]]></description>
				  <note:creatorImage>https://assets.st-note.com/production/uploads/images/105349869/profile_b3f6e12a7abf0c6918581c9648d1139b.png?fit=bounds&amp;format=jpeg&amp;quality=85&amp;width=330</note:creatorImage>
				  <note:creatorName>綾小路</note:creatorName>
				  <pubDate>Thu, 26 Oct 2023 20:18:24 +0900</pubDate>
				  <link>https://note.com/swingmammav2/n/n6e38203a479c</link>
				  <guid>https://note.com/swingmammav2/n/n6e38203a479c</guid>
				</item>
				<item>
				  <title>身の回りのユニバーサルデザイン〜ボードゲーム編〜</title>
				  <media:thumbnail>https://assets.st-note.com/production/uploads/images/120379021/rectangle_large_type_2_86243038e791f1a48c1ea786a2ec5b00.png?width=800</media:thumbnail>
				  <description><![CDATA[<p name="d2e267cb-0571-445b-b490-6369bacfa171" id="d2e267cb-0571-445b-b490-6369bacfa171">こんにちは、アジケのブログチームです。</p><p name="f2aa9082-826f-448f-9cf3-dc85fa23a3cd" id="f2aa9082-826f-448f-9cf3-dc85fa23a3cd">前回の<a href="https://note.com/ajike/n/n9a086e08bf7c" target="_blank" rel="nofollow noopener">「ボトルデザイン」のUI考察</a>から、身の回りには色々な工夫を凝らした商品があることを実感しました。<br>そこで、今回はアクセシビリティ観点に着目し、年齢や障がいの有無などにかかわらず、誰もが楽しめる「ユニバーサルデザインなボードゲーム」を紹介します！</p><br/><a href='https://note.com/ajike/n/n94f04dca99a8'>続きをみる</a>]]></description>
				  <note:creatorImage>https://assets.st-note.com/production/uploads/images/7135456/profile_e16250bc3e667a0c459159c4047bc262.jpg?fit=bounds&amp;format=jpeg&amp;quality=85&amp;width=330</note:creatorImage>
				  <note:creatorName>ajike丨UX Design</note:creatorName>
				  <pubDate>Thu, 02 Nov 2023 10:00:18 +0900</pubDate>
				  <link>https://note.com/ajike/n/n94f04dca99a8</link>
				  <guid>https://note.com/ajike/n/n94f04dca99a8</guid>
				</item>
			  </channel>
			</rss>
			`,
			expected: []*timeseriesdata.TimeSeriesDataBlogFeed{
				{
					ID:          "https---note.com-swingmammav2-n-n6e38203a479c",
					PublishedAt: time.Date(2023, 10, 26, 20, 18, 24, 0, constant.JST),
					Title:       "もしも猫展",
					URL:         "https://note.com/swingmammav2/n/n6e38203a479c",
				},
				{
					ID:          "https---note.com-ajike-n-n94f04dca99a8",
					PublishedAt: time.Date(2023, 11, 2, 10, 0, 18, 0, constant.JST),
					Title:       "身の回りのユニバーサルデザイン〜ボードゲーム編〜",
					URL:         "https://note.com/ajike/n/n94f04dca99a8",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			parser := RSSParser{
				FP: &gofeed.Parser{},
			}
			data, err := parser.Do(
				context.Background(),
				bytes.NewBufferString(tC.inputR),
				nil,
			)
			test_helper.AssertError(t, tC.expectedError, err)
			if err == nil {
				assert.Equal(t, len(tC.expected), len(data))
				for i, expected := range tC.expected {
					real := data[i].(*timeseriesdata.TimeSeriesDataBlogFeed)
					assert.Equal(t, expected.ID, real.ID)
					assert.Equal(t, expected.PublishedAt, real.PublishedAt)
					assert.Equal(t, expected.Title, real.Title)
					assert.Equal(t, expected.URL, real.URL)
				}
			}
		})
	}
}
