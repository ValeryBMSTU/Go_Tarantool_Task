package repository

import (
	"flag"
	tarantool "github.com/tarantool/go-tarantool"
)

var ConnStr string = ""

func (rep *Repository) NewRepository() error {
	tarantoolAddr := flag.String("addr", "127.0.0.1:3301", "tarantool addr")

	flag.Parse()

	var err error
	opts := tarantool.Opts{User: "guest"}

	tConn, err := tarantool.Connect(*tarantoolAddr, opts)
	if err != nil {
		return err
	}

	rep.tar = tConn

	return nil
}