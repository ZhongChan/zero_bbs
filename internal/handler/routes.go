// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "zero-mall/zero_bbs/internal/handler/user"
	"zero-mall/zero_bbs/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/captcha",
				Handler: user.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/verificationCodes",
				Handler: user.VerificationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/users",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserCheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/info",
					Handler: user.InfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
