package cmd

import (
	"github.com/spf13/cobra"

	lib "github.com/n3k0lai/golf/cmd/lib"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Controls the client storage of the config",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		jsonString := lib.GetChatterinoConfigJSON()
		chatterinoConfig := lib.ParseChatterinoConfigJSON(jsonString)
		//fmt.Println(chatterinoConfig)
		//fmt.Printf("json map: %v\n", &chatterinoConfig)

		lib.AuditConfig(chatterinoConfig)
		// save locally

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	clientCmd.Flags().BoolP("toggle", "a", false, "Audit config")
}
