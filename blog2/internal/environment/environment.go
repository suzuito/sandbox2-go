package environment

type Environment struct {
	DirPathArticleHTML string `required:"false" split_words:"true"`
	DBUser             string `required:"false" split_words:"true"`
	DBPassword         string `required:"false" split_words:"true"`
	DBGCPRegion        string `required:"false" envconfig:"DB_GCP_REGION"`
	ArticleHTMLBucket  string `required:"false" envconfig:"ARTICLE_HTML_BUCKET"`
	SiteOrigin         string `required:"false" split_words:"true"`
	Env                string `required:"true" split_words:"true"`
	AdminToken         string `required:"true" split_words:"true"`
}
