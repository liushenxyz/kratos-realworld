package service

import (
	"context"
	"fmt"

	v1 "realword/api/realword/v1"
	"realword/internal/biz"
)

// RealWordService is a realword service.
type RealWordService struct {
	v1.UnimplementedRealWordServer

	uc *biz.RealWordUsecase
}

// NewRealWordService new a realword service.
func NewRealWordService(uc *biz.RealWordUsecase) *RealWordService {
	return &RealWordService{uc: uc}
}

func (s *RealWordService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	fmt.Println(in)
	return &v1.LoginReply{User: nil}, nil
}
