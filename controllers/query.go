package controllers

import (
	"fmt"

	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/database"
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/models"
)

// QueryPatchByID adalah fungsi untuk memperbarui pesanan berdasarkan ID yang diberikan.
func QueryPatchByID(updatedOrder models.Order, id uint) (models.Order, error) {
	db := database.GetDB()

	// Cek apakah pesanan dengan ID yang diberikan ada dalam database
	var existingOrder models.Order
	err := db.Preload("Items").First(&existingOrder, id).Error
	if err != nil {
		return models.Order{}, err // Mengembalikan error jika pesanan tidak ditemukan
	}

	// Perbarui data pesanan yang sesuai
	existingOrder.Customer_name = updatedOrder.Customer_name
	existingOrder.Ordered_at = updatedOrder.Ordered_at

	// Perbarui data item (Item) jika diberikan dalam JSON request
	for i := range updatedOrder.Items {
		err = db.Model(&existingOrder.Items[i]).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			return models.Order{}, err // Mengembalikan error jika terjadi kesalahan saat memperbarui item
		}
	}

	// Simpan perubahan dalam database
	err = db.Save(&existingOrder).Error
	if err != nil {
		return models.Order{}, err // Mengembalikan error jika terjadi kesalahan saat menyimpan perubahan pesanan
	}

	return existingOrder, nil
}

// QueryCreate adalah fungsi untuk membuat pesanan baru dalam database.
func QueryCreate(orderInput models.Order) models.Order {
	db := database.GetDB()

	newOrder := orderInput

	dberr := db.Debug().Create(&newOrder).Error

	if dberr != nil {
		panic(dberr)
	}

	return newOrder
}

// QueryGetByID adalah fungsi untuk mendapatkan pesanan berdasarkan ID yang diberikan.
func QueryGetByID(id uint) (models.Order, error) {
	db := database.GetDB()

	var order models.Order

	// Mencari pesanan berdasarkan ID
	err := db.Preload("Items").First(&order, id).Error
	if err != nil {
		return models.Order{}, err // Mengembalikan error jika pesanan tidak ditemukan
	}

	return order, nil
}

// QueryGetAll adalah fungsi untuk mendapatkan semua pesanan dari database.
func QueryGetAll() []models.Order {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		panic(dberr)
	}

	return orders
}

// QueryDeleteByID adalah fungsi untuk menghapus pesanan berdasarkan ID yang diberikan.
func QueryDeleteByID(id uint) {
	db := database.GetDB()

	dberr := db.Where("Order_id=?", id).Delete(&models.Item{}).Error

	if dberr != nil {
		panic(dberr)
	}

	dberr = db.Delete(&models.Order{}, id).Error

	if dberr != nil {
		panic(dberr)
	}

	fmt.Println("Data Deleted")
}

// QueryUpdateByID adalah fungsi untuk memperbarui pesanan berdasarkan ID yang diberikan.
func QueryUpdateByID(orderInput models.Order, id uint) models.Order {
	db := database.GetDB()

	updatedOrder := orderInput
	var err error

	for i := range updatedOrder.Items {
		err = db.Model(&updatedOrder.Items[i]).Where("Item_id=?", updatedOrder.Items[i].Item_id).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			panic(err)
		}
	}

	var updatedOnlyOrder models.Order
	updatedOnlyOrder.Customer_name = updatedOrder.Customer_name
	updatedOnlyOrder.Ordered_at = updatedOrder.Ordered_at

	err = db.Model(&updatedOnlyOrder).Where("Order_id=?", id).Updates(&updatedOnlyOrder).Error

	if err != nil {
		panic(err)
	}

	err = db.Preload("Items").Where("Order_id=?", id).Find(&updatedOrder).Error

	if err != nil {
		panic(err)
	}

	return updatedOrder
}
