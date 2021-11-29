package usecase

import (
	"context"

	"github.com/teris-io/shortid"

	"github.com/lmiedzinski/cutt-ovh-backend/internal/entity"
)

type RedirectUseCase struct {
	repository RedirectRepository
}

func NewRedirectUseCase(repository RedirectRepository) *RedirectUseCase {
	return &RedirectUseCase{repository: repository}
}

func (uc *RedirectUseCase) CreateRedirect(ctx context.Context, url string) (entity.Redirect, error) {
	slug, err := shortid.Generate()
	if err != nil {
		return entity.Redirect{}, err
	}
	er := entity.Redirect{Url: url, Slug: slug}
	err = uc.repository.CreateRedirect(context.Background(), er)
	if err != nil {
		return entity.Redirect{}, err
	}
	return er, nil
}
