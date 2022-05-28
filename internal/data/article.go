package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"realword/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) Save(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	return g, nil
}

func (r *articleRepo) Update(ctx context.Context, g *biz.Article) (*biz.Article, error) {
	return g, nil
}

func (r *articleRepo) FindByID(context.Context, int64) (*biz.Article, error) {
	return nil, nil
}

func (r *articleRepo) ListByHello(context.Context, string) ([]*biz.Article, error) {
	return nil, nil
}

func (r *articleRepo) ListAll(context.Context) ([]*biz.Article, error) {
	return nil, nil
}
