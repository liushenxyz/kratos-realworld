package service

import (
	"context"
	v1 "realword/api/realword/v1"
)

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
