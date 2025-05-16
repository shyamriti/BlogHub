package repo

import (
	"BlogHub/pkg/db"
	"BlogHub/pkg/models"
)

func CheckForValidUser(userID uint) bool {
	var user models.User
	result := db.DB.Where("id = ?", userID).First(&user)

	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func IsUserExist(email string) (bool, models.User) {
	var user models.User
	result := db.DB.Where("email=?", email).First(&user)

	if result.RowsAffected != 0 {
		return true, user
	}

	return false, user
}

func SaveUser(user *models.User) (uint, error) {
	if err := db.DB.Save(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func FindUserNameByID(userId uint)(string, error){
	var user models.User
if err:= db.DB.Where("id=?", userId).First(&user).Error; err!=nil{
	return "", err
}
return user.Name, nil
}