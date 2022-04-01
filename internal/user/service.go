package user

import (
	"context"
	"rest-api-tutorial/cmd/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	//TODO
	return
}
