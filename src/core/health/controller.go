package health

import "github.com/kataras/iris"

type HealthController struct {
}

func (h *HealthController) CheckHealth(ctx iris.Context) {
	ctx.Writef("Server OK!")
}
