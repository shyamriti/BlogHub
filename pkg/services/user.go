package services

import (
	"BlogHub/pkg/dto"
	"BlogHub/pkg/models"
	"BlogHub/pkg/repo"
	"BlogHub/pkg/utils"
	"fmt"
)

func RegisterUserService(reqUser dto.RegisterUserRequest) (dto.ResponseUser, error) {
	var err error
	user := models.User{
		Name:     reqUser.Name,
		Email:    reqUser.Email,
		Password: reqUser.Password,
	}

	user.ID, err = repo.SaveUser(&user)

	if err != nil {
		return dto.ResponseUser{}, err
	}
	return dto.ResponseUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func LoginUserService(reqUser dto.LoginUserRequest) (string, error) {
	exist, user := repo.IsUserExist(reqUser.Email)

	if !exist {
		return "", fmt.Errorf("user not found")
	}

	if err := utils.CheckPasswordHash(reqUser.Password, user.Password); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("token generation failed")
	}

	return token, nil
}
