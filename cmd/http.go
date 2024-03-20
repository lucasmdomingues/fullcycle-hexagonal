/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"

	productdb "github.com/lucasmdomingues/hexagonal/adapters/db"
	"github.com/lucasmdomingues/hexagonal/adapters/web"
	"github.com/lucasmdomingues/hexagonal/application"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running cli command...")

		db, err = sql.Open("sqlite3", "./db.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		productDB := productdb.NewProductDB(db)
		productService := application.NewProductService(productDB)

		server := web.NewServer(productService)

		log.Println("webserver has been started...")
		server.Serve()
		log.Println("webserver has been stopped...")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
