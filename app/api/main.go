package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/trihatmaja/Azan-Schedule/calculation"
	"github.com/trihatmaja/Azan-Schedule/database"
	"github.com/trihatmaja/Azan-Schedule/handler"
)

func main() {
	gotenv.Load("../../.env")

	dbOpt := database.OptionMySQL{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Charset:  os.Getenv("MYSQL_CHARSET"),
	}

	mysql, _ := database.NewMySQL(dbOpt)
	tdjamaluddin := calculation.NewTDjamaluddin()

	az := azan.New(mysql, tdjamaluddin)

	azHandler := handler.NewHandler(az)

	router := httprouter.New()
	router.GET("/healthz", azHandler.Healthz)
	router.POST("/generate", azHandler.Generate)
	router.GET("/", azHandler.All)

	http.ListenAndServe(":1234", router)
}
