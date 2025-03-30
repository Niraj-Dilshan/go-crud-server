package startup

import (
	"crud-server/config"
	"log"

	"github.com/gin-gonic/gin"
)

// InitTestServer initializes a server for testing
func InitTestServer() (*gin.Engine, *Module, error) {
	// Load test configuration
	testConfig := config.LoadConfig(".test.env")

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize module with test configuration
	module := NewModule(testConfig)

	// Initialize database
	if err := module.InitDB(); err != nil {
		log.Printf("Failed to initialize test database: %v", err)
		return nil, nil, err
	}

	// Initialize database schema
	if err := InitIndexes(module.GetDB()); err != nil {
		log.Printf("Failed to initialize test database schema: %v", err)
		return nil, nil, err
	}

	// Create test server
	server := NewServer(testConfig)
	router := server.GetRouter()

	// Initialize API
	api := router.Group("/api")
	module.InitTaskAPI(api)

	return router, module, nil
}
