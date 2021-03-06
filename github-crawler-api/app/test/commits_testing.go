// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": commits TestHelpers
//
// Command:
// $ goagen
// --design=github.com\VincentRbbmnd\in4334-sa-report\github-crawler-api\design
// --out=$(GOPATH)\src\github.com\VincentRbbmnd\in4334-sa-report\github-crawler-api
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"
)

// ListCommitsBadRequest runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCommitsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, from *time.Time, limit *int, until *time.Time) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		query["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		query["until"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/repositories/%v/commits/list", repoID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		prms["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		prms["until"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	listCtx, _err := app.NewListCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListCommitsNoContent runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCommitsNoContent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, from *time.Time, limit *int, until *time.Time) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		query["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		query["until"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/repositories/%v/commits/list", repoID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		prms["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		prms["until"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	listCtx, _err := app.NewListCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// ListCommitsNotFound runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCommitsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, from *time.Time, limit *int, until *time.Time) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		query["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		query["until"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/repositories/%v/commits/list", repoID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		prms["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		prms["until"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	listCtx, _err := app.NewListCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListCommitsOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCommitsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, from *time.Time, limit *int, until *time.Time) (http.ResponseWriter, app.CommitCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		query["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		query["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		query["until"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/repositories/%v/commits/list", repoID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	if from != nil {
		sliceVal := []string{(*from).Format(time.RFC3339)}
		prms["from"] = sliceVal
	}
	if limit != nil {
		sliceVal := []string{strconv.Itoa(*limit)}
		prms["limit"] = sliceVal
	}
	if until != nil {
		sliceVal := []string{(*until).Format(time.RFC3339)}
		prms["until"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	listCtx, _err := app.NewListCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.CommitCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.CommitCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.CommitCollection", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowCommitsBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCommitsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, sha string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/repositories/%v/commits/%v", repoID, sha),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	prms["sha"] = []string{fmt.Sprintf("%v", sha)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	showCtx, _err := app.NewShowCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowCommitsNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCommitsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, sha string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/repositories/%v/commits/%v", repoID, sha),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	prms["sha"] = []string{fmt.Sprintf("%v", sha)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	showCtx, _err := app.NewShowCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ShowCommitsOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowCommitsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CommitsController, repoID int, sha string) (http.ResponseWriter, *app.Commit) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/repositories/%v/commits/%v", repoID, sha),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["repoID"] = []string{fmt.Sprintf("%v", repoID)}
	prms["sha"] = []string{fmt.Sprintf("%v", sha)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CommitsTest"), rw, req, prms)
	showCtx, _err := app.NewShowCommitsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Commit
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Commit)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Commit", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
