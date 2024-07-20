package environment

type Environment struct {
	Env      string `required:"true" split_words:"true"`
	Port     int    `required:"true" split_words:"true"`
	LogLevel string `required:"false" split_words:"true"`

	DBName               string `required:"false" split_words:"true"`
	DBUser               string `required:"false" split_words:"true"`
	DBPassword           string `required:"false" split_words:"true"`
	DBInstanceUnixSocket string `required:"false" split_words:"true"`

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

	OAuth2ProviderLINEClientID     string `envconfig:"OAUTH2_PROVIDER_LINE_CLIENT_ID" required:"true"`
	OAuth2ProviderLINEClientSecret string `envconfig:"OAUTH2_PROVIDER_LINE_CLIENT_SECRET" required:"true"`
	OAuth2ProviderLINERedirectURL  string `envconfig:"OAUTH2_PROVIDER_LINE_CLIENT_REDIRECT_URL" required:"true"`

	FrontBaseURL string `required:"true" split_words:"true"`

	WebPushAPIUserVAPIDPrivateKey string `envconfig:"WEB_PUSH_API_USER_VAPID_PRIVATE_KEY" required:"true"`
	WebPushAPIUserVAPIDPublicKey  string `envconfig:"WEB_PUSH_API_USER_VAPID_PUBLIC_KEY" required:"true"`
}
