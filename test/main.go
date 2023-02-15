package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yaameen/faster"
)

func x(c *faster.Ctx) error {
	c.Set("X-Api-Version", "1.0.0")
	return c.Next()
}

func s(app faster.FastRouter) faster.FastRouter {
	v1 := faster.New()
	v1.Get("/hell", func(c *faster.Ctx) error {
		return c.SendString("Why??")
	})
	return app.Mount("v1", v1)
}

func index(c *faster.Ctx) error {
	return c.SendString("Island Index")
}
func ro() *fiber.App {
	v1 := fiber.New()
	v1.Get("/", index)

	return v1
}

func main() {

	app := fiber.New()
	// micro := faster.New()

	// micro.Use(x)
	// x := app.Mount("john", micro) // GET /john/doe -> 200 OK

	// x.Get("/x", func(c *faster.Ctx) error {
	// 	return c.SendString("Hola") // GET /john/x -> 200 OK
	// }).Prefix("v1").Get("y", func(c *faster.Ctx) error {
	// 	return c.SendString("Hola") // GET /john/v1/y -> 200 OK
	// })

	// x.Get("/y", func(c *faster.Ctx) error {
	// 	return c.SendString("Hola") // GET /john/y -> 200 OK
	// })

	// micro.Get("/doe", func(c *faster.Ctx) error {
	// 	return c.SendStatus(200) // GET /john/doe -> 200 OK
	// })

	// s(app)

	// v1 := faster.New()
	app.Mount("islands", ro())
	// v1.Get("/hello", func(c *faster.Ctx) error {
	// 	return c.SendString("Why??")
	// })

	app.Get("/", func(c *faster.Ctx) error {
		return c.SendString("Hola") // GET /john/y -> 200 OK
	})

	app.All("*", func(c *faster.Ctx) error {
		return c.SendStatus(404) // <- 404
	})

	log.Fatal(app.Listen(":3010"))
}
