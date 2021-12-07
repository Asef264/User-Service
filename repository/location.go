package repository

import (
	"context"
	"log"
	"user/models"
)

type LocBaseRepo interface {
	SaveLocationBasedUser(ctx context.Context, locationData models.Location) error
}

type location struct {
}

func NewLocBasedRepo() LocBaseRepo {
	return new(location)
}

func (l *location) SaveLocationBasedUser(ctx context.Context, locationData models.Location) error {
	stmt, names := models.LocationTable.Insert()

	err := DBS.Query(stmt, names).BindStruct(locationData).ExecRelease()
	if err != nil {
		log.Println("error on queriing or binding the struct in repository/location ")
		return err
	}

	return nil
}
