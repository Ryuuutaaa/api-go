package models

import (
	"context"
	"demo/config"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Create(ctx context.Context, data User) error {
	return config.Connection.WithContext(ctx).Create(&data).Error
}

func Read(ctx context.Context, id uint) (User, error) {
	var user User
	if err := config.Connection.WithContext(ctx).First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func Update(ctx context.Context, id uint, data User) (User, error) {
	var user User
	if err := config.Connection.WithContext(ctx).Where("id = ?", id).Updates(&data).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func Delete(ctx context.Context, id uint) error {
	return config.Connection.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error
}
