package runserver

import (
	"fmt"
	"github.com/cyberkarma/chatserver/configs"

	//"github.com/cyberkarma/chatserver/configs"
	"github.com/cyberkarma/chatserver/server"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"net/http"
)

var RunServer = &cobra.Command{
	Use:   "runServer",
	Short: "Run the server",
	Long:  "Blablabla",
	Run: func(cmd *cobra.Command, args []string) {
		router := server.RouterBuilder{}
		config, err := configs.LoadConfig()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(config.Pc.Port, "-config")
		}

		runErr := http.ListenAndServe(config.Pc.Port, router.Build())
		if runErr != nil {
			fmt.Println(runErr)
		}
		fmt.Println("Server is working now")
	},
}
