package database

import (
	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
	response "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/response"
)

type UserRepository struct {
	SqlHandler
}

//DB操作のための関数
//同階層の./interface/database/user_repogitory.goから呼び出している
// func (db *UserRepository) Store(u domain.User) {
// 	db.Create(&u)
// }

func (db *UserRepository) Store(u json.CreateUserRequest) errors.MyError {
	errorTemp := db.Create(u)
	return errorTemp
}

// func (db *UserRepository) Select() []model.User {
// 	user := []model.User{}
// 	db.FindAll(user) //&
// 	return user
// }

func (db *UserRepository) Select() response.AllUserResponse {
	user := []model.User{}
	result := db.FindAll(&user)
	return result
}

func (db *UserRepository) Delete(id string) {
	user := []model.User{}
	db.DeleteById(&user, id)
}
func (db *UserRepository) Update(u model.User, name string) {
	db.UpdateById(u, name)
}
