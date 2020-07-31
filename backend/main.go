package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"log"
	"musicplayer/api"
	"musicplayer/sql"
	"os"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		getenv("MYSQL_USER", "root"),
		getenv("MYSQL_PASSWORD", "12345678"),
		getenv("MYSQL_HOST", "localhost"),
		getenv("MYSQL_DATABASE", "musicplayer"),
	)

	db, err := sql.ConnectMySQL(dsn)
	if err != nil {
		log.Fatalf("error connecting to the MySQL database: %v", err)
	}

	secret := getenv("JWT_SECRET", "EbnJO3bwmX")

	app := iris.New()

	app.Use(Cors)

	subRouter := api.Router(db, secret)
	app.PartyFunc("/", subRouter)

	addr := fmt.Sprintf(":%s", getenv("PORT", "8080"))
	app.Listen(addr)
}

func getenv(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

// Cors
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}