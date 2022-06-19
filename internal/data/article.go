package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"realworld/internal/biz"
	"realworld/internal/data/model"
	"realworld/internal/pkg/middleware/auth"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *articleRepo) CreateArticle(ctx context.Context, bizArticle *biz.Article) (*biz.Article, error) {
	var article model.Article
	article.Slug = bizArticle.Slug
	article.Title = bizArticle.Title
	article.Description = bizArticle.Description
	article.Body = bizArticle.Body
	article.AuthorID = bizArticle.AuthorID
	var tags []model.Tag
	for _, t := range bizArticle.TagList {
		tags = append(tags, model.Tag{Tag: t})
	}

	tx := r.data.db.Begin()
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, t := range tags {
		err := tx.Where(&model.Tag{Tag: t.Tag}).First(&t).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(&article).Association("Tags").Append(&t); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Where(article.ID).Preload("Favorites").Preload("Tags").Preload("Author").First(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	bizArticle.ID = article.ID
	bizArticle.CreatedAt = article.CreatedAt
	bizArticle.UpdatedAt = article.UpdatedAt
	bizArticle.FavoritesCount = uint(len(article.Favorites))
	cu := auth.FromContext(ctx)
	for _, u := range article.Favorites {
		if u.ID == cu.ID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: false,
		//TODO Author.Following
	}
	return bizArticle, nil
}

func (r *articleRepo) UpdateArticleBySlug(ctx context.Context, id uint, argsMap map[string]interface{}) (*biz.Article, error) {
	var article model.Article
	if err := r.data.db.First(&article, id).Updates(argsMap).Error; err != nil {
		return nil, errors.InternalServer("article", err.Error())
	}

	if err := r.data.db.Where(article.ID).Preload("Favorites").Preload("Tags").Preload("Author").First(&article).Error; err != nil {
		return nil, err
	}

	var bizArticle biz.Article
	bizArticle.ID = article.ID
	bizArticle.CreatedAt = article.CreatedAt
	bizArticle.UpdatedAt = article.UpdatedAt
	bizArticle.Slug = article.Slug
	bizArticle.Title = article.Title
	bizArticle.Description = article.Description
	bizArticle.Body = article.Body
	bizArticle.FavoritesCount = uint(len(article.Favorites))
	cu := auth.FromContext(ctx)
	for _, u := range article.Favorites {
		if u.ID == cu.ID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: false,
		//TODO Author.Following
	}
	for _, t := range article.Tags {
		bizArticle.TagList = append(bizArticle.TagList, t.Tag)
	}
	return &bizArticle, nil
}

func (r *articleRepo) DeleteArticleBySlug(ctx context.Context, slug string) error {
	//TODO 软删除与唯一索引冲突

	err := r.data.db.Where(&model.Article{Slug: slug}).Delete(&model.Article{}).Error
	if err != nil {
		return errors.InternalServer("article", err.Error())
	}
	return nil
}

func (r *articleRepo) GetArticleBySlug(ctx context.Context, slug string) (*biz.Article, error) {
	var article model.Article
	err := r.data.db.Where(&model.Article{Slug: slug}).Preload("Favorites").Preload("Tags").Preload("Author").First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("article", "not found by slug")
		}
		return nil, err
	}

	var bizArticle biz.Article
	bizArticle.ID = article.ID
	bizArticle.CreatedAt = article.CreatedAt
	bizArticle.UpdatedAt = article.UpdatedAt
	bizArticle.Slug = article.Slug
	bizArticle.Title = article.Title
	bizArticle.Description = article.Description
	bizArticle.Body = article.Body
	bizArticle.FavoritesCount = uint(len(article.Favorites))
	cu := auth.FromContext(ctx)
	for _, u := range article.Favorites {
		if u.ID == cu.ID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: false,
		//TODO Author.Following
	}
	for _, t := range article.Tags {
		bizArticle.TagList = append(bizArticle.TagList, t.Tag)
	}
	return &bizArticle, nil
}

func (r *articleRepo) ListArticle(ctx context.Context, limit, offset int) ([]*biz.Article, int64, error) {
	var (
		articles    []model.Article
		count       int64
		bizArticles []*biz.Article
	)
	r.data.db.Model(&articles).Count(&count)
	r.data.db.Preload("Favorites").
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tag asc")
		}).
		Preload("Author").
		Limit(limit).
		Offset(offset).
		Order("created_at desc").Find(&articles)

	for _, article := range articles {
		bizArticle := new(biz.Article)
		bizArticle.ID = article.ID
		bizArticle.CreatedAt = article.CreatedAt
		bizArticle.UpdatedAt = article.UpdatedAt
		bizArticle.Slug = article.Slug
		bizArticle.Title = article.Title
		bizArticle.Description = article.Description
		bizArticle.Body = article.Body
		bizArticle.FavoritesCount = uint(len(article.Favorites))
		cu := auth.FromContext(ctx)
		for _, u := range article.Favorites {
			if u.ID == cu.ID {
				bizArticle.Favorited = true
			}
		}
		bizArticle.Author = &biz.Profile{
			Username:  article.Author.Username,
			Bio:       article.Author.Bio,
			Image:     article.Author.Image,
			Following: false,
			//TODO Author.Following
			//bizArticle.Author.Following, _ = us.IsFollower(a.AuthorID, cu.ID)
		}
		for _, t := range article.Tags {
			bizArticle.TagList = append(bizArticle.TagList, t.Tag)
		}
		bizArticles = append(bizArticles, bizArticle)
	}
	return bizArticles, count, nil
}

func (r *articleRepo) ListArticleByTag(ctx context.Context) ([]*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) ListArticleByAuthor(ctx context.Context) ([]*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) ListArticleByWhoFavorited(ctx context.Context) ([]*biz.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) FavoriteArticle(ctx context.Context, slug string) error {
	//TODO implement me
	panic("implement me")
}

func (r *articleRepo) UnfavoriteArticle(ctx context.Context, slug string) error {
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

func (c commentRepo) DeleteComment(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (c commentRepo) GetCommentByArticle(ctx context.Context, id uint) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentRepo) ListComment(ctx context.Context, slug string) ([]*biz.Comment, error) {
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
