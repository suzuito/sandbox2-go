package environment

type Environment struct {
	DirPathArticleHTML                      string `required:"false" split_words:"true"`
	DBUser                                  string `required:"false" split_words:"true"`
	DBPassword                              string `required:"false" split_words:"true"`
	DBGCPRegion                             string `required:"false" envconfig:"DB_GCP_REGION"`
	ArticleMarkdownBucket                   string `required:"false" envconfig:"ARTICLE_MARKDOWN_BUCKET"`
	ArticleFileUploadedBucket               string `required:"false" envconfig:"ARTICLE_FILE_UPLOADED_BUCKET"` // 変数名が微妙。ユーザーがアップロードした画像を置く場所
	FunctionTriggerTopicIDStartImageProcess string `required:"false" envconfig:"FUNCTION_TRIGGER_TOPIC_ID_START_IMAGE_PROCESS"`
	SiteOrigin                              string `required:"false" split_words:"true"`
	Env                                     string `required:"true" split_words:"true"`
	AdminToken                              string `required:"true" split_words:"true"`
}
