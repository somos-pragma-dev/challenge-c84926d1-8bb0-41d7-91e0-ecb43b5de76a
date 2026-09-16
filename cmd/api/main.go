package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/empresa/product-api/internal/domain"
	"github.com/empresa/product-api/internal/dto"
	"github.com/empresa/product-api/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No se encontró archivo .env, usando variables de entorno del sistema")
	}

	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	if err := runMigrations(db); err != nil {
		log.Fatalf("Error al ejecutar migraciones: %v", err)
	}

	router := initRouter(db)

	port := getServerPort()
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Servidor iniciando en el puerto %s", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func initDatabase() (*gorm.DB, error) {
	dbPath := getDatabasePath()
	dsn := fmt.Sprintf("%s?cache=shared", dbPath)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error al verificar conexión con la base de datos: %w", err)
	}

	log.Println("Conexión a la base de datos establecida correctamente")
	return db, nil
}

func runMigrations(db *gorm.DB) error {
	log.Println("Ejecutando migraciones de base de datos...")

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		return fmt.Errorf("error en migración de productos: %w", err)
	}

	log.Println("Migraciones ejecutadas correctamente")
	return nil
}

func initRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(errorMiddleware())

	productRepo := repository.NewProductRepository(db)
	productHandler := dto.NewProductHandler(productRepo)

	api := router.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.ListProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "API de gestión de productos funcionando correctamente",
		})
	})

	return router
}

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Printf("Error en solicitud: %v", e.Err)
			}
		}
	}
}

func getDatabasePath() string {
	path := os.Getenv("DATABASE_PATH")
	if path == "" {
		return "./data/products.db"
	}
	return path
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}