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

type PendidikanController struct {
	DB *gorm.DB
}

func NewPendidikanController(DB *gorm.DB) PendidikanController {
	return PendidikanController{DB}
}

func (pc *PendidikanController) CreatePendidikan(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreatePendidikanRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newPendidikan := m.Pendidikan{
		ID_Pendidikan:   payload.ID_Pendidikan,
		Nama_Pendidikan: payload.Nama_Pendidikan,
		Info:            payload.Info,
		Status:          payload.Status,
		Serial_Number:   payload.Serial_Number,
		Created_At:      now,
		Creator:         currentUser.Fullname,
		Updated_At:      now,
		Updater:         currentUser.Fullname,
	}

	result := pc.DB.Create(&newPendidikan)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPendidikan})
}

func (pc *PendidikanController) UpdatePendidikan(c *gin.Context) {
	pendidikanId := c.Param("pendidikanId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdatePendidikanRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedPendidikan m.Pendidikan
	result := pc.DB.First(&updatedPendidikan, "kd_pendidikan = ?", pendidikanId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Pendidikan not found"})
		return
	}

	now := time.Now()
	pendidikanToUpdate := m.Pendidikan{
		Nama_Pendidikan: payload.Nama_Pendidikan,
		Status:          payload.Status,
		Info:            payload.Info,
		Serial_Number:   payload.Serial_Number,
		Updated_At:      now,
		Updater:         currentUser.Fullname,
	}

	pc.DB.Model(&updatedPendidikan).Updates(pendidikanToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPendidikan})
}

func (pc *PendidikanController) FindPendidikanById(c *gin.Context) {
	pendidikanId := c.Param("pendidikanId")

	var pendidikan m.Pendidikan
	result := pc.DB.First(&pendidikan, "kd_pendidikan = ?", pendidikanId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Pendidikan not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": pendidikan})
}

func (pc *PendidikanController) FindPendidikans(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var pendidikans []m.Pendidikan
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&pendidikans)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(pendidikans), "data": pendidikans})
}

func (pc *PendidikanController) DeletePendidikan(c *gin.Context) {
	pendidikanId := c.Param("pendidikanId")

	result := pc.DB.Delete(&m.Pendidikan{}, "kd_pendidikan = ?", pendidikanId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Pendidikan not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
