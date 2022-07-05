package usecase

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
)

type UserRepository interface {
	Store(model.User)
	Select() []model.User
	Delete(id string)
	Update(u model.User, name string)
}
