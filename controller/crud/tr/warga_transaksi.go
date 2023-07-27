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

type WargaTransaksiController struct {
	DB *gorm.DB
}

func NewWargaTransaksiController(DB *gorm.DB) WargaTransaksiController {
	return WargaTransaksiController{DB}
}

func (wtc *WargaTransaksiController) CreateWargaTransaksi(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.CreateWargaTransaksiRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newWargaTransaksi := m.WargaTransaksi{
		KD_Warga:           payload.KD_Warga,
		KD_Sandi:           payload.KD_Sandi,
		Tgl_Transaksi:      payload.Tgl_Transaksi,
		Serial_Number:      payload.Serial_Number,
		TransactionNominal: payload.TransactionNominal,
		Status:             payload.Status,
		Created_At:         now,
		Creator:            currentUser.Fullname,
		Updated_At:         now,
		Updater:            currentUser.Fullname,
	}

	result := wtc.DB.Create(&newWargaTransaksi)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newWargaTransaksi})
}

func (wtc *WargaTransaksiController) UpdateWargaTransaksi(c *gin.Context) {
	wargaTransaksiId := c.Param("wargaTransaksiId")
	sandiId := c.Param("sandiId")
	tglTrans := c.Param("tglTrans")
	serialId := c.Param("serialId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateWargaTransaksiRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedWargaTransaksi m.WargaTransaksi
	result := wtc.DB.First(&updatedWargaTransaksi, "warga_id = ?", "kd_sandi = ?", "tgl_transaksi = ?", "no_urut = ?", wargaTransaksiId, sandiId, tglTrans, serialId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	now := time.Now()
	WargaTransaksiToUpdate := m.WargaTransaksi{
		KD_Warga:           payload.KD_Warga,
		KD_Sandi:           payload.KD_Sandi,
		Tgl_Transaksi:      payload.Tgl_Transaksi,
		TransactionNominal: payload.TransactionNominal,
		Status:             payload.Status,
		Serial_Number:      payload.Serial_Number,
		Created_At:         now,
		Creator:            currentUser.Fullname,
		Updated_At:         now,
		Updater:            currentUser.Fullname,
	}

	wtc.DB.Model(&updatedWargaTransaksi).Updates(WargaTransaksiToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedWargaTransaksi})
}

func (wtc *WargaTransaksiController) FindWargaTransaksiById(c *gin.Context) {
	wargaTransaksiId := c.Param("wargaTransaksiId")
	sandiId := c.Param("sandiId")
	tglTrans := c.Param("tglTrans")
	serialId := c.Param("serialId")

	var wargaTransaksi m.WargaTransaksi
	result := wtc.DB.Preload("Warga").Preload("SandiTransaksi").First(&wargaTransaksi, "warga_id = ?", "kd_sandi = ?", "tgl_transaksi = ?", "no_urut = ?", wargaTransaksiId, sandiId, tglTrans, serialId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Warga Transaksi not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": wargaTransaksi})
}

func (wtc *WargaTransaksiController) FindWargaTransaksis(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var wargaTransaksis []m.WargaTransaksi
	results := wtc.DB.Limit(intLimit).Offset(offset).Select("DISTINCT ON (warga_id) *").Order("warga_id, kd_sandi, tgl_transaksi, no_urut").Preload("Warga").Preload("SandiTransaksi").Find(&wargaTransaksis)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(wargaTransaksis), "data": wargaTransaksis})
}

func (wtc *WargaTransaksiController) DeleteWargaTransaksi(c *gin.Context) {
	wargaTransaksiId := c.Param("wargaTransaksiId")
	sandiId := c.Param("sandiId")
	tglTrans := c.Param("tglTrans")
	serialId := c.Param("serialId")

	result := wtc.DB.Delete(&m.WargaTransaksi{}, "warga_id = ?", "kd_sandi = ?", "tgl_transaksi = ?", "no_urut = ?", wargaTransaksiId, sandiId, tglTrans, serialId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Warga Transaksi not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
