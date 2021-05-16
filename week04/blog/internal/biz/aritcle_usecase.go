package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

func NewBlogUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper("usecase/article", logger)}
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

type ArticleRepo interface {
	CreateArticle(ctx context.Context, title string, content string) (*Article, error)
	DeleteArticle(ctx context.Context, id int64) error
	UpdateArticle(ctx context.Context) (*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	ListArticles(ctx context.Context, page, perPage int64) ([]*Article, error)
}

func (uc *ArticleUsecase) Create(ctx context.Context, title, content string) (*Article, error) {
	return uc.repo.CreateArticle(ctx, title, content)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) (err error) {
	_, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		return err
	}
	return uc.repo.DeleteArticle(ctx, id)
}

func (uc *ArticleUsecase) Update(ctx context.Context) (*Article, error) {
	return uc.repo.UpdateArticle(ctx)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (*Article, error) {
	return uc.repo.GetArticle(ctx, id)
}

func (uc *ArticleUsecase) List(ctx context.Context, page, perPage int64) ([]*Article, error) {
	return uc.repo.ListArticles(ctx, page, perPage)
}
