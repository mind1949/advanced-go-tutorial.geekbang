package data

import (
	"blog/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ArticleRepo = (*articleRepo)(nil)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper("data/article", logger),
	}
}

func (a *articleRepo) CreateArticle(ctx context.Context, title, content string) (*biz.Article, error) {
	// TODO:
	return nil, nil
}

func (a *articleRepo) DeleteArticle(ctx context.Context, id int64) (err error) {
	// TODO:
	return nil
}

func (a *articleRepo) UpdateArticle(ctx context.Context) (*biz.Article, error) {
	// TODO:
	return nil, nil
}

func (a *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	// TODO:
	return nil, nil
}

func (a *articleRepo) ListArticles(ctx context.Context, page, perPage int64) ([]*biz.Article, error) {
	// TODO:
	return nil, nil
}
