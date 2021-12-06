package user

import (
	"context"

	"zero-mall/zero_bbs/user/rpc/userclient"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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

	//使用 bcrypt 进行密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}

	resp, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.UserRequest{
		Name:             req.Name,
		Password:         string(hash), 
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
