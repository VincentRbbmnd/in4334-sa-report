package main

import (
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"
	"github.com/goadesign/goa"
)

// DevelopersController implements the developers resource.
type DevelopersController struct {
	*goa.Controller
}

// NewDevelopersController creates a developers controller.
func NewDevelopersController(service *goa.Service) *DevelopersController {
	return &DevelopersController{Controller: service.NewController("DevelopersController")}
}

// List runs the list action.
func (c *DevelopersController) List(ctx *app.ListDevelopersContext) error {
	// DevelopersController_List: start_implement
	res := commitDB.ListDevelopers(ctx, ctx.RepoID, ctx.From, ctx.Until, ctx.Limit)

	// DevelopersController_List: end_implement
	return ctx.OK(res)
}
