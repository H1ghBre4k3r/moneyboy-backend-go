package variables

import "os"

type tokenVariables struct {
	RefreshTokenSecret string
	AccessTokenSecret  string
	VerifyTokenSecret  string
}

type databaseVariables struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type mailAuthVariables struct {
	User string
	Pass string
}

type mailVariables struct {
	Host string
	Port string
	Addr string
	Auth mailAuthVariables
}

var (
	HOST  = os.Getenv("HOST")
	TOKEN = tokenVariables{
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		VerifyTokenSecret:  os.Getenv("VERIFY_TOKEN_SECRET"),
	}
	DATABASE = databaseVariables{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Name:     os.Getenv("MYSQL_DATABASE"),
	}
	MAIL = mailVariables{
		Host: os.Getenv("SMTP_HOST"),
		Port: os.Getenv("SMTP_PORT"),
		Addr: os.Getenv("SMTP_ADDR"),
		Auth: mailAuthVariables{
			User: os.Getenv("SMTP_USER"),
			Pass: os.Getenv("SMTP_PASS"),
		},
	}
)
