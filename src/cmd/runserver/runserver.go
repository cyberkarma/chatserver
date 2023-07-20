package runserver

import (
	"context"
	"fmt"
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
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.User, config.Password, config.Host, config.Dbport, config.Dbname)
		pool, err := pgxpool.Connect(context.Background(), dsn)
		if err != nil {
			return errors.Wrap(err, "Pool connection error")
		}
		err = pool.Ping(context.Background())
		if err != nil {
			return errors.Wrap(err, "Pool ping error")
		}

		//Run server part
		err = http.ListenAndServe("localhost:"+strconv.Itoa(config.Port), router.Build())
		if err != nil {
			return errors.Wrap(err, "RunServer error")
		}
		return nil
	},
}
