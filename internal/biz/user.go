package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}

type UserRepo interface {
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) Registration(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) GetProfile(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) FollowUser(ctx context.Context, u *User) error {
	return nil
}

func (uc *UserUsecase) UnfollowUser(ctx context.Context, u *User) error {
	return nil
}
