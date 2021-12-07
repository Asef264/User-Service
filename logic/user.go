package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"user/models"
	"user/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserLogicInterface interface {
	SaveNewUser(ctx context.Context, userData models.User) (newUser *models.User, err error)
	GetUserData(ctx context.Context, id string) (*models.UserDb, error)
}

var passwordCost int

type user struct {
	repo repository.UserRepoInterface
}

func NewLogicUserInterface(repo repository.UserRepoInterface) UserLogicInterface {
	return &user{
		repo: repo,
	}
}

func (u *user) SaveNewUser(ctx context.Context, userData models.User) (*models.User, error) {

	if err := userData.Validate(); err != nil {
		return nil, err
	}

	userData.Id = uuid.New().String()

	userData.CreatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userData.Password),
		10)
	if err != nil {
		return nil, errors.New("Internal Error")
	}

	stringLocation, err := json.Marshal(userData.Location)
	if err != nil {
		return nil, errors.New("error on marshaling the location struct")
	}
	userData.Password = string(hashedPassword)

	err = u.repo.SaveNewUser(ctx, models.UserDb{
		Fullname:     userData.Fullname,
		Id:           userData.Id,
		Username:     userData.Username,
		Password:     userData.Password,
		Email:        userData.Email,
		ProfileImage: userData.ProfileImage,
		Country:      userData.Country,
		City:         userData.City,
		Location:     string(stringLocation),
		PhoneNumber:  userData.PhoneNumber,
		CreatedAt:    userData.CreatedAt,
		Deleted:      userData.Deleted,
		IsAdmin:      userData.IsAdmin,
		Bio:          userData.Bio,
	})
	if err != nil {
		return nil, err
	}

	userData.Password = "*****"
	return &userData, nil

}

func (u *user) GetUserData(ctx context.Context, id string) (*models.UserDb, error) {

	if id == "" {
		return nil, errors.New("we cant retrive a user without an id or a location or a username")
	}
	userData, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
