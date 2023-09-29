package models

import "time"

// Order adalah struktur data untuk merepresentasikan entitas pesanan.
type Order struct {
	Order_id      uint      `gorm:"primaryKey" json:"orderID"`                                    // Order_id adalah ID unik dari pesanan, ditandai sebagai primary key di basis data.
	Customer_name string    `gorm:"notnull" json:"customerName"`                                  // Customer_name adalah nama pelanggan yang melakukan pesanan, tidak boleh kosong di basis data.
	Items         []Item    `json:"Items" gorm:"foreignKey:Order_id;constraint:OnDelete:CASCADE"` // Items adalah daftar item dalam pesanan yang terkait dengan pesanan ini.
	Ordered_at    time.Time `json:"orderedAt"`                                                    // Ordered_at adalah waktu kapan pesanan ini dibuat.
}
