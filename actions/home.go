package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

// HomeHealth default implementation.
func HomeHealth(c buffalo.Context) error {
	// return c.Render(http.StatusOK, r.HTML("home/health.html"))
	return c.Render(http.StatusOK, r.String("OK"))
}
