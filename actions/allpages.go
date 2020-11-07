package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// AllpagesIndex default implementation.
func AllpagesIndex(c buffalo.Context) error {
	// c.Flash().Add("success", "You have been to index")
	return c.Render(http.StatusOK, r.HTML("allpages\\index.html"))
	// return c.Render(http.StatusOK, r.JSON("{success}"))
}

// AllpagesContact default implementation.
func AllpagesContact(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("allpages\\contact.html"))
}

// AllpagesBlogs default implementation.
func AllpagesBlogs(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("allpages\\blogs.html"))
}

// AllpagesWorks default implementation.
func AllpagesWorks(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("allpages\\works.html"))
}

// AllpagesAbout default implementation.
func AllpagesAbout(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("allpages\\about.html"))
}

// AllpagesLogin default implementation.
func AllpagesLogin(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("allpages\\login.html"))
}
