package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
)

type User struct {
	Email    string
	Username string
	Bio      string
	Token    string
	Password string
	Image    *anypb.Any
}

type Profile struct {
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	VerifyPassword(ctx context.Context, user *User) (bool, error)
}

type ProfileRepo interface {
	GetProfileByUsername(ctx context.Context, username string) (*Profile, error)
	FollowUserByUsername(ctx context.Context, username string) (*Profile, error)
	UnfollowUserByUsername(ctx context.Context, username string) (*Profile, error)
}

type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	//uc.ur.GetUserByEmail(ctx context.Context, email string) (*User, error)
	//uc.ur.VerifyPassword(ctx context.Context, user *User) (bool, error)
	return nil, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	user := &User{
		Email:    u.Email,
		Username: u.Username,
		Token:    "Token",
		Bio:      "Bio",
		Image:    nil,
	}
	return user, nil
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
