package tr

import (
	"net/http"
	m "smapurv1_api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WargaKKController struct {
	DB *gorm.DB
}

func NewWargaKKController(DB *gorm.DB) WargaKKController {
	return WargaKKController{DB}
}

func (wkc *WargaKKController) CreateWargaKK(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.CreateWargaKKRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newWargaKK := m.WargaKK{
		KD_Warga:      payload.KD_Warga,
		Nama_Warga:    payload.Nama_Warga,
		NIK:           payload.NIK,
		Tempat_Lahir:  payload.Tempat_Lahir,
		Tgl_Lahir:     payload.Tgl_Lahir,
		Jenis_Kelamin: payload.Jenis_Kelamin,
		Pekerjaan:     payload.Pekerjaan,
		KD_Pendidikan: payload.KD_Pendidikan,
		KD_Agama:      payload.KD_Agama,
		KD_Hub:        payload.KD_Hub,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := wkc.DB.Create(&newWargaKK)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newWargaKK})
}

func (wkc *WargaKKController) UpdateWargaKK(c *gin.Context) {
	wargaKKId := c.Param("wargaKKId")
	nik := c.Param("nik")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateWargaKKRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedWargaKK m.WargaKK
	result := wkc.DB.First(&updatedWargaKK, "warga_id = ?", "nik = ?", wargaKKId, nik)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	now := time.Now()
	wargaKKToUpdate := m.WargaKK{
		KD_Warga:      payload.KD_Warga,
		Nama_Warga:    payload.Nama_Warga,
		NIK:           payload.NIK,
		Tempat_Lahir:  payload.Tempat_Lahir,
		Tgl_Lahir:     payload.Tgl_Lahir,
		Jenis_Kelamin: payload.Jenis_Kelamin,
		Pekerjaan:     payload.Pekerjaan,
		KD_Pendidikan: payload.KD_Pendidikan,
		KD_Agama:      payload.KD_Agama,
		KD_Hub:        payload.KD_Hub,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	wkc.DB.Model(&updatedWargaKK).Updates(wargaKKToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedWargaKK})
}

func (wkc *WargaKKController) FindWargaKKById(c *gin.Context) {
	wargaId := c.Param("wargaId")
	nik := c.Param("nik")

	var wargaKK m.WargaKK
	result := wkc.DB.Preload("Warga").First(&wargaKK, "warga_id = ?", "nik = ?", wargaId, nik)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "WargaKK not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": wargaKK})
}

func (wkc *WargaKKController) FindWargaKKs(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var wargaKKs []m.WargaKK
	results := wkc.DB.Limit(intLimit).Offset(offset).Select("DISTINCT ON (warga_id) *").Order("warga_id, nik").Preload("Warga").Find(&wargaKKs)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(wargaKKs), "data": wargaKKs})
}

func (wkc *WargaKKController) DeleteWargaKK(c *gin.Context) {
	wargaId := c.Param("wargaId")
	nik := c.Param("nik")

	result := wkc.DB.Delete(&m.WargaKK{}, "warga_id = ?", "nik = ?", wargaId, nik)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "WargaKK not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
