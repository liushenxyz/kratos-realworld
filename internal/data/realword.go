package data

import (
	"context"

	"realword/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type realwordRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealWordRepo .
func NewRealWordRepo(data *Data, logger log.Logger) biz.RealWordRepo {
	return &realwordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realwordRepo) Save(ctx context.Context, g *biz.RealWord) (*biz.RealWord, error) {
	return g, nil
}

func (r *realwordRepo) Update(ctx context.Context, g *biz.RealWord) (*biz.RealWord, error) {
	return g, nil
}

func (r *realwordRepo) FindByID(context.Context, int64) (*biz.RealWord, error) {
	return nil, nil
}

func (r *realwordRepo) ListByHello(context.Context, string) ([]*biz.RealWord, error) {
	return nil, nil
}

func (r *realwordRepo) ListAll(context.Context) ([]*biz.RealWord, error) {
	return nil, nil
}
