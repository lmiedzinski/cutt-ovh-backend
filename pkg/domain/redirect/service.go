package redirect

import (
	"context"

	"github.com/teris-io/shortid"
)

type RedirectService struct {
	repository *RedirectRepo
}

func NewRedirectService(repository *RedirectRepo) *RedirectService {
	return &RedirectService{repository: repository}
}

func (s *RedirectService) createRedirect(ctx context.Context, url string) (redirect, error) {
	slug, err := shortid.Generate()
	if err != nil {
		return redirect{}, err
	}
	er := redirect{Url: url, Slug: slug}
	err = s.repository.createRedirect(context.Background(), er)
	if err != nil {
		return redirect{}, err
	}
	return er, nil
}

func (s *RedirectService) getRedirect(ctx context.Context, slug string) (redirect, error) {
	er, err := s.repository.getRedirectBySlug(context.Background(), slug)
	if err != nil {
		return redirect{}, err
	}
	return er, nil
}
