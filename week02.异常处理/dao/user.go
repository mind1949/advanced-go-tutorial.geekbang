package dao

import (
	"context"
	"database/sql"
	"fmt"
	"week02/biz"
	"week02/bizerr"

	"github.com/pkg/errors"
)

var (
	_ biz.UserRepo = NewUserRepo()
)

func NewUserRepo() biz.UserRepo {
	return &userRepo{
		db: globalDB,
	}
}

type userRepo struct {
	db *sql.DB
}

func (*userRepo) table() string {
	return "users"
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	var (
		user  biz.User
		query = fmt.Sprintf("select id, nickname, email, pwd from %s where id = ?", u.table())
	)
	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Nickname, &user.Email, &user.Pwd)
	if err == sql.ErrNoRows {
		return nil, errors.Wrapf(bizerr.ErrNotFound, "err: %s, query: %q, args: {id: %d}", err, query, id)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "query: %q, args: {id: %d}", query, id)
	}

	return &user, nil
}
