package svc

import (
	"errors"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/syncx"
	"zero-mall/zero_bbs/user/model"
	"zero-mall/zero_bbs/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Cache     cache.Cache
}

var ErrNotFound = errors.New("data not find")

func NewServiceContext(c config.Config) *ServiceContext {

	ca := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("dc"), ErrNotFound)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
		Cache:     ca,
	}
}
