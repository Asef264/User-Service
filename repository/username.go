package repository

import (
	"context"
	"log"
	"user/models"
)

type UsernameRepo interface {
	SaveUsernameBasedUser(ctx context.Context, usernameData models.Username) error
}

type username struct {
}

func NewRepoUsername() UsernameRepo {
	return new(username)
}

func (u *username) SaveUsernameBasedUser(ctx context.Context, usernameData models.Username) error {
	stmt, names := models.UsernameTable.Insert()
	err := DBS.Query(stmt, names).BindStruct(usernameData).ExecRelease()
	if err != nil {
		log.Println("error on saving userData in username table", err.Error())
		return err
	}
	return nil
}
