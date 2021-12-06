package logic

import (
	"context"
	"database/sql"
	"errors"
	"zero-mall/zero_bbs/user/model"
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

type VerificationCode struct {
	Phone string
	Code  string
}

func (l *RegisterLogic) Register(in *user.UserRequest) (*user.UserResponse, error) {

	var data VerificationCode
	err := l.svcCtx.Cache.Get(in.VerificationKey, &data)
	if err != nil {
		return nil, err
	}

	if in.VerificationCode != data.Code {
		return nil, errors.New("短信验证码错误！")
	}

	//TODO 添加更多字段
	//TODO 优化 sql 生成的 model 字段值

	//register user
	_, err = l.svcCtx.UserModel.Insert(&model.Users{
		Name: in.Name,
		Password: sql.NullString{
			String: in.Password,
			Valid:  true,
		},
		Phone: sql.NullString{
			String: data.Phone,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Name:  in.Name,
		Phone: data.Phone,
	}, nil
}
