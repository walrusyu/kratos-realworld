package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.UserRepo = &userRepo{}

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	newUser := User{Email: user.Email, Username: user.Username, Bio: user.Bio, Image: user.Image, PasswordHash: user.PasswordHash}
	result := r.data.db.Create(&newUser)
	return result.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := new(User)
	result := r.data.db.Where("email = ?", email).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not found by email")
	}
	return &biz.User{Email: u.Email, PasswordHash: u.PasswordHash, Username: u.Username, Bio: u.Bio, Image: u.Image}, nil
}

var _ biz.ProfileRepo = &profileRepo{}

type profileRepo struct {
}

func NewProfileRepo() biz.ProfileRepo {
	return &profileRepo{}
}
