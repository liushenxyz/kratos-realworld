package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:256;uniqueIndex;not nul"`
	Email        string `gorm:"size:256;uniqueIndex;not null"`
	PasswordHash string
	Bio          string
	Image        string
	Followers    []Follow  `gorm:"foreignKey:FollowingID"`
	Followings   []Follow  `gorm:"foreignKey:FollowerID"`
	Favorites    []Article `gorm:"many2many:favorites;"`
}

type Follow struct {
	Follower    User
	FollowerID  uint `gorm:"primaryKey" sql:"type:int not null"`
	Following   User
	FollowingID uint `gorm:"primaryKey" sql:"type:int not null"`
}

func (u *User) IsFollowing(id uint) bool {
	if u.Followings == nil {
		return false
	}
	for _, f := range u.Followings {
		if f.FollowingID == id {
			return true
		}
	}
	return false
}

func (u *User) IsFollower(id uint) bool {
	if u.Followers == nil {
		return false
	}
	for _, f := range u.Followers {
		if f.FollowerID == id {
			return true
		}
	}
	return false
}
