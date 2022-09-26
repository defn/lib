package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
}

type QueryResult struct {
	Results []User `json:"results"`
}

func main() {
	app := fiber.New()

	users := []User{{"defn"}, {"Tolan"}, {"lamda"}, {"Hana"}}

	results := QueryResult{users}

	app.Get("/meh", func(c *fiber.Ctx) error {
		b, err := json.Marshal(results)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.SendString(string(b))
	})

	app.Listen(":3000")
}
