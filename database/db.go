package database

import (
	"fmt"
	"log"

	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment2_Hafidzurrohman-Saifullah.git/models" // Mengimpor model data dari paket models.

	_ "github.com/lib/pq"     // Mengimpor driver PostgreSQL tanpa secara eksplisit menggunakan kode ini.
	"gorm.io/driver/postgres" // Mengimpor driver GORM untuk PostgreSQL.
	"gorm.io/gorm"            // Mengimpor GORM, ORM (Object-Relational Mapping) untuk Go.
)

// Variabel global untuk menyimpan instance database dan error yang terkait dengannya.
var (
	db  *gorm.DB
	err error
)

// Konstanta yang mendefinisikan informasi koneksi ke database PostgreSQL.
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Hafidzurr1"
	dbname   = "postgres"
)

// Fungsi StartDB digunakan untuk menginisialisasi dan memulai koneksi ke database.
func StartDB() {
	// Mengonstruksi string koneksi berdasarkan konstanta yang didefinisikan sebelumnya.
	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Membuka koneksi ke database PostgreSQL menggunakan GORM dan konfigurasi default.
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err) // Jika terjadi kesalahan, keluarkan pesan kesalahan dan hentikan program.
	}

	// Mengaktifkan mode debugging dan otomatis melakukan migrasi tabel Order dan Item jika belum ada.
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
	fmt.Println("Connected to database") // Pesan konfirmasi bahwa koneksi berhasil dibuat.
}

// Fungsi GetDB digunakan untuk mendapatkan instance database yang sudah dibuat.
func GetDB() *gorm.DB {
	return db
}
