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

type BanjarController struct {
	DB *gorm.DB
}

func NewBanjarController(DB *gorm.DB) BanjarController {
	return BanjarController{DB}
}

func (bc *BanjarController) CreateBanjar(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(m.Users)
	var payload *m.CreateBanjarRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	newBanjar := m.Banjar{
		ID_Banjar:     payload.ID_Banjar,
		Nama_Banjar:   payload.Nama_Banjar,
		Info:          payload.Info,
		Status:        payload.Status,
		Serial_Number: payload.Serial_Number,
		Created_At:    now,
		Creator:       currentUser.Fullname,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	result := bc.DB.Create(&newBanjar)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Data is already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBanjar})
}

func (bc *BanjarController) UpdateBanjar(c *gin.Context) {
	banjarId := c.Param("banjarId")
	currentUser := c.MustGet("currentUser").(m.Users)

	var payload *m.UpdateBanjarRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedBanjar m.Banjar
	result := bc.DB.First(&updatedBanjar, "kd_banjar = ?", banjarId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Banjar not found"})
		return
	}

	now := time.Now()
	banjarToUpdate := m.Banjar{
		Nama_Banjar:   payload.Nama_Banjar,
		Status:        payload.Status,
		Info:          payload.Info,
		Serial_Number: payload.Serial_Number,
		Updated_At:    now,
		Updater:       currentUser.Fullname,
	}

	bc.DB.Model(&updatedBanjar).Updates(banjarToUpdate)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedBanjar})
}

func (bc *BanjarController) FindBanjarById(c *gin.Context) {
	banjarId := c.Param("banjarId")

	var banjar m.Banjar
	result := bc.DB.First(&banjar, "kd_banjar = ?", banjarId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Banjar not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": banjar})
}

func (bc *BanjarController) FindBanjars(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var banjars []m.Banjar
	results := bc.DB.Limit(intLimit).Offset(offset).Find(&banjars)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(banjars), "data": banjars})
}

func (bc *BanjarController) DeleteBanjar(c *gin.Context) {
	banjarId := c.Param("banjarId")

	result := bc.DB.Delete(&m.Banjar{}, "kd_banjar = ?", banjarId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Banjar not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
