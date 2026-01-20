package repository

import "context"

type FolderRepo interface {
	Create(ctx context.Context) error
}

type repo struct{}

func FolderRepoMethod() FolderRepo {
	return &repo{}
}

func (r *repo) Create(ctx context.Context) error {
	return nil
}
