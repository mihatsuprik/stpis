package routes

import (
	"errors"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Requestid int  `json:"requestid"`
	Userid    int  `json:"userid"`
	Notice    int  `json:"notice"`
}

func CreateResponseMessage(messageModel models.Message) Message {
	return Message{ID: messageModel.ID, Requestid: messageModel.Requestid, Userid: messageModel.Userid, Notice: messageModel.Notice}
}

func FindMessageHelper(id int, message *models.Message) error {
	database.Database.Db.Find(&message, "id = ?", id)
	if message.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func CreateMessage(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&message)
	responseDocument := CreateResponseMessage(message)

	return c.Status(200).JSON(responseDocument)
}

func GetMessages(c *fiber.Ctx) error {
	message := []models.Message{}

	database.Database.Db.Find(&message)
	responseDocument := []Message{}

	for _, message := range message {
		responseDocuments := CreateResponseMessage(message)
		responseDocument = append(responseDocument, responseDocuments)
	}

	return c.Status(200).JSON(responseDocument)
}

func GetMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var message models.Message

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindMessageHelper(id, &message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseFinance := CreateResponseMessage(message)
	return c.Status(200).JSON(responseFinance)
}

func UpdateMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var message models.Message

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindMessageHelper(id, &message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedMessage struct {
		ID        uint `json:"id" gorm:"primaryKey"`
		Requestid int  `json:"requestid"`
		Userid    int  `json:"userid"`
		Notice    int  `json:"notice"`
	}

	var updateData UpdatedMessage

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	message.Requestid = updateData.Requestid
	message.Userid = updateData.Userid
	message.Notice = updateData.Notice

	database.Database.Db.Save(&message)

	responseDocument := CreateResponseMessage(message)
	return c.Status(200).JSON(responseDocument)

}

// delete product
func DeleteMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var message models.Message

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindMessageHelper(id, &message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&message).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the message")
}
