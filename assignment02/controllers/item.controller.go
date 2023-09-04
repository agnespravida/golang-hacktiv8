package controllers

import (
	"net/http"
	"time"

	"assignment02/databases"
	"assignment02/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type OrderRequest struct {
	CustomerName string        `json:"customerName"`
	OrderedAt    string        `json:"orderedAt"`
	Items        []models.Item `json:"items"`
}

func CreateOrder(ctx *gin.Context) {
	//ini proses request
	var newOrder OrderRequest
	db := databases.GetDB()

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ordered_date, _ = time.Parse(time.RFC3339, newOrder.OrderedAt)
	Order := &models.Order{
		CustomerName: newOrder.CustomerName,
		OrderedAt:    ordered_date,
		Items:        newOrder.Items,
	}

	db.Create(&Order)

	//ini response
	ctx.JSON(http.StatusCreated, gin.H{
		"data":         Order,
		"kodeResponse": http.StatusCreated,
		"message":      "Data berhasil ditambahkan",
	})
}

func GetAllOrders(ctx *gin.Context) {
	db := databases.GetDB()
	var orders []models.Order
	db.Preload("Items").Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{
		"data":         orders,
		"kodeResponse": http.StatusOK,
		"message":      "Data berhasil dittemukan",
	})
}

func GetOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := databases.GetDB()

	Order := &models.Order{}
	err := db.Preload("Items").First(Order, "order_id = ?", orderID).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"kodeResponse": http.StatusNotFound,
			"message":      "Data gagal didapat",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data":         Order,
			"kodeResponse": http.StatusOK,
			"message":      "Data berhasil dittemukan",
		})
	}
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := databases.GetDB()

	Order := &models.Order{}
	Item := &models.Item{}
	errItem := db.Delete(Item, "order_id = ?", orderID).Error
	errOrder := db.Delete(Order, "order_id = ?", orderID).Error

	if errItem != nil || errOrder != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"kodeResponse": http.StatusBadRequest,
			"message":      "Gagal menghapus data",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"kodeResponse": http.StatusOK,
			"message":      "Berhasil meghapus data",
		})
	}
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := databases.GetDB()
	var newOrder OrderRequest

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ordered_date, _ = time.Parse(time.RFC3339, newOrder.OrderedAt)
	Order := &models.Order{
		CustomerName: newOrder.CustomerName,
		OrderedAt:    ordered_date,
		Items:        newOrder.Items,
	}

	// Update order by condition order_id = param order_id
	result := db.Model(&models.Order{}).Clauses(clause.Returning{}).Where("order_id = ?", orderID).Updates(&Order)

	// Update items by condition item_id = request item_id
	if len(Order.Items) > 0 {

		for i := 0; i < len(Order.Items); i++ {
			fmt.Println(Order.Items[i])
			db.Model(&models.Item{}).Clauses(clause.Returning{}).Where("item_id = ?", Order.Items[i].ItemID).Updates(&Order.Items[i])
		}

	}
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"kodeResponse": http.StatusBadRequest,
			"message":      "Gagal update data",
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data":         Order,
			"kodeResponse": http.StatusOK,
			"message":      "Berhasil update data",
		})
	}
}
