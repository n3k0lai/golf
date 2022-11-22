/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/n3k0lai/golf/cmd/lib"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backup called")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		chatterinoSettings := lib.GetChatterinoConfigJSON()

		// Create the file
		out, err := os.Create(homeDir + "/settings.json")
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		// Writer the body to file
		_, err = io.Copy(out, strings.NewReader(chatterinoSettings))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("backup complete")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
