package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/service"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"

	"zero-mall/zero_bbs/user/rpc/internal/svc"
	"zero-mall/zero_bbs/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerificationCodesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerificationCodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationCodesLogic {
	return &VerificationCodesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerificationCodesLogic) VerificationCodes(in *user.VerificationCodeRequest) (*user.VerificationCodeResponse, error) {
	var code string
	if l.svcCtx.Config.Mode == service.DevMode {
		code = "1234"
	} else {
		//TODO random code and send to user phone
	}

	key := "verificationCode_" + rand.String(15)
	expiredAt := time.Now().Add(time.Minute * 5)

	//缓存验证码 5 分钟过期
	err := l.svcCtx.Cache.SetWithExpire(key, map[string]interface{}{
		"phone": in.Phone,
		"code":  code,
	}, time.Minute*5)

	if err != nil {
		return nil, err
	}

	return &user.VerificationCodeResponse{
		Key:       key,
		ExpiredAt: expiredAt.String(),
	}, nil
}
