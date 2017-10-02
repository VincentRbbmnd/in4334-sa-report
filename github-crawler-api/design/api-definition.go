package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	_ "lab.weave.nl/forks/jsgen"
)

// HOPPA API
var _ = API("GHCrawler", func() {
	Title("Visualization API for github crawler")
	Description("API for retrieving specific data from the github crawler")
	Contact(func() {
		Name("Rick Proost, Wim Spaargaren, Vincent Robbemond")
		Email("rpjproost@gmail.com")
		URL("http://127.0.0.1")
	})
	Host("86.87.235.82:8081")
	Scheme("http")
	BasePath("/v1")
	Origin("*", func() {
		Headers("Authorization, X-Auth, X-Pin", "X-Platform", "content-type")
		Methods("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS")
		MaxAge(600)
		Credentials()
	})
	Consumes("application/json")
	Consumes("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
