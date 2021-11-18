package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"time"
	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

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

var store = base64Captcha.DefaultMemStore

func (l *CaptchaLogic) Captcha(req types.CaptchaRequest) (*types.CaptchaResponse, error) {
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
		"phone":  req.Phone,
		"answer": answer,
	}, time.Second*2)

	if err != nil {
		return nil, err
	}

	return &types.CaptchaResponse{
		Captcha_key:           id,
		Expired_at:            expiredAt.String(),
		Captcha_image_content: item.EncodeB64string(),
	}, nil
}
