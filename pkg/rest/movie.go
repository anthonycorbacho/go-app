package rest

import (
	"github.com/anthonycorbacho/go-app/pkg/movie"
	"github.com/labstack/echo"
	"net/http"
)

// GetAllMovies returns all movies or empty.
func GetAllMovies(ms movie.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, ms.GetAll())
	}
}

// GetMovie returns a movie by its ID.
func GetMovie(ms movie.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := c.Param("id")

		m, err := ms.Get(movie.MovieID(ID))
		if err != nil {
			return parseError(c, err)
		}
		return c.JSON(http.StatusOK, m)
	}
}

// CreateMovie create a movie.
func CreateMovie(ms movie.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(movie.Movie)
		if err := c.Bind(m); err != nil {
			return parseError(c, err)
		}

		err := ms.Create(m)
		if err != nil {
			return parseError(c, err)
		}
		return c.JSON(http.StatusOK, m)
	}
}

// UpdateMovie update a movie.
func UpdateMovie(ms movie.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(movie.Movie)
		if err := c.Bind(m); err != nil {
			return parseError(c, err)
		}

		err := ms.Update(m)
		if err != nil {
			return parseError(c, err)
		}
		return c.JSON(http.StatusOK, m)
	}
}

// DeleteMovie create a flag.
func DeleteMovie(ms movie.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := c.Param("id")

		err := ms.Delete(movie.MovieID(ID))
		if err != nil {
			return parseError(c, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
