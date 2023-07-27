package auth

import (
	"net/http"
	m "smapurv1_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfilerController struct {
	DB *gorm.DB
}

func NewProfilerController(DB *gorm.DB) ProfilerController {
	return ProfilerController{DB}
}

func (prc *ProfilerController) GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)

	// userResponse := &m.UserResponse{
	// 	ID_User:       currentUser.ID_User,
	// 	Fullname:      currentUser.Fullname,
	// 	Username:      currentUser.Username,
	// 	Last_Login:    currentUser.Last_Login,
	// 	Last_Logout:   currentUser.Last_Logout,
	// 	Photo:         currentUser.Photo,
	// 	Email:         currentUser.Email,
	// 	NIK:           currentUser.NIK,
	// 	Info:          currentUser.Info,
	// 	Status:        currentUser.Status,
	// 	Serial_Number: currentUser.Serial_Number,
	// 	Creator:       currentUser.Creator,
	// 	Updater:       currentUser.Updater,
	// }

	var userWithRole []m.UserRole
	if err := prc.DB.Where("kd_user = ?", currentUser.ID_User).Preload("User").Preload("Role").Find(&userWithRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userWithRole}})
}

func (prc *ProfilerController) UpdateProfile(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUser m.Users
	result := prc.DB.First(&updatedUser, "kd_user = ?", currentUser.ID_User)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	now := time.Now()
	userToUpdate := m.Users{
		Fullname:   payload.Fullname,
		Username:   payload.Username,
		Photo:      payload.Photo,
		Email:      payload.Email,
		NIK:        payload.NIK,
		Info:       payload.Info,
		Status:     payload.Status,
		Updated_At: now,
		Updater:    currentUser.Fullname,
	}

	prc.DB.Model(&updatedUser).Updates(userToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

func (prc *ProfilerController) UpdatePassword(c *gin.Context) {
	//Implement Update Password Method
}
