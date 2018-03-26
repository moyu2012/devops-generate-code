package controller

import (
	generate_service "devops-generate-code/src/server/application_config/service"

	"github.com/kataras/iris"
)

type ApplicationController struct {
}

func (a *ApplicationController) GenerateVueCode(ctx iris.Context) {
	name := ctx.Params().Get("name")
	service := &generate_service.ApplicationConfigService{}
	println("---start to generate vue init code--")
	file, dirName, err := service.CMDGenerateVue(name)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.SendFile(file, "generate-vue.tar.gz")
	err = service.DeleteGZ(dirName)
	println("---End to generate vue init code--")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
}

func (a *ApplicationController) GenerateJavaCode(ctx iris.Context) {
	name := ctx.Params().Get("name")
	group_id := ctx.Params().Get("group_id")
	artifact_id := ctx.Params().Get("artifact_id")
	version := ctx.Params().Get("version")
	package_name := ctx.Params().Get("package_name")
	service := &generate_service.ApplicationConfigService{}
	println("---start to generate java init code--")
	file, dirName, err := service.CMDGenerateJava(name, group_id, artifact_id, version, package_name)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.SendFile(file, "generate-java.tar.gz")
	err = service.DeleteGZ(dirName)
	println("---End to generate vue init code--")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

}

func (a *ApplicationController) GenerateAngularCode(ctx iris.Context) {
	name := ctx.Params().Get("name")
	service := &generate_service.ApplicationConfigService{}
	println("---start to generate vue init code--")
	file, dirName, err := service.CMDGenerateAngular(name)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.SendFile(file, "generate-angular.tar.gz")
	err = service.DeleteGZ(dirName)
	println("---End to generate angular init code--")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
}
