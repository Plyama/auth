package main

import (
	"github.com/auth/internal/config"
	lgr "github.com/auth/internal/logger"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	appConfig, err := config.NewAppConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal("error while init config", zap.Error(err))
	}

	_, err = lgr.New(
		"auth",
		appConfig.Logger.Version,
		appConfig.Logger.Env,
		appConfig.Logger.Level,
	)
	if err != nil {
		log.Fatal("error while init logger", zap.Error(err))
	}
}
