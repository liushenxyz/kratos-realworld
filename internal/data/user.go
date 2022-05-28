package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realword/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) VerifyPassword(ctx context.Context, user *biz.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func (p profileRepo) GetProfileByUsername(ctx context.Context, username string) (*biz.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p profileRepo) FollowUserByUsername(ctx context.Context, username string) (*biz.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p profileRepo) UnfollowUserByUsername(ctx context.Context, username string) (*biz.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
