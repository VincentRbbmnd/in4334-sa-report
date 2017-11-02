// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Application Contexts
//
// Command:
// $ goagen
// --design=github.com\VincentRbbmnd\in4334-sa-report\github-crawler-api\design
// --out=$(GOPATH)\src\github.com\VincentRbbmnd\in4334-sa-report\github-crawler-api
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
	"time"
)

// ListCommitsContext provides the commits list action context.
type ListCommitsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	From   *time.Time
	Limit  *int
	RepoID int
	Until  *time.Time
}

// NewListCommitsContext parses the incoming request URL and body, performs validations and creates the
// context used by the commits controller list action.
func NewListCommitsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListCommitsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCommitsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFrom := req.Params["from"]
	if len(paramFrom) > 0 {
		rawFrom := paramFrom[0]
		if from, err2 := time.Parse(time.RFC3339, rawFrom); err2 == nil {
			tmp1 := &from
			rctx.From = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("from", rawFrom, "datetime"))
		}
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp3 := limit
			tmp2 := &tmp3
			rctx.Limit = tmp2
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramRepoID := req.Params["repoID"]
	if len(paramRepoID) > 0 {
		rawRepoID := paramRepoID[0]
		if repoID, err2 := strconv.Atoi(rawRepoID); err2 == nil {
			rctx.RepoID = repoID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("repoID", rawRepoID, "integer"))
		}
		if rctx.RepoID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`repoID`, rctx.RepoID, 1, true))
		}
	}
	paramUntil := req.Params["until"]
	if len(paramUntil) > 0 {
		rawUntil := paramUntil[0]
		if until, err2 := time.Parse(time.RFC3339, rawUntil); err2 == nil {
			tmp5 := &until
			rctx.Until = tmp5
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("until", rawUntil, "datetime"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCommitsContext) OK(r CommitCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.commit+json; type=collection")
	if r == nil {
		r = CommitCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *ListCommitsContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListCommitsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCommitsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowCommitsContext provides the commits show action context.
type ShowCommitsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RepoID int
	Sha    string
}

// NewShowCommitsContext parses the incoming request URL and body, performs validations and creates the
// context used by the commits controller show action.
func NewShowCommitsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCommitsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCommitsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRepoID := req.Params["repoID"]
	if len(paramRepoID) > 0 {
		rawRepoID := paramRepoID[0]
		if repoID, err2 := strconv.Atoi(rawRepoID); err2 == nil {
			rctx.RepoID = repoID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("repoID", rawRepoID, "integer"))
		}
		if rctx.RepoID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`repoID`, rctx.RepoID, 1, true))
		}
	}
	paramSha := req.Params["sha"]
	if len(paramSha) > 0 {
		rawSha := paramSha[0]
		rctx.Sha = rawSha
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCommitsContext) OK(r *Commit) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.commit+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowCommitsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCommitsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListDevelopersContext provides the developers list action context.
type ListDevelopersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	From   *time.Time
	Limit  *int
	RepoID int
	Until  *time.Time
}

// NewListDevelopersContext parses the incoming request URL and body, performs validations and creates the
// context used by the developers controller list action.
func NewListDevelopersContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListDevelopersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListDevelopersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFrom := req.Params["from"]
	if len(paramFrom) > 0 {
		rawFrom := paramFrom[0]
		if from, err2 := time.Parse(time.RFC3339, rawFrom); err2 == nil {
			tmp7 := &from
			rctx.From = tmp7
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("from", rawFrom, "datetime"))
		}
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) > 0 {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			tmp9 := limit
			tmp8 := &tmp9
			rctx.Limit = tmp8
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
	}
	paramRepoID := req.Params["repoID"]
	if len(paramRepoID) > 0 {
		rawRepoID := paramRepoID[0]
		if repoID, err2 := strconv.Atoi(rawRepoID); err2 == nil {
			rctx.RepoID = repoID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("repoID", rawRepoID, "integer"))
		}
		if rctx.RepoID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`repoID`, rctx.RepoID, 1, true))
		}
	}
	paramUntil := req.Params["until"]
	if len(paramUntil) > 0 {
		rawUntil := paramUntil[0]
		if until, err2 := time.Parse(time.RFC3339, rawUntil); err2 == nil {
			tmp11 := &until
			rctx.Until = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("until", rawUntil, "datetime"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListDevelopersContext) OK(r GhuserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.ghuser+json; type=collection")
	if r == nil {
		r = GhuserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *ListDevelopersContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListDevelopersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListDevelopersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListRepositoriesContext provides the repositories list action context.
type ListRepositoriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRepositoriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the repositories controller list action.
func NewListRepositoriesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListRepositoriesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListRepositoriesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRepositoriesContext) OK(r RepositoryCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.repository+json; type=collection")
	if r == nil {
		r = RepositoryCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListRepositoriesContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// ShowRepositoriesContext provides the repositories show action context.
type ShowRepositoriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RepoID int
}

// NewShowRepositoriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the repositories controller show action.
func NewShowRepositoriesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowRepositoriesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowRepositoriesContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRepoID := req.Params["repoID"]
	if len(paramRepoID) > 0 {
		rawRepoID := paramRepoID[0]
		if repoID, err2 := strconv.Atoi(rawRepoID); err2 == nil {
			rctx.RepoID = repoID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("repoID", rawRepoID, "integer"))
		}
		if rctx.RepoID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`repoID`, rctx.RepoID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRepositoriesContext) OK(r *Repository) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.repository+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowRepositoriesContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRepositoriesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
