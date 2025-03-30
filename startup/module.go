package startup

import (
	"crud-server/api/task"
	"crud-server/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Module handles application dependency initialization
type Module struct {
	config   *config.Config
	db       *gorm.DB
	validate *validator.Validate
}

// NewModule creates a new module instance
func NewModule(config *config.Config) *Module {
	return &Module{
		config:   config,
		validate: validator.New(),
	}
}

// InitDB initializes the database connection
func (m *Module) InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.config.DBHost,
		m.config.DBPort,
		m.config.DBUser,
		m.config.DBPassword,
		m.config.DBName,
	)

	log.Println("Database connection string:", dsn)

	var err error
	m.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

// InitTaskAPI initializes the task API with all dependencies and registers routes
func (m *Module) InitTaskAPI(router *gin.RouterGroup) {
	service := task.NewService(m.db, m.validate)
	controller := task.NewController(service)
	controller.RegisterRoutes(router)
}

// GetDB returns the database connection
func (m *Module) GetDB() *gorm.DB {
	return m.db
}
