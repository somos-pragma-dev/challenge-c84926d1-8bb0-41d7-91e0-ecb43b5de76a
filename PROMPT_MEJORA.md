# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener un ZIP con el proyecto completo y arrancable.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Como saber que terminaste

```bash
el comando de build o arranque canonico del stack elegido
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
API REST con Go, Gin framework y GORM

### Reto
- Tema: Go Gin API
- Seniority: junior-l1
- Tipo: practical
- Título: Desarrollo de una API REST para gestión de productos
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Creación de la estructura básica de la API — objetivo: Definir y crear la estructura básica de la API, incluyendo la configuración del enrutamiento y la conexión a la base de datos. — entregable (NO resolver): API con rutas básicas y conexión a la base de datos.
- Fase 2: Implementación de validaciones y manejo de errores — objetivo: Implementar las validaciones necesarias para los productos y manejar adecuadamente los errores que puedan ocurrir durante las operaciones. — entregable (NO resolver): API con validaciones y manejo de errores implementados.
- Fase 3: Optimización y refactorización de la API — objetivo: Optimizar y refactorizar la API para mejorar su rendimiento y mantenibilidad. — entregable (NO resolver): API optimizada y refactorizada.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/empresa/product-api

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.18.0
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.7
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// === ARCHIVO: cmd/api/main.go ===
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

// === ARCHIVO: internal/domain/product.go ===
package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"size:255;not null;uniqueIndex" json:"name"`
	Price     float64        `gorm:"not null" json:"price"`
	Stock     int            `gorm:"not null;default:0" json:"stock"`
	Category  string         `gorm:"size:100;not null" json:"category"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("el nombre del producto es requerido")
	}

	if len(p.Name) < 3 {
		return errors.New("el nombre del producto debe tener al menos 3 caracteres")
	}

	if len(p.Name) > 255 {
		return errors.New("el nombre del producto no puede exceder 255 caracteres")
	}

	if p.Price < 0 {
		return errors.New("el precio del producto no puede ser negativo")
	}

	if p.Stock < 0 {
		return errors.New("el stock del producto no puede ser negativo")
	}

	if p.Category == "" {
		return errors.New("la categoría del producto es requerida")
	}

	return nil
}

func (p *Product) UpdateStock(quantity int) error {
	newStock := p.Stock + quantity
	if newStock < 0 {
		return errors.New("el stock resultante no puede ser negativo")
	}
	p.Stock = newStock
	return nil
}

func (p *Product) IsAvailable() bool {
	return p.Stock > 0
}

func (p *Product) CanDecreaseStock(quantity int) bool {
	return p.Stock >= quantity
}

type ProductRepository interface {
	Create(product *Product) error
	FindByID(id uint) (*Product, error)
	FindAll() ([]Product, error)
	FindByName(name string) (*Product, error)
	FindByCategory(category string) ([]Product, error)
	Update(product *Product) error
	Delete(id uint) error
	ExistsByName(name string, excludeID uint) (bool, error)
}

// === ARCHIVO: internal/repository/product_repository.go ===
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/empresa/product-api/internal/domain"
)

var (
	ErrProductNotFound       = errors.New("producto no encontrado")
	ErrProductAlreadyExists = errors.New("el producto ya existe")
	ErrDatabaseOperation    = errors.New("error en la operación de base de datos")
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return fmt.Errorf("validación fallida: %w", err)
	}

	exists, err := r.ExistsByName(product.Name, 0)
	if err != nil {
		return err
	}
	if exists {
		return ErrProductAlreadyExists
	}

	result := r.db.Create(product)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return nil
}

func (r *ProductRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product

	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return &product, nil
}

func (r *ProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return products, nil
}

func (r *ProductRepository) FindByName(name string) (*domain.Product, error) {
	var product domain.Product

	result := r.db.Where("name = ?", name).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return &product, nil
}

func (r *ProductRepository) FindByCategory(category string) ([]domain.Product, error) {
	var products []domain.Product

	result := r.db.Where("category = ?", category).Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return products, nil
}

func (r *ProductRepository) Update(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return fmt.Errorf("validación fallida: %w", err)
	}

	exists, err := r.ExistsByName(product.Name, product.ID)
	if err != nil {
		return err
	}
	if exists {
		return ErrProductAlreadyExists
	}

	result := r.db.Save(product)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (r *ProductRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (r *ProductRepository) ExistsByName(name string, excludeID uint) (bool, error) {
	var count int64

	query := r.db.Model(&domain.Product{}).Where("name = ?", name)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return count > 0, nil
}

// === ARCHIVO: internal/dto/product_dto.go ===
package dto

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/empresa/product-api/internal/domain"
	"github.com/empresa/product-api/internal/repository"
)

type ProductDTO struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required,min=3,max=255"`
	Price    float64 `json:"price" binding:"required,gte=0"`
	Stock    int     `json:"stock" binding:"gte=0"`
	Category string  `json:"category" binding:"required,min=1,max=100"`
}

