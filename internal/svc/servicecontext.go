package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"time"
	"zero-mall/zero_bbs/internal/config"
	"zero-mall/zero_bbs/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	Cache      cache.Cache
	UserClient userclient.User
}

var ErrNotFound = errors.New("data not find")

func timeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	startTime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("调用 %s 方法 耗时: %v\n", method, time.Now().Sub(startTime))
	return nil
}

func NewServiceContext(c config.Config) *ServiceContext {

	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), ErrNotFound)
	ur := userclient.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))

	return &ServiceContext{
		Config:     c,
		Cache:      ca,
		UserClient: ur,
	}
}
