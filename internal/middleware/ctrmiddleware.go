package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

type CTRMiddleWareMiddleware struct {
}

func NewCTRMiddleWareMiddleware() *CTRMiddleWareMiddleware {
	return &CTRMiddleWareMiddleware{}
}

func (m *CTRMiddleWareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		next(w, r)
		duration := time.Now().Sub(start)
		logx.Info("响应时间为", duration)
		// assemble kpi
		// todo send kpi
	}
}
