package service

import (
	"context"

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

// SayHello implements realword.RealWordServer.
func (s *RealWordService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateRealWord(ctx, &biz.RealWord{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