type UpdateProductRequest struct {
	Name     string  `json:"name" binding:"omitempty,min=3,max=255"`
	Price    float64 `json:"price" binding:"omitempty,gte=0"`
	Stock    int     `json:"stock" binding:"omitempty,gte=0"`
	Category string  `json:"category" binding:"omitempty,min=1,max=100"`
}

type ProductListResponse struct {
	Success bool        `json:"success"`
	Data    []ProductDTO `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ProductResponse struct {
	Success bool       `json:"success"`
	Data    ProductDTO `json:"data,omitempty"`
	Message string     `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ProductHandler struct {
	repository *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: repo}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.repository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener la lista de productos",
		})
		return
	}

	productDTOs := make([]ProductDTO, 0, len(products))
	for _, p := range products {
		productDTOs = append(productDTOs, ProductDTO{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	c.JSON(http.StatusOK, ProductListResponse{
		Success: true,
		Data:    productDTOs,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	product, err := h.repository.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
		},
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   validationErrors,
		})
		return
	}

	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	if err := h.repository.Create(product); err != nil {
		if errors.Is(err, repository.ErrProductAlreadyExists) {
			c.JSON(http.StatusConflict, ErrorResponse{
				Success: false,
				Error:   "Ya existe un producto con ese nombre",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al crear el producto",
		})
		return
	}

	c.JSON(http.StatusCreated, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
		},
		Message: "Producto creado exitosamente",
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	existingProduct, err := h.repository.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener el producto",
		})
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   validationErrors,
		})
		return
	}

	if req.Name != "" {
		existingProduct.Name = req.Name
	}
	if req.Price != 0 {
		existingProduct.Price = req.Price
	}
	if req.Stock != 0 {
		existingProduct.Stock = req.Stock
	}
	if req.Category != "" {
		existingProduct.Category = req.Category
	}

	if err := h.repository.Update(existingProduct); err != nil {
		if errors.Is(err, repository.ErrProductAlreadyExists) {
			c.JSON(http.StatusConflict, ErrorResponse{
				Success: false,
				Error:   "Ya existe un producto con ese nombre",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al actualizar el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       existingProduct.ID,
			Name:     existingProduct.Name,
			Price:    existingProduct.Price,
			Stock:    existingProduct.Stock,
			Category: existingProduct.Category,
		},
		Message: "Producto actualizado exitosamente",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	if err := h.repository.Delete(uint(id)); err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al eliminar el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Producto eliminado exitosamente",
	})
}

func formatValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			switch e.Field() {
			case "Name":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El nombre es requerido")
				case "min":
					errorMessages = append(errorMessages, "El nombre debe tener al menos 3 caracteres")
				case "max":
					errorMessages = append(errorMessages, "El nombre no puede exceder 255 caracteres")
				}
			case "Price":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El precio es requerido")
				case "gte":
					errorMessages = append(errorMessages, "El precio no puede ser negativo")
				}
			case "Stock":
				switch e.Tag() {
				case "gte":
					errorMessages = append(errorMessages, "El stock no puede ser negativo")
				}
			case "Category":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "La categoría es requerida")
				case "min":
					errorMessages = append(errorMessages, "La categoría debe tener al menos 1 caracter")
				case "max":
					errorMessages = append(errorMessages, "La categoría no puede exceder 100 caracteres")
				}
			}
		}
		if len(errorMessages) > 0 {
			return errorMessages[0]
		}
	}
	return "Error de validación"
}


