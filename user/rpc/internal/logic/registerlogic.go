package logic

import (
	"context"

	"zero-mall/zero_bbs/user/rpc/internal/svc"
	"zero-mall/zero_bbs/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.UserRequest) (*user.UserResponse, error) {

	var data interface{}
	err := l.svcCtx.Cache.Get(in.VerificationKey, data)
	if err != nil {
		return nil, err
	}

	return &user.UserResponse{}, nil
}
