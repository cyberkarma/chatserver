package runserver

import (
	"fmt"
	"github.com/cyberkarma/chatserver/server"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

var RunServer = &cobra.Command{
	Use:   "runServer",
	Short: "Run the server",
	Long:  "Blablabla",
	Run: func(cmd *cobra.Command, args []string) {
		router := server.RouterBuilder{}
		err := http.ListenAndServe(viper.GetString("prod.port"), router.Build())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Server is working now")
	},
}
