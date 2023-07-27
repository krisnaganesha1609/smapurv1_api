package ms

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	m "smapurv1_api/models"
)

type TempekController struct {
	DB *gorm.DB
}

func NewTempekController(DB *gorm.DB) TempekController {
	return TempekController{DB}
}

func (tc *TempekController) CreateTempek(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateTempekRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newTempek := m.Tempek{
		ID_Tempek:     payload.ID_Tempek,
		Nama_Tempek:   payload.Nama_Tempek,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := tc.DB.Create(&newTempek)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTempek})
}

func (tc *TempekController) UpdateTempek(c *gin.Context) {
	tempekId := c.Param("tempekId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateTempekRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedTempek m.Tempek
	result := tc.DB.First(&updatedTempek, "kd_tempek = ?", tempekId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Tempek not found"})
		return
	}

	now := time.Now()
	tempekToUpdate := m.Tempek{
		Nama_Tempek:   payload.Nama_Tempek,
		Status:        payload.Status,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	tc.DB.Model(&updatedTempek).Updates(tempekToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTempek})
}

func (tc *TempekController) FindTempekById(c *gin.Context) {
	tempekId := c.Param("tempekId")

	var tempek m.Tempek
	result := tc.DB.First(&tempek, "kd_tempek = ?", tempekId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Tempek not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": tempek})
}

func (tc *TempekController) FindTempeks(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tempeks []m.Tempek
	results := tc.DB.Limit(intLimit).Offset(offset).Find(&tempeks)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tempeks), "data": tempeks})
}

func (tc *TempekController) DeleteTempek(c *gin.Context) {
	tempekId := c.Param("tempekId")

	result := tc.DB.Delete(&m.Tempek{}, "kd_tempek = ?", tempekId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Tempek not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
