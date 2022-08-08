package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/guochenmeinian/shorten-url-project/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	// close connection at the end
	r := database.CreateClient(0)
	defer r.Close()

	// redis is a key-value paired database
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})
		// if the error is not "can't found", then it's likely a connection problem
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}

	// increment and also have to close connection
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	// redirect user to actual URL
	return c.Redirect(value, 301)
}