// === ARCHIVO: go.sum ===
github.com/bytedance/sonic v1.11.6 h1:FK+iAjmtomE2T8B9M4P5iVIXI4V6VJkT0Y9O4N9kIZJw=
github.com/bytedance/sonic v1.11.6/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/bytedance/sonic/loader v0.1.1 h1:zipNp5WEFF5/LRBw2v3pFNpJ0o9xV5VN/F1E6C6lF4E=
github.com/bytedance/sonic/loader v0.1.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/cloudwego/base64x v0.1.4 h1:1BDyGB1LDi1j0AKz1X3aQfd4qTyaxeV5K3f8H8F7kws=
github.com/cloudwego/base64x v0.1.4/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/cloudwego/iasm v0.2.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/cloudwego/iasm v0.2.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/gabriel-vasile/mimetype v1.4.3 h1:w4qZ3uR5+q5t4vZ5+AU4A4qU84ZOIokp9Y4Z5Y9O0F0=
github.com/gabriel-vasile/mimetype v1.4.3/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/gin-contrib/sse v0.1.0 h1:/KfnngQIYve1fAwzzYbMp+7yD2V8E1f3E7l5Y5U0Jxk=
github.com/gin-contrib/sse v0.1.0/go.mod h1:Byte彩虹1Lq1L1QeV3vZ8zK1S1r4V7L8a4V0E0F9Y0eGcM=
github.com/gin-gonic/gin v1.10.0 h1:qK0T1r7t1yJv7aB3y6J0T3LqL8V7Y8F5G9A0B5C0D1E=
github.com/gin-gonic/gin v1.10.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/locales v0.14.1 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/locales v0.14.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/universal-translator v0.18.1 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/universal-translator v0.18.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/validator/v10 v10.18.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/go-playground/validator/v10 v10.18.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/goccy/go-json v0.10.2 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/goccy/go-json v0.10.2/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/jinzhu/inflection v1.0.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/jinzhu/inflection v1.0.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/jinzhu/now v1.1.5 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/jinzhu/now v1.1.5/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/joho/godotenv v1.5.1 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/joho/godotenv v1.5.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/json-iterator/go v1.1.12 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/json-iterator/go v1.1.12/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/klauspost/cpuid/v2 v2.2.7 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/klauspost/cpuid/v2 v2.2.7/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/leodido/go-urn v1.4.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/leodido/go-urn v1.4.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/mattn/go-isatty v0.0.20 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/mattn/go-isatty v0.0.20/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/mattn/go-sqlite3 v1.14.24 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/mattn/go-sqlite3 v1.14.24/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/modern-go/reflect2 v1.0.2 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/modern-go/reflect2 v1.0.2/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/pelletier/go-toml/v2 v2.2.2 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/pelletier/go-toml/v2 v2.2.2/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/twitchyliquid64/golang-asm v0.15.1 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/twitchyliquid64/golang-asm v0.15.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/ugorji/go/codec v1.2.12 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
github.com/ugorji/go/codec v1.2.12/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/arch v0.8.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/arch v0.8.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/crypto v0.23.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/crypto v0.23.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/net v0.25.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/net v0.25.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/sys v0.20.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/sys v0.20.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/text v0.15.0 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
golang.org/x/text v0.15.0/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
google.golang.org/protobuf v1.34.1 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
google.golang.org/protobuf v1.34.1/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gorm.io/driver/sqlite v1.5.5 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
gorm.io/driver/sqlite v1.5.5/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
gorm.io/gorm v1.25.7 h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
gorm.io/gorm v1.25.7/go.mod h1:0wK3C4k7ZJYh4CNbxsJ1N5q1z8E0Z0H3LqLEv7Y0eGcM=
// === ARCHIVO: config/config.go ===
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


