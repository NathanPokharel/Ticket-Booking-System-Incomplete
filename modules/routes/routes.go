package routes

import (
	"strconv"

	"github.com/1-Zed-1/Bus-Booking/modules/data"
	"github.com/1-Zed-1/Bus-Booking/templates"
	"github.com/labstack/echo"
)

func HandleRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		err := templates.Home().Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			ctx.Logger().Errorf("Failed to render template: %v", err)
			return echo.NewHTTPError(500, "Failed Rendering Template")
		}
		return nil
	})

	e.GET("/data-api/get-groups/", func(ctx echo.Context) error {
		groups, err := data.GetGroups()
		if err != nil {
			ctx.Logger().Errorf("database connection failed: %v", err)
			return echo.NewHTTPError(500, "database connection failed")
		}
		err = templates.RenderGroups(groups).Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			ctx.Logger().Errorf("Failed to render template: %v", err)
			return echo.NewHTTPError(500, "Failed Rendering Template")
		}
		return nil
	})

	e.GET("/data-api/get-buses/:groupID", func(ctx echo.Context) error {
		groupID, err := strconv.Atoi(ctx.Param("groupID"))
		if err != nil {
			ctx.Logger().Errorf("Failed converting string to int", err)
			return echo.NewHTTPError(500, "Internal server error")
		}
		buses, err := data.GetBuses(uint(groupID))
		if err != nil {
			ctx.Logger().Errorf("Failed Reading From Database", err)
			return echo.NewHTTPError(500, "Internal server error")
		}
		err = templates.RenderBuses(buses).Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			ctx.Logger().Errorf("Failed to render template: %v", err)
			return echo.NewHTTPError(500, "Failed Rendering Template")
		}
		return nil
	})

	e.GET("/data-api/get-seats/:busID", func(ctx echo.Context) error {
		busID, err := strconv.Atoi(ctx.Param("busID"))
		if err != nil {
			ctx.Logger().Errorf("Failed converting string to int", err)
			return echo.NewHTTPError(500, "Internal server error")
		}
		seats, err := data.GetSeats(uint(busID))
		if err != nil {
			ctx.Logger().Errorf("Failed Reading From Database", err)
			return echo.NewHTTPError(500, "Internal server error")
		}
		err = templates.RenderSeats(seats).Render(ctx.Request().Context(), ctx.Response().Writer)
		if err != nil {
			ctx.Logger().Errorf("Failed to render template: %v", err)
			return echo.NewHTTPError(500, "Failed Rendering Template")
		}
		return nil
	})
	/*data.GetGroups()
	data.GetBuses(1)
	data.GetBusesSeats(3)*/
}
