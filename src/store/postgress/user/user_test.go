package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"testing"
)

var db interface{}

func init() {
	viper.AutomaticEnv()
	fmt.Println("Start")
	db = viper.Get("DB_PG")
	fmt.Println(db, "TESTTEST")
}

var pool, errP = pgxpool.Connect(context.Background(), db.(string))

var repo = &PGUserRepository{pool: pool}
var user = &User{
	ID:   21,
	Name: "Artem",
}

func TestPGUserRepository_Create(t *testing.T) {
	err := repo.Create(context.Background(), user)

	if err != nil {
		t.Error("Error while testing pg create")
	}

	defer pool.Close()

}
