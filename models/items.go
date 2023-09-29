package models

// Item adalah struktur data yang digunakan untuk merepresentasikan suatu item dalam sistem.
type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"lineItemID"` // ID unik untuk setiap item, ditandai sebagai kunci primer dalam database GORM dan diidentifikasi sebagai "lineItemID" dalam respons JSON.
	Item_code   string `json:"itemCode"`                     // Kode item, digunakan untuk mengidentifikasi item.
	Description string `json:"description"`                  // Deskripsi item, berisi informasi tentang item tersebut.
	Quantity    int    `json:"quantity"`                     // Kuantitas item, menunjukkan berapa banyak item yang ada dalam pesanan.
	Order_id    uint   `json:"orderID"`                      // ID pesanan yang terkait dengan item ini.
}
