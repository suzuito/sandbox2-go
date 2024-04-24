package environment

type Environment struct {
	DirPathArticleHTML    string `required:"false" split_words:"true"`
	DBName                string `required:"false" split_words:"true"`
	DBUser                string `required:"false" split_words:"true"`
	DBPassword            string `required:"false" split_words:"true"`
	DBGCPRegion           string `required:"false" envconfig:"DB_GCP_REGION"`
	DBInstanceUnixSocket  string `required:"false" split_words:"true"`
	ArticleMarkdownBucket string `required:"false" envconfig:"ARTICLE_MARKDOWN_BUCKET"`
	FileUploadedBucket    string `required:"false" envconfig:"FILE_UPLOADED_BUCKET"` // 変数名が微妙。ユーザーがアップロードした画像を置く場所
	FileBucket            string `required:"false" envconfig:"FILE_BUCKET"`
	SiteOrigin            string `required:"false" split_words:"true"`
	Env                   string `required:"true" split_words:"true"`
	AdminToken            string `required:"false" split_words:"true"`
}
