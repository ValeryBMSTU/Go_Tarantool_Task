package models


type PostKeyValue struct {
	Key string			`json:"key"`
	Value interface{} 	`json:"value"`
}

type KeyValue struct {
	Key string			`json:"key"`
	Value interface{}	`json:"value"`
}

type OutKeyValue struct {
	Key interface{}		`json:"key"`
	Value interface{}	`json:"value"`
}

type PutValue struct {
	Value interface{}	`json:"value"`
}