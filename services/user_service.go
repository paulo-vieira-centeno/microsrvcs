package services

import "paulocenteno/microsrvcs/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
