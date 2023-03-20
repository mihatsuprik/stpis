package routes

import (
	"errors"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Documents struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Creater int    `json:"creater"`
}

func CreateResponseDocument(documentModel models.Documents) Documents {
	return Documents{ID: documentModel.ID, Name: documentModel.Name, Content: documentModel.Content, Creater: documentModel.Creater}
}

func FindDocumentHelper(id int, document *models.Documents) error {
	database.Database.Db.Find(&document, "id = ?", id)
	if document.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func CreateDocument(c *fiber.Ctx) error {
	var document models.Documents

	if err := c.BodyParser(&document); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&document)
	responseDocument := CreateResponseDocument(document)

	return c.Status(200).JSON(responseDocument)
}

func GetDocuments(c *fiber.Ctx) error {
	document := []models.Documents{}

	database.Database.Db.Find(&document)
	responseDocument := []Documents{}

	for _, document := range document {
		responseDocuments := CreateResponseDocument(document)
		responseDocument = append(responseDocument, responseDocuments)
	}

	return c.Status(200).JSON(responseDocument)
}

func GetDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var document models.Documents

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindDocumentHelper(id, &document); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseFinance := CreateResponseDocument(document)
	return c.Status(200).JSON(responseFinance)
}

func UpdateDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var document models.Documents

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindDocumentHelper(id, &document); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedDocuments struct {
		ID      uint   `json:"id" gorm:"primaryKey"`
		Name    string `json:"name"`
		Content string `json:"content"`
		Creater int    `json:"creater"`
	}

	var updateData UpdatedDocuments

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	document.Name = updateData.Name
	document.Content = updateData.Content
	document.Creater = updateData.Creater

	database.Database.Db.Save(&document)

	responseDocument := CreateResponseDocument(document)
	return c.Status(200).JSON(responseDocument)

}

// delete product
func DeleteDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var document models.Documents

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindDocumentHelper(id, &document); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&document).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the Document")
}
