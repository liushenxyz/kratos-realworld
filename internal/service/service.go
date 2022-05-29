package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUsecase
	ac *biz.ArticleUsecase

	log *log.Helper
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.UserUsecase, ac *biz.ArticleUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{
		uc:  uc,
		ac:  ac,
		log: log.NewHelper(logger),
	}
}
