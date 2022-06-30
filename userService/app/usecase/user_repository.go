package usecase

import (
	"github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
	Update(u domain.User, name string)
}
