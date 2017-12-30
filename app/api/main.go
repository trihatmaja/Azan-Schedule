/*
	The api implement mysql db, memcached, and Dr. T. Djamaluddin calculation
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"

	azan "github.com/trihatmaja/Azan-Schedule"

	"github.com/trihatmaja/Azan-Schedule/cache"
	"github.com/trihatmaja/Azan-Schedule/calculation"
	"github.com/trihatmaja/Azan-Schedule/database"
	"github.com/trihatmaja/Azan-Schedule/handler"
)

var (
	Version string
	Build   string
	AppPort string
)

func version(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	retString := fmt.Sprintf("{Version: %s, Build: %s}", Version, Build)
	fmt.Fprintf(w, retString)
}

func main() {
	gotenv.Load(".env")

	AppPort = os.Getenv("APP_PORT")

	dbOpt := database.OptionMySQL{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Charset:  os.Getenv("MYSQL_CHARSET"),
	}

	mysql, _ := database.NewMySQL(dbOpt)

	cOpt := cache.OptionsMemcached{
		Server:    strings.Split(os.Getenv("MEMCACHED_HOST"), ","),
		PrefixKey: os.Getenv("MEMCACHED_PREFIX_KEY"),
	}

	mcached := cache.NewMemcached(cOpt)

	tdjamaluddin := calculation.NewTDjamaluddin()

	az := azan.New(mysql, mcached, tdjamaluddin)

	azHandler := handler.NewHandler(az)

	router := httprouter.New()
	router.GET("/api/healthz", azHandler.Healthz)
	router.GET("/api/metrics", azHandler.Metrics)
	router.POST("/api/generate", azHandler.Generate)
	router.GET("/api/cities/", azHandler.ByCities)
	router.GET("/api/cities/:city", azHandler.ByCity)
	router.GET("/api/cities/:city/date/:date", azHandler.ByCityDate)
	router.GET("/api/cities/:city/month/:month", azHandler.ByCityMonth)
	router.GET("/version.txt", version)
	//router.GET("/api/", azHandler.All)

	handler := cors.Default().Handler(router)

	log.Println("application ready to receive request at port " + AppPort)
	http.ListenAndServe(":"+AppPort, handler)
}
