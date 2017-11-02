// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Application Controllers
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
	"github.com/goadesign/goa/cors"
	"github.com/goadesign/goa/encoding/form"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(form.NewDecoder, "application/x-www-form-urlencoded")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CommitsController is the controller interface for the Commits actions.
type CommitsController interface {
	goa.Muxer
	List(*ListCommitsContext) error
	Show(*ShowCommitsContext) error
}

// MountCommitsController "mounts" a Commits resource controller on the given service.
func MountCommitsController(service *goa.Service, ctrl CommitsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/repositories/:repoID/commits/list", ctrl.MuxHandler("preflight", handleCommitsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v1/repositories/:repoID/commits/:sha", ctrl.MuxHandler("preflight", handleCommitsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCommitsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleCommitsOrigin(h)
	service.Mux.Handle("GET", "/v1/repositories/:repoID/commits/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Commits", "action", "List", "route", "GET /v1/repositories/:repoID/commits/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCommitsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleCommitsOrigin(h)
	service.Mux.Handle("GET", "/v1/repositories/:repoID/commits/:sha", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Commits", "action", "Show", "route", "GET /v1/repositories/:repoID/commits/:sha")
}

// handleCommitsOrigin applies the CORS response headers corresponding to the origin.
func handleCommitsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Auth, X-Pin, X-Platform, content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// DevelopersController is the controller interface for the Developers actions.
type DevelopersController interface {
	goa.Muxer
	List(*ListDevelopersContext) error
}

// MountDevelopersController "mounts" a Developers resource controller on the given service.
func MountDevelopersController(service *goa.Service, ctrl DevelopersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/repositories/:repoID/developers/list", ctrl.MuxHandler("preflight", handleDevelopersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListDevelopersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleDevelopersOrigin(h)
	service.Mux.Handle("GET", "/v1/repositories/:repoID/developers/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Developers", "action", "List", "route", "GET /v1/repositories/:repoID/developers/list")
}

// handleDevelopersOrigin applies the CORS response headers corresponding to the origin.
func handleDevelopersOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Auth, X-Pin, X-Platform, content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// RepositoriesController is the controller interface for the Repositories actions.
type RepositoriesController interface {
	goa.Muxer
	List(*ListRepositoriesContext) error
	Show(*ShowRepositoriesContext) error
}

// MountRepositoriesController "mounts" a Repositories resource controller on the given service.
func MountRepositoriesController(service *goa.Service, ctrl RepositoriesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/repositories/list", ctrl.MuxHandler("preflight", handleRepositoriesOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v1/repositories/:repoID", ctrl.MuxHandler("preflight", handleRepositoriesOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRepositoriesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleRepositoriesOrigin(h)
	service.Mux.Handle("GET", "/v1/repositories/list", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Repositories", "action", "List", "route", "GET /v1/repositories/list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRepositoriesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleRepositoriesOrigin(h)
	service.Mux.Handle("GET", "/v1/repositories/:repoID", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Repositories", "action", "Show", "route", "GET /v1/repositories/:repoID")
}

// handleRepositoriesOrigin applies the CORS response headers corresponding to the origin.
func handleRepositoriesOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Auth, X-Pin, X-Platform, content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