// === ARCHIVO: internal/handlers/product_handler.go ===
package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	dto "github.com/empresa/product-api/internal/dto"
	repository "github.com/empresa/product-api/internal/repository"
	"github.com/empresa/product-api/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		service: service.NewProductService(repo),
	}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al obtener productos",
			Message: err.Error(),
		})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, dto.ProductListResponse{
			Products: []dto.ProductResponse{},
			Total:    0,
		})
		return
	}

	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	c.JSON(http.StatusOK, dto.ProductListResponse{
		Products: productResponses,
		Total:    len(productResponses),
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		if err.Error() == "producto no encontrado" {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al obtener producto",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Datos inválidos",
			Message: formatValidationErrors(err),
		})
		return
	}

	product, err := h.service.CreateProduct(req)
	if err != nil {
		switch err.Error() {
		case "el nombre del producto ya existe":
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error:   "Conflicto de datos",
				Message: err.Error(),
			})
		case "el precio no puede ser negativo":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		case "el nombre no puede estar vacío":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error al crear producto",
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Datos inválidos",
			Message: formatValidationErrors(err),
		})
		return
	}

	product, err := h.service.UpdateProduct(uint(id), req)
	if err != nil {
		switch err.Error() {
		case "producto no encontrado":
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
		case "el nombre del producto ya existe":
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error:   "Conflicto de datos",
				Message: err.Error(),
			})
		case "el precio no puede ser negativo":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error al actualizar producto",
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	err = h.service.DeleteProduct(uint(id))
	if err != nil {
		if err.Error() == "producto no encontrado" {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al eliminar producto",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Producto eliminado exitosamente",
	})
}

func formatValidationErrors(err error) string {
	if err == nil {
		return "Error de validación desconocido"
	}
	return err.Error()
}
// === ARCHIVO: internal/service/product_service.go ===
package service

import (
	"errors"
	repository "github.com/empresa/product-api/internal/repository"
	dto "github.com/empresa/product-api/internal/dto"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductDTO, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("error al consultar la base de datos")
	}

	var dtos []dto.ProductDTO
	for _, p := range products {
		dtos = append(dtos, dto.ProductDTO{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	return dtos, nil
}

func (s *ProductService) GetProductByID(id uint) (dto.ProductDTO, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return dto.ProductDTO{}, errors.New("producto no encontrado")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) CreateProduct(req dto.CreateProductRequest) (dto.ProductDTO, error) {
	if req.Name == "" {
		return dto.ProductDTO{}, errors.New("el nombre no puede estar vacío")
	}

	if req.Price < 0 {
		return dto.ProductDTO{}, errors.New("el precio no puede ser negativo")
	}

	exists, err := s.repo.ExistsByName(req.Name, 0)
	if err != nil {
		return dto.ProductDTO{}, errors.New("error al verificar el nombre del producto")
	}
	if exists {
		return dto.ProductDTO{}, errors.New("el nombre del producto ya existe")
	}

	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	if err := s.repo.Create(product); err != nil {
		return dto.ProductDTO{}, errors.New("error al guardar el producto")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) UpdateProduct(id uint, req dto.UpdateProductRequest) (dto.ProductDTO, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return dto.ProductDTO{}, errors.New("producto no encontrado")
	}

	if req.Name != nil && *req.Name != "" {
		exists, err := s.repo.ExistsByName(*req.Name, id)
		if err != nil {
			return dto.ProductDTO{}, errors.New("error al verificar el nombre del producto")
		}
		if exists {
			return dto.ProductDTO{}, errors.New("el nombre del producto ya existe")
		}
		product.Name = *req.Name
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return dto.ProductDTO{}, errors.New("el precio no puede ser negativo")
		}
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	if req.Category != nil {
		product.Category = *req.Category
	}

	if err := s.repo.Update(product); err != nil {
		return dto.ProductDTO{}, errors.New("error al actualizar el producto")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("producto no encontrado")
	}

	if err := s.repo.Delete(product.ID); err != nil {
		return errors.New("error al eliminar el producto")
	}

	return nil
}
// === ARCHIVO: internal/handlers/product_handler_test.go ===
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	dto "github.com/empresa/product-api/internal/dto"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type MockProductService struct {
	products []dto.ProductDTO
	err      error
}

func (m *MockProductService) GetAllProducts() ([]dto.ProductDTO, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.products, nil
}

func (m *MockProductService) GetProductByID(id uint) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			return p, nil
		}
	}
	return dto.ProductDTO{}, errors.New("producto no encontrado")
}

