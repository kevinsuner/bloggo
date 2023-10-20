// SPDX-License-Identifier: MIT
package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Start an HTTP server pointing to the public directory",
	Long: `
Starts a Fiber HTTP server to serve the content contained in the 
public directory.
	`,
	Run: test,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func test(cmd *cobra.Command, args []string) {
	var baseURL string = viper.GetString("base-url")

	app := fiber.New()
	app.Static("/public", "./public")
	app.Static("/public/assets", "./public/assets")
	app.Static("/public/css", "./public/css")
	app.Static("/public/posts", "./public/posts")

	port, _, _ := strings.Cut(strings.SplitAfterN(baseURL, ":", -1)[2], "/")
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
