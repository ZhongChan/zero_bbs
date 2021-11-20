package user

import (
	"context"
	"fmt"
	"strconv"

	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) InfoLogic {
	return InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(uid string) (*types.UserResponse, error) {

	userInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	fmt.Println("userInt:", userInt)
	//TODO get user by userInt

	return &types.UserResponse{
		Name:  "caesar",
		Phone: "13818403077",
	}, nil
}
