package user

import (
	"net/http"
	"zero-mall/zero_bbs/internal/logic/user"

	"github.com/tal-tech/go-zero/rest/httpx"
	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"
)

func CaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewCaptchaLogic(r.Context(), ctx)
		resp, err := l.Captcha(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
