package usecase

import (
	"context"

	"github.com/lmiedzinski/cutt-ovh-backend/internal/entity"
)

type (
	Redirect interface {
		CreateRedirect(ctx context.Context, url string) (entity.Redirect, error)
	}

	RedirectRepository interface {
		CreateRedirect(ctx context.Context, er entity.Redirect) error
	}
)
