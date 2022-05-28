package service

import (
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
}

// NewRealWordService new a realword service.
func NewRealWordService(uc *biz.UserUsecase, ac *biz.ArticleUsecase) *RealWordService {
	return &RealWordService{uc: uc, ac: ac}
}
