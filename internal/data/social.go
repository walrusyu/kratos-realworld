package data

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ArticleRepo = &articleRepo{}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *articleRepo) Save(context.Context, *v1.Article) (*v1.Article, error) {
	return nil, nil
}
func (r *articleRepo) Update(context.Context, *v1.Article) (*v1.Article, error) {
	return nil, nil
}
func (r *articleRepo) FindByID(context.Context, int64) (*v1.Article, error) {
	return nil, nil
}
func (r *articleRepo) ListByHello(context.Context, string) ([]*v1.Article, error) {
	return nil, nil
}
func (r *articleRepo) ListAll(context.Context) ([]*v1.Article, error) {
	return nil, nil
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var _ biz.CommentRepo = &commentRepo{}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (r *commentRepo) Save(context.Context, *v1.Comment) (*v1.Comment, error) {
	return nil, nil
}
func (r *commentRepo) Update(context.Context, *v1.Comment) (*v1.Comment, error) {
	return nil, nil
}
func (r *commentRepo) FindByID(context.Context, int64) (*v1.Comment, error) {
	return nil, nil
}
func (r *commentRepo) ListByHello(context.Context, string) ([]*v1.Comment, error) {
	return nil, nil
}
func (r *commentRepo) ListAll(context.Context) ([]*v1.Comment, error) {
	return nil, nil
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var _ biz.TagRepo = &tagRepo{}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func (r *tagRepo) Save(context.Context, *v1.Tag) (*v1.Tag, error) {
	return nil, nil
}
func (r *tagRepo) Update(context.Context, *v1.Tag) (*v1.Tag, error) {
	return nil, nil
}
func (r *tagRepo) FindByID(context.Context, int64) (*v1.Tag, error) {
	return nil, nil
}
func (r *tagRepo) ListByHello(context.Context, string) ([]*v1.Tag, error) {
	return nil, nil
}
func (r *tagRepo) ListAll(context.Context) ([]*v1.Tag, error) {
	return nil, nil
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var _ biz.ProfileRepo = &profileRepo{}
