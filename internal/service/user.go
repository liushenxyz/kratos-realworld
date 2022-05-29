package service

import (
	"context"
	pb "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	s.log.Infof("input data %v", in)
	u, err := s.uc.Login(ctx, in.User.Email, in.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		User: &pb.User{
			User: &pb.User_User{
				Email:    u.Email,
				Token:    u.Token,
				Username: u.Username,
				Bio:      u.Bio,
				Image:    u.Image,
			},
		},
	}, nil
}

func (s *RealWorldService) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.RegistrationReply, error) {
	s.log.Infof("input data %v", in)
	u, err := s.uc.CreateUser(ctx, in.User.Username, in.User.Email, in.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationReply{
		User: &pb.User{
			User: &pb.User_User{
				Email:    u.Email,
				Token:    u.Token,
				Username: u.Username,
				Bio:      u.Bio,
				Image:    u.Image,
			},
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	return &pb.GetCurrentUserReply{}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	return &pb.GetProfileReply{}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserReply, error) {
	return &pb.FollowUserReply{}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {
	return &pb.UnfollowUserReply{}, nil
}
