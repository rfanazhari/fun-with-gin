package config

import (
	"fun-with-gin/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

type DatabaseConfig struct {
	UserName          string
	Password          string
	Database          string
	Host              string
	Port              string
	TimeZone          string
	MaxConnectionIdle int
	MaxConnectionOpen int
	Schema            string
	Debug             bool
	TimeOut           time.Duration
	SslMode           string
}

func DatabasePGSQL() DatabaseConfig {
	schema := "public"
	if os.Getenv("DB_SCHEMA") != "" {
		schema = os.Getenv("DB_SCHEMA")
	}

	return DatabaseConfig{
		UserName:          os.Getenv("DB_USERNAME"),
		Password:          os.Getenv("DB_PASSWORD"),
		Database:          os.Getenv("DB_NAME"),
		Host:              os.Getenv("DB_HOST"),
		Port:              os.Getenv("DB_PORT"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		MaxConnectionIdle: utils.ConvertInt("DB_MAX_CON_IDLE"),
		MaxConnectionOpen: utils.ConvertInt("DB_MAX_CON_OPEN"),
		Schema:            schema,
		Debug:             utils.ConvertBool("DEBUG"),
		TimeOut:           time.Duration(utils.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}

func NewConnectionPGSQL() (*gorm.DB, error) {
	cfg := DatabasePGSQL()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	configs := []string{
		"host=" + cfg.Host,
		"user=" + cfg.UserName,
		"password=" + cfg.Password,
		"dbname=" + cfg.Database,
		"port=" + cfg.Port,
		"TimeZone=" + cfg.TimeZone,
		"sslmode=disable",
	}
	dsn := strings.Join(configs, " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err.Error())

	}
	return db, nil
}
