package logic

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"zero-mall/zero_bbs/user/model"

	"zero-mall/zero_bbs/user/rpc/internal/svc"
	"zero-mall/zero_bbs/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.UserReply, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByPhone(sql.NullString{
		String: in.Phone,
		Valid:  true,
	})
	switch err {
	case nil:
		if !userInfo.Password.Valid {
			return nil, errorIncorrectPassword
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errorIncorrectPassword
		}

		err = bcrypt.CompareHashAndPassword(hash, []byte(userInfo.Password.String))
		if err != nil {
			return nil, errorIncorrectPassword
		}

		return &user.UserReply{
			Uid:      userInfo.Id,
			Username: userInfo.Name,
		}, nil
	case model.ErrNotFound:
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}

}
