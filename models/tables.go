package models

import (
	"user/constants"

	"github.com/scylladb/gocqlx/table"
)

// definition of user table
var UserTableMetadata = table.Metadata{
	Name: constants.KeyspaceName + "." + constants.UserTableName,
	Columns: []string{
		"id",
		"deleted",
		"username",
		"fullname",
		"password",
		"email",
		"phone_number",
		"bio",
		"profile_image",
		"location",
		"country",
		"city",
		"is_admin",
		"created_at",
	},
	PartKey: []string{"id"},
	SortKey: []string{"deleted", "username"},
}

//difination of username table
var usernameTableMetadata = table.Metadata{
	Name: constants.KeyspaceName + "." + constants.UsernameTableName,
	Columns: []string{
		"username",
		"deleted",
	},
	PartKey: []string{"username"},
	SortKey: []string{"deleted"},
}

//difination of location table
var locationTableMetadata = table.Metadata{
	Name: constants.KeyspaceName + "." + constants.LocationTableName,
	Columns: []string{
		"country",
		"deleted",
		"city",
	},
	PartKey: []string{"country"},
	SortKey: []string{"deleted", "city"},
}

var UsersTable = table.New(UserTableMetadata)

var UsernameTable = table.New(usernameTableMetadata)

var LocationTable = table.New(locationTableMetadata)
