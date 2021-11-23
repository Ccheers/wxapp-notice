package service

import (
	"context"
	"net/http"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/notice/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type NoticeService struct {
	pb.UnimplementedNoticeServer
	noticeUseCase *biz.NoticeUseCase
	log           *log.Helper
}

func NewNoticeService(noticeUseCase *biz.NoticeUseCase, logger log.Logger) *NoticeService {
	return &NoticeService{
		noticeUseCase: noticeUseCase,
		log:           log.NewHelper(logger),
	}
}

func (s *NoticeService) SubscribeNotice(ctx context.Context, req *pb.SubscribeNoticeRequest) (*pb.SubscribeNoticeReply, error) {
	err := s.noticeUseCase.SubscribeNotice(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SubscribeNoticeReply{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
func (s *NoticeService) CheckSubscribeStatus(ctx context.Context, req *pb.CheckSubscribeStatusRequest) (*pb.CheckSubscribeStatusReply, error) {
	exist, err := s.noticeUseCase.CheckOpenID(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CheckSubscribeStatusReply{
		Subscribed: exist,
		Code:       http.StatusOK,
		Msg:        "success",
	}, nil
}

func (s *NoticeService) DeleteNotice(ctx context.Context, request *pb.DeleteNoticeRequest) (*pb.DeleteNoticeReply, error) {
	panic("implement me")
}

func (s *NoticeService) ListNotice(ctx context.Context, request *pb.ListNoticeRequest) (*pb.ListNoticeReply, error) {
	panic("implement me")
}

func (s *NoticeService) PutNotice(ctx context.Context, request *pb.PutNoticeRequest) (*pb.PutNoticeReply, error) {
	_, err := s.noticeUseCase.AddNotice(ctx, &biz.Job{
		Openid:          request.GetNotice().GetOpenid(),
		Content:         request.GetNotice().GetContent(),
		Time:            request.GetNotice().GetTime(),
		Frequency:       request.GetNotice().GetFrequency(),
		FrequencyDetail: request.GetNotice().GetFrequencyDetail(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PutNoticeReply{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
