package usecase

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
	response "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/response"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u json.CreateUserRequest) errors.MyError {
	tempError := interactor.UserRepository.Store(u)
	return tempError
}

// func (interactor *UserInteractor) GetInfo() []model.User {
// 	return interactor.UserRepository.Select()
// }

func (interactor *UserInteractor) GetInfo() response.AllUserResponse {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}

func (interactor *UserInteractor) Update(u model.User, name string) {
	interactor.UserRepository.Update(u, name)
}
