package ad

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	m "smapurv1_api/models"
	u "smapurv1_api/utils"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateUserRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if payload.Password != payload.Password_Confirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Password does not match"})
		return
	}

	hashedPassword, err := u.HashingPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := m.Users{
		ID_User:  payload.ID_User,
		Fullname: payload.Fullname,
		Username: payload.Username,
		//Extract last login & logout info from auth
		// Last_Login: ,
		// Last_Logout: ,
		Photo:         payload.Photo,
		Email:         payload.Email,
		NIK:           payload.NIK,
		Password:      hashedPassword,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := uc.DB.Create(&newUser)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUser m.Users
	result := uc.DB.First(&updatedUser, "kd_user = ?", userId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	now := time.Now()
	userToUpdate := m.Users{
		Fullname:      payload.Fullname,
		Username:      payload.Username,
		Photo:         payload.Photo,
		Email:         payload.Email,
		NIK:           payload.NIK,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	uc.DB.Model(&updatedUser).Updates(userToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

func (uc *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	var user m.Users
	result := uc.DB.First(&user, "kd_user = ?", userId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "USer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (uc *UserController) FindUsers(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []m.Users
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	result := uc.DB.Delete(&m.Users{}, "kd_user = ?", userId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
