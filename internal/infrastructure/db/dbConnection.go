package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	"github.com/ashkarax/ciao-socialmedia/internal/domain"
	hashpassword "github.com/ashkarax/ciao-socialmedia/pkg/hash_password"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config.DataBase) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBPort)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set the timezone to UTC
		},
	})
	if dberr != nil {
		return DB, nil
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain.Admin{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Users{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.OtpInfo{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Post{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.PostMedia{}); err != nil {
		return DB, err
	}

	CheckAndCreateAdmin(DB)
	return DB, nil

}
func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "ciao"
		Email    = "ciao@gmail.com"
		Password = "ciaociao"
	)
	HashedPassword := hashpassword.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
