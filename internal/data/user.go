package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"realworld/internal/biz"
	"realworld/internal/data/model"
	"realworld/internal/pkg/util"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, errors.InternalServer("user", err.Error())
	}
	result := r.data.db.Create(&model.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: ph,
	})
	if result.Error != nil {
		return nil, errors.InternalServer("user", result.Error.Error())
	}

	return user, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := new(model.User)
	result := r.data.db.First(&u, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not found by email")
	}
	if result.Error != nil {
		return nil, errors.InternalServer("user", result.Error.Error())
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        nil,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) VerifyPassword(password, passwordhash string) bool {
	return util.CheckPasswordHash(password, passwordhash)
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
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
