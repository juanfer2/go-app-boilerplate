package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/controllers"
)

// "github.com/juanfer2/api-rest-go/controllers"
func helloWorld(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/tasks", controllers.GetTaks)
	// app.Post("/tasks", controllers.CreateTaks)
}

func StartServerApp() {
	// db := databases.Conn()
	// log.Println(db.AutoMigrate(&models.Task{}).Error)
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":4000")
}
