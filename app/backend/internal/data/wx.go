package data

import (
	"context"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/transport/http"
)

const wxServerEndpoint = "https://api.weixin.qq.com"

type WxRepoImpl struct {
	log *log.Helper
}

func NewWxRepoImpl(logger log.Logger) biz.WxRepo {
	return &WxRepoImpl{
		log: log.NewHelper(logger),
	}
}

func (w *WxRepoImpl) Code2Session(ctx context.Context, request *pb.Code2SessionRequest) (*pb.Code2SessionReply, error) {
	cli, err := http.NewClient(ctx, http.WithEndpoint(wxServerEndpoint))
	if err != nil {
		return nil, err
	}
	return pb.NewWxHTTPClient(cli).Code2Session(ctx, request)
}

func (w *WxRepoImpl) ImplWxRepo() {
	panic("implement me")
}
