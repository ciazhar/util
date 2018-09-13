package util

import (
	"github.com/ciazhar/db"
	"gopkg.in/mgo.v2/bson"
)

type ok interface {
	Validate() error
}

type ErrMissingField string
type ErrNotFound string
type ErrAlreadyExist string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}

func (e ErrNotFound) Error() string {
	return string(e) + " is not found"
}

func (e ErrAlreadyExist) Error() string {
	return string(e) + " is already exists"
}

func IsValueAlreadyExists(collection string, variable string, value interface{}) error {
	count,_ := db.Mongo.C(collection).Find(bson.M{variable:value}).Count()
	if count>=1{
		return ErrAlreadyExist(collection)
	}
	return nil
}

func IsValueNotFound(collection string, variable string, value interface{}) error {
	count,_ := db.Mongo.C(collection).Find(bson.M{variable:value}).Count()
	if count==0{
		return ErrNotFound(collection)
	}
	return nil
}
