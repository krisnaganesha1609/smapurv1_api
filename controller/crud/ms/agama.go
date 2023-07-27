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

type AgamaController struct {
	DB *gorm.DB
}

func NewAgamaController(DB *gorm.DB) AgamaController {
	return AgamaController{DB}
}

func (ac *AgamaController) CreateAgama(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateAgamaRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newAgama := m.Agama{
		ID_Agama:      payload.ID_Agama,
		Nama_Agama:    payload.Nama_Agama,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := ac.DB.Create(&newAgama)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newAgama})
}

func (ac *AgamaController) UpdateAgama(c *gin.Context) {
	agamaId := c.Param("agamaId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateAgamaRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedAgama m.Agama
	result := ac.DB.First(&updatedAgama, "kd_agama = ?", agamaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Agama not found"})
		return
	}

	now := time.Now()
	agamaToUpdate := m.Agama{
		Nama_Agama:    payload.Nama_Agama,
		Status:        payload.Status,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	ac.DB.Model(&updatedAgama).Updates(agamaToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedAgama})
}

func (ac *AgamaController) FindAgamaById(c *gin.Context) {
	agamaId := c.Param("agamaId")

	var agama m.Agama
	result := ac.DB.First(&agama, "kd_agama = ?", agamaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Agama not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": agama})
}

func (ac *AgamaController) FindAgamas(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var agamas []m.Agama
	results := ac.DB.Limit(intLimit).Offset(offset).Find(&agamas)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(agamas), "data": agamas})
}

func (ac *AgamaController) DeleteAgama(c *gin.Context) {
	agamaId := c.Param("agamaId")

	result := ac.DB.Delete(&m.Agama{}, "kd_agama = ?", agamaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Agama not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
