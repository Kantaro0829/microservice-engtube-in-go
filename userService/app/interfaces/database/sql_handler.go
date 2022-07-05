package database

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
)

// type SqlHandler interface {
// 	Create(object interface{})
// 	FindAll(object interface{})
// 	DeleteById(object interface{}, id string)
// 	UpdateById(object domain.User, name string)
// }

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	DeleteById(object interface{}, id string)
	UpdateById(object model.User, name string)
}
