package routes

import (
	"errors"

	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Role    string `json:"role"`
	Email   string `json:"email"`
}

func CreateResponseUsers(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Surname: userModel.Surname, Age: userModel.Age, Role: userModel.Role, Email: userModel.Email}
}

func CreateUsers(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseProduct := CreateResponseUsers(user)

	return c.Status(200).JSON(responseProduct)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUsers(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func FindUserHelper(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func FindUserName(name string, user *models.User) error {
	database.Database.Db.Find(&user, "name = ?", name)
	if user.Name == "" {
		return errors.New("user does not exist")
	}

	return nil
}

func FindNameUs(c *fiber.Ctx) error {
	name, err := c.ParamsInt("name")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindUserName(string(rune(name)), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUsers(user)
	return c.Status(200).JSON(responseUser)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("user must be an integer")
	}

	if err := FindUserHelper(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUsers(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}

	if err := FindUserHelper(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedUser struct {
		ID      uint   `json:"id" gorm:"primaryKey"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Age     int    `json:"age"`
		Role    string `json:"role"`
		Email   string `json:"email"`
	}

	var updateData UpdatedUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.Name = updateData.Name
	user.Surname = updateData.Surname
	user.Age = updateData.Age
	user.Role = updateData.Role
	user.Email = updateData.Email

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUsers(user)
	return c.Status(200).JSON(responseUser)

}
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("product_id must be an integer")
	}
	if err := FindUserHelper(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted the User")
}
