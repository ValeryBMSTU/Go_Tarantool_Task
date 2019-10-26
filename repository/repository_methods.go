package repository

import (
	"encoding/json"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/models"
	tarantool "github.com/tarantool/go-tarantool"
)

var ConnStr string = ""

func (rep *Repository) NewRepository() error {
	opts := tarantool.Opts{User: "guest"}
	tConn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		return err
	}
	rep.tar = tConn

	return nil
}

func (rep *Repository) Insert(keyValue models.PostKeyValue) (Resp interface{}, Err error) {
	value, _ := json.Marshal(keyValue.Value)

	var temp interface{}
	_ = json.Unmarshal(value, &temp)
	println(temp)

	resp, err := rep.tar.Insert("test",
		[]interface{}{keyValue.Key, string(value)})
	if err != nil {
		return models.KeyValue{}, err
	}

	str2 := string(value)
	_ = json.Unmarshal([]byte(str2), &temp)
	println(temp)

	str3 := `{"2321r":"kmefqwe","werwqf":"dferfewe"}`
	_ = json.Unmarshal([]byte(str3), &temp)
	println(temp)

	return resp.Data[0], nil
}

func (rep *Repository) Select(key string) (Resp interface{}, Err error) {
	resp, err := rep.tar.Select("test", "primary", 0, 1, tarantool.IterEq, []interface{}{key})
	if err != nil {
		return models.KeyValue{}, err
	}

	if len(resp.Data) == 0 {
		return models.KeyValue{}, tarantool.Error{Code: 4, Msg: "key not found"}
	}

	return resp.Data[0], nil
}

func (rep *Repository) Delete(key string) (Resp interface{}, Err error) {
	resp, err := rep.tar.Delete("test", "primary", []interface{}{key})
	if err != nil {
		return models.KeyValue{}, err
	}

	if len(resp.Data) == 0 {
		return models.KeyValue{}, tarantool.Error{Code: 4, Msg: "key not found"}
	}


	return resp.Data[0], nil
}

func (rep *Repository) Update(keyValue models.PostKeyValue) (Resp interface{}, Err error) {
	resp, err := rep.tar.Update("test", "primary", []interface{}{keyValue.Key},
		[]interface{}{[]interface{}{"=", 1, keyValue.Value}})
	if err != nil {
		return models.KeyValue{}, err
	}

	if len(resp.Data) == 0 {
		return models.KeyValue{}, tarantool.Error{Code: 4, Msg: "key not found"}
	}


	return resp.Data[0], nil
}