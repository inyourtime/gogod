package delivery

import "gogod/domain"

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: uu,
	}
}
