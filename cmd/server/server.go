package main

import (
	"github.com/anthonycorbacho/go-app/pkg"
	"github.com/anthonycorbacho/go-app/pkg/movie/couch"
	"github.com/anthonycorbacho/go-app/pkg/rest"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"time"
)

func main() {
	// Get the configuration
	var c pkg.Configuration
	err := envconfig.Process("goApp", &c)
	if err != nil {
		panic(err)
	}

	movie := couch.New(c.Couch.Url)

	// Init web framework
	e := echo.New()
	e.HideBanner = true
	e.Server.ReadTimeout = 10 * time.Second
	e.Server.WriteTimeout = 2 * time.Minute

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("32M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Api
	api := e.Group("/api")

	movies := api.Group("/movies")
	movies.GET("", rest.GetAllMovies(movie))
	movies.GET("/:id", rest.GetMovie(movie))
	movies.POST("", rest.CreateMovie(movie))
	movies.PATCH("", rest.UpdateMovie(movie))
	movies.DELETE("/:id", rest.DeleteMovie(movie))

	log.Fatal(e.Start(c.Port))
}
