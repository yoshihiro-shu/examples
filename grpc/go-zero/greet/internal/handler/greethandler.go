package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/examples/grpc/go-zero/greet/internal/logic"
	"github.com/yoshihiro-shu/examples/grpc/go-zero/greet/internal/svc"
	"github.com/yoshihiro-shu/examples/grpc/go-zero/greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx)
		resp, err := l.Greet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
