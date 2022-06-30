package usecase

import "github.com/Kantaro0829/microservice-engtube-in-go/userService/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}

func (interactor *UserInteractor) Update(u domain.User, name string) {
	interactor.UserRepository.Update(u, name)
}