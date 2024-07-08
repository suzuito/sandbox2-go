package environment

type Environment struct {
	Port     int    `required:"true" split_words:"true"`
	Env      string `required:"true" split_words:"true"`
	LogLevel string `required:"false" split_words:"true"`

	DBName string `required:"false" split_words:"true"`
	DBUser string `required:"false" split_words:"true"`

	CorsAllowOrigins  []string `required:"false" split_words:"true"`
	CorsAllowMethods  []string `required:"false" split_words:"true"`
	CorsAllowHeaders  []string `required:"false" split_words:"true"`
	CorsExposeHeaders []string `required:"false" split_words:"true"`

	JWTAdminRefreshTokenSigningPrivateKey string `required:"true" split_words:"true"`
	JWTAdminAccessTokenSigningPrivateKey  string `required:"true" split_words:"true"`
	JWTAdminAccessTokenSigningPublicKey   string `required:"true" split_words:"true"`

	JWTUserRefreshTokenSigningPrivateKey string `required:"true" split_words:"true"`
	JWTUserAccessTokenSigningPrivateKey  string `required:"true" split_words:"true"`
	JWTUserAccessTokenSigningPublicKey   string `required:"true" split_words:"true"`
}
