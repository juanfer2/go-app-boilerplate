package servers

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"

	"github.com/juanfer2/api-rest-go/controllers"
	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

// "github.com/juanfer2/api-rest-go/controllers"
func helloWorld(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/tasks", controllers.GetTasks)
	app.Post("/tasks", controllers.CreateTasks)
	// app.Post("/tasks", controllers.CreateTaks)
}

func StartServerApp() {
	db := databases.Conn()
	// Migrate the schema
	db.AutoMigrate(&models.Task{})

	// db := databases.Conn()
	// log.Println(db.AutoMigrate(&models.Task{}).Error)
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	app := fiber.New()
	app.Use("/api-doc", swagger.Handler) // default

	app.Use("/api-doc", swagger.New(swagger.Config{ // custom
		URL:         "./docs/doc.json",
		DeepLinking: true,
	}))
	setupRoutes(app)
	app.Listen(":4000")
}
