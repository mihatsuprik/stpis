package routes

import (
	"errors"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Approve struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Requestid string `json:"requestid"`
	Docid     int    `json:"docid"`
}

func CreateResponseApprove(approveModel models.Approve) Approve {
	return Approve{ID: approveModel.ID, Requestid: approveModel.Requestid, Docid: approveModel.Docid}
}

func FindApproveHelper(id int, approve *models.Approve) error {
	database.Database.Db.Find(&approve, "id = ?", id)
	if approve.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func ApproveDocument(c *fiber.Ctx) error {
	var approve models.Approve

	if err := c.BodyParser(&approve); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&approve)
	responseAprove := CreateResponseApprove(approve)

	return c.Status(200).JSON(responseAprove)
}

// delete product
func DeleteApproveDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var approve models.Approve

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindApproveHelper(id, &approve); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&approve).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the Requset")
}
