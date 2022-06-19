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
	err := r.data.db.Create(&model.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: util.HashPassword(user.Password),
	}).Error
	if err != nil {
		return nil, errors.InternalServer("user", err.Error())
	}
	return user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, id uint, argsMap map[string]interface{}) (*biz.User, error) {
	var m model.User
	if argsMap["Password"] != nil {
		argsMap["PasswordHash"] = util.HashPassword(argsMap["Password"].(string))
	}
	if err := r.data.db.First(&m, id).Updates(argsMap).Error; err != nil {
		return nil, errors.InternalServer("user", err.Error())
	}
	return &biz.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		Bio:          m.Bio,
		Image:        m.Image,
		PasswordHash: m.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*biz.User, error) {
	var m model.User
	if err := r.data.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("user", "not found by id")
		}
		return nil, errors.InternalServer("user", err.Error())
	}
	return &biz.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		Bio:          m.Bio,
		Image:        m.Image,
		PasswordHash: m.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	var m model.User
	if err := r.data.db.Where(&model.User{Email: email}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("user", "not found by email")
		}
		return nil, errors.InternalServer("user", err.Error())
	}
	return &biz.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		Bio:          m.Bio,
		Image:        m.Image,
		PasswordHash: m.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	var m model.User
	if err := r.data.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("user", "not found by username")
		}
		return nil, errors.InternalServer("user", err.Error())
	}
	return &biz.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		Bio:          m.Bio,
		Image:        m.Image,
		PasswordHash: m.PasswordHash,
	}, nil
}

func (r *userRepo) VerifyPassword(password, passwordhash string) bool {
	return util.CheckPasswordHash(password, passwordhash)
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func (p profileRepo) FollowingsList(ctx context.Context, followerID uint) ([]*biz.User, error) {
	var (
		u              model.User
		followingsList []model.Follow
		userList       []*biz.User
	)
	err := p.data.db.First(&u, followerID).Error
	if err != nil {
		return nil, err
	}
	if err := p.data.db.Model(&u).Preload("Following").Association("Followings").Find(&followingsList); err != nil {
		return nil, err
	}
	for _, f := range followingsList {
		userList = append(userList, &biz.User{
			ID:       f.Following.ID,
			Email:    f.Follower.Email,
			Username: f.Follower.Username,
			Bio:      f.Follower.Bio,
			Image:    f.Follower.Image,
		})
	}
	return userList, nil
}

func (p profileRepo) FollowersList(ctx context.Context, followingID uint) ([]*biz.User, error) {
	var (
		u             model.User
		followersList []model.Follow
		userList      []*biz.User
	)
	err := p.data.db.First(&u, followingID).Error
	if err != nil {
		return nil, err
	}
	if err := p.data.db.Model(&u).Preload("Follower").Association("Followers").Find(&followersList); err != nil {
		return nil, err
	}
	for _, f := range followersList {
		userList = append(userList, &biz.User{
			ID:       f.Following.ID,
			Email:    f.Following.Email,
			Username: f.Following.Username,
			Bio:      f.Following.Bio,
			Image:    f.Following.Image,
		})
	}
	return userList, nil
}

func (p profileRepo) IsFollowing(ctx context.Context, followerID, followingID uint) (bool, error) {
	var m model.User
	p.data.db.First(&m, followerID)
	var followingsList []model.Follow
	if err := p.data.db.Model(&m).Association("Followings").Find(&followingsList); err != nil {
		return false, errors.InternalServer("user", err.Error())
	}
	for _, f := range followingsList {
		if f.FollowingID == followingID {
			return true, nil
		}
	}
	return false, nil
}

func (p profileRepo) FollowUser(ctx context.Context, followerID, followingID uint) error {
	var m model.User
	p.data.db.First(&m, followerID)
	if err := p.data.db.Model(&m).Association("Followings").Append(&model.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}); err != nil {
		return errors.InternalServer("user", err.Error())
	}
	return nil
}

func (p profileRepo) UnFollowUser(ctx context.Context, followerID, followingID uint) error {
	if err := p.data.db.Delete(&model.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}).Error; err != nil {
		return errors.InternalServer("user", err.Error())
	}
	return nil
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
