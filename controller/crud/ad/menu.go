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

type MenuController struct {
	DB *gorm.DB
}

func NewMenuController(DB *gorm.DB) MenuController {
	return MenuController{DB}
}

func (mc *MenuController) CreateMenu(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateMenuRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newMenu := m.Menu{
		ID_Menu:        payload.ID_Menu,
		Menu_Name:      payload.Menu_Name,
		KD_Parent_Menu: payload.KD_Parent_Menu,
		Level_Menu:     payload.Level_Menu,
		Group_Menu:     payload.Group_Menu,
		Link:           payload.Link,
		Icon_Menu:      payload.Icon_Menu,
		Info:           payload.Info,
		Status:         payload.Status,
		Serial_Number:  payload.Serial_Number,
		Created_At:     now,
		Creator:        currentUser.Fullname,
		Updated_At:     now,
		Updater:        currentUser.Fullname,
	}

	result := mc.DB.Create(&newMenu)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Menu is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMenu})
}

func (mc *MenuController) UpdateMenu(c *gin.Context) {
	menuId := c.Param("menuId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateMenuRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedMenu m.Menu
	result := mc.DB.First(&updatedMenu, "kd_menu = ?", menuId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Menu not found"})
		return
	}

	now := time.Now()
	menuToUpdate := m.Menu{
		Menu_Name:      payload.Menu_Name,
		Icon_Menu:      payload.Icon_Menu,
		Group_Menu:     payload.Group_Menu,
		KD_Parent_Menu: payload.KD_Parent_Menu,
		Level_Menu:     payload.Level_Menu,
		Link:           payload.Link,
		Status:         payload.Status,
		Info:           payload.Info,
		Serial_Number:  payload.Serial_Number,
		Updated_At:     now,
		Updater:        currentUser.Fullname,
	}

	mc.DB.Model(&updatedMenu).Updates(menuToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedMenu})
}

func (mc *MenuController) FindMenuById(c *gin.Context) {
	menuId := c.Param("menuId")

	var menu m.Menu
	result := mc.DB.First(&menu, "kd_menu = ?", menuId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Menu not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": menu})
}

func (mc *MenuController) FindMenus(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var menus []m.Menu
	results := mc.DB.Limit(intLimit).Offset(offset).Find(&menus)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(menus), "data": menus})
}

func (mc *MenuController) DeleteMenu(c *gin.Context) {
	menuId := c.Param("menuId")

	result := mc.DB.Delete(&m.Menu{}, "kd_menu = ?", menuId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Menu not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
