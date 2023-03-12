package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()

		results := QueryResult{
			[]User{{"defn"}, {"Tolan"}, {"lamda"}, {"Hana"}},
		}

		app.Get("/meh", func(c *fiber.Ctx) error {
			b, err := json.Marshal(results)
			if err != nil {
				fmt.Println(err)
				return err
			}

			return c.SendString(string(b))
		})

		app.Listen(":3000")

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
