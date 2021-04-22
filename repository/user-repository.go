package repository

import (
	"log"

	"github.com/raismaulana/simple_todo/entity"
	"github.com/raismaulana/simple_todo/helper"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	DeleteUser(UserID string) (bool, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(userID string) (entity.User, error)
	GetAll() []entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func StaticUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (db *userRepository) InsertUser(user entity.User) entity.User {
	user.Password = helper.PasswordHash(user.Password)
	db.connection.Create(&user)
	return user
}

func (db *userRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = helper.PasswordHash(user.Password)
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	log.Println(user)
	db.connection.Save(&user)
	return user
}

func (db *userRepository) DeleteUser(userID string) (bool, error) {
	return true, nil
}

func (db *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (db *userRepository) FindByID(userID string) (entity.User, error) {
	var user entity.User
	res := db.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (db *userRepository) GetAll() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}
