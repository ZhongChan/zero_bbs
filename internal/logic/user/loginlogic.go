package user

import (
	"context"
	"time"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {
	//TODO 获取用户信息

	jwtToken, err := l.GetJwtToken(1)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Name:     "caesar",
		JwtToken: jwtToken,
	}, nil
}

func (l *LoginLogic) GetJwtToken(uid int64) (types.JwtToken, error) {
	iat := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + l.svcCtx.Config.Auth.AccessExpire
	claims["iat"] = iat
	claims["uid"] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtStr, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return types.JwtToken{}, err
	}
	return types.JwtToken{
		AccessToken:  jwtStr,
		AccessExpire: iat + accessExpire,
		RefreshAfter: iat + accessExpire/2,
	}, nil
}
