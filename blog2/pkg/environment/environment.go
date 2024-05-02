package environment

type Environment struct {
	DirPathArticleHTML    string `required:"false" split_words:"true"`
	DBName                string `required:"false" split_words:"true"`
	DBUser                string `required:"false" split_words:"true"`
	DBPassword            string `required:"false" split_words:"true"`
	DBGCPRegion           string `required:"false" envconfig:"DB_GCP_REGION"`
	DBInstanceUnixSocket  string `required:"false" split_words:"true"`
	ArticleMarkdownBucket string `required:"false" envconfig:"ARTICLE_MARKDOWN_BUCKET"`
	FileBucket            string `required:"false" envconfig:"FILE_BUCKET"`
	FileThumbnailBucket   string `required:"false" envconfig:"FILE_THUMBNAIL_BUCKET"`
	SiteOrigin            string `required:"false" split_words:"true"`
	Env                   string `required:"true" split_words:"true"`
	AdminToken            string `required:"false" split_words:"true"`
	BaseURLFile           string `required:"true" envconfig:"BASE_URL_FILE"`
	BaseURLFileThumbnail  string `required:"true" envconfig:"BASE_URL_FILE_THUMBNAIL"`
}
