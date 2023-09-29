package routers

import (
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/controllers"

	"github.com/gin-gonic/gin"
)

// StartServer adalah fungsi yang digunakan untuk membuat dan mengkonfigurasi router Gin.
func StartServer() *gin.Engine {
	// Membuat instance router Gin dengan pengaturan default.
	router := gin.Default()

	// Menambahkan rute HTTP untuk mengambil semua pesanan (orders).
	router.GET("/orders", controllers.GetAllOrders)

	// Menambahkan rute HTTP untuk mengambil pesanan berdasarkan ID.
	router.GET("/orders/:orderID", controllers.GetOrderByID)

	// Menambahkan rute HTTP untuk membuat pesanan baru.
	router.POST("/orders", controllers.CreateOrders)

	// Menambahkan rute HTTP untuk menghapus pesanan berdasarkan ID.
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	// Menambahkan rute HTTP untuk memperbarui pesanan berdasarkan ID.
	router.PUT("/orders/:orderID", controllers.UpdateOrderByID)

	// Menambahkan rute HTTP untuk memperbarui sebagian pesanan berdasarkan ID.
	router.PATCH("/orders/:orderID", controllers.PatchOrderByID)

	// Mengembalikan instance router yang sudah dikonfigurasi.
	return router
}
