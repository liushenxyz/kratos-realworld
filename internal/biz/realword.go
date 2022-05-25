package biz

import (
	"context"

	v1 "realword/api/realword/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// RealWord is a RealWord model.
type RealWord struct {
	Hello string
}

// RealWordRepo is a Greater repo.
type RealWordRepo interface {
	Save(context.Context, *RealWord) (*RealWord, error)
	Update(context.Context, *RealWord) (*RealWord, error)
	FindByID(context.Context, int64) (*RealWord, error)
	ListByHello(context.Context, string) ([]*RealWord, error)
	ListAll(context.Context) ([]*RealWord, error)
}

// RealWordUsecase is a RealWord usecase.
type RealWordUsecase struct {
	repo RealWordRepo
	log  *log.Helper
}

// NewRealWordUsecase new a RealWord usecase.
func NewRealWordUsecase(repo RealWordRepo, logger log.Logger) *RealWordUsecase {
	return &RealWordUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRealWord creates a RealWord, and returns the new RealWord.
func (uc *RealWordUsecase) CreateRealWord(ctx context.Context, g *RealWord) (*RealWord, error) {
	uc.log.WithContext(ctx).Infof("CreateRealWord: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
