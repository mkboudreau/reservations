package config

import (
	"github.com/mkboudreau/loggo"
	"os"
	"strconv"
)

var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()

const (
	ENV_SMTPHOST string = "SMTP_HOST"
	ENV_SMTPPORT        = "SMTP_PORT"
	ENV_SMTPUSER        = "SMTP_USER"
	ENV_SMTPPASS        = "SMTP_PASS"
)

type SmtpConfig struct {
	Host string
	Port int
	User string
	Pass string
}

// returns the config using values from the environment: SMTP.HOST, SMTP.PORT, SMTP.USER, SMTP.PASS
func NewSmtpConfigFromEnvironment() *SmtpConfig {
	host := os.Getenv(ENV_SMTPHOST)
	port, err := strconv.Atoi(os.Getenv(ENV_SMTPPORT))
	user := os.Getenv(ENV_SMTPUSER)
	pass := os.Getenv(ENV_SMTPPASS)

	if err != nil {
		logger.Fatal("Could not convert smtp port to integer")
	}

	return NewSmtpConfigWithValues(host, port, user, pass)
}

func NewSmtpConfigWithValues(host string, port int, user string, password string) *SmtpConfig {
	return &SmtpConfig{
		Host: host,
		Port: port,
		User: user,
		Pass: password,
	}
}
