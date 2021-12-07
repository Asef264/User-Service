package logic

import (
	"context"
	"user/models"
	"user/repository"
)

type UsernameLogicInterface interface {
	SaveUsernameBasedUser(ctx context.Context, usernameData models.Username) (*models.Username, error)
}

type username struct {
	repo repository.UsernameRepo
}

func NewUsernameLogic(repo repository.UsernameRepo) UsernameLogicInterface {
	return &username{
		repo: repo,
	}
}

func (u *username) SaveUsernameBasedUser(ctx context.Context, usernameData models.Username) (*models.Username, error) {

	if err := usernameData.Validation(); err != nil {
		return nil, err
	}
	err := u.repo.SaveUsernameBasedUser(ctx, models.Username{
		Username: usernameData.Username,
		Deleted:  false,
		Id:       usernameData.Id,
	})
	if err != nil {
		return nil, err
	}
	return &usernameData, nil

}
