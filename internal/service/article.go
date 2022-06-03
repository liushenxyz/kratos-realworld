package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) ListArticles(ctx context.Context, in *v1.ListArticlesRequest) (*v1.ListArticlesReply, error) {
	return &v1.ListArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, in *v1.FeedArticlesRequest) (*v1.FeedArticlesReply, error) {
	return &v1.FeedArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	return &v1.GetArticleReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	return &v1.CreateArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	return &v1.UpdateArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	return &v1.DeleteArticleReply{}, nil
}

func (s *RealWorldService) AddComments(ctx context.Context, in *v1.AddCommentsRequest) (*v1.AddCommentsReply, error) {
	return &v1.AddCommentsReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, in *v1.GetCommentsRequest) (*v1.GetCommentsReply, error) {
	return &v1.GetCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComments(ctx context.Context, in *v1.DeleteCommentsRequest) (*v1.DeleteCommentsReply, error) {
	return &v1.DeleteCommentsReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, in *v1.FavoriteArticleRequest) (*v1.FavoriteArticleReply, error) {
	return &v1.FavoriteArticleReply{}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, in *v1.UnfavoriteArticleRequest) (*v1.UnfavoriteArticleReply, error) {
	return &v1.UnfavoriteArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, in *v1.GetTagsRequest) (*v1.GetTagsReply, error) {
	return &v1.GetTagsReply{}, nil
}
