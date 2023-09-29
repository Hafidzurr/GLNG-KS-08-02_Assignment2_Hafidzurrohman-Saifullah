package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/models"

	"github.com/gin-gonic/gin"
)

// PatchOrderByID mengubah pesanan berdasarkan ID yang diberikan dengan data yang dikirim dalam permintaan JSON.
func PatchOrderByID(c *gin.Context) {
	// Mendapatkan ID pesanan dari parameter URL
	orderID := c.Param("orderID")
	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		// Jika ID tidak valid, kembalikan respons JSON dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	var existingOrder models.Order
	if err := c.ShouldBindJSON(&existingOrder); err != nil {
		// Jika data JSON tidak dapat di-bind ke struct Order, kembalikan respons dengan status BadRequest
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Panggil fungsi QueryPatchByID untuk mengubah pesanan berdasarkan ID
	updatedOrder, err := QueryPatchByID(existingOrder, uint(convertedOrderID))
	if err != nil {
		// Jika pesanan tidak ditemukan, kembalikan respons dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Order with ID %v not found", orderID),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	// Kembalikan pesanan yang telah diubah dalam respons JSON dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v has been successfully updated", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

// CreateOrders membuat pesanan baru berdasarkan data yang diberikan dalam permintaan JSON.
func CreateOrders(c *gin.Context) {
	var newOrder models.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		// Jika data JSON tidak dapat di-bind ke struct Order, kembalikan respons dengan status BadRequest
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Panggil fungsi QueryCreate untuk membuat pesanan baru
	newOrder = QueryCreate(newOrder)

	// Kembalikan pesanan baru dalam respons JSON dengan status Created
	c.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Data successfully created",
		"status":  http.StatusCreated,
	})
}

// GetOrderByID mengambil pesanan berdasarkan ID yang diberikan dalam parameter URL.
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		// Jika ID tidak valid, kembalikan respons JSON dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	// Panggil fungsi QueryGetByID untuk mengambil pesanan berdasarkan ID
	order, err := QueryGetByID(uint(convertedOrderID))

	if err != nil {
		// Jika terjadi kesalahan saat mengambil pesanan, kembalikan respons dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Error: %v", err), // Mengembalikan pesan error jika terjadi kesalahan saat mengambil pesanan
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	if order.Order_id == 0 {
		// Jika pesanan tidak ditemukan, kembalikan respons dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Order with ID %s not found", orderID),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	// Kembalikan pesanan yang ditemukan dalam respons JSON dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": fmt.Sprintf("Order with ID %s has been successfully retrieved", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

// GetAllOrders mengambil semua pesanan yang ada.
func GetAllOrders(c *gin.Context) {
	// Panggil fungsi QueryGetAll untuk mengambil semua pesanan
	orders := QueryGetAll()

	// Kembalikan semua pesanan dalam respons JSON dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "Orders fetched successfully",
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})

}

// DeleteOrder menghapus pesanan berdasarkan ID yang diberikan dalam parameter URL.
func DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		// Jika ID tidak valid, kembalikan respons JSON dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	// Panggil fungsi QueryDeleteByID untuk menghapus pesanan berdasarkan ID
	QueryDeleteByID(uint(convertedOrderID))

	// Kembalikan respons bahwa pesanan telah berhasil dihapus dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": fmt.Sprintf("Order with ID %v Has been successfully deleted", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}

// UpdateOrderByID memperbarui pesanan berdasarkan ID yang diberikan dengan data yang dikirim dalam permintaan JSON.
func UpdateOrderByID(c *gin.Context) {
	var updatedOrder models.Order
	orderID := c.Param("orderID")

	convertedOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		// Jika ID tidak valid, kembalikan respons JSON dengan status NotFound
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("Invalid Params"),
			"status":  fmt.Sprintf("%d", http.StatusNotFound),
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		// Jika data JSON tidak dapat di-bind ke struct Order, kembalikan respons dengan status BadRequest
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Panggil fungsi QueryUpdateByID untuk memperbarui pesanan berdasarkan ID
	updatedOrder = QueryUpdateByID(updatedOrder, uint(convertedOrderID))

	// Kembalikan pesanan yang telah diperbarui dalam respons JSON dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": fmt.Sprintf("Order with ID %v Has been successfully updated", orderID),
		"status":  fmt.Sprintf("%d", http.StatusOK),
	})
}
