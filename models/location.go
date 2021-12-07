package models

import "errors"

type Location struct {
	Country string `json:"country"`
	Deleted bool   `json:"deleted"`
	City    string `json:"city"`
}

func (l *Location) LocaValidation() error {
	if err := l.CountryValidation(); err != nil {
		return err
	}
	return nil
}

func (l *Location) CountryValidation() error {
	if l.Country == "" {
		return errors.New("کشور خود را انتخاب کنید")
	}
	return nil
}
