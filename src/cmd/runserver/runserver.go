package runserver

import (
	"fmt"
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
		err := http.ListenAndServe(":3000", router.Build())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Server is working now")
	},
}
