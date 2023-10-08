package main

import (
	"fmt"
	"par/config"
	"par/driver/mysql"
	fr "par/faktoryandroute"
	"par/migrasi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	migrasi.MigrateDB(db)

	e := echo.New()

	fr.FaktoryAndRoute(e, db)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", cfg.SERVER_PORT)))
}
