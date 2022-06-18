package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/conf"
	"realworld/internal/pkg/middleware/auth"
)

type User struct {
	ID           uint
	Email        string
	Username     string
	Bio          string
	Image        string
	Password     string
	PasswordHash string
	Token        string
}

type Profile struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, id uint, argsMap map[string]interface{}) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	VerifyPassword(password, passwordhash string) bool
}

type ProfileRepo interface {
	IsFollowing(ctx context.Context, followerID, followingID uint) (bool, error)
	FollowUser(ctx context.Context, followerID, followingID uint) error
	UnfollowUser(ctx context.Context, followerID, followingID uint) error
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
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !uc.ur.VerifyPassword(password, u.PasswordHash) {
		return nil, errors.Unauthorized("user", "login failed password error")
	}
	token := auth.CreateTokenString(uc.confAuth.Secret, u.Email, u.Username, u.ID)
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

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*User, error) {
	cu := auth.FromContext(ctx)
	u, err := uc.ur.GetUserByID(ctx, cu.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, argsMap map[string]interface{}) (*User, error) {
	cu := auth.FromContext(ctx)
	u, err := uc.ur.UpdateUser(ctx, cu.ID, argsMap)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUsecase) GetProfile(ctx context.Context, username string) (*Profile, error) {
	cu := auth.FromContext(ctx)
	fu, err := uc.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	result, err := uc.pr.IsFollowing(ctx, cu.ID, fu.ID)
	if err != nil {
		return nil, err
	}
	return &Profile{
		Username:  fu.Username,
		Bio:       fu.Bio,
		Image:     fu.Image,
		Following: result,
	}, nil
}

func (uc *UserUsecase) FollowUser(ctx context.Context, username string) (*Profile, error) {
	cu := auth.FromContext(ctx)
	fu, err := uc.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if err := uc.pr.FollowUser(ctx, cu.ID, fu.ID); err != nil {
		return nil, err
	}
	return &Profile{
		Username:  fu.Username,
		Bio:       fu.Bio,
		Image:     fu.Image,
		Following: true,
	}, nil
}

func (uc *UserUsecase) UnfollowUser(ctx context.Context, username string) (*Profile, error) {
	cu := auth.FromContext(ctx)
	fu, err := uc.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if err := uc.pr.UnfollowUser(ctx, cu.ID, fu.ID); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &Profile{
		Username:  fu.Username,
		Bio:       fu.Bio,
		Image:     fu.Image,
		Following: false,
	}, nil
}
