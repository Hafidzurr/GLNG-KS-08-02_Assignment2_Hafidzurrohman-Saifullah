package main

import (
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/database"
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/routers"
)

func main() {
	// Mulai database. Ini menginisialisasi koneksi database dan tabel yang diperlukan.
	database.StartDB()

	// Tentukan PORT yang akan digunakan untuk server web.
	var PORT = ":8080"

	// Mulai server web menggunakan router yang telah dikonfigurasi.
	routers.StartServer().Run(PORT)
}
