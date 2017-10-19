package main

import (
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"

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

	repos := repoDB.ListRepository(ctx)
	return ctx.OK(repos)

	// RepositoriesController_List: end_implement
}

// Show runs the show action.
func (c *RepositoriesController) Show(ctx *app.ShowRepositoriesContext) error {
	// RepositoriesController_Show: start_implement

	repo, err := repoDB.OneRepository(ctx, ctx.RepoID)
	if err != nil {
		return ctx.BadRequest(err)
	}
	return ctx.OK(repo)
	// RepositoriesController_Show: end_implement
	res := &app.Repository{}
	return ctx.OK(res)
}
