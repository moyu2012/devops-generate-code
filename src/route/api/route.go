package api

import (
	healthController "devops-config-agent/src/core/health"
	applicationConfigController "devops-config-agent/src/server/application_config/controller"

	"github.com/kataras/iris"
)

func InitRoute(api *iris.Application) *iris.Application {
	root := api.Party("/")
	{
		health := &healthController.HealthController{}
		root.Get("/health", health.CheckHealth)
	}
	v1 := api.Party("/api/v1")
	{
		appConfigController := &applicationConfigController.ApplicationController{}
		v1.Post("/config", appConfigController.GenerateVMConfig)
	}

	return api
}
