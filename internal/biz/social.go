package biz

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type ArticleRepo interface {
	Save(context.Context, *v1.Article) (*v1.Article, error)
	Update(context.Context, *v1.Article) (*v1.Article, error)
	FindByID(context.Context, int64) (*v1.Article, error)
	ListByHello(context.Context, string) ([]*v1.Article, error)
	ListAll(context.Context) ([]*v1.Article, error)
}

type CommentRepo interface {
	Save(context.Context, *v1.Comment) (*v1.Comment, error)
	Update(context.Context, *v1.Comment) (*v1.Comment, error)
	FindByID(context.Context, int64) (*v1.Comment, error)
	ListByHello(context.Context, string) ([]*v1.Comment, error)
	ListAll(context.Context) ([]*v1.Comment, error)
}

type TagRepo interface {
	Save(context.Context, *v1.Tag) (*v1.Tag, error)
	Update(context.Context, *v1.Tag) (*v1.Tag, error)
	FindByID(context.Context, int64) (*v1.Tag, error)
	ListByHello(context.Context, string) ([]*v1.Tag, error)
	ListAll(context.Context) ([]*v1.Tag, error)
}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo

	log *log.Helper
}

func NewSocialUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (sc *SocialUsecase) CreateArticle(ctx context.Context, article *v1.Article) (*v1.Article, error) {
	sc.log.WithContext(ctx).Infof("CreateArticle: %v", article.Title)
	return sc.ar.Save(ctx, article)
}
