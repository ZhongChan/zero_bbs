// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"zero-mall/zero_bbs/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CaptchaRequest  = user.CaptchaRequest
	CaptchaResponse = user.CaptchaResponse

	User interface {
		GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}