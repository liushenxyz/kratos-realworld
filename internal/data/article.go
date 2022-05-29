package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) UpdateArticleBySlug(ctx context.Context, slug string, article *biz.Article) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) DeleteArticleBySlug(ctx context.Context, slug string) error {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) GetArticleBySlug(ctx context.Context, slug string) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) FavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) UnfavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (c commentRepo) CreateComment(ctx context.Context, comment *biz.Comment) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentRepo) DeleteComment(ctx context.Context, slug string, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c commentRepo) GetCommentByArticle(ctx context.Context, article *biz.Article) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func (t tagRepo) ListTag(ctx context.Context) ([]*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
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
