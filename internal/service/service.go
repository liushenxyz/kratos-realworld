package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "realword/api/realword/v1"
	"realword/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWordService)

// RealWordService is a realword service.
type RealWordService struct {
	v1.UnimplementedRealWordServer

	uc *biz.UserUsecase
	ac *biz.ArticleUsecase

	log *log.Helper
}

// NewRealWordService new a realword service.
func NewRealWordService(uc *biz.UserUsecase, ac *biz.ArticleUsecase, logger log.Logger) *RealWordService {
	return &RealWordService{
		uc:  uc,
		ac:  ac,
		log: log.NewHelper(logger),
	}
}