func (m *MockProductService) CreateProduct(req dto.CreateProductRequest) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	return dto.ProductDTO{
		ID:       1,
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}, nil
}

func (m *MockProductService) UpdateProduct(id uint, req dto.UpdateProductRequest) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			result := p
			if req.Name != nil {
				result.Name = *req.Name
			}
			if req.Price != nil {
				result.Price = *req.Price
			}
			if req.Stock != nil {
				result.Stock = *req.Stock
			}
			if req.Category != nil {
				result.Category = *req.Category
			}
			return result, nil
		}
	}
	return dto.ProductDTO{}, errors.New("producto no encontrado")
}

func (m *MockProductService) DeleteProduct(id uint) error {
	if m.err != nil {
		return m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

func setupTestRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	return router
}

func TestListProducts_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
			{ID: 2, Name: "Producto 2", Price: 200.00, Stock: 5, Category: "Ropa"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products", handler.ListProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Total != 2 {
		t.Errorf("Expected 2 products, got %d", response.Total)
	}

	if len(response.Products) != 2 {
		t.Errorf("Expected 2 products in array, got %d", len(response.Products))
	}
}

func TestListProducts_Empty(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products", handler.ListProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Total != 0 {
		t.Errorf("Expected 0 products, got %d", response.Total)
	}
}

func TestGetProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products/:id", handler.GetProduct)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Name != "Producto 1" {
		t.Errorf("Expected product name 'Producto 1', got '%s'", response.Name)
	}
}

func TestGetProduct_InvalidID(t *testing.T) {
	router := setupTestRouter()

	handler := &ProductHandler{service: &ProductServiceMock{}}
	router.GET("/products/:id", handler.GetProduct)

	req, _ := http.NewRequest("GET", "/products/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response dto.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Error != "ID inválido" {
		t.Errorf("Expected error message 'ID inválido', got '%s'", response.Error)
	}
}

func TestCreateProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.POST("/products", handler.CreateProduct)

	body := dto.CreateProductRequest{
		Name:     "Nuevo Producto",
		Price:    150.00,
		Stock:    20,
		Category: "Hogar",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response dto.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Name != "Nuevo Producto" {
		t.Errorf("Expected product name 'Nuevo Producto', got '%s'", response.Name)
	}
}

func TestCreateProduct_InvalidJSON(t *testing.T) {
	router := setupTestRouter()

	handler := &ProductHandler{service: &ProductServiceMock{}}
	router.POST("/products", handler.CreateProduct)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte("{invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestCreateProduct_DuplicateName(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		err: errors.New("el nombre del producto ya existe"),
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.POST("/products", handler.CreateProduct)

	body := dto.CreateProductRequest{
		Name:     "Producto Existente",
		Price:    100.00,
		Stock:    10,
		Category: "Electrónica",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409, got %d", w.Code)
	}

	var response dto.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Error != "Conflicto de datos" {
		t.Errorf("Expected error 'Conflicto de datos', got '%s'", response.Error)
	}
}

func TestDeleteProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.DELETE("/products/:id", handler.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteProduct_NotFound(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		err: errors.New("producto no encontrado"),
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.DELETE("/products/:id", handler.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// === ARCHIVO: pkg/validators/product_validator.go ===
package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/empresa/product-api/internal/domain"
	"github.com/empresa/product-api/internal/repository"
)

type ProductValidator struct {
	repo *repository.ProductRepository
}

func NewProductValidator(repo *repository.ProductRepository) *ProductValidator {
	return &ProductValidator{repo: repo}
}

func (v *ProductValidator) ValidateProduct(product *domain.Product) error {
	validate := validator.New()

	err := validate.Struct(product)
	if err != nil {
		return v.formatValidationErrors(err)
	}

	if err := v.validateUniqueName(product); err != nil {
		return err
	}

	return nil
}

func (v *ProductValidator) validateUniqueName(product *domain.Product) error {
	exists, err := v.repo.ExistsByName(product.Name, product.ID)
	if err != nil {
		return fmt.Errorf("error al verificar nombre del producto: %w", err)
	}

	if exists {
		return fmt.Errorf("el nombre '%s' ya existe en la base de datos", product.Name)
	}

	return nil
}

func (v *ProductValidator) ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("el nombre del producto no puede estar vacío")
	}

	if len(name) < 3 {
		return fmt.Errorf("el nombre del producto debe tener al menos 3 caracteres")
	}

	if len(name) > 100 {
		return fmt.Errorf("el nombre del producto no puede exceder 100 caracteres")
	}

	return nil
}

func (v *ProductValidator) ValidatePrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("el precio no puede ser negativo")
	}

	if price > 999999.99 {
		return fmt.Errorf("el precio excede el límite permitido")
	}

	return nil
}

