package actions

import (
	"fmt"
	"net/http"

	"callofduty/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	incidents := &models.Incidents{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Incidents from the DB
	if err := q.All(incidents); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("incidents", incidents)
		return c.Render(http.StatusOK, r.HTML("/incidents/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(incidents))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(incidents))
	}).Respond(c)
}

// RoutesHandler displays the available app site map
func RoutesHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("route.html"))
}

// HomeHealth default implementation.
func HomeHealth(c buffalo.Context) error {
	// return c.Render(http.StatusOK, r.HTML("home/health.html"))
	return c.Render(http.StatusOK, r.String("OK"))
}
