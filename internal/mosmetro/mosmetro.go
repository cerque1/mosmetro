package mosmetro

import (
	"database/sql"
	"fmt"
	"log"
	"mosmetro/internal/controller"
	"mosmetro/internal/repo/persistent"
	mosmetro_usecase "mosmetro/internal/usecase/mosmetro"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Run() {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_ADDR"),
			os.Getenv("POSTGRES_DB")))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	us := mosmetro_usecase.New(persistent.New(db))

	log.Println("Start serve")
	http.ListenAndServe(":8080", controller.NewRouter(us))
}
