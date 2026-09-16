package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	App      AppConfig
}

type DatabaseConfig struct {
	Path        string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
	LogMode     bool
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConfig struct {
	Name        string
	Environment string
	Debug       bool
}

var appConfig *Config

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No se encontró archivo .env, usando valores por defecto")
	}

	appConfig = &Config{
		Database: DatabaseConfig{
			Path:        getDatabasePath(),
			MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
			MaxLifetime:  getEnvAsDuration("DB_MAX_LIFETIME", time.Hour),
			LogMode:      getEnvAsBool("DB_LOG_MODE", false),
		},
		Server: ServerConfig{
			Port:         getServerPort(),
			ReadTimeout:  getEnvAsDuration("SERVER_READ_TIMEOUT", 30*time.Second),
			WriteTimeout: getEnvAsDuration("SERVER_WRITE_TIMEOUT", 30*time.Second),
		},
		App: AppConfig{
			Name:        getEnv("APP_NAME", "Product API"),
			Environment: getEnv("APP_ENV", "development"),
			Debug:       getEnvAsBool("APP_DEBUG", true),
		},
	}

	return appConfig, nil
}

func Get() *Config {
	if appConfig == nil {
		var err error
		appConfig, err = Load()
		if err != nil {
			panic("Error al cargar la configuración: " + err.Error())
		}
	}
	return appConfig
}

func getDatabasePath() string {
	dbPath := getEnv("DB_PATH", "./data/products.db")

	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Advertencia: No se pudo crear el directorio de la base de datos: %v\n", err)
	}

	return dbPath
}

func getServerPort() string {
	return getEnv("SERVER_PORT", "8080")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Valor inválido para %s: %s, usando valor por defecto %d\n", key, value, defaultValue)
		return defaultValue
	}
	return intValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Valor inválido para %s: %s, usando valor por defecto %v\n", key, value, defaultValue)
		return defaultValue
	}
	return boolValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		fmt.Printf("Valor inválido para %s: %s, usando valor por defecto %v\n", key, value, defaultValue)
		return defaultValue
	}
	return duration
}

func InitDatabase(cfg *DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.Path

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener la conexión de la base de datos: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	if cfg.LogMode {
		db = db.Session(&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	}

	return db, nil
}

func GetDatabaseConfig() DatabaseConfig {
	return Get().Database
}

func GetServerConfig() ServerConfig {
	return Get().Server
}

func GetAppConfig() AppConfig {
	return Get().App
}