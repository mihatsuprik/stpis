package routes

import (
	"errors"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	User   int    `json:"user"`
	Assign string `json:"assign"`
}

func CreateResponseRequest(requestModel models.Request) Request {
	return Request{ID: requestModel.ID, Name: requestModel.Name, User: requestModel.User, Assign: requestModel.Assign}
}

func FindRequestHelper(id int, request *models.Request) error {
	database.Database.Db.Find(&request, "id = ?", id)
	if request.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func CreateRequset(c *fiber.Ctx) error {
	var request models.Request

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&request)
	responseDocument := CreateResponseRequest(request)

	return c.Status(200).JSON(responseDocument)
}

func UpdateRequest(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var request models.Request

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindRequestHelper(id, &request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedRequest struct {
		ID     uint   `json:"id" gorm:"primaryKey"`
		Name   string `json:"name"`
		User   int    `json:"user"`
		Assign string `json:"assign"`
	}

	var updateData UpdatedRequest

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	request.Name = updateData.Name
	request.User = updateData.User
	request.Assign = updateData.Assign

	database.Database.Db.Save(&request)

	responseRequest := CreateResponseRequest(request)
	return c.Status(200).JSON(responseRequest)

}

func GetRequest(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var request models.Request

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindRequestHelper(id, &request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseRequest := CreateResponseRequest(request)
	return c.Status(200).JSON(responseRequest)
}

// delete product
func DeleteRequest(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var request models.Request

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindRequestHelper(id, &request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&request).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the Requset")
}
