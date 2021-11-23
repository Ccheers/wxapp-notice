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

func NewWxUseCase(wxRepo WxRepo, l log.Logger, conf *conf.WxAppConfig) *WxUseCase {
	return &WxUseCase{wxRepo: wxRepo, log: log.NewHelper(l), conf: conf}
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
func (receiver *WxUseCase) GetAccessToken(ctx context.Context) (string, error) {
	res, err := receiver.wxRepo.GetAccessToken(ctx, &pb.GetAccessTokenRequest{
		GrantType: "client_credential",
		Appid:     receiver.conf.AppId,
		Secret:    receiver.conf.AppSecret,
	})
	if err != nil {
		return "", err
	}
	return res.AccessToken, nil
}
