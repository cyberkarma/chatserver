package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "printCmd21",
	Short: "Run the print",
	Long:  "Blablabla",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print is Running")
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
