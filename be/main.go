package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	customerFirstNames = []string{"Wolf", "Rogue", "Reaper", "Shadow", "Viper", "Hunter", "Stryker", "Ghost", "Warrior", "Ace", "Raider", "Blade", "Thunder", "Steele", "Falcon", "Phoenix", "Grizzly", "Hawk", "Ranger", "Savage", "Scorpion", "Talon", "Jaguar", "Patriot", "Raptor", "Spartan", "Trojan", "Valor", "Echo", "Cobra", "Gunner", "Rex", "Steel", "Titan", "Vandal", "Rogue", "Bolt", "Bullet", "Bandit", "Commando", "Delta", "Devil", "Fox", "Gator", "Gunner", "Marshall", "Maverick", "Neptune", "Raven", "Shark", "Vortex", "Warlock", "Wildcat", "Zeus"}
	customerLastNames  = []string{"Black", "Steel", "Fury", "Hunter", "Storm", "Wolf", "Blade", "Ranger", "Shadow", "Reaper", "Bolt", "Stryker", "Phoenix", "Falcon", "Viper", "Thunder", "Crusher", "Commander", "Ghost", "Steele", "Hawk", "Cobra", "Griffin", "Rogue", "Valor", "Jaguar", "Spartan", "Raider", "Trojan", "Talon", "Raptor", "Vandal", "Barracuda", "Rex", "Tiger", "Vortex", "Warlock", "Maverick", "Gunner", "Fox", "Scorpion", "Neptune", "Vulture", "Raven", "Zeus", "Gator", "Delta", "Wildcat", "Marshall", "Mako", "Maverick"}
	nationNames        = []string{"American", "British", "Canadian", "Australian", "German", "French", "Italian", "Japanese", "Russian", "Chinese"}
)
var db *gorm.DB
var err error

// Customer model
type Customer struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Note      string `json:"note"`
	NationID  uint   `json:"nationId"`
	Nation    Nation `json:"nation"`
}

type CustomerDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Nationality model
type Nation struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name"`
}

func main() {
	// Initialize Gin
	r := gin.Default()
	// var err
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	// Initialize GORM
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Customer{}, &Nation{})

	// Seed data
	seedData(db)

	// CRUD operations for customers
	r.GET("/customers", getCustomers)
	r.POST("/customers", createCustomer)
	r.GET("/customers/:id", getCustomer)
	r.PUT("/customers/:id", updateCustomer)
	r.DELETE("/customers/:id", deleteCustomer)

	// CRUD operations for nation
	r.GET("/nations", getNations)
	r.POST("/nations", createNation)
	r.GET("/nations/:id", getNation)
	r.PUT("/nations/:id", updateNation)
	r.DELETE("/nations/:id", deleteNation)

	// SSE handler to invalidate cache
	r.GET("/sse", sseHandler)

	// Start server
	r.Run(":3000")
	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

// CRUD operations for customers
func getCustomers(c *gin.Context) {
	var customers []CustomerDTO
	db.Model(&Customer{}).Preload("Nations").Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func getCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := db.Preload("Nation").First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	customerDTO := customer.toDTO()
	time.Sleep(time.Second * 1)
	c.JSON(http.StatusOK, customerDTO)
}

func updateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer Customer
	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer Customer
	fmt.Println("id delete", id)
	if id == "7" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Some purposeful error..."})
		return
	}
	db.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// CRUD operations for nation
func getNations(c *gin.Context) {
	var nations []Nation
	db.Order("name").Find(&nations)
	time.Sleep(time.Second * 5)
	c.JSON(http.StatusOK, nations)
}

func createNation(c *gin.Context) {
	var nation Nation
	if err := c.ShouldBindJSON(&nation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&nation)
	c.JSON(http.StatusOK, nation)
}

func getNation(c *gin.Context) {
	var nation Nation
	id := c.Param("id")
	if err := db.First(&nation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nation not found"})
		return
	}
	c.JSON(http.StatusOK, nation)
}

func updateNation(c *gin.Context) {
	id := c.Param("id")
	var nation Nation
	if err := db.First(&nation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nation not found"})
		return
	}
	if err := c.ShouldBindJSON(&nation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&nation)
	c.JSON(http.StatusOK, nation)
}

func deleteNation(c *gin.Context) {
	id := c.Param("id")
	var nation Nation
	db.First(&nation, id)
	db.Delete(&nation)
	c.JSON(http.StatusOK, gin.H{"message": "Nation deleted successfully"})
}

// SSE handler to invalidate cache
func sseHandler(c *gin.Context) {
	// Server-Sent Events logic to notify frontend of changes in the "Nation" table
}

// toDTO converts Customer into the desired DTO format
func (c *Customer) toDTO() map[string]interface{} {
	// Construct the name from first name and last name
	name := fmt.Sprintf("%s %s", c.FirstName, c.LastName)

	// Construct the fields array with desired properties
	fields := []map[string]interface{}{
		{"name": "nationId", "type": "select", "query": "nations", "value": c.NationID},
		// {"name": "note", "type": "text", "value": c.Note},
		{"name": "firstName", "type": "text", "value": c.FirstName},
		{"name": "lastName", "type": "text", "value": c.LastName},
		// Add more fields as needed
	}

	// Return the DTO object
	return map[string]interface{}{
		"id":     c.ID,
		"name":   name,
		"fields": fields,
	}
}

func seedData(db *gorm.DB) {
	// Delete old data

	db.Exec("DELETE FROM customers")
	db.Exec("DELETE FROM nations")
	db.Exec("DELETE FROM SQLITE_SEQUENCE")
	// Seed nations
	for _, name := range nationNames {
		db.Create(&Nation{Name: name})
	}

	// Seed customers
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		firstName := customerFirstNames[rand.Intn(len(customerFirstNames))]
		lastName := customerLastNames[rand.Intn(len(customerLastNames))]
		idNation := rand.Intn(len(nationNames))
		nation := nationNames[idNation]
		db.Create(&Customer{
			FirstName: firstName,
			LastName:  lastName,
			NationID:  uint(idNation),
			Nation:    Nation{ID: uint(idNation), Name: nation},
		})
	}
}
