package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"realworld/internal/biz"
	"realworld/internal/data/model"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *articleRepo) CreateArticle(ctx context.Context, userID uint, bizArticle *biz.Article) (*biz.Article, error) {
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
	for _, u := range article.Favorites {
		if u.ID == userID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: article.Author.IsFollower(userID),
	}
	return bizArticle, nil
}

func (r *articleRepo) UpdateArticleBySlug(ctx context.Context, userID, id uint, argsMap map[string]interface{}) (*biz.Article, error) {
	var article model.Article
	if err := r.data.db.First(&article, id).Updates(argsMap).Error; err != nil {
		return nil, errors.InternalServer("article", err.Error())
	}
	err := r.data.db.Where(article.ID).
		Preload("Favorites").
		Preload("Tags").
		Preload("Author").
		Preload("Author.Followers").
		First(&article).Error
	if err != nil {
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
	for _, u := range article.Favorites {
		if u.ID == userID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: article.Author.IsFollower(userID),
	}
	for _, t := range article.Tags {
		bizArticle.TagList = append(bizArticle.TagList, t.Tag)
	}
	return &bizArticle, nil
}

func (r *articleRepo) DeleteArticleBySlug(ctx context.Context, slug string) error {
	// TODO 软删除与唯一索引冲突

	err := r.data.db.Where(&model.Article{Slug: slug}).Delete(&model.Article{}).Error
	if err != nil {
		return errors.InternalServer("article", err.Error())
	}
	return nil
}

func (r *articleRepo) GetArticleBySlug(ctx context.Context, userID uint, slug string) (*biz.Article, error) {
	var article model.Article
	err := r.data.db.Where(&model.Article{Slug: slug}).
		Preload("Favorites").
		Preload("Tags").
		Preload("Author").
		Preload("Author.Followers").
		First(&article).Error
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
	for _, u := range article.Favorites {
		if u.ID == userID {
			bizArticle.Favorited = true
		}
	}
	bizArticle.Author = &biz.Profile{
		Username:  article.Author.Username,
		Bio:       article.Author.Bio,
		Image:     article.Author.Image,
		Following: article.Author.IsFollower(userID),
	}
	for _, t := range article.Tags {
		bizArticle.TagList = append(bizArticle.TagList, t.Tag)
	}
	return &bizArticle, nil
}

func (r *articleRepo) FeedArticles(ctx context.Context, limit, offset int, userID uint, userList []*biz.User) ([]*biz.Article, int64, error) {
	var (
		articles    []model.Article
		count       int64
		bizArticles []*biz.Article
	)
	ids := make([]uint, len(userList))
	for i, u := range userList {
		ids[i] = u.ID
	}
	r.data.db.Where("author_id in (?)", ids).
		Preload("Favorites").
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tag asc")
		}).
		Preload("Author").
		Preload("Author.Followers").
		Offset(offset).
		Limit(limit).
		Order("created_at desc").
		Find(&articles)
	r.data.db.Where("author_id in (?)", ids).Model(&articles).Count(&count)

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
		for _, u := range article.Favorites {
			if u.ID == userID {
				bizArticle.Favorited = true
			}
		}
		bizArticle.Author = &biz.Profile{
			Username:  article.Author.Username,
			Bio:       article.Author.Bio,
			Image:     article.Author.Image,
			Following: article.Author.IsFollower(userID),
		}
		for _, t := range article.Tags {
			bizArticle.TagList = append(bizArticle.TagList, t.Tag)
		}
		bizArticles = append(bizArticles, bizArticle)
	}
	return bizArticles, count, nil
}

func (r *articleRepo) ListArticle(ctx context.Context, limit, offset int, userID uint) ([]*biz.Article, int64, error) {
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
		Preload("Author.Followers").
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
		for _, u := range article.Favorites {
			if u.ID == userID {
				bizArticle.Favorited = true
			}
		}
		bizArticle.Author = &biz.Profile{
			Username:  article.Author.Username,
			Bio:       article.Author.Bio,
			Image:     article.Author.Image,
			Following: article.Author.IsFollower(userID),
		}
		for _, t := range article.Tags {
			bizArticle.TagList = append(bizArticle.TagList, t.Tag)
		}
		bizArticles = append(bizArticles, bizArticle)
	}
	return bizArticles, count, nil
}

func (r *articleRepo) FavoriteArticle(ctx context.Context, userID uint, bizArticle *biz.Article) (*biz.Article, error) {
	var (
		article model.Article
		user    model.User
	)
	article.ID = bizArticle.ID
	user.ID = userID
	err := r.data.db.Model(&article).Association("Favorites").Append(&user)
	if err != nil {
		return nil, err
	}

	bizArticle.FavoritesCount = uint(len(article.Favorites))
	bizArticle.Favorited = true
	return bizArticle, nil
}

func (r *articleRepo) UnFavoriteArticle(ctx context.Context, userID uint, bizArticle *biz.Article) (*biz.Article, error) {
	var (
		article model.Article
		user    model.User
	)
	article.ID = bizArticle.ID
	user.ID = userID
	err := r.data.db.Model(&article).Association("Favorites").Delete(&user)
	if err != nil {
		return nil, err
	}

	bizArticle.FavoritesCount = uint(len(article.Favorites))
	bizArticle.Favorited = false
	return bizArticle, nil
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (c commentRepo) CreateComment(ctx context.Context, userID uint, bizComment *biz.Comment) (*biz.Comment, error) {
	var comment model.Comment
	comment.Body = bizComment.Body
	comment.UserID = bizComment.AuthorID
	comment.ArticleID = bizComment.ArticleID
	err := c.data.db.Create(&comment).Error
	if err != nil {
		return nil, errors.InternalServer("comment", err.Error())
	}
	if err := c.data.db.Where(comment.ID).Preload("User").First(&comment).Error; err != nil {
		return nil, errors.InternalServer("comment", err.Error())
	}

	bizComment.ID = comment.ID
	bizComment.CreatedAt = comment.CreatedAt
	bizComment.UpdatedAt = comment.UpdatedAt
	bizComment.Body = comment.Body
	bizComment.Author = &biz.Profile{
		Username:  comment.User.Username,
		Bio:       comment.User.Bio,
		Image:     comment.User.Image,
		Following: comment.User.IsFollower(userID),
	}
	return bizComment, nil
}

func (c commentRepo) DeleteComment(ctx context.Context, id uint) error {
	err := c.data.db.Delete(&model.Comment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c commentRepo) GetCommentByArticle(ctx context.Context, userID uint, bizArticle *biz.Article) ([]*biz.Comment, error) {
	var (
		comments    []*model.Comment
		bizComments []*biz.Comment
	)
	err := c.data.db.Where("article_id = ?", bizArticle.ID).Preload("User").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		bizComment := new(biz.Comment)
		bizComment.ID = comment.ID
		bizComment.CreatedAt = comment.CreatedAt
		bizComment.UpdatedAt = comment.UpdatedAt
		bizComment.Body = comment.Body
		bizComment.Author = &biz.Profile{
			Username:  comment.User.Username,
			Bio:       comment.User.Bio,
			Image:     comment.User.Image,
			Following: comment.User.IsFollower(userID),
		}
		bizComments = append(bizComments, bizComment)
	}
	return bizComments, nil
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func (t tagRepo) ListTag(ctx context.Context) ([]biz.Tag, error) {
	var (
		tags    []model.Tag
		bizTags []biz.Tag
	)
	err := t.data.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		bizTags = append(bizTags, biz.Tag(tag.Tag))
	}
	return bizTags, nil
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
