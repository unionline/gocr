package user_service

import (
	"gocr/models"
	"gocr/resposity"
	"gocr/utils"
)

// UserService ...
type UserService struct {
	User *models.User
}

// NewUserService ...
func NewUserService() *UserService {
	return &UserService{
		User: models.NewUser(),
	}
}

var userRepo *resposity.UserRepo

func init() {
	userRepo = resposity.NewUserRepo()
}

// GetUserInfo ...
func (service *UserService) GetUserInfo(id uint) (resp models.User, err error) {
	resp, err = userRepo.GetUserInfo(id)
	return
}

// GetUserInfo ...
func (service *UserService) CreateUser(username, password string) (err error) {

	user := models.NewUser()
	user.Username = username

	user.Salt = utils.GetSalt()
	user.Password = utils.MD5Encode(password + user.Salt + "!@#$%^&*")

	err = userRepo.CreateUser(user)
	return
}
