package database

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
	response "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/response"
)

type SqlHandler interface {
	Create(object json.CreateUserRequest) errors.MyError
	FindAll(object interface{}) response.AllUserResponse
	DeleteById(object interface{}, id string)
	UpdateById(object model.User, name string)
}
