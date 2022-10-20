package controller

import (
	"net/http"

	"github.com/B6202385/G04-Farmmart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Orders
func CreateOrder(c *gin.Context) {
	var Order entity.Order
	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Order})
}

func ListOrderbyUser(c *gin.Context) {
	var Order []entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("User").Raw("SELECT * FROM orders WHERE user_id = ? ",id).Find(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Order})
}

// GET /Order/:id
func GetOrder(c *gin.Context) {
	var Order entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("User").Raw("SELECT * FROM orders WHERE id = ?", id).Find(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Order})
}

// GET /Orders
func ListOrders(c *gin.Context) {
	var Orders []entity.Order
	if err := entity.DB().Preload("User").Raw("SELECT * FROM orders").Find(&Orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Orders})
}

// DELETE /Orders/:id
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Orders
func UpdateOrder(c *gin.Context) {
	var Order entity.Order
	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Order.ID).First(&Order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	if err := entity.DB().Save(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Order})
}