package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"zero-mall/zero_bbs/internal/logic/user"
	"zero-mall/zero_bbs/internal/svc"
	"zero-mall/zero_bbs/internal/types"
)

func VerificationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerificationCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewVerificationLogic(r.Context(), ctx)
		resp, err := l.Verification(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
