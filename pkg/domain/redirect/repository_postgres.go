package redirect

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/postgres"
)

const _defaultEntityCap = 64

type RedirectRepo struct {
	*postgres.Postgres
}

var (
	notFoundError = fmt.Errorf("NOT FOUND")
)

func NewRedirectPostgresRepository(pg *postgres.Postgres) *RedirectRepo {
	return &RedirectRepo{pg}
}

func (r *RedirectRepo) createRedirect(ctx context.Context, er redirect) error {
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

func (r *RedirectRepo) getRedirectBySlug(ctx context.Context, slug string) (redirect, error) {
	sql, args, err := r.Builder.
		Select("slug, url").
		From("redirects").
		Where(squirrel.Eq{"slug": slug}).
		ToSql()
	if err != nil {
		return redirect{}, fmt.Errorf("RedirectRepo - GetRedirectBySlug - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return redirect{}, fmt.Errorf("RedirectRepo - GetRedirectBySlug - Query: %w", err)
	}
	defer rows.Close()

	entities := make([]redirect, 0, _defaultEntityCap)

	for rows.Next() {
		e := redirect{}

		err = rows.Scan(&e.Slug, &e.Url)
		if err != nil {
			return redirect{}, fmt.Errorf("RedirectRepo - GetRedirectBySlug - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	if len(entities) == 0 {
		return redirect{}, nil
	}

	return entities[0], nil
}
