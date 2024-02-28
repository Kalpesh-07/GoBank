package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretsecret@localhost:5432/simple_bank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {
	//conn, err := sql.Open(dbDriver, dbSource)
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
