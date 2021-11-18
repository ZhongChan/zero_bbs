package svc

import (
	"errors"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
	"zero-mall/zero_bbs/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
}

var ErrNotFound = errors.New("data not find")

func NewServiceContext(c config.Config) *ServiceContext {

	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), ErrNotFound)

	return &ServiceContext{
		Config: c,
		Cache:  ca,
	}
}
