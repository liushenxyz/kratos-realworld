package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

func (s *RealWorldService) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.ListArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, in *pb.FeedArticlesRequest) (*pb.FeedArticlesReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.FeedArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	s.log.Infof("input data %v", in)
	bizArticle, err := s.ac.GetArticle(ctx, in.Slug)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{
		Article: &pb.Article{
			Slug:           bizArticle.Slug,
			Title:          bizArticle.Title,
			Description:    bizArticle.Description,
			Body:           bizArticle.Body,
			TagList:        bizArticle.TagList,
			CreatedAt:      timestamppb.New(bizArticle.CreatedAt),
			UpdatedAt:      timestamppb.New(bizArticle.UpdatedAt),
			Favorited:      bizArticle.Favorited,
			FavoritesCount: uint32(bizArticle.FavoritesCount),
			Author: &pb.Author{
				Username:  bizArticle.Author.Username,
				Bio:       bizArticle.Author.Bio,
				Image:     bizArticle.Author.Image,
				Following: bizArticle.Author.Following,
			},
		},
	}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	s.log.Infof("input data %v", in)
	bizArticle, err := s.ac.CreateArticle(ctx, &biz.Article{
		Title:       in.Article.Title,
		Description: in.Article.Description,
		Body:        in.Article.Body,
		TagList:     in.Article.TagList,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateArticleReply{
		Article: &pb.Article{
			Slug:           bizArticle.Slug,
			Title:          bizArticle.Title,
			Description:    bizArticle.Description,
			Body:           bizArticle.Body,
			TagList:        bizArticle.TagList,
			CreatedAt:      timestamppb.New(bizArticle.CreatedAt),
			UpdatedAt:      timestamppb.New(bizArticle.UpdatedAt),
			Favorited:      bizArticle.Favorited,
			FavoritesCount: uint32(bizArticle.FavoritesCount),
			Author: &pb.Author{
				Username:  bizArticle.Author.Username,
				Bio:       bizArticle.Author.Bio,
				Image:     bizArticle.Author.Image,
				Following: bizArticle.Author.Following,
			},
		},
	}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, in *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	s.log.Infof("input data %v", in)
	var argsMap = map[string]interface{}{}
	if in.Article.Title != nil {
		argsMap["Title"] = *in.Article.Title
	}
	if in.Article.Description != nil {
		argsMap["Description"] = *in.Article.Description
	}
	if in.Article.Body != nil {
		argsMap["Body"] = *in.Article.Body
	}
	bizArticle, err := s.ac.UpdateArticle(ctx, in.Slug, argsMap)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateArticleReply{
		Article: &pb.Article{
			Slug:           bizArticle.Slug,
			Title:          bizArticle.Title,
			Description:    bizArticle.Description,
			Body:           bizArticle.Body,
			TagList:        bizArticle.TagList,
			CreatedAt:      timestamppb.New(bizArticle.CreatedAt),
			UpdatedAt:      timestamppb.New(bizArticle.UpdatedAt),
			Favorited:      bizArticle.Favorited,
			FavoritesCount: uint32(bizArticle.FavoritesCount),
			Author: &pb.Author{
				Username:  bizArticle.Author.Username,
				Bio:       bizArticle.Author.Bio,
				Image:     bizArticle.Author.Image,
				Following: bizArticle.Author.Following,
			},
		},
	}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, in *pb.DeleteArticleRequest) (*empty.Empty, error) {
	s.log.Infof("input data %v", in)
	err := s.ac.DeleteArticle(ctx, in.Slug)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *RealWorldService) AddComments(ctx context.Context, in *pb.AddCommentsRequest) (*pb.AddCommentsReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.AddCommentsReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, in *pb.GetCommentsRequest) (*pb.GetCommentsReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.GetCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComments(ctx context.Context, in *pb.DeleteCommentsRequest) (*empty.Empty, error) {
	s.log.Infof("input data %v", in)
	return nil, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, in *pb.FavoriteArticleRequest) (*pb.FavoriteArticleReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.FavoriteArticleReply{}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, in *pb.UnfavoriteArticleRequest) (*pb.UnfavoriteArticleReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.UnfavoriteArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, in *empty.Empty) (*pb.GetTagsReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.GetTagsReply{}, nil
}
