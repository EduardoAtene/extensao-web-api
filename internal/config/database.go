package config

import (
	"fmt"
	"log"

	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeRewards(db *gorm.DB) error {
	rewards := []models.RewardType{
		{
			Name:        "Bronze",
			Description: "Completou 10 descartes",
			Requirement: 10,
		},
		{
			Name:        "Viajante",
			Description: "Descarte em 5 pontos diferentes",
			Requirement: 5,
		},
		{
			Name:        "Prata",
			Description: "Completou 20 descartes",
			Requirement: 20,
		},
		{
			Name:        "Ouro",
			Description: "Completou 30 descartes",
			Requirement: 30,
		},
		{
			Name:        "Variado",
			Description: "Descarte em 20 pontos diferentes",
			Requirement: 20,
		},
		{
			Name:        "Diamante",
			Description: "Completou 50 descartes",
			Requirement: 50,
		},
	}

	for _, reward := range rewards {
		if err := db.FirstOrCreate(&reward, models.RewardType{Name: reward.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}

func SetupDB() *gorm.DB {
	// Substituindo o carregamento do .env por valores diretos
	dbUser := "admin"
	dbPassword := "admin"
	dbHost := "db" // Nome do serviço no docker-compose
	dbName := "extensao_database"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrações e inicialização de dados
	db.AutoMigrate(&models.User{}, &models.RewardType{}, &models.UserReward{}, &models.Discard{}, &models.Location{})

	if err := InitializeRewards(db); err != nil {
		log.Fatal(err)
	}

	return db
}

// func SetupDB() *gorm.DB {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db.AutoMigrate(&models.User{}, &models.RewardType{}, &models.UserReward{}, &models.Discard{}, &models.Location{})

// 	if err := InitializeRewards(db); err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }
