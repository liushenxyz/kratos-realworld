package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	gs "github.com/gosimple/slug"
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

	AuthorID uint
	Author   *Profile

	ArticleID uint
	Article   *Article
}

type Tag string

type ArticleRepo interface {
	CreateArticle(ctx context.Context, userID uint, article *Article) (*Article, error)
	UpdateArticleBySlug(ctx context.Context, userID, id uint, argsMap map[string]interface{}) (*Article, error)
	DeleteArticleBySlug(ctx context.Context, slug string) error
	GetArticleBySlug(ctx context.Context, userID uint, slug string) (*Article, error)
	ListArticle(ctx context.Context, limit, offset int, userID uint) ([]*Article, int64, error)
	FeedArticles(ctx context.Context, limit, offset int, userID uint, userList []*User) ([]*Article, int64, error)

	FavoriteArticle(ctx context.Context, userID uint, article *Article) (*Article, error)
	UnFavoriteArticle(ctx context.Context, userID uint, article *Article) (*Article, error)
}

type CommentRepo interface {
	CreateComment(ctx context.Context, userID uint, comment *Comment) (*Comment, error)
	DeleteComment(ctx context.Context, id uint) error
	GetCommentByArticle(ctx context.Context, userID uint, article *Article) ([]*Comment, error)
}

type TagRepo interface {
	ListTag(ctx context.Context) ([]Tag, error)
}

type ArticleUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

func NewArticleUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, ur UserRepo, pr ProfileRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ar: ar, cr: cr, tr: tr, ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (ac *ArticleUsecase) GetArticle(ctx context.Context, slug string) (*Article, error) {
	var userID uint
	cu := auth.FromContext(ctx)
	if cu != nil {
		userID = cu.ID
	} else {
		userID = 0
	}
	// TODO ????????????, Author.Following??????????????????

	return ac.ar.GetArticleBySlug(ctx, userID, slug)
}

func (ac *ArticleUsecase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	cu := auth.FromContext(ctx)
	article.AuthorID = cu.ID
	article.Slug = gs.Make(article.Title)

	return ac.ar.CreateArticle(ctx, cu.ID, article)
}

func (ac *ArticleUsecase) UpdateArticle(ctx context.Context, slug string, argsMap map[string]interface{}) (*Article, error) {
	cu := auth.FromContext(ctx)
	article, err := ac.ar.GetArticleBySlug(ctx, cu.ID, slug)
	if err != nil {
		return nil, err
	}
	if argsMap["Title"] != nil {
		argsMap["Slug"] = gs.Make(argsMap["Title"].(string))
	}
	// TODO UpdateArticle.TagList
	// TODO ??????????????????

	return ac.ar.UpdateArticleBySlug(ctx, cu.ID, article.ID, argsMap)
}

func (ac *ArticleUsecase) DeleteArticle(ctx context.Context, slug string) error {
	// TODO ??????????????????

	return ac.ar.DeleteArticleBySlug(ctx, slug)
}

func (ac *ArticleUsecase) ListArticles(ctx context.Context, limit, offset int) ([]*Article, int64, error) {
	if limit == 0 {
		limit = 20
	}
	var userID uint
	cu := auth.FromContext(ctx)
	if cu != nil {
		userID = cu.ID
	} else {
		userID = 0
	}
	// TODO ????????????, Author.Following??????????????????

	articles, count, err := ac.ar.ListArticle(ctx, limit, offset, userID)
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

func (ac *ArticleUsecase) FeedArticles(ctx context.Context, limit, offset int) ([]*Article, int64, error) {
	if limit == 0 {
		limit = 20
	}
	cu := auth.FromContext(ctx)
	userList, err := ac.pr.FollowingsList(ctx, cu.ID)
	if err != nil {
		return nil, 0, err
	}
	articles, count, err := ac.ar.FeedArticles(ctx, limit, offset, cu.ID, userList)
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

func (ac *ArticleUsecase) FavoriteArticle(ctx context.Context, slug string) (*Article, error) {
	cu := auth.FromContext(ctx)
	article, err := ac.ar.GetArticleBySlug(ctx, cu.ID, slug)
	if err != nil {
		return nil, err
	}
	article, err = ac.ar.FavoriteArticle(ctx, cu.ID, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (ac *ArticleUsecase) UnFavoriteArticle(ctx context.Context, slug string) (*Article, error) {
	cu := auth.FromContext(ctx)
	article, err := ac.ar.GetArticleBySlug(ctx, cu.ID, slug)
	if err != nil {
		return nil, err
	}
	article, err = ac.ar.UnFavoriteArticle(ctx, cu.ID, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (ac *ArticleUsecase) AddComments(ctx context.Context, slug string, comment *Comment) (*Comment, error) {
	cu := auth.FromContext(ctx)
	userID := cu.ID
	article, err := ac.ar.GetArticleBySlug(ctx, userID, slug)
	if err != nil {
		return nil, err
	}
	comment.ArticleID = article.ID
	comment.AuthorID = userID
	return ac.cr.CreateComment(ctx, userID, comment)
}

func (ac *ArticleUsecase) GetComments(ctx context.Context, slug string) ([]*Comment, error) {
	var userID uint
	cu := auth.FromContext(ctx)
	if cu != nil {
		userID = cu.ID
	} else {
		userID = 0
	}
	// TODO ????????????, Author.Following??????????????????
	article, err := ac.ar.GetArticleBySlug(ctx, userID, slug)
	if err != nil {
		return nil, err
	}
	return ac.cr.GetCommentByArticle(ctx, userID, article)
}

func (ac *ArticleUsecase) DeleteComments(ctx context.Context, slug string, id uint) error {
	// TODO ??????????????????

	return ac.cr.DeleteComment(ctx, id)
}

func (ac *ArticleUsecase) GetTags(ctx context.Context) ([]Tag, error) {
	return ac.tr.ListTag(ctx)
}
