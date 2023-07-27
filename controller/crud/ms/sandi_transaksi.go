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

type SandiTransaksiController struct {
	DB *gorm.DB
}

func NewSandiTransaksiController(DB *gorm.DB) SandiTransaksiController {
	return SandiTransaksiController{DB}
}

func (stc *SandiTransaksiController) CreateSandiTransaksi(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateSandiTransaksiRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newSandiTransaksi := m.SandiTransaksi{
		ID_Sandi:      payload.ID_Sandi,
		Nama_Sandi:    payload.Nama_Sandi,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := stc.DB.Create(&newSandiTransaksi)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newSandiTransaksi})
}

func (stc *SandiTransaksiController) UpdateSandiTransaksi(c *gin.Context) {
	sandiTransaksiId := c.Param("sandiTransaksiId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateSandiTransaksiRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedSandiTransaksi m.SandiTransaksi
	result := stc.DB.First(&updatedSandiTransaksi, "kd_sandi = ?", sandiTransaksiId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Sandi Transaksi not found"})
		return
	}

	now := time.Now()
	sandiTransaksiToUpdate := m.SandiTransaksi{
		Nama_Sandi:    payload.Nama_Sandi,
		Status:        payload.Status,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	stc.DB.Model(&updatedSandiTransaksi).Updates(sandiTransaksiToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedSandiTransaksi})
}

func (stc *SandiTransaksiController) FindSandiTransaksiById(c *gin.Context) {
	sandiTransaksiId := c.Param("sandiTransaksiId")

	var sandiTransaksi m.SandiTransaksi
	result := stc.DB.First(&sandiTransaksi, "kd_sandi = ?", sandiTransaksiId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Sandi Transaksi not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": sandiTransaksi})
}

func (stc *SandiTransaksiController) FindSandiTransaksis(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var sandiTransaksis []m.SandiTransaksi
	results := stc.DB.Limit(intLimit).Offset(offset).Find(&sandiTransaksis)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(sandiTransaksis), "data": sandiTransaksis})
}

func (stc *SandiTransaksiController) DeleteSandiTransaksi(c *gin.Context) {
	sandiTransaksiId := c.Param("sandiTransaksiId")

	result := stc.DB.Delete(&m.SandiTransaksi{}, "kd_sandi = ?", sandiTransaksiId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Sandi Transaksi not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
