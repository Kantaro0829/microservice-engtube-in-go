package usecase

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
	response "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/response"
)

type UserRepository interface {
	Store(json.CreateUserRequest) errors.MyError
	Select() response.AllUserResponse //[]model.User
	Delete(id string)
	Update(u model.User, name string)
}
