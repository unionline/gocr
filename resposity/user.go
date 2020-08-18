package resposity

import (
	"gocr/initializers/db"
	"gocr/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// GetUserInfo ...
func (*UserRepo) GetUserInfo(id uint) (user models.User, err error) {
	user, err = Query.GetUserInfo(db.Db, id)
	return
}

// CreateUserInfo ...
func (*UserRepo) CreateUser(user *models.User) (err error) {
	err = Insert.CreateUser(db.Db, user)
	return
}
