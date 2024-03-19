/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"

	"github.com/lucasmdomingues/hexagonal/adapters/cli"
	productdb "github.com/lucasmdomingues/hexagonal/adapters/db"
	"github.com/lucasmdomingues/hexagonal/application"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

var (
	action       string
	productID    string
	productName  string
	productPrice float64
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("args", args)
		defer db.Close()

		productDB := productdb.NewProductDB(db)
		productService := application.NewProductService(productDB)

		result, err := cli.Run(productService, action, productID, productName, productPrice)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(result)
	},
}

func init() {
	log.Println("running cli command...")

	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable or Disable a product")
	cliCmd.Flags().StringVarP(&productID, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "Product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product price")

	db, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connected...")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
