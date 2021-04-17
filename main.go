package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/internal/controller"
	"github.com/starptech/go-web/internal/core"
	"github.com/starptech/go-web/internal/models"
)
var (
	sqlType = "postgres"
	username = "postgres"
	password = "$Gear2021"
	host = "127.0.0.1"
	port = 5432
	dbname = "goweb"
	schema =""
	sslmode = "disable"
	connect_timeout = 10000
	sqlTable = ""

)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	config.ConnectionString, err = BuildDatasourceName(sqlType, host, port, username, password, dbname, sslmode, schema, "")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	config.RedisAddr = "127.0.0.1:6379"
	config.RedisPwd = "newbared"

	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	userCtrl := &controller.User{}
	userListCtrl := &controller.UserList{}
	healthCtrl := &controller.Healthcheck{}

	main := server.Echo.Group("/FileManager")
	main.GET("/home", controller.GetStartPage)

	// api endpoints
	g := server.Echo.Group("/api")
	g.GET("/users/:id", userCtrl.GetUserJSON)

	// pages
	u := server.Echo.Group("/users")
	u.GET("", userListCtrl.GetUsers)
	u.GET("/:id", userCtrl.GetUser)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	user := models.User{Name: "Peter"}
	mr := server.GetModelRegistry()
	err = mr.Register(user)

	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Create(&user)
	// Start server
	go func() {
		if err := server.Start(config.Address); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}

// Build the datasource name required to connect to a database using the specified driver.
func BuildDatasourceName(driver string, host string, port int, user string, password string, name string,
	sslMode string, schema string, sslrootcert string) (string, error) {

	switch driver {
	case "postgres":
		dsn := []string{}
		dsn = appendNotEmpty(dsn, "host=%s", host)
		dsn = append(dsn, fmt.Sprintf("port=%d", port))
		dsn = appendNotEmpty(dsn, "user=%s", user)
		dsn = appendNotEmpty(dsn, "password=%s", password)
		dsn = appendNotEmpty(dsn, "dbname=%s", name)
		dsn = appendNotEmpty(dsn, "sslmode=%s", sslMode)
		dsn = appendNotEmpty(dsn, "search_path=%s", schema)
		dsn = appendNotEmpty(dsn, "sslrootcert=%s", sslrootcert)
		return strings.Join(dsn, " "), nil
	case "sqlite3":
		// Only support
		return ":memory:", nil
	}
	return "", fmt.Errorf("unsupported driver: %v", driver)
}

func appendNotEmpty(s []string, format string, value string) []string {
	if value != "" {
		s = append(s, fmt.Sprintf(format, value))
	}
	return s
}
