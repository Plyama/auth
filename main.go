// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BearerAuth
// @in header
// @name Authorization
// @x-extension-openapi {"example": "value on a json format"}

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
