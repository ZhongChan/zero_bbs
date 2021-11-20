package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"zero-mall/zero_bbs/internal/logic/user"
	"zero-mall/zero_bbs/internal/svc"
)

func InfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")

		l := user.NewInfoLogic(r.Context(), ctx)
		resp, err := l.Info(userId)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
