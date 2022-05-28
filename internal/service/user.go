package service

import (
	"context"
	pb "realword/api/realword/v1"
	"realword/internal/biz"
)

func (s *RealWordService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	s.log.Infof("input data %v", in)
	return &pb.LoginReply{}, nil
}

func (s *RealWordService) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.RegistrationReply, error) {
	s.log.Infof("input data %v", in)
	user, err := s.uc.CreateUser(ctx, &biz.User{
		Username: in.User.Username,
		Email:    in.User.Email,
		Password: in.User.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationReply{User: &pb.User{
		User: &pb.User_User{
			Email:    user.Email,
			Token:    user.Token,
			Username: user.Username,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}}, nil
}

func (s *RealWordService) GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	return &pb.GetCurrentUserReply{}, nil
}

func (s *RealWordService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

func (s *RealWordService) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	return &pb.GetProfileReply{}, nil
}

func (s *RealWordService) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserReply, error) {
	return &pb.FollowUserReply{}, nil
}

func (s *RealWordService) UnfollowUser(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {
	return &pb.UnfollowUserReply{}, nil
}
