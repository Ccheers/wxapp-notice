package server

import (
	pb "github.com/Ccheers/wxapp-notice/app/backend/api/notice/v1"
	v12 "github.com/Ccheers/wxapp-notice/app/backend/api/signin/v1"
	v1 "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, server *service.WxService, noticeServer *service.NoticeService, signInService *service.SigninService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterWxHTTPServer(srv, server)
	pb.RegisterNoticeHTTPServer(srv, noticeServer)
	v12.RegisterSigninHTTPServer(srv, signInService)
	return srv
}
