package runserver

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

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
		viper.AutomaticEnv()
		db := viper.Get("DB_PG")
		serv := viper.Get("SERVER")
		fmt.Println(db, serv, "DB AND SERVER")
		//Router part
		router := server.RouterBuilder{}

		//DB connection pool part
		pool, err := pgxpool.Connect(context.Background(), db.(string))
		if err != nil {
			return errors.Wrap(err, "Pool connection error")
		}
		err = pool.Ping(context.Background())
		if err != nil {
			return errors.Wrap(err, "Pool ping error")
		}

		//Run server part
		err = http.ListenAndServe("localhost:"+serv.(string), router.Build())
		if err != nil {
			return errors.Wrap(err, "RunServer error")
		}
		fmt.Println("Server is running!")
		return nil
	},
}
