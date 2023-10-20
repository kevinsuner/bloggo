// SPDX-License-Identifier: MIT
package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "bloggo",
	Short: "Bloggo is a free and simplistic static site generator",
	Long: `
Bloggo is a free and simplistic static site generator for individuals that
prioritize content over features, it lets you use custom themes by running
a simple command, and uses Go, Markdown and Frontmatter to generate HTML
pages that are optimized for performance and SEO.
	`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("bloggo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
