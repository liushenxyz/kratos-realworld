package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "realworld/api/realworld/v1"
)

func (s *RealWorldService) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
	return &pb.ListArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, in *pb.FeedArticlesRequest) (*pb.FeedArticlesReply, error) {
	return &pb.FeedArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	return &pb.CreateArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, in *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, in *pb.DeleteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *RealWorldService) AddComments(ctx context.Context, in *pb.AddCommentsRequest) (*pb.AddCommentsReply, error) {
	return &pb.AddCommentsReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, in *pb.GetCommentsRequest) (*pb.GetCommentsReply, error) {
	return &pb.GetCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComments(ctx context.Context, in *pb.DeleteCommentsRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, in *pb.FavoriteArticleRequest) (*pb.FavoriteArticleReply, error) {
	return &pb.FavoriteArticleReply{}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, in *pb.UnfavoriteArticleRequest) (*pb.UnfavoriteArticleReply, error) {
	return &pb.UnfavoriteArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, in *empty.Empty) (*pb.GetTagsReply, error) {
	return &pb.GetTagsReply{}, nil
}
