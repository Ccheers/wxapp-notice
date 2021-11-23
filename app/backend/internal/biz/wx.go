package biz

import (
	"context"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type WxUseCase struct {
	wxRepo WxRepo
	log    *log.Helper
	conf   *conf.WxAppConfig
}

func NewWxUseCase(wxRepo WxRepo, l log.Logger) *WxUseCase {
	return &WxUseCase{wxRepo: wxRepo, log: log.NewHelper(l)}
}

type WxRepo interface {
	Code2Session(ctx context.Context, request *pb.Code2SessionRequest) (*pb.Code2SessionReply, error)
	SubscribeSend(ctx context.Context, request *pb.SubscribeSendRequest) (*pb.SubscribeSendReply, error)
	GetAccessToken(ctx context.Context, request *pb.GetAccessTokenRequest) (*pb.GetAccessTokenReply, error)
}

func (receiver *WxUseCase) Code2Session(ctx context.Context, code string) (*pb.Code2SessionReply, error) {
	return receiver.wxRepo.Code2Session(ctx, &pb.Code2SessionRequest{
		Appid:     receiver.conf.AppId,
		Secret:    receiver.conf.AppSecret,
		JsCode:    code,
		GrantType: "authorization_code",
	})
}
