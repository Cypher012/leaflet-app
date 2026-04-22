package config

type AppConfig struct {
	Port               string
	AppEnv             string
	DatabaseURL        string
	FrontendURL        string
	BackendURL         string
	CookieDomain       string
	GithubClientID     string
	GithubClientSecret string
	GithubCallbackURL  string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleCallbackURL  string
	SessionSecret      string
	R2AccountID        string
	R2Bucket           string
	R2AccessKey        string
	R2SecretKey        string
	R2PublicURL        string
}

func LoadConfig() (AppConfig, error) {
	envs, err := GetEnvs(
		"PORT",
		"APP_ENV",
		"DATABASE_URL",
		"FRONTEND_URL",
		"BACKEND_URL",
		"COOKIE_DOMAIN",
		"GITHUB_CLIENT_ID",
		"GITHUB_CLIENT_SECRET",
		"GITHUB_CALLBACK_URL",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"GOOGLE_CALLBACK_URL",
		"SESSION_SECRET",
		"R2_ACCOUNT_ID",
		"R2_BUCKET",
		"R2_ACCESS_KEY",
		"R2_SECRET_KEY",
		"R2_PUBLIC_URL",
	)
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Port:               envs["PORT"],
		AppEnv:             envs["APP_ENV"],
		DatabaseURL:        envs["DATABASE_URL"],
		FrontendURL:        envs["FRONTEND_URL"],
		BackendURL:         envs["BACKEND_URL"],
		CookieDomain:       envs["COOKIE_DOMAIN"],
		GithubClientID:     envs["GITHUB_CLIENT_ID"],
		GithubClientSecret: envs["GITHUB_CLIENT_SECRET"],
		GithubCallbackURL:  envs["GITHUB_CALLBACK_URL"],
		GoogleClientID:     envs["GOOGLE_CLIENT_ID"],
		GoogleClientSecret: envs["GOOGLE_CLIENT_SECRET"],
		GoogleCallbackURL:  envs["GOOGLE_CALLBACK_URL"],
		SessionSecret:      envs["SESSION_SECRET"],
		R2AccountID:        envs["R2_ACCOUNT_ID"],
		R2Bucket:           envs["R2_BUCKET"],
		R2AccessKey:        envs["R2_ACCESS_KEY"],
		R2SecretKey:        envs["R2_SECRET_KEY"],
		R2PublicURL:        envs["R2_PUBLIC_URL"],
	}, nil
}
