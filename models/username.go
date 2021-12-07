package models

import "errors"

type Username struct {
	Username string `json:"username"`
	Deleted  bool   `json:"deleted"`
	Id       string `json:"id"`
}

func (u *Username) Validation() error {
	err := u.UsernameValidation()
	if err != nil {
		return err
	}
	return nil
}

func (u *Username) UsernameValidation() error {
	if len(u.Username) <= 2 || len(u.Username) >= 15 {
		return errors.New("محدوده نام کاربری ۳ تا ۱۴ کاراکتر است")
	}
	if u.Username == "" {
		return errors.New("نام کاربری باید انتخاب شود")
	}
	return nil
}
