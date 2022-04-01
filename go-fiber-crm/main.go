package main

import (
	"fmt"
	"go-fiber-crm/database"
	"go-fiber-crm/lead"
	"log"

	"github.com/gofiber/fiber/v2"
)



func main() {
	app := fiber.New()
	initDatabase()
	setRoutes(app)
	log.Fatal(app.Listen(":8000"))

	db, err := database.DBConn.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}


func setRoutes(app *fiber.App) {
	app.Get("api/v1",func (c *fiber.Ctx) error {
		return c.SendString("hello world! from api")
	})
	app.Get("api/v1/lead",lead.GetLeads)
	app.Post("api/v1/lead",lead.NewLead)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
} 

func initDatabase() {
		fmt.Println("init main")
		database.DBConn.AutoMigrate(&lead.Lead{})
}