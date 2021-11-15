package repository

import (
	"github.com/titoyudha/Go_Gin_RestAPI/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Contract what can USer Repo can do in db
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx gorm.DB) {
	var user entity.User
	return *db.connection.Where("email: ?").Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email: ?").Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
