package api

import (
	healthController "devops-generate-code/src/core/health"
	applicationConfigController "devops-generate-code/src/server/application_config/controller"

	"github.com/kataras/iris"
)

func InitRoute(api *iris.Application) *iris.Application {
	root := api.Party("/")
	{
		health := &healthController.HealthController{}
		root.Get("/health", health.CheckHealth)
	}
	vue := api.Party("/vue")
	{
		appConfigController := &applicationConfigController.ApplicationController{}
		vue.Get("/{name}/generate-vue.tar.gz", appConfigController.GenerateVueCode)
	}
	java := api.Party("/java")
	{
		appConfigController := &applicationConfigController.ApplicationController{}
		java.Get("/{name}/{group_id}/{artifact_id}/{version}/{package_name}/generate-java.tar.gz", appConfigController.GenerateJavaCode)
		//?name={name}&&group_id={group_id}&&artifact_id={artifact_id}&&version={version}&&package_name={package_name}
	}
	return api
}
