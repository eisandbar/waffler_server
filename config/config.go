package config

import (
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type AppConf struct {
	AppName     string `yaml:"app_name"`
	Server      Server `yaml:"server"`
	Logger      Logger `yaml:"logger"`
	DatabaseURL string `yaml:"database_url"`
	Token       Token  `yaml:"token"`
}

type Token struct {
	AccessTTL     time.Duration `yaml:"access_ttl"`
	RefreshTTL    time.Duration `yaml:"refresh_ttl"`
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

func NewAppConf() AppConf {
	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
		},
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func (a *AppConf) Init(logger *zap.Logger) {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		logger.Fatal("config: parse server shutdown timeout error")
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second
	if err != nil {
		logger.Fatal("config: parse rpc server shutdown timeout error")
	}

	a.Server.ShutdownTimeout = shutDownTimeout
}
