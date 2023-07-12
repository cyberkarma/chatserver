package runserver

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
)

var RunServer = &cobra.Command{
	Use:   "printCmd21",
	Short: "Run the print",
	Long:  "Blablabla",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print is Running")
	},
}
