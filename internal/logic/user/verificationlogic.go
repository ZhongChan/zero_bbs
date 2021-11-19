package user

import (
	"context"
	"zero-mall/zero_bbs/user/rpc/userclient"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerificationLogic {
	return VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req types.VerificationCodeRequest) (*types.VerificationCodeResponse, error) {
	resp, err := l.svcCtx.UserClient.VerificationCodes(l.ctx, &userclient.VerificationCodeRequest{
		CaptchaKey:  req.CaptchaKey,
		CaptchaCode: req.CaptchaCode,
	})
	if err != nil {
		return nil, err
	}

	return &types.VerificationCodeResponse{
		Key:       resp.Key,
		ExpiredAt: resp.ExpiredAt,
	}, nil
}
