package lead

import (
	"fmt"
	"go-fiber-crm/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name 				string	`json:"name"`
	Company 		string	`json:"company"`
	Email 			string	`json:"email"`
	Phone 			int			`json:"phone"`
}

var db *gorm.DB

func init() {
	fmt.Println("init lead")
	db = database.DBConn
}

func GetLeads(ctx *fiber.Ctx) error {
	db = database.DBConn
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) error{
	id := ctx.Params("id")
	var lead Lead

	if err := db.First(&lead, id).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) error {
	lead := new(Lead)

	if err := ctx.BodyParser(lead); err != nil {
		return ctx.Status(http.StatusServiceUnavailable).SendString(err.Error())
		
	}

	if err := db.Create(lead).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		
	}

	return ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := db.Delete(&Lead{}, id).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.SendString("lead successfully deleted")
}