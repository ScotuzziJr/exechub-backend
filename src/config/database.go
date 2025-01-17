package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ScotuzziJr/lead-service-exechub/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Carregar variáveis de ambiente
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Conectar ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// Realizar migração para garantir que a tabela 'leads' será criada ou atualizada
	err = db.AutoMigrate(&models.Lead{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}

	return db, nil
}
