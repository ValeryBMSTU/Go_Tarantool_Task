package repository

import(
	tarantool "github.com/tarantool/go-tarantool"
)

type Repository struct {
	connectionString string
	tar tarantool.Connector
}

type IRepository interface {

}
