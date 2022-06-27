package biz

import (
	"context"
	"errors"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
	"unsafe"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Token        string
	Bio          string
	Image        string
	PasswordHash string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	uc UserRepo
	pr ProfileRepo

	log  *log.Helper
	jwtc *conf.JWT
}

func NewUserUsecase(uc UserRepo, pr ProfileRepo, logger log.Logger, jwtc *conf.JWT) *UserUsecase {
	return &UserUsecase{uc: uc, pr: pr, log: log.NewHelper(logger), jwtc: jwtc}
}

func (uu *UserUsecase) Register(ctx context.Context, username, email, password string) (*User, error) {
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashed,
	}
	if err := uu.uc.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Username: username,
		Token:    uu.generateToken(username),
	}, nil
}

func (uu *UserUsecase) Login(ctx context.Context, email string, password string) (*User, error) {
	user, err := uu.uc.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(user.PasswordHash, password) {
		return nil, errors.New("login failed")
	}
	return &User{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    "",
	}, nil
}

func (uu *UserUsecase) generateToken(username string) string {
	token, err := auth.GenerateJWT(uu.jwtc.Token, username)
	if err != nil {
		return ""
	}
	return token
}

func hashPassword(pwd string) (string, error) {
	bytes := *(*[]byte)(unsafe.Pointer(&pwd))
	newBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&newBytes)), nil
}

func verifyPassword(hashed, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), *(*[]byte)(unsafe.Pointer(&pwd)))
	return err == nil
}
