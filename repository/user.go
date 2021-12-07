package repository

import (
	"context"
	"log"
	"user/constants"
	"user/models"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
)

type UserRepoInterface interface {
	SaveNewUser(ctx context.Context, userData models.UserDb) error
	GetUser(ctx context.Context, id string) (*models.UserDb, error)
}

type user struct {
}

func NewRepoUserInterface() UserRepoInterface {
	return new(user)
}

func (u *user) SaveNewUser(ctx context.Context, userData models.UserDb) error {
	stmt, names := models.UsersTable.Insert()
	err := DBS.Query(stmt, names).BindStruct(userData).ExecRelease()
	if err != nil {
		log.Println("error on saving new user", err.Error())
		return err
	}
	return nil
}

func (u *user) GetUser(ctx context.Context, id string) (*models.UserDb, error) {

	stmt, names := qb.Select(constants.KeyspaceName+"."+constants.UserTableName).
		Where(qb.Eq("id"), qb.Eq("deleted")).ToCql()

	query := DBS.Query(stmt, names).BindStruct(&models.UserDb{
		Id:      id,
		Deleted: false,
	})

	defer query.Release()

	var user models.UserDb
	err := query.Get(&user)
	if err != nil {
		if err == gocql.ErrNotFound || err == gocql.ErrNoMetadata {
			return nil, nil
		}

		log.Printf("error on reading id")
		return nil, err
	}

	return &user, nil
}
