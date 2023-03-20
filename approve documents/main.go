package main

import (
	"log"

	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/routes"
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Approval Management")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api", home)
	app.Post("/user/createUser", routes.CreateUsers)
	app.Post("/document/CreateDocument", routes.CreateDocument)
	app.Post("/group/CreateGroup", routes.CreateGroups)
	app.Post("/document/CreateRequest", routes.CreateRequset)
	app.Post("/message/CreateMessage", routes.CreateMessage)
	app.Post("/document/CreateApproveDocument", routes.ApproveDocument)
	app.Get("/user/GetUsers", routes.GetUsers)
	app.Get("/document/GetDocuments", routes.GetDocuments)
	app.Get("/group/GetGroups", routes.GetGroups)
	app.Get("/message/GetMessage", routes.GetMessages)
	app.Get("/user/GetUser/:id", routes.GetUser)
	app.Get("/group/GetGroup/:id", routes.GetGroup)
	app.Get("/document/GetRequest/:id", routes.GetRequest)
	app.Get("/message/GetMessage/:id", routes.GetMessage)
	app.Get("/user/GetUser/:id", routes.GetUser)
	app.Put("/updateUser/:id", routes.UpdateUser)
	app.Put("/group/updateGroup/:id", routes.UpdateGroup)
	app.Put("/document/updateRequest/:id", routes.UpdateRequest)
	app.Put("/document/updateDocument/:id", routes.UpdateDocument)
	app.Put("/message/updateMessage/:id", routes.UpdateMessage)
	app.Delete("/user/deleteUser/:id", routes.DeleteUser)
	app.Delete("/document/deleteRequest/:id", routes.DeleteRequest)
	app.Delete("/group/deleteGroup/:id", routes.DeleteGroup)
	app.Delete("/document/deleteDocument/:id", routes.DeleteDocument)
	app.Delete("/document/deleteApproveDocument/:id", routes.DeleteApproveDocument)
	app.Delete("/message/deleteMessage/:id", routes.DeleteMessage)
	app.Get("/Doc/403", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusForbidden)
	})
	app.Get("/bad-request", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	})
	app.Get("/500", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusInternalServerError).SendString("Bad Request")
	})
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
