package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"time"

	"zero-mall/zero_bbs/user/rpc/internal/svc"
	"zero-mall/zero_bbs/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *user.CaptchaRequest) (*user.CaptchaResponse, error) {

	var driver base64Captcha.Driver
	driver = base64Captcha.DefaultDriverDigit

	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}

	//缓存手机号码 和 验证码答案
	expiredAt := time.Now().Add(time.Minute * 2)
	err = l.svcCtx.Cache.SetWithExpire(id, map[string]interface{}{
		"phone":  in.Phone,
		"answer": answer,
	}, time.Minute*2)

	if err != nil {
		return nil, err
	}

	return &user.CaptchaResponse{
		CaptchaKey:          id,
		ExpiredAt:           expiredAt.String(),
		CaptchaImageContent: item.EncodeB64string(),
	}, nil
}
