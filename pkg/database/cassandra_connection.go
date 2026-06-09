package database

import "github.com/gocql/gocql"

type CassandraConnection interface {
	Query(string, ...any) *gocql.Query
	Exec(string, ...any) error
}

type cassandraConnection struct {
	session *gocql.Session
}

func NewCassandraConnection() *cassandraConnection {
	cluster := gocql.NewCluster("localhost")
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return &cassandraConnection{session}
}

func (c *cassandraConnection) Query(query string, args ...any) *gocql.Query {
	return c.session.Query(query, args...)
}

func (c *cassandraConnection) Exec(query string, args ...any) error {
	return c.session.Query(query, args...).Exec()
}
