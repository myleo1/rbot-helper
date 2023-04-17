package controller

import (
	"github.com/mizuki1412/go-core-kit/service/restkit/context"
	"github.com/mizuki1412/go-core-kit/service/restkit/router"
	"robot-helper/middleware"
	"robot-helper/service"
)

func Init(router *router.Router) {
	r := router.Use(middleware.AuthToken())
	{
		//下发profile
		r.Post("/profile", profile)
	}
}

type profileParams struct {
	OracleName string `validate:"required"`
	AzureName  string `validate:"required"`
}

func profile(ctx *context.Context) {
	params := &profileParams{}
	ctx.BindForm(params)
	c := service.NewConfig(params.OracleName, params.AzureName)
	c.MakeConfig()
	c.RestartRBot()
	ctx.JsonSuccess("ok")
}
