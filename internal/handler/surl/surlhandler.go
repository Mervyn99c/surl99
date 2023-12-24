package surl

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"starry/internal/logic/surl"
	"starry/internal/svc"
	"starry/internal/types"
)

func SurlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	start := time.Now()
	// todo 记录整个接口响应时间， 中间件：异步发送kafka？
	// todo 长链为微信、支付宝小程序时，ios直接跳转，安卓通过中转页面跳转
	// todo 布隆过滤器
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SurlRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 过滤空码或metric
		if req.Surl == "" || len(req.Surl) != 7 || "metrix" == req.Surl {
			httpx.ErrorCtx(r.Context(), w, sqlx.ErrNotFound)
		}

		l := surl.NewSurlLogic(r.Context(), svcCtx)
		lurl, err := l.Router(&req)

		//重定向前记录 记录接口响应时间
		duration := time.Now().Sub(start)
		fmt.Println(duration)
		// todo generate sendkpi

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			http.Redirect(w, r, lurl, 302)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
