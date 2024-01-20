package db

import (
	"fmt"

	"dateApp/config"

	// product "dateApp/pkg/product/modules/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	DBSchema string
	Port     int
	SSLMode  string
	TimeZone string
}

// NewDatabaseConnection Create new database connection based on given config
func NewDatabaseConnection(c PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v search_path=%v", c.Host, c.Username, c.Password, c.DBName, c.Port, c.SSLMode, c.DBSchema)

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}

func DatabaseConnection() *gorm.DB {
	postgresConfig := PostgresConfig{
		Host:     config.GetConfig().DBHost,
		Username: config.GetConfig().DBUsername,
		Password: config.GetConfig().DBPassword,
		DBName:   config.GetConfig().DBName,
		DBSchema: config.GetConfig().DBSchema,
		Port:     config.GetConfig().DBPort,
		SSLMode:  "disable",
	}

	db, err := NewDatabaseConnection(postgresConfig)
	if err != nil {
		panic(err)
	}

	return db.Debug()
}
