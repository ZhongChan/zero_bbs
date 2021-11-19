package logic

import (
	"context"
	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"
	"zero-mall/zero_bbs/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *CaptchaLogic) Captcha(req types.CaptchaRequest) (*types.CaptchaResponse, error) {
	resp, err := l.svcCtx.UserClient.GetCaptcha(l.ctx, &userclient.CaptchaRequest{
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &types.CaptchaResponse{
		Captcha_key:           resp.CaptchaKey,
		Expired_at:            resp.ExpiredAt,
		Captcha_image_content: resp.CaptchaImageContent,
	}, nil
}
