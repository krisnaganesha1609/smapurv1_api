package ad

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	m "smapurv1_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoleController struct {
	DB *gorm.DB
}

func NewUserRoleController(DB *gorm.DB) UserRoleController {
	return UserRoleController{DB}
}

func (urc *UserRoleController) CreateUserRole(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateUserWithRoleRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newUserRole := m.UserRole{
		KD_User:       payload.KD_User,
		KD_Role:       payload.KD_Role,
		Default_Role:  payload.Default_Role,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := urc.DB.Create(&newUserRole)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "There's conflict with other data"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUserRole})
}

func (urc *UserRoleController) UpdateUserRole(c *gin.Context) {
	userId := c.Param("userId")
	roleId := c.Param("roleId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.CreateUserWithRoleRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUserRole m.UserRole
	result := urc.DB.First(&updatedUserRole, "kd_user = ?", "kd_role = ?", userId, roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	now := time.Now()
	userRoleToUpdate := m.UserRole{
		KD_User:       payload.KD_User,
		KD_Role:       payload.KD_Role,
		Default_Role:  payload.Default_Role,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	urc.DB.Model(&updatedUserRole).Updates(userRoleToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUserRole})
}

func (urc *UserRoleController) FindUserRoleById(c *gin.Context) {
	userId := c.Param("userId")
	roleId := c.Param("roleId")

	var userRole m.UserRole
	result := urc.DB.Preload("User").Preload("Role").First(&userRole, "kd_user = ?", " kd_role = ?", userId, roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": userRole})
}

func (urc *UserRoleController) FindUserRoles(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var userRoles []m.UserRole
	results := urc.DB.Limit(intLimit).Offset(offset).Select("DISTINCT ON (kd_user) *").Order("kd_user, kd_role").Preload("User").Preload("Role").Find(&userRoles)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(userRoles), "data": userRoles})
}

func (urc *UserRoleController) DeleteUserRole(c *gin.Context) {
	userId := c.Param("userId")
	roleId := c.Param("roleId")

	result := urc.DB.Delete(&m.UserRole{}, "kd_user = ?", "kd_role = ?", userId, roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
