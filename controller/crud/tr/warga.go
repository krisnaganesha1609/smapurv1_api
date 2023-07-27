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

type WargaController struct {
	DB *gorm.DB
}

func NewWargaController(DB *gorm.DB) WargaController {
	return WargaController{DB}
}

func (wc *WargaController) CreateWarga(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.CreateWargaRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newWarga := m.Warga{
		ID_Warga:      payload.ID_Warga,
		KD_Banjar:     payload.KD_Banjar,
		KD_Tempek:     payload.KD_Tempek,
		Nama_Warga:    payload.Nama_Warga,
		NIK:           payload.NIK,
		Tempat_Lahir:  payload.Tempat_Lahir,
		Tgl_Lahir:     payload.Tgl_Lahir,
		Jenis_Kelamin: payload.Jenis_Kelamin,
		Pekerjaan:     payload.Pekerjaan,
		KD_Pendidikan: payload.KD_Pendidikan,
		Alamat1:       payload.Alamat1,
		Alamat2:       payload.Alamat2,
		No_Telp:       payload.No_Telp,
		No_Hp:         payload.No_Hp,
		Email:         payload.Email,
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

	result := wc.DB.Create(&newWarga)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newWarga})
}

func (wc *WargaController) UpdateWarga(c *gin.Context) {
	wargaId := c.Param("wargaId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateWargaRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedWarga m.Warga
	result := wc.DB.First(&updatedWarga, "warga_id = ?", wargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	now := time.Now()
	wargaToUpdate := m.Warga{
		KD_Banjar:     payload.KD_Banjar,
		KD_Tempek:     payload.KD_Tempek,
		Nama_Warga:    payload.Nama_Warga,
		NIK:           payload.NIK,
		Tempat_Lahir:  payload.Tempat_Lahir,
		Tgl_Lahir:     payload.Tgl_Lahir,
		Jenis_Kelamin: payload.Jenis_Kelamin,
		Pekerjaan:     payload.Pekerjaan,
		KD_Pendidikan: payload.KD_Pendidikan,
		Alamat1:       payload.Alamat1,
		Alamat2:       payload.Alamat2,
		No_Telp:       payload.No_Telp,
		No_Hp:         payload.No_Hp,
		Email:         payload.Email,
		KD_Agama:      payload.KD_Agama,
		KD_Hub:        payload.KD_Hub,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	wc.DB.Model(&updatedWarga).Updates(wargaToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedWarga})
}

func (wc *WargaController) FindWargaById(c *gin.Context) {
	wargaId := c.Param("wargaId")

	var warga m.Warga
	result := wc.DB.Preload("Banjar").Preload("Tempek").Preload("Pendidikan").Preload("HubKeluarga").Preload("Agama").First(&warga, "warga_id = ?", wargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Warga not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": warga})
}

func (wc *WargaController) FindWargas(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var wargas []m.Warga
	results := wc.DB.Limit(intLimit).Offset(offset).Preload("Banjar").Preload("Tempek").Preload("Pendidikan").Preload("HubKeluarga").Preload("Agama").Find(&wargas)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(wargas), "data": wargas})
}

func (wc *WargaController) DeleteWarga(c *gin.Context) {
	wargaId := c.Param("wargaId")

	result := wc.DB.Delete(&m.Warga{}, "warga_id = ?", wargaId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Warga not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
