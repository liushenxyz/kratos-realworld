package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gosimple/slug"
	"realworld/internal/pkg/middleware/auth"
	"time"
)

type Article struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Slug        string
	Title       string
	Description string
	Body        string

	Favorited      bool
	FavoritesCount uint

	AuthorID uint
	Author   *Profile

	CommentList []Comment
	TagList     []string
}

type Comment struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Body string

	ArticleID uint
	Author    *Profile

	Article *Article
}

type Tag string

type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
	UpdateArticleBySlug(ctx context.Context, article *Article) (*Article, error)
	DeleteArticleBySlug(ctx context.Context, slug string) error
	GetArticleBySlug(ctx context.Context, slug string) (*Article, error)
	ListArticle(ctx context.Context) ([]*Article, error)

	FavoriteArticle(ctx context.Context, slug string) error
	UnfavoriteArticle(ctx context.Context, slug string) error
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id uint) error
	GetCommentByArticle(ctx context.Context, id uint) (*Comment, error)
	ListComment(ctx context.Context, slug string) ([]*Comment, error)
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

func (ac *ArticleUsecase) ListArticles(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) FeedArticles(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) GetArticle(ctx context.Context, slug string) (*Article, error) {
	//TODO Author.Following

	return ac.ar.GetArticleBySlug(ctx, slug)
}

func (ac *ArticleUsecase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	cu := auth.FromContext(ctx)
	article.AuthorID = cu.ID
	article.Slug = slug.Make(article.Title)
	//TODO Author.Following

	return ac.ar.CreateArticle(ctx, article)
}

func (ac *ArticleUsecase) UpdateArticle(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) DeleteArticle(ctx context.Context, slug string) error {
	return ac.ar.DeleteArticleBySlug(ctx, slug)
}

func (ac *ArticleUsecase) AddComments(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) GetComments(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) DeleteComments(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) FavoriteArticle(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) UnfavoriteArticle(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}

func (ac *ArticleUsecase) GetTags(ctx context.Context, a *Article) (*Article, error) {
	return nil, nil
}
