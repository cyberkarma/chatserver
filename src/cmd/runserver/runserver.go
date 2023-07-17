package runserver

import (
	"github.com/cyberkarma/chatserver/configs"
	"github.com/pkg/errors"
	"strconv"

	"github.com/cyberkarma/chatserver/server"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"net/http"
)

var RunServer = &cobra.Command{
	Use:   "runServer",
	Short: "Run the server",
	Long:  "Blablabla",
	RunE: func(cmd *cobra.Command, args []string) error {
		router := server.RouterBuilder{}
		config, err := configs.LoadConfig()
		if err != nil {
			return errors.Wrap(err, "LoadConfig error")
		}

		runErr := http.ListenAndServe("localhost:"+strconv.Itoa(config.Port), router.Build())
		if runErr != nil {
			return errors.Wrap(runErr, "RunServer error")
		}
		return nil
	},
}
