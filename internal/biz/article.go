package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
}

type Comment struct {
}

type Tag struct {
}

type ArticleRepo interface {
	ListArticle(ctx context.Context) ([]*Article, error)
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
	UpdateArticleBySlug(ctx context.Context, slug string, article *Article) (*Article, error)
	DeleteArticleBySlug(ctx context.Context, slug string) error
	GetArticleBySlug(ctx context.Context, slug string) (*Article, error)
	FavoriteArticle(ctx context.Context, slug string) (*Article, error)
	UnfavoriteArticle(ctx context.Context, slug string) (*Article, error)
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, slug string, id int64) error
	GetCommentByArticle(ctx context.Context, article *Article) ([]*Comment, error)
}

type TagRepo interface {
	ListTag(ctx context.Context) ([]*Tag, error)
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
