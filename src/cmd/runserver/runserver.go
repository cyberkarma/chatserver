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

const (
	host     = "localhost"
	port     = 54320
	user     = "postgres"
	password = "password"
	dbname   = "test"
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
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, dbname)
		pool, err := pgxpool.Connect(context.Background(), dsn)
		if err != nil {
			fmt.Println(err, "27 строка")
		} else {
			fmt.Println(pool, "31 строка")
		}
		pingErr := pool.Ping(context.Background())
		if pingErr != nil {
			fmt.Println(pingErr, "33 строка")
		} else {
			fmt.Println("Ошибки нет 35 строка")
		}

		//Run server part
		runErr := http.ListenAndServe("localhost:"+strconv.Itoa(config.Port), router.Build())
		if runErr != nil {
			return errors.Wrap(runErr, "RunServer error")
		}
		return nil
	},
}
