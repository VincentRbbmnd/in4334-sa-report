package main

import (
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"

	"github.com/goadesign/goa"
)

// CommitsController implements the commits resource.
type CommitsController struct {
	*goa.Controller
}

// NewCommitsController creates a commits controller.
func NewCommitsController(service *goa.Service) *CommitsController {
	return &CommitsController{Controller: service.NewController("CommitsController")}
}

// List runs the list action.
func (c *CommitsController) List(ctx *app.ListCommitsContext) error {
	// CommitsController_List: start_implement
	res := commitDB.ListCommitWithUsersWithLocationForRepo(ctx, 41881900, ctx.Payload.From, ctx.Payload.Till, *ctx.Payload.Limit)
	// CommitsController_List: end_implement
	// res := app.CommitCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CommitsController) Show(ctx *app.ShowCommitsContext) error {
	// CommitsController_Show: start_implement

	// Put your logic here

	// CommitsController_Show: end_implement
	res := &app.Commit{}
	return ctx.OK(res)
}
