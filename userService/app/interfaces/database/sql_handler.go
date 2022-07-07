package database

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
)

type SqlHandler interface {
	Create(object json.CreateUserRequest) errors.MyError
	FindAll(object interface{})
	DeleteById(object interface{}, id string)
	UpdateById(object model.User, name string)
}
