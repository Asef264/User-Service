package scyllaQueries

import (
	"user/constants"
)

const (
	//create keyspace
	CreateUserKeyspaceQuery = `CREATE KEYSPACE IF NOT EXISTS ` +
		constants.KeyspaceName + ` WITH  
		replication = {'class': 'SimpleStrategy', 'replication_factor': '1' } `

	// create user links table
	CreateUserTableQuery = `
CREATE TABLE IF NOT EXISTS  ` + constants.KeyspaceName + `.` + constants.UserTableName + ` (
                id text,
				deleted boolean,
                username text,						 
				fullname text,
				phone_number text,
                password  text,
				email text,
				bio text,
				profile_image text,
				location text,
				country text,
				city text,
				is_admin boolean,
				created_at timestamp,
				
                   PRIMARY KEY ((id), deleted));
`

	// create usernames table
	CreateUsernamesTableQuery = `
CREATE TABLE IF NOT EXISTS  ` + constants.KeyspaceName + `.` + constants.UsernameTableName + ` (
                   username text,
				   deleted boolean,
                   user_id text,
                   PRIMARY KEY ((username), deleted));
`

	// create locations table
	CreateLocationsTableQuery = `
CREATE TABLE IF NOT EXISTS  ` + constants.KeyspaceName + `.` + constants.LocationTableName + ` (
                   country text,
				   deleted boolean,
                   city text, 
                   PRIMARY KEY ((country), deleted, city));
`
)
