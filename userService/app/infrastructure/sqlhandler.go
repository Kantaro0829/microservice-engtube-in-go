package infrastructure

import (
	"fmt"

	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	errors "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/error"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"
	"github.com/Kantaro0829/microservice-engtube-in-go/userService/interfaces/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:ecc@tcp(user-db:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

//データベースが変わった場合や使用しているフレームワークが
//変更された場合などはここを変更する
//interface層内の./database配下にinterfaceを定義する
func (handler *SqlHandler) Create(obj json.CreateUserRequest) errors.MyError {
	//Gorm.Createメソッド
	userTable := model.User{}
	userTable.Email = obj.Email
	userTable.Name = obj.Name
	userTable.Password = obj.Password
	userTable.YoutubeApiKey = obj.YoutubeApiKey
	userTable.LastWatchedVideoId = obj.LastWatchedVideoId
	userTable.ID = 0

	errors := errors.MyError{}

	if err := handler.db.Create(&userTable).Error; err != nil {
		errors.Error = err
		errors.Message = "メールアドレスの重複"
		return errors
	}

	errors.Error = nil
	errors.Message = ""
	return errors
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	//Gorm.Findメソッド
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	//Gorm.Deleteメソッド
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) UpdateById(obj model.User, name string) {
	//Gorm.Updateメソッド
	handler.db.First(&obj)
	fmt.Println("objの中身")
	fmt.Println(obj.ID)
	fmt.Println(obj.Name)
	obj.Name = name
	handler.db.Save(&obj)

}
