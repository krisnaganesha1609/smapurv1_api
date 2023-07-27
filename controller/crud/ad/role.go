package ad

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	m "smapurv1_api/models"
)

type RoleController struct {
	DB *gorm.DB
}

func NewRoleController(DB *gorm.DB) RoleController {
	return RoleController{DB}
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateRoleRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newRole := m.Role{
		ID_Role:       payload.ID_Role,
		Role_Name:     payload.Role_Name,
		Icon_Role:     payload.Icon_Role,
		Group_Role:    payload.Group_Role,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := rc.DB.Create(&newRole)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Role is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newRole})
}

func (rc *RoleController) UpdateRole(c *gin.Context) {
	roleId := c.Param("roleId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateRoleRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedRole m.Role
	result := rc.DB.First(&updatedRole, "kd_role = ?", roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Role not found"})
		return
	}

	now := time.Now()
	roleToUpdate := m.Role{
		Role_Name:     payload.Role_Name,
		Icon_Role:     payload.Icon_Role,
		Group_Role:    payload.Group_Role,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	rc.DB.Model(&updatedRole).Updates(roleToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedRole})
}

func (rc *RoleController) FindRoleById(c *gin.Context) {
	roleId := c.Param("roleId")

	var role m.Role
	result := rc.DB.First(&role, "kd_role = ?", roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": role})
}

func (rc *RoleController) FindRoles(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var roles []m.Role
	results := rc.DB.Limit(intLimit).Offset(offset).Find(&roles)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(roles), "data": roles})
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	roleId := c.Param("roleId")

	result := rc.DB.Delete(&m.Role{}, "kd_role = ?", roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Role not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
