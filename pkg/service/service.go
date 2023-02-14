package service

import (
	_ "evendo-viator/docs"
	"evendo-viator/pkg/handlers"
	HandlerAPI "evendo-viator/pkg/handlers/api"
	"evendo-viator/pkg/repo"
	"github.com/caarlos0/env/v6"
	"gitlab.com/jfcore/common/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

// NewService
// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /api/v1
func NewService() *Service {
	s := &Service{
		service.NewApp("evendo-viator", "v1.0"),
		&extraSetting{},
	}
	_ = env.Parse(s.setting)

	repoMongoProduct := repo.NewProductMongoRepo()
	handleAPI := HandlerAPI.NewViatorAPIHandlers()
	handleCronJob := handlers.NewCronJobHandlers(
		handleAPI,
		repoMongoProduct,
	)
	handleCronJob.StartCron()
	return s
}
