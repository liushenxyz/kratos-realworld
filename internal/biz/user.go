package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
	"realworld/internal/conf"
	"realworld/internal/pkg/middleware/auth"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Token        string
	Image        *anypb.Any
	Password     string
	PasswordHash string
}

type Profile struct {
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	VerifyPassword(password, passwordhash string) bool
}

type ProfileRepo interface {
	GetProfileByUsername(ctx context.Context, username string) (*Profile, error)
	FollowUserByUsername(ctx context.Context, username string) (*Profile, error)
	UnfollowUserByUsername(ctx context.Context, username string) (*Profile, error)
}

type UserUsecase struct {
	ur       UserRepo
	pr       ProfileRepo
	confAuth *conf.Auth
	log      *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, confAuth *conf.Auth, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, confAuth: confAuth, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !uc.ur.VerifyPassword(password, u.PasswordHash) {
		return nil, errors.Unauthorized("user", "login failed")
	}

	token, err := auth.CreateTokenString(uc.confAuth.Secret, u.Username)
	if err != nil {
		return nil, err
	}
	return &User{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    token,
	}, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, username, email, password string) (*User, error) {
	u, err := uc.ur.CreateUser(ctx, &User{
		Email:    email,
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
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
