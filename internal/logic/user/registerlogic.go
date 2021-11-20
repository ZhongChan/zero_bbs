package user

import (
	"context"
	"zero-mall/zero_bbs/user/rpc/userclient"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserRequest) (*types.UserResponse, error) {
	resp, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.UserRequest{
		Name:             req.Name,
		Password:         req.Password,
		VerificationKey:  req.VerificationKey,
		VerificationCode: req.VerificationCode,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResponse{
		Name:  resp.Name,
		Phone: resp.Phone,
	}, nil
}
