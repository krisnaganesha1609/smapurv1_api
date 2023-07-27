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

type HubKeluargaController struct {
	DB *gorm.DB
}

func NewHubKeluargaController(DB *gorm.DB) HubKeluargaController {
	return HubKeluargaController{DB}
}

func (hkc *HubKeluargaController) CreateHubKeluarga(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateHubKeluargaRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newHubKeluarga := m.HubKeluarga{
		ID_Hubungan:   payload.ID_Hubungan,
		Nama_Hubungan: payload.Nama_Hubungan,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := hkc.DB.Create(&newHubKeluarga)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newHubKeluarga})
}

func (hkc *HubKeluargaController) UpdateHubKeluarga(c *gin.Context) {
	hubKeluargaId := c.Param("hubKeluargaId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateHubKeluargaRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedHubKeluarga m.HubKeluarga
	result := hkc.DB.First(&updatedHubKeluarga, "kd_hub_keluarga = ?", hubKeluargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Hub. Keluarga not found"})
		return
	}

	now := time.Now()
	hubKeluargaToUpdate := m.HubKeluarga{
		Nama_Hubungan: payload.Nama_Hubungan,
		Status:        payload.Status,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	hkc.DB.Model(&updatedHubKeluarga).Updates(hubKeluargaToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedHubKeluarga})
}

func (hkc *HubKeluargaController) FindHubKeluargaById(c *gin.Context) {
	hubKeluargaId := c.Param("hubKeluargaId")

	var hubKeluarga m.HubKeluarga
	result := hkc.DB.First(&hubKeluarga, "kd_hub_keluarga = ?", hubKeluargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Hub. Keluarga not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": hubKeluarga})
}

func (hkc *HubKeluargaController) FindHubKeluargas(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var hubKeluargas []m.HubKeluarga
	results := hkc.DB.Limit(intLimit).Offset(offset).Find(&hubKeluargas)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(hubKeluargas), "data": hubKeluargas})
}

func (hkc *HubKeluargaController) DeleteHubKeluarga(c *gin.Context) {
	hubKeluargaId := c.Param("hubKeluargaId")

	result := hkc.DB.Delete(&m.HubKeluarga{}, "kd_hub_keluarga = ?", hubKeluargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Hub. Keluarga not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
