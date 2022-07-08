package response

import (
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
)

type UserRespnse struct {
	Data  []model.User
	Error errors.MyError
}

type AllUserResponse struct {
	Data  interface{}
	Error errors.MyError
}
