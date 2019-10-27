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

	outKeyValue, err := use.Cast(res)

	return outKeyValue, nil
}

func (use *Usecase) GetValue(Key string) (models.OutKeyValue, error) {

	res, err := use.Repository.Select(Key)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	outKeyValue, err := use.Cast(res)

	return outKeyValue, nil
}

func (use *Usecase) Delete(Key string) (models.OutKeyValue, error) {
	res, err := use.Repository.Delete(Key)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	outKeyValue, err := use.Cast(res)

	return outKeyValue, nil
}

func (use *Usecase) Set(keyValue models.PostKeyValue) (models.OutKeyValue, error) {
	res, err := use.Repository.Update(keyValue)
	if err != nil {
		return models.OutKeyValue{}, err
	}

	outKeyValue, err := use.Cast(res)

	return outKeyValue, nil
}

func (use *Usecase) Cast(keyValue interface{}) (models.OutKeyValue, error) {
	resList, ok := keyValue.([]interface{})
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}
	key, ok := resList[0].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	///
	tempValue, ok := resList[1].(string)
	if !ok {
		return models.OutKeyValue{}, errors.New("incorrect res in GetValue")
	}

	var value interface{}
	_ = json.Unmarshal([]byte(tempValue), &value)

	var outKeyValue models.OutKeyValue
	outKeyValue.Key = key
	outKeyValue.Value = value

	return outKeyValue, nil
}

