package services

import (
	"paulocenteno/microsrvcs/domain"
	error_utils "paulocenteno/microsrvcs/utils"
)

func GetUser(userId int64) (*domain.User, *error_utils.AppError) {
	return domain.GetUser(userId)
}
