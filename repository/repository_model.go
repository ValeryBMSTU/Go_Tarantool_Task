package repository

import(
	"github.com/ValeryBMSTU/Go_Tarantool_Task/models"
	tarantool "github.com/tarantool/go-tarantool"
)

type Repository struct {
	connectionString string
	tar tarantool.Connector
}

type IRepository interface {
	Insert(keyValue models.PostKeyValue) (interface{}, error)
	Select(key string) (interface{}, error)
	Delete(key string) (interface{}, error)
	Update(keyValue models.PostKeyValue) (interface{}, error)
}
