package auth

import (
	"net/http"
	m "smapurv1_api/models"
	u "smapurv1_api/utils"
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

	var userWithRole []m.UserRole
	if err := prc.DB.Select("DISTINCT ON (kd_user) *").Where("kd_user = ?", currentUser.ID_User).Preload("User").Preload("Role").Find(&userWithRole).Error; err != nil {
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
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdatePasswordByProfileRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	hashedOldPassword, _ := u.HashingPassword(payload.OldPassword)

	if err := u.CheckPasswordHash(hashedOldPassword, currentUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid old password"})
		return
	}

	if payload.NewPassword != payload.NewPasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Password doesn't match"})
		return
	}

	hashedNewPassword, _ := u.HashingPassword(payload.NewPassword)

	var updatedUser m.Users
	result := prc.DB.First(&updatedUser, "kd_user = ?", currentUser.ID_User)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	updatedUser.Password = hashedNewPassword
	updatedUser.Updated_At = time.Now()
	updatedUser.Updater = currentUser.Fullname
	prc.DB.Save(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password updated successfully"})
}