func (v *ProductValidator) ValidateStock(stock int) error {
	if stock < 0 {
		return fmt.Errorf("el stock no puede ser negativo")
	}

	if stock > 999999 {
		return fmt.Errorf("el stock excede el límite permitido")
	}

	return nil
}

func (v *ProductValidator) ValidateCategory(category string) error {
	validCategories := []string{"electronics", "clothing", "food", "books", "home", "sports", "toys", "other"}

	categoryLower := strings.ToLower(category)

	for _, valid := range validCategories {
		if categoryLower == valid {
			return nil
		}
	}

	return fmt.Errorf("categoría '%s' no válida. Las categorías válidas son: %s", category, strings.Join(validCategories, ", "))
}

func (v *ProductValidator) formatValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string

		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case "Name":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El nombre del producto es obligatorio")
				case "min":
					errorMessages = append(errorMessages, "El nombre debe tener al menos 3 caracteres")
				case "max":
					errorMessages = append(errorMessages, "El nombre no puede exceder 100 caracteres")
				}
			case "Price":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El precio es obligatorio")
				case "gte":
					errorMessages = append(errorMessages, "El precio debe ser mayor o igual a 0")
				}
			case "Stock":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El stock es obligatorio")
				case "gte":
					errorMessages = append(errorMessages, "El stock debe ser mayor o igual a 0")
				}
			case "Category":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "La categoría es obligatoria")
				}
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("Error en el campo %s", fieldErr.Field()))
			}
		}

		return strings.Join(errorMessages, "; ")
	}

	return err.Error()
}

func (v *ProductValidator) ValidateProductForUpdate(existingProduct *domain.Product, updatedProduct *domain.Product) error {
	validate := validator.New()

	err := validate.Struct(updatedProduct)
	if err != nil {
		return v.formatValidationErrors(err)
	}

	if updatedProduct.Name != existingProduct.Name {
		if err := v.validateUniqueName(updatedProduct); err != nil {
			return err
		}
	}

	return nil
}

// === ARCHIVO: pkg/errors/error_handler.go ===
package errors

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ValidationErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Details map[string]string `json:"details,omitempty"`
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	errMsg := err.Error()
	log.Printf("[ERROR] %s", errMsg)

	switch {
	case strings.Contains(errMsg, "not found"):
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "NOT_FOUND",
			Message: errMsg,
			Code:    http.StatusNotFound,
		})
	case strings.Contains(errMsg, "ya existe"):
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "CONFLICT",
			Message: errMsg,
			Code:    http.StatusConflict,
		})
	case strings.Contains(errMsg, "no puede estar vacío"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no puede ser negativo"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "al menos"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no puede exceder"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no válida"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "error al verificar"):
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Error al procesar la solicitud",
			Code:    http.StatusInternalServerError,
		})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Ha ocurrido un error interno. Por favor, intente más tarde.",
			Code:    http.StatusInternalServerError,
		})
	}
}

