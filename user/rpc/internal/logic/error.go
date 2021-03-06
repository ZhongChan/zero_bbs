package logic

import (
	"zero-mall/zero_bbs/shared"
)

var (
	errorDuplicateUsername  = shared.NewDefaultError("用户名已经注册")
	errorDuplicateMobile    = shared.NewDefaultError("手机号已经被占用")
	errorUsernameUnRegister = shared.NewDefaultError("用户未注册")
	errorIncorrectPassword  = shared.NewDefaultError("用户密码错误")
	errorUserNotFound       = shared.NewDefaultError("用户不存在")
)
