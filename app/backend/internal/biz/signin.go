package biz

import (
	"context"
	"time"
)

// SignInTimeFormat 签到日期格式化
const SignInTimeFormat = "2006-1-02"

type SignInUseCase struct {
	signInRepo SignInRepo
}

func NewSignInUseCase(signInRepo SignInRepo) *SignInUseCase {
	return &SignInUseCase{signInRepo: signInRepo}
}

type SignInRepo interface {
	SignIn(ctx context.Context, openid string, date time.Time) error
	GetSignInList(ctx context.Context, openid string) ([]time.Time, error)
}

func (s *SignInUseCase) SignIn(ctx context.Context, openid string) error {
	return s.signInRepo.SignIn(ctx, openid, time.Now())
}

func (s SignInUseCase) SignInList(ctx context.Context, openid string) (map[string]bool, error) {
	times, err := s.signInRepo.GetSignInList(ctx, openid)
	if err != nil {
		return nil, err
	}
	m := make(map[string]bool, len(times))
	for _, t := range times {
		m[t.Format(SignInTimeFormat)] = true
	}
	return m, nil
}
