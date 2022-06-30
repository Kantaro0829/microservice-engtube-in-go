package controller

import (
	"net/http"

	"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
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
	u := domain.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []domain.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}

func (controller *UserController) Update(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.Interactor.Update(user, user.Name)
}
