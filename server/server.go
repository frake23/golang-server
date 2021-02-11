package server

import (
	"./services"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"os/signal"
	"time"
)

// App ...
type App struct {
	server *echo.Echo
	db     *JsonDb
}

// InitApp ...
func InitApp(connectionString string) *App {
	return &App{
		server: echo.New(),
		db:     &JsonDb{connectionString: connectionString},
	}
}

// Run ...
func (a *App) Run(port string) error {
	// Init middleware
	a.server.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := &ServerContext{c, &services.UserService{a.db}}
				return next(cc)
			}
		})
	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())

	// Init routes
	u := a.server.Group("/users")
	u.GET("", getAllUsers)
	u.GET("/:id", getUser)
	u.DELETE("/:id", deleteUser)
	u.POST("/add", addUser)
	u.PUT("/update", updateUser)

	// Start listening
	a.server.Logger.Fatal(a.server.Start(":" + port))

	go func() {
		if err := a.server.Start(":" + port); err != nil {
			a.server.Logger.Info("Failed to listen and serve:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.server.Shutdown(ctx)
}
