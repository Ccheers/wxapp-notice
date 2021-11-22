package service

import (
	"context"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
)

type WxService struct {
	pb.UnimplementedWxServer
	wxUseCase *biz.WxUseCase
}

func NewWxService(wxUseCase *biz.WxUseCase) *WxService {
	return &WxService{
		wxUseCase: wxUseCase,
	}
}

func (s *WxService) Code2Session(ctx context.Context, req *pb.Code2SessionRequest) (*pb.Code2SessionReply, error) {
	return s.wxUseCase.Code2Session(ctx, req.GetJsCode())
}
