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

type RoleMenuController struct {
	DB *gorm.DB
}

func NewRoleMenuController(DB *gorm.DB) RoleMenuController {
	return RoleMenuController{DB}
}

//Implement Role Menu CRUD Method

func (rmc *RoleMenuController) CreateRoleMenu(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateRoleMenuRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newRoleMenu := m.RoleMenu{
		KD_Role:        payload.KD_Role,
		KD_Menu:        payload.KD_Menu,
		Allow_Add:      payload.Allow_Add,
		Allow_Edit:     payload.Allow_Edit,
		Allow_View:     payload.Allow_View,
		Allow_Del:      payload.Allow_Del,
		Allow_Approve:  payload.Allow_Approve,
		Allow_Download: payload.Allow_Download,
		Allow_Upload:   payload.Allow_Upload,
		Visible:        payload.Visible,
		Group_Role:     payload.Group_Role,
		Info:           payload.Info,
		Status:         payload.Status,
		Serial_Number:  payload.Serial_Number,
		Created_At:     now,
		Creator:        currentUser.Fullname,
		Updated_At:     now,
		Updater:        currentUser.Fullname,
	}

	result := rmc.DB.Create(&newRoleMenu)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "succcess", "data": newRoleMenu})
}

func (rmc *RoleMenuController) UpdateRoleMenu(c *gin.Context) {
	roleId := c.Param("roleId")
	menuId := c.Param("menuId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateRoleMenuRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedRoleMenu m.RoleMenu
	result := rmc.DB.First(&updatedRoleMenu, "kd_role = ?", "kd_menu = ?", roleId, menuId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	now := time.Now()
	roleMenuToUpdate := m.RoleMenu{
		KD_Role:        payload.KD_Role,
		KD_Menu:        payload.KD_Menu,
		Allow_Add:      payload.Allow_Add,
		Allow_Edit:     payload.Allow_Edit,
		Allow_View:     payload.Allow_View,
		Allow_Del:      payload.Allow_Del,
		Allow_Approve:  payload.Allow_Approve,
		Allow_Download: payload.Allow_Download,
		Allow_Upload:   payload.Allow_Upload,
		Visible:        payload.Visible,
		Group_Role:     payload.Group_Role,
		Info:           payload.Info,
		Status:         payload.Status,
		Serial_Number:  payload.Serial_Number,
		Updated_At:     now,
		Updater:        currentUser.Fullname,
	}

	rmc.DB.Model(&updatedRoleMenu).Updates(roleMenuToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedRoleMenu})
}

func (rmc *RoleMenuController) FindRoleMenuById(c *gin.Context) {
	menuId := c.Param("menuId")
	roleId := c.Param("roleId")

	var roleMenu m.RoleMenu
	result := rmc.DB.Preload("Menu").Preload("Role").First(&roleMenu, "kd_menu = ?", " kd_role = ?", menuId, roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": roleMenu})
}

func (rmc *RoleMenuController) FindRoleMenus(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var roleMenus []m.RoleMenu
	results := rmc.DB.Limit(intLimit).Offset(offset).Select("DISTINCT ON (kd_role) *").Order("kd_menu, kd_role").Preload("Menu").Preload("Role").Find(&roleMenus)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(roleMenus), "data": roleMenus})
}

func (rmc *RoleMenuController) DeleteRoleMenu(c *gin.Context) {
	menuId := c.Param("menuId")
	roleId := c.Param("roleId")

	result := rmc.DB.Delete(&m.RoleMenu{}, "kd_menu = ?", "kd_role = ?", menuId, roleId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Data not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
