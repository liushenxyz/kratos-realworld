package model

import "gorm.io/gorm"

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
