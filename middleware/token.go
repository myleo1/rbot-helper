package middleware

import (
	"github.com/mizuki1412/go-core-kit/class"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/mizuki1412/go-core-kit/service/restkit/context"
	"github.com/mizuki1412/go-core-kit/service/restkit/router"
	"robot-helper/constant"
)

var token string

func AuthToken() router.Handler {
	if token == "" {
		token = configkit.GetStringD(constant.ConfigKeyToken)
	}
	return func(ctx *context.Context) {

		if getToken := ctx.Request.Header.Get("token"); getToken != token {
			ctx.Json(context.RestRet{
				Result: context.ResultAuthErr,
				Message: class.String{
					String: "token验证失败",
					Valid:  true,
				},
			})
			ctx.Proxy.Abort()
		}
		ctx.Proxy.Next()
	}
}
