package usecase

import (
	"encoding/json"
	"errors"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/models"
	"github.com/ValeryBMSTU/Go_Tarantool_Task/repository"
	"sync"
)

func (use *Usecase) NewUsecase(mu *sync.Mutex, rep repository.IRepository) {
	use.Repository = rep
	use.Mu = mu
}

func (use *Usecase) AddKeyValue(KeyValue models.PostKeyValue) (models.OutKeyValue, error) {
	res, err := use.Repository.Insert(KeyValue)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	resList, ok := res.([]interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	key, ok := resList[0].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	///
	value, ok := resList[1].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	var temp interface{}
	str2 := string(value)
	_ = json.Unmarshal([]byte(str2), &temp)
	println(temp)

	var outKeyValue models.OutKeyValue
	outKeyValue.Key = key
	outKeyValue.Value = temp

	return outKeyValue, nil
}

func (use *Usecase) GetValue(Key string) (models.OutKeyValue, error) {

	res, err := use.Repository.Select(Key)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	resList, ok := res.([]interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	key, ok := resList[0].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	value := resList[1].(interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	var outKeyValue models.OutKeyValue
	outKeyValue.Key = key
	outKeyValue.Value = value

	return outKeyValue, nil
}

func (use *Usecase) Delete(Key string) (models.OutKeyValue, error) {
	res, err := use.Repository.Delete(Key)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	resList, ok := res.([]interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in Set")
	}
	key, ok := resList[0].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	value := resList[1].(interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	var outKeyValue models.OutKeyValue
	outKeyValue.Key = key
	outKeyValue.Value = value

	return outKeyValue, nil
}

func (use *Usecase) Set(keyValue models.PostKeyValue) (models.OutKeyValue, error) {
	res, err := use.Repository.Update(keyValue)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	resList, ok := res.([]interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in Set")
	}
	key, ok := resList[0].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	value := resList[1].(interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	var outKeyValue models.OutKeyValue
	outKeyValue.Key = key
	outKeyValue.Value = value

	return outKeyValue, nil
}
