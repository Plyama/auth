package main

import (
	"github.com/plyama/auth/internal"
	"github.com/plyama/auth/internal/db"
	lgr "github.com/plyama/auth/internal/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	//appConfig, err := config.NewAppConfig(os.Getenv("CONFIG_PATH"))
	//if err != nil {
	//	log.Fatal("error while init config", zap.Error(err))
	//}

	logger, err := lgr.New(
		"auth",
		"1",
		"1",
		"1",
	)
	if err != nil {
		logger.Fatal("error while init logger", zap.Error(err))
	}

	err = godotenv.Load("env")
	if err != nil {
		logger.Fatal("failed to load env", zap.Error(err))
	}

	DB, err := db.NewPgGorm()
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	err = db.Migrate(DB)
	if err != nil {
		logger.Fatal("failed to migrate", zap.Error(err))
	}

	internal.Run(DB, logger)
}
