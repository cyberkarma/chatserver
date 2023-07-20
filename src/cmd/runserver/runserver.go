package runserver

import (
	"context"
	"github.com/cyberkarma/chatserver/configs"
	"github.com/jackc/pgx/v4/pgxpool"
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
		//Router part
		router := server.RouterBuilder{}
		config, err := configs.LoadConfig()
		if err != nil {
			return errors.Wrap(err, "LoadConfig error")
		}

		//DB connection pool part
		pool, err := pgxpool.Connect(context.Background(), config.DB)
		if err != nil {
			return errors.Wrap(err, "Pool connection error")
		}
		err = pool.Ping(context.Background())
		if err != nil {
			return errors.Wrap(err, "Pool ping error")
		}

		//Run server part
		err = http.ListenAndServe("localhost:"+strconv.Itoa(config.Server), router.Build())
		if err != nil {
			return errors.Wrap(err, "RunServer error")
		}
		return nil
	},
}
