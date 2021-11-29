package repository

import (
	"context"

	"github.com/lmiedzinski/cutt-ovh-backend/internal/entity"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/postgres"
)

type RedirectRepo struct {
	*postgres.Postgres
}

func NewRedirectPostgresRepository(pg *postgres.Postgres) *RedirectRepo {
	return &RedirectRepo{pg}
}

func (r *RedirectRepo) CreateRedirect(ctx context.Context, er entity.Redirect) error {
	sql, args, err := r.Builder.
		Insert("redirects").
		Columns("slug, url").
		Values(er.Slug, er.Url).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
