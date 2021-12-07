package models

import (
	"errors"
	"net/mail"
	"time"
)

type User struct {
	Fullname     string    `json:"fullname"`
	Username     string    `json:"username"`
	Id           string    `json:"id"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Bio          string    `json:"bio"`
	ProfileImage string    `json:"profile_image"`
	Location     location  `json:"location"`
	Country      string    `json:"contry"`
	City         string    `json:"city"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	Deleted      bool      `json:"deleted"`
}

//a user for save in db
type UserDb struct {
	Fullname     string    `json:"fullname"`
	Username     string    `json:"username"`
	Id           string    `json:"id"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Bio          string    `json:"bio"`
	ProfileImage string    `json:"profile_image"`
	Location     string    `json:"location"`
	Country      string    `json:"contry"`
	City         string    `json:"city"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	Deleted      bool      `json:"deleted"`
}

type location struct {
	N string `json:"n"`
	S string `json:"s"`
}

type CountryBasedUser struct {
	Country string `json:"country"`
	Deleted bool   `json:"deleted"`
	City    string `json:"city"`
}

type IdBasedUser struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

//creating a struct for login request
type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//creating an array to maintin usernames avoiding to be choosed again
var Usernames []string

//login response struct
type LoginResponse struct {
	AccessKey string `json:"access_key"`
}

func (u *User) Validate() error {
	if err := u.NameValidation(); err != nil {
		return err
	}
	if err := u.UsernameValidation(); err != nil {
		return err
	}
	if err := u.PasswordValidation(); err != nil {
		return err
	}
	if err := u.EmailValidation(); err != nil {
		return nil
	}
	return nil
}

//NameValidation function emplementation
func (u *User) NameValidation() error {
	if len(u.Fullname) <= 2 {
		return errors.New("نامی که وارد کرده اید بسیار کوتاه است / نام شما باید بیش از ۲ حرف داشته باشد")
	}
	if len(u.Fullname) >= 20 {
		return errors.New("بنظر میرسد نام شما از ۲۰ حرف بیشتر است / شما میتوانید تنها ۲۰ حرف به عنوان نام وارد کنید")
	}
	if u.Fullname == "" {
		return errors.New("لطفا یک نام انتخاب کنید")
	}
	return nil
}

//UsernameValidation function emplementation
func (u *User) UsernameValidation() error {
	//Usernames is an arrey that keeps all picked usernames from users
	//and here by this range we check about existing usernames avoiding to be chosed again
	for _, exist := range Usernames {
		if exist == u.Username {
			return errors.New("این نام کاربری توسط شخص دیگری اختیار شده است نام دیگری انتخاب کنید")
		}
	}

	if len(u.Username) <= 3 {
		return errors.New("نام کاربری شما بایدبیش از ۳ حرف داشته باشد")
	}
	if len(u.Username) >= 10 {
		return errors.New("نام کاربری شما میتواند تنها شامل ۹ حرف باشد")
	}
	if u.Username == "" {
		return errors.New("لطفا یک نام کاربری انتخاب کنید!")
	}
	return nil
}

// PasswordValidation function emplementation
func (u *User) PasswordValidation() error {
	if u.Password == u.Fullname || u.Password == u.Email || u.Password == u.Username || u.Password == u.PhoneNumber {
		return errors.New("برای کلمه ی عبور نمی توانید  شماره تماس . نام انتخابیتان.ایمیل یا نام کاربریتان را وارد کنید")
	}
	if len(u.Password) <= 3 {
		return errors.New("کلمه ی عبور حداقل باید ۴ کاراکتر باشد")
	}
	if len(u.Password) >= 20 {
		return errors.New(" محدوده ی مجاز کلمه عبور ۴ و ۲۰ کاراکتر است")
	}
	return nil
}

//EmailValidation function emplementation
func (u *User) EmailValidation() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("آدرس ایمیل را اشتباه وارد کرده اید")
	}
	return nil
}
