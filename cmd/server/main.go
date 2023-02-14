package main

import (
	"context"
	"evendo-viator/conf"
	"evendo-viator/pkg/service"
	"gitlab.com/jfcore/common/logger"
	"os"
)

func main() {
	logger.Init("evendo-viator")
	conf.LoadConfig()
	_ = os.Setenv("PORT", conf.GetConfig().Port)
	_ = os.Setenv("ENABLE_DB", conf.GetConfig().EnableDB)
	_ = os.Setenv("CRONJOB_THREAD", conf.GetConfig().CronJobThread)
	_ = os.Setenv("NUMBER_OF_PRODUCT_BY_THREAD", conf.GetConfig().NumberOfProductPerThread)
	_ = os.Setenv("MONGO_HOST", conf.GetConfig().MongoHost)
	_ = os.Setenv("MONGO_PORT", conf.GetConfig().MongoPort)
	_ = os.Setenv("MONGO_USERNAME", conf.GetConfig().MongoUsername)
	_ = os.Setenv("MONGO_PASSWORD", conf.GetConfig().MongoPassword)

	app := service.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