func HandleValidationErrors(c *gin.Context, err error) {
	errMsg := err.Error()
	log.Printf("[VALIDATION ERROR] %s", errMsg)

	details := parseValidationDetails(errMsg)

	response := ValidationErrorResponse{
		Error:   "VALIDATION_ERROR",
		Message: "Error de validación en los datos proporcionados",
		Code:    http.StatusBadRequest,
	}

	if len(details) > 0 {
		response.Details = details
	}

	c.JSON(http.StatusBadRequest, response)
}

func parseValidationDetails(errMsg string) map[string]string {
	details := make(map[string]string)

	errors := strings.Split(errMsg, "; ")
	for _, e := range errors {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}

		if strings.Contains(e, "nombre") {
			details["name"] = e
		} else if strings.Contains(e, "precio") {
			details["price"] = e
		} else if strings.Contains(e, "stock") {
			details["stock"] = e
		} else if strings.Contains(e, "categoría") || strings.Contains(e, "categoria") {
			details["category"] = e
		} else {
			details["general"] = e
		}
	}

	return details
}

func HandleNotFound(c *gin.Context, resource string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   "NOT_FOUND",
		Message: fmt.Sprintf("%s no encontrado", resource),
		Code:    http.StatusNotFound,
	})
}

func HandleBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   "BAD_REQUEST",
		Message: message,
		Code:    http.StatusBadRequest,
	})
}

func HandleUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error:   "UNAUTHORIZED",
		Message: "No autorizado para realizar esta acción",
		Code:    http.StatusUnauthorized,
	})
}

func HandleMethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, ErrorResponse{
		Error:   "METHOD_NOT_ALLOWED",
		Message: "Método no permitido para este recurso",
		Code:    http.StatusMethodNotAllowed,
	})
}

// GlobalErrorHandler returns a Gin middleware that handles panics
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %v", err)
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "INTERNAL_ERROR",
					Message: "Se ha producido un error inesperado",
					Code:    http.StatusInternalServerError,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// === ARCHIVO: migrations/0001_create_products_table.sql ===
-- Migration: 0001_create_products_table.sql
-- Description: Creates the products table for the e-commerce product management API
-- Created at: 2024

-- Enable foreign keys support (SQLite specific)
PRAGMA foreign_keys = ON;

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(50) NOT NULL,
    sku VARCHAR(50) UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index on name for unique validation queries
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- Create index on category for filtering queries
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);

-- Create index on is_active for active product queries
CREATE INDEX IF NOT EXISTS idx_products_is_active ON products(is_active);

-- Create index on created_at for ordering queries
CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);

-- Insert sample data for testing
INSERT INTO products (name, description, price, stock, category, sku, is_active) VALUES
    ('Laptop Pro 15', 'Laptop de alta gama con procesador i7 y 16GB RAM', 1299.99, 25, 'electronics', 'LAP-PRO-15', 1),
    ('Camiseta Algodón', 'Camiseta de algodón orgánica color negro', 19.99, 150, 'clothing', 'CAM-alg-001', 1),
    ('Café Premium 1kg', 'Café premium tostado medio de Colombia', 24.50, 80, 'food', 'CAF-pre-1kg', 1),
    ('Novela El Principito', 'Edición especial del clásico de Saint-Exupéry', 12.99, 200, 'books', 'NOV-PRI-001', 1),
    ('Lámpara LED Escritorio', 'Lámpara LED regulable para escritorio', 35.00, 60, 'home', 'LAM-LED-001', 1);

-- Create trigger to update updated_at timestamp on row update
CREATE TRIGGER IF NOT EXISTS update_products_timestamp
AFTER UPDATE ON products
BEGIN
    UPDATE products SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Create view for active products
CREATE VIEW IF NOT EXISTS active_products AS
SELECT 
    id,
    name,
    description,
    price,
    stock,
    category,
    sku,
    created_at,
    updated_at
FROM products
WHERE is_active = 1;

-- Create view for low stock products (stock < 10)
CREATE VIEW IF NOT EXISTS low_stock_products AS
SELECT 
    id,
    name,
    price,
    stock,
    category,
    sku
FROM products
WHERE is_active = 1 AND stock < 10
ORDER BY stock ASC;

```
