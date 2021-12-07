package logic

import (
	"context"
	"user/models"
	"user/repository"
)

type LocationBasedUserLogic interface {
	SaveLocationBasedUser(ctx context.Context, locationData models.Location) (*models.Location, error)
}

type location struct {
	repo repository.LocBaseRepo
}

func NewLocBaseUserLogic(repo repository.LocBaseRepo) LocationBasedUserLogic {
	return &location{
		repo: repo,
	}
}

func (l *location) SaveLocationBasedUser(ctx context.Context, locationData models.Location) (*models.Location, error) {
	err := locationData.LocaValidation()
	if err != nil {
		return nil, err
	}

	err = l.repo.SaveLocationBasedUser(ctx, models.Location{
		Country: locationData.Country,
		Deleted: false,
		City:    locationData.City,
	})
	if err != nil {
		return nil, err
	}
	return &locationData, nil
}
