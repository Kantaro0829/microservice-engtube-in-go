package controller

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
	json "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/json"
	model "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain/model"

	"github.com/Kantaro0829/microservice-engtube-in-go/userService/interfaces/database"
	"github.com/Kantaro0829/microservice-engtube-in-go/userService/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

//ルーティングのハンドラ
func (controller *UserController) Create(c *gin.Context) {

	user := json.CreateUserRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempError := controller.Interactor.Add(user)

	if tempError.Message != "" {
		c.JSON(http.StatusConflict, gin.H{"message": tempError.Message, "status": http.StatusConflict})
		return
	}
	createdUsers := controller.Interactor.GetInfo()
	fmt.Println(createdUsers)

	c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
	return

}

func (controller *UserController) GetUser(c *gin.Context) {
	res := controller.Interactor.GetInfo()
	if res.Error.Message != "" {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}

func (controller *UserController) Update(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.Interactor.Update(user, user.Name)
}
