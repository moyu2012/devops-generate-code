package controller

import (
	"devops-config-agent/src/server/application_config/service"

	"github.com/kataras/iris"
)

type ApplicationController struct {
}

type ApplicationConfig struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Config struct {
	Path                  string              `json:"path"`
	ApplicationConfigList []ApplicationConfig `json:"config"`
}

func (a *ApplicationController) GenerateVMConfig(ctx iris.Context) {
	config := &Config{}
	if err := ctx.ReadJSON(config); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	fileOperate := &service.ApplicationConfigService{}
	for _, configData := range config.ApplicationConfigList {
		if err := fileOperate.WriteToFile(configData.Name, config.Path, configData.Content); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}
	}
	ctx.Writef("Received: OK!")
}
