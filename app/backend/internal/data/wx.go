package data

import (
	"context"
	"sync"
	"time"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/pkg/trylock"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/transport/http"
)

const wxServerEndpoint = "https://api.weixin.qq.com"

const lockTimeOut = time.Second * 5

type WxRepoImpl struct {
	log               *log.Helper
	cache             sync.Map
	accessTokenLocker trylock.TryMutexLocker
}

func NewWxRepoImpl(logger log.Logger) biz.WxRepo {
	return &WxRepoImpl{
		log:               log.NewHelper(logger),
		cache:             sync.Map{},
		accessTokenLocker: trylock.NewAtomicMutex(),
	}
}

func (w *WxRepoImpl) Code2Session(ctx context.Context, request *pb.Code2SessionRequest) (*pb.Code2SessionReply, error) {
	cli, err := http.NewClient(ctx, http.WithEndpoint(wxServerEndpoint))
	if err != nil {
		return nil, err
	}
	return pb.NewWxHTTPClient(cli).Code2Session(ctx, request)
}

func (w *WxRepoImpl) SubscribeSend(ctx context.Context, request *pb.SubscribeSendRequest) (*pb.SubscribeSendReply, error) {
	cli, err := http.NewClient(ctx, http.WithEndpoint(wxServerEndpoint))
	if err != nil {
		return nil, err
	}
	return pb.NewWxHTTPClient(cli).SubscribeSend(ctx, request)
}

func (w *WxRepoImpl) GetAccessToken(ctx context.Context, request *pb.GetAccessTokenRequest) (*pb.GetAccessTokenReply, error) {
	err := w.accessTokenLocker.TryLock(lockTimeOut)
	defer w.accessTokenLocker.Unlock()
	if err != nil {
		return nil, err
	}

	if val, ok := w.cache.Load(request.Appid); ok {
		if reply, ok := val.(*pb.GetAccessTokenReply); ok {
			if time.Now().Before(time.Unix(reply.ExpiresIn, 0)) {
				return reply, nil
			}
		}
	}

	cli, err := http.NewClient(ctx, http.WithEndpoint(wxServerEndpoint))
	if err != nil {
		return nil, err
	}
	reply, err := pb.NewWxHTTPClient(cli).GetAccessToken(ctx, request)
	if err != nil {
		return nil, err
	}
	w.cache.Store(request.Appid, reply)
	return reply, nil
}
