package handlers

import (
	"fmt"
	"jobin/db"
	"jobin/utils"
	"jobin/views"
	"math/rand"

	"time"

	"github.com/gofiber/fiber/v2"
)

func generateRandomNumber() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

type Content struct {
	Content string `form:"content"`
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, views.Home("Abdulkader"))
	})

	app.Get("/new", func(c *fiber.Ctx) error {
		print(c.Query("content"))

		return Render(c, views.NewPage(""))
	})

	app.Get("/test12", func(c *fiber.Ctx) error {
		randomNumber := generateRandomNumber()
		// sleep for 2 seconds
		time.Sleep(2 * time.Second)


		return c.Redirect("/" + randomNumber, fiber.StatusSeeOther)
	})

	app.Get("/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		if len(slug) != 8 {
			return c.Status(fiber.StatusNotFound).SendString("Invalid")
		}

		paste, err := db.GetPaste(slug)

		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}
		return Render(c, views.SavedPaste(paste.Content, slug))
	})

	app.Get("/:slug/copy", func(c *fiber.Ctx) error {

		slug := c.Params("slug")
		if len(slug) != 8 {
			return c.Status(fiber.StatusNotFound).SendString("Invalid")
		}

		

		paste, err := db.GetPaste(slug)
		
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		}

		return Render(c, views.NewPage(paste.Content))
	})

	app.Get("/:slug/raw", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		if len(slug) != 8 {
			return c.Status(fiber.StatusNotFound).SendString("Invalid")
		}

		paste, err := db.GetPaste(slug)

		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}
		return Render(c, views.RawPaste(paste.Content))
	})



	app.Post("/new", func(c *fiber.Ctx) error {

		con := new(Content)

		if err := c.BodyParser(con); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Request")
		}

		slug, err := utils.GenerateSlug()

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		}

		err = db.CreatePaste(con.Content, slug)
		
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		}

		c.Set("HX-Redirect", "/" + slug)
		
		return c.Redirect("/" + slug, 201)
	})

	

}	