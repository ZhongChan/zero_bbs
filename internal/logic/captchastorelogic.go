package logic

import (
	"context"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaStoreLogic {
	return CaptchaStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaStoreLogic) CaptchaStore(req types.CaptchaRequest) (*types.CaptchaResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CaptchaResponse{}, nil
}
