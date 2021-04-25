package biz

import "context"

var (
	_ UserUsecase = new(userUsecase)
)

type UserUsecase interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}

type User struct {
	ID       int64
	Nickname string
	Email    string
	Pwd      string
}

type userUsecase struct {
	repo UserRepo
}

func (u *userUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	return u.repo.GetUser(ctx, id)
}
