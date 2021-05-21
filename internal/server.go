package internal

import (
	"net/http"
	"os"

	"github.com/plyama/auth/internal/app"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, logger *zap.Logger) {

	core := app.NewApp(db)
	r := NewRouter(core.Services)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		logger.Fatal("failed to start http server")
	}
}
