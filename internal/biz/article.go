package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
}

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type ArticleUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

func NewArticleUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) ListArticles(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) FeedArticles(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) GetArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) CreateArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) UpdateArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) DeleteArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) AddComments(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) GetComments(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) DeleteComments(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) FavoriteArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) UnfavoriteArticle(ctx context.Context, a *Article) error {
	return nil
}

func (uc *UserUsecase) GetTags(ctx context.Context, a *Article) error {
	return nil
}
