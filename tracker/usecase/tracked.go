package usecase

import (
	//"tracker/models"
	"tracker/repository"
)

type trackedUsecase struct {
	trackedRepository repository.TrackedRepository
}

type TrackedUsecase interface {
}

func InitUserUsecase(trackedRepository repository.TrackedRepository) TrackedUsecase {
	return &trackedUsecase{
		trackedRepository,
	}
}
