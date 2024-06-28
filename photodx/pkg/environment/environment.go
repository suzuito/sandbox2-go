package environment

type Environment struct {
	Port          int    `required:"true" split_words:"true"`
	Env           string `required:"true" split_words:"true"`
	LogLevel      string `required:"false" split_words:"true"`
	Auth0Domain   string `required:"true" split_words:"true"`
	Auth0Audience string `required:"true" split_words:"true"`

	DBName string `required:"false" split_words:"true"`
	DBUser string `required:"false" split_words:"true"`

	CorsAllowOrigins  []string `required:"false" split_words:"true"`
	CorsAllowMethods  []string `required:"false" split_words:"true"`
	CorsAllowHeaders  []string `required:"false" split_words:"true"`
	CorsExposeHeaders []string `required:"false" split_words:"true"`

	JWTRefreshTokenSigningPrivateKey string `required:"true" split_words:"true"`
	JWTAccessTokenSigningPrivateKey  string `required:"true" split_words:"true"`
	JWTAccessTokenSigningPublicKey   string `required:"true" split_words:"true"`
}
