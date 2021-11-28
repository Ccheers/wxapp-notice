package service

import (
	"context"
	"net/http"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/signin/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
)

type SigninService struct {
	pb.UnimplementedSigninServer
	useCase *biz.SignInUseCase
}

func NewSigninService(useCase *biz.SignInUseCase) *SigninService {
	return &SigninService{
		useCase: useCase,
	}
}

func (s *SigninService) Signin(ctx context.Context, req *pb.SigninRequest) (*pb.SigninReply, error) {
	err := s.useCase.SignIn(ctx, req.GetOpenid())
	if err != nil {
		return nil, err
	}
	return &pb.SigninReply{
		Code: http.StatusOK,
		Msg:  "OK",
	}, nil
}
func (s *SigninService) ListSignin(ctx context.Context, req *pb.ListSigninRequest) (*pb.ListSigninReply, error) {
	list, err := s.useCase.SignInList(ctx, req.GetOpenid())
	if err != nil {
		return nil, err
	}
	return &pb.ListSigninReply{
		Code: http.StatusOK,
		Msg:  "OK",
		List: list,
	}, nil
}
