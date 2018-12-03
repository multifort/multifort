package log

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %s", err.Error())
	}
	defer logger.Sync()

}
