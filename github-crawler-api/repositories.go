package main

import (
	"github-crawler-api/app"
	"github.com/goadesign/goa"
)

// RepositoriesController implements the repositories resource.
type RepositoriesController struct {
	*goa.Controller
}

// NewRepositoriesController creates a repositories controller.
func NewRepositoriesController(service *goa.Service) *RepositoriesController {
	return &RepositoriesController{Controller: service.NewController("RepositoriesController")}
}

// List runs the list action.
func (c *RepositoriesController) List(ctx *app.ListRepositoriesContext) error {
	// RepositoriesController_List: start_implement

	// Put your logic here

	// RepositoriesController_List: end_implement
	res := app.CommitCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *RepositoriesController) Show(ctx *app.ShowRepositoriesContext) error {
	// RepositoriesController_Show: start_implement

	// Put your logic here

	// RepositoriesController_Show: end_implement
	res := &app.Repository{}
	return ctx.OK(res)
}
