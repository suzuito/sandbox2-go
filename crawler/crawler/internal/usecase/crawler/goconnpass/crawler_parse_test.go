package goconnpass

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/internal/constant"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

var body = `
{
	"results_start": 1,
	"results_returned": 10,
	"results_available": 10000,
	"events": [
	  {
		"event_id": 299921,
		"title": "多様な集団でも分析活用できる！高い分類精度と再現性をもつクラスタリング",
		"catch": "～非階層クラスター分析 k-umeyama～",
		"description": "<h2>イベント概要</h2>\n<p>本イベントでは、k-umeyamaという革新的なクラスタリング手法に焦点を当て、既存手法と比較したときの、その精度の高さを実証します。\nk-umeyamaは、データ分析における新たな可能性を切り拓く手法で、データの相関性を排除せずに、高度なクラスタリングを実現します。\nまた、シグモイド関数を活用してシード選択過程を改良し、データ分類の精度を向上させました。\n本ウェビナーでは、Pythonを用いて記述したk-umeyamaの挙動を確認し、高い成果が期待できる分野、課題を確認します。</p>\n<h2>このような方にオススメです！</h2>\n<p>＜データ分析や機械学習に興味を持つ、以下の対象者向けです＞</p>\n<p>●データサイエンティスト: データ分析の専門家、データサイエンスの実務家</p>\n<p>●エンジニア：クラスタリング手法に関する知識を深めたいエンジニアや研究者</p>\n<p>●マーケティング関係者：マーケティング戦略の最適化に興味を持つマーケッターや担当者</p>\n<p>●アナリスト: データ分析を通じてビジネスの意思決定に貢献したいアナリストや経営者</p>\n<p>※Pythonでの開発や実装経験がある方であれば、より理解が深まります</p>\n<h2>講演者</h2>\n<p>梅山貴彦</p>\n<p>日本マーケティング・リサーチ協会（JMRA）リサーチ・イノベーション委員会委員長 マーケティングリサーチ歴20年以上。</p>\n<p>日本行動計量学会所属。</p>\n<p>株式会社クロス・マーケティング 取締役</p>\n<p>リサーチ・ソリューション本部副本部長</p>\n<p>クロス・マーケティング社リサーチ部門の部門長を務めつつ、Pythonを用いた分析手法の開発も主導する。</p>\n<p>リサーチにおいては、IT分野、EC、消費財、自動車等の業界を広範囲に担当。</p>\n<p>各種の多変量解析に精通し、購買行動モデルの構築等を通じて、多数のクライアントの意思決定に貢献している。</p>\n<h2>司会者</h2>\n<p>水原 亮</p>\n<p>株式会社クロス・マーケティング</p>\n<p>リサーチ・ソリューション本部RDX推進部マネージャー</p>\n<p>マーケティングリサーチ歴15年。</p>\n<p>通信、IT、エンタメを中心としたマーケティングリサーチ以外にも、行政案件、学術調査等の社会調査に広く精通。</p>\n<p>リサーチ実務以外で、新規サービス開発や100万人規模のリサーチパネルの改善・運用・品質管理の担当経験あり。</p>\n<p>現在は、社内のDX推進担当。</p>",
		"event_url": "https://connpass.com/event/299921/",
		"started_at": "2023-11-02T14:00:00+09:00",
		"ended_at": "2023-11-02T15:00:00+09:00",
		"limit": 50,
		"hash_tag": "",
		"event_type": "participation",
		"accepted": 0,
		"waiting": 0,
		"updated_at": "2023-10-19T19:37:52+09:00",
		"owner_id": 818418,
		"owner_nickname": "cross-marke",
		"owner_display_name": "cross-marke",
		"place": "オンライン",
		"address": "オンライン",
		"lat": null,
		"lon": null,
		"series": null
	  },
	  {
		"event_id": 299434,
		"title": "Gストラング線形代数勉強会",
		"catch": "",
		"description": "<h1>概要</h1>\n<h2>データ分析のスキルと知識を向上したい人のために楽しく学びあうグループを目指します。</h2>\n<h2>教科書：線形代数イントロダクション 第4版 G. ストラング</h2>\n<p>G. ストラングの「線形代数イントロダクション」</p>\n<p>大まかな課題・内容はグループで決め、細かな内容を運営者が配信します。QiitaとGithubを活用します。</p>\n<p>主催者が教科書を一行一行読み上げ、参加者の皆さんに理解の状況を確認します。疑問点は参加者全員で明確にしていきます。</p>\n<h2>進行の仕方</h2>\n<p>Zoomを用いてオンラインで会議形式で進めます。\n1人でも参加者がいれば勉強会は実施いたします。大体1週間に1度、1時間半行います。</p>\n<p>第3章 ベクトル空間と部分空間  </p>\n<p><del>3.6 4つの部分空間の次元 2021/7/１</del>   <br>\n第4章 直交性<br>\n<del>4.1 4つの部分空間の直交性2021/7/8,15,22,29,8/5,12,26,9/2</del>  <br>\n<del>4.2 射影  2021/9/2,9,16</del><br>\n<del>4.3 最小二乗近似 2021/9/22,30,10/7,/14</del> <br>\n<del>4.4 直交基底とグラム-シュミット法 2021/10/14,21,28,11/4,11,18</del>        </p>\n<p>第5章 行列式<br>\n<del>5.1 行列式の性質 2021/11/25,12/2,9,16,23</del>  <br>\n<del>5.2 置換と余因子 2021/12/23,2022/1/13,20,27</del> <br>\n<del>5.3 クラメルの定理、逆行列、体積  2021/2/3,10,17,24</del>  </p>\n<p>第6章 固有値と固有ベクトル<br>\n<del>6.1 固有値と固有ベクトル 2022/3/3,10,17,24,31</del>  <br>\n<del>6.2 行列の対角化 2022/3/31, /4/7,14,21,28,5/12</del>    <br>\n<del>6.3 微分方程式への応用 5/12,19,26, 6/2/9,16,23,30,7/7,7/14,21,28</del>  <br>\n<del>6.4 対称行列 2022/8/4,18,25,9/1,8</del>       <br>\n<del>4.2 射影と4.3 最小二乗近似にしばらく戻ります。2022/9/15,22,29,2022/10/6,13,20,27</del>    <br>\n<del>6.5 正定値行列 2022/11/3,10,17,24,2022/12/1,8</del><br>\n<del>6.6 相似行列2022/12/8,15,2023/2/2,9,16,23,2023/3/2,9,16</del>    <br>\n<del>6.7 特異値分解2023/0309,16,23,30,5/4,11</del>  <br>\n<del>7.1 線形変換の概念 2023/5/18,25, 6/1</del>  <br>\n<del>7.2 線形変換の行列 2023/6/7,15,22,29,7/6,13,20,27,8/3,10,17,24,31,9/7,14,21</del><br>\n7.3 対角化と疑似逆行列 2012/10/12,19</p>\n<p><a href=\"https://qiita.com/innovation1005/items/cb2a8d724aea6f915904\" rel=\"nofollow\">Python3ではじめるシステムトレード: 行列の構成</a></p>\n<p><a href=\"https://qiita.com/innovation1005/items/50558d102bcc86efe698\" rel=\"nofollow\">Python3ではじめるシステムトレード：特異値分解</a></p>\n<p><a href=\"https://qiita.com/innovation1005/items/2e8960ef6fa775f3e201\" rel=\"nofollow\">Python3ではじめるシステムトレード：固有値と固有ベクトル入門</a></p>\n<p><a href=\"https://qiita.com/innovation1005/items/99b3159627d6932e22b3\" rel=\"nofollow\">python3ではじめるシステムトレード：射影と最小二乗法</a></p>\n<p><a href=\"https://qiita.com/innovation1005/items/4770d7f693c83420dbc6\" rel=\"nofollow\">Python3ではじめるシステムトレード: PCA</a></p>\n<h3>チャットに参加、聴講のみも可能です。</h3>\n<p>また、不明な点などについてはh.moriya@quasars22.co.jp にご連絡いただければと思います。</p>\n<h2>準備</h2>\n<p>開始前までにZoomというオンライ会議システムをインストールしてください。\n開始時刻前に参加IDとパスワードを\"参加者への情報\"の欄に張ります。無料のzoom用のIDとパスワードを発行します。セッションは３回行います。一回のセッションは４０分で切れますが、３回のセッションは同じIDとパスワードを使います。</p>\n<h2>対象者</h2>\n<p>線形代数について理論的な背景をきちんと学びたい方  </p>\n<h2>運営</h2>\n<p>森谷博之</p>",
		"event_url": "https://study-data-analysis.connpass.com/event/299434/",
		"started_at": "2023-10-19T19:00:00+09:00",
		"ended_at": "2023-10-19T20:30:00+09:00",
		"limit": 20,
		"hash_tag": "",
		"event_type": "participation",
		"accepted": 6,
		"waiting": 0,
		"updated_at": "2023-10-19T17:44:33+09:00",
		"owner_id": 577436,
		"owner_nickname": "full_of_flower",
		"owner_display_name": "full_of_flower",
		"place": "zoom",
		"address": "",
		"lat": null,
		"lon": null,
		"series": {
		  "id": 11670,
		  "title": "データの分析方法を学ぶ会",
		  "url": "https://study-data-analysis.connpass.com/"
		}
	  }
	]
}
`

func Test(t *testing.T) {
	ctx := context.Background()
	crwl := NewCrawler(nil, nil)
	data, err := crwl.Parse(ctx, bytes.NewReader([]byte(body)), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Fail()
		return
	}
	assert.Equal(t, 2, len(data))
	expected := []struct {
		ID   timeseriesdata.TimeSeriesDataID
		Date time.Time
	}{
		{
			ID:   timeseriesdata.TimeSeriesDataID("connpass-299921"),
			Date: time.Date(2023, time.November, 2, 14, 0, 0, 0, constant.JST),
		},
		{
			ID:   timeseriesdata.TimeSeriesDataID("connpass-299434"),
			Date: time.Date(2023, time.October, 19, 19, 0, 0, 0, constant.JST),
		},
	}
	for i := range expected {
		assert.Equal(t, expected[i].ID, data[i].GetID())
		assert.Equal(t, expected[i].Date, data[i].GetPublishedAt())
	}
}
