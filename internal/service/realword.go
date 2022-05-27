package service

import (
	"context"
	v1 "realword/api/realword/v1"
	"realword/internal/biz"
)

// RealWordService is a realword service.
type RealWordService struct {
	v1.UnimplementedRealWordServer

	uc *biz.RealWordUsecase
}

// NewRealWordService new a realword service.
func NewRealWordService(uc *biz.RealWordUsecase) *RealWordService {
	return &RealWordService{uc: uc}
}

func (s *RealWordService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}

func (s *RealWordService) Registration(ctx context.Context, in *v1.RegistrationRequest) (*v1.RegistrationReply, error) {
	return &v1.RegistrationReply{}, nil
}

func (s *RealWordService) GetCurrentUser(ctx context.Context, in *v1.GetCurrentUserRequest) (*v1.GetCurrentUserReply, error) {
	return &v1.GetCurrentUserReply{}, nil
}

func (s *RealWordService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

func (s *RealWordService) GetProfile(ctx context.Context, in *v1.GetProfileRequest) (*v1.GetProfileReply, error) {
	return &v1.GetProfileReply{}, nil
}

func (s *RealWordService) FollowUser(ctx context.Context, in *v1.FollowUserRequest) (*v1.FollowUserReply, error) {
	return &v1.FollowUserReply{}, nil
}

func (s *RealWordService) UnfollowUser(ctx context.Context, in *v1.UnfollowUserRequest) (*v1.UnfollowUserReply, error) {
	return &v1.UnfollowUserReply{}, nil
}

func (s *RealWordService) ListArticles(ctx context.Context, in *v1.ListArticlesRequest) (*v1.ListArticlesReply, error) {
	return &v1.ListArticlesReply{}, nil
}

func (s *RealWordService) FeedArticles(ctx context.Context, in *v1.ListArticlesRequest) (*v1.ListArticlesReply, error) {
	return &v1.ListArticlesReply{}, nil
}

func (s *RealWordService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	return &v1.GetArticleReply{}, nil
}

func (s *RealWordService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	return &v1.CreateArticleReply{}, nil
}

func (s *RealWordService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	return &v1.UpdateArticleReply{}, nil
}

func (s *RealWordService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	return &v1.DeleteArticleReply{}, nil
}

func (s *RealWordService) AddComments(ctx context.Context, in *v1.AddCommentsRequest) (*v1.AddCommentsReply, error) {
	return &v1.AddCommentsReply{}, nil
}

func (s *RealWordService) GetComments(ctx context.Context, in *v1.GetCommentsRequest) (*v1.GetCommentsReply, error) {
	return &v1.GetCommentsReply{}, nil
}

func (s *RealWordService) DeleteComments(ctx context.Context, in *v1.DeleteCommentsRequest) (*v1.DeleteCommentsReply, error) {
	return &v1.DeleteCommentsReply{}, nil
}

func (s *RealWordService) FavoriteArticle(ctx context.Context, in *v1.FavoriteArticleRequest) (*v1.FavoriteArticleReply, error) {
	return &v1.FavoriteArticleReply{}, nil
}

func (s *RealWordService) UnfavoriteArticle(ctx context.Context, in *v1.UnfavoriteArticleRequest) (*v1.UnfavoriteArticleReply, error) {
	return &v1.UnfavoriteArticleReply{}, nil
}

func (s *RealWordService) GetTags(ctx context.Context, in *v1.GetTagsRequest) (*v1.GetTagsReply, error) {
	return &v1.GetTagsReply{}, nil
}
