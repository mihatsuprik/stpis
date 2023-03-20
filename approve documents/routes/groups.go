package routes

import (
	"errors"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type Groups struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func CreateResponseGroups(groupsModel models.Groups) Groups {
	return Groups{ID: groupsModel.ID, Name: groupsModel.Name}
}

func FindGroupsHelper(id int, groups *models.Groups) error {
	database.Database.Db.Find(&groups, "id = ?", id)
	if groups.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func CreateGroups(c *fiber.Ctx) error {
	var groups models.Groups

	if err := c.BodyParser(&groups); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&groups)
	responseGroups := CreateResponseGroups(groups)

	return c.Status(200).JSON(responseGroups)
}

func GetGroups(c *fiber.Ctx) error {
	groups := []models.Groups{}

	database.Database.Db.Find(&groups)
	responseGroups := []Groups{}

	for _, groups := range groups {
		responseGroup := CreateResponseGroups(groups)
		responseGroups = append(responseGroups, responseGroup)
	}

	return c.Status(200).JSON(responseGroups)
}

func GetGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var group models.Groups

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindGroupsHelper(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseFinance := CreateResponseGroups(group)
	return c.Status(200).JSON(responseFinance)
}

func UpdateGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var group models.Groups

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindGroupsHelper(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedGroup struct {
		ID   uint   `json:"id" gorm:"primaryKey"`
		Name string `json:"name"`
	}

	var updateData UpdatedGroup

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	group.Name = updateData.Name

	database.Database.Db.Save(&group)

	responseDocument := CreateResponseGroups(group)
	return c.Status(200).JSON(responseDocument)

}

// delete product
func DeleteGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var group models.Groups

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindGroupsHelper(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&group).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the group")
}
