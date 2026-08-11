
package config

import (
	"fmt"
	"log"

	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Employee{},
		&models.Attendance{},
		&models.Notification{},
		&models.Query{},
		&models.QueryReply{},
		&models.Leave{},
		&models.LeaveBalance{},
		&models.Payroll{},
		&models.Document{},
		&models.Company{},
		&models.SystemSettings{},
	); err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	DB = db
	log.Println("✅ PostgreSQL Connected & Migrated")
}