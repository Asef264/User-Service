package repository

import (
	"fmt"
	"log"
	"user/config"
	"user/constants"
	"user/repository/scyllaQueries"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var DBS gocqlx.Session

func Init() {
	scylladbConnection()
}

func scylladbConnection() {
	fmt.Println("address: ", config.Configs.Scylladb.Address)
	cluster := gocql.NewCluster(config.Configs.Scylladb.Address...)
	log.Println("connecting with scylladb server...")
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Println("unable to create session", err)
	}
	DBS = session

	createScyllaTables()

	cluster.Keyspace = constants.KeyspaceName
}

func createScyllaTables() {
	err := DBS.Query(scyllaQueries.CreateUserKeyspaceQuery, []string{}).Exec()
	if err != nil {
		log.Printf("error on scylla key space creation, error: %s \n", err.Error())
	}
	err = DBS.Query(scyllaQueries.CreateUserTableQuery, []string{}).Exec()
	if err != nil {
		log.Printf("error on scylla  users table creation, error: %s \n", err.Error())
	}

	err = DBS.Query(scyllaQueries.CreateUsernamesTableQuery, []string{}).Exec()
	if err != nil {
		log.Printf("error on scylla  usernames table creation, error: %s \n", err.Error())
	}

	err = DBS.Query(scyllaQueries.CreateLocationsTableQuery, []string{}).Exec()
	if err != nil {
		log.Printf("error on scylla  links table creation, error: %s \n", err.Error())
	}

}
