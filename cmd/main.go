package main


import (
	"github.com/1-Zed-1/Bus-Booking/modules/data"
	"github.com/1-Zed-1/Bus-Booking/modules/routes"
	"github.com/labstack/echo"
)

func main() {
	data.InitDatabase()

	//data.Populate()
	e := echo.New()
	routes.HandleRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
