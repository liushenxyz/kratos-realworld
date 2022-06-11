package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
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

func (s *RealWorldService) GetCurrentUser(ctx context.Context, in *empty.Empty) (*pb.GetCurrentUserReply, error) {
	s.log.Infof("input data %v", in)
	u, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetCurrentUserReply{
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

func (s *RealWorldService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	s.log.Infof("input data %v", in)
	var argsMap = map[string]interface{}{}
	if in.User.Email != nil {
		argsMap["Email"] = *in.User.Email
	}
	if in.User.Username != nil {
		argsMap["Username"] = *in.User.Username
	}
	if in.User.Password != nil {
		argsMap["Password"] = *in.User.Password
	}
	if in.User.Image != nil {
		argsMap["Image"] = *in.User.Image
	}
	if in.User.Bio != nil {
		argsMap["Bio"] = *in.User.Bio
	}
	u, err := s.uc.UpdateUser(ctx, argsMap)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{
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

func (s *RealWorldService) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	s.log.Infof("input data %v", in)
	p, err := s.uc.GetProfile(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileReply{
		Profile: &pb.Profile{
			Profile: &pb.Profile_Profile{
				Username:  p.Username,
				Bio:       p.Bio,
				Image:     p.Image,
				Following: p.Following,
			},
		},
	}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserReply, error) {
	s.log.Infof("input data %v", in)
	p, err := s.uc.FollowUser(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	return &pb.FollowUserReply{
		Profile: &pb.Profile{
			Profile: &pb.Profile_Profile{
				Username:  p.Username,
				Bio:       p.Bio,
				Image:     p.Image,
				Following: p.Following,
			},
		},
	}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {
	s.log.Infof("input data %v", in)
	p, err := s.uc.UnfollowUser(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	return &pb.UnfollowUserReply{
		Profile: &pb.Profile{
			Profile: &pb.Profile_Profile{
				Username:  p.Username,
				Bio:       p.Bio,
				Image:     p.Image,
				Following: p.Following,
			},
		},
	}, nil
}
