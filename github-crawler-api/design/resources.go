package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("commits", func() {
	DefaultMedia(CommitMedia)
	Parent("repositories")
	BasePath("/commits")
	Action("list", func() {
		Routing(
			GET("/list"),
		)
		Params(func() {
			Param("from", DateTime, "From date", func() {
			})
			Param("until", DateTime, "Till ID", func() {
			})
			Param("limit", Integer, "Limit the results")
		})
		Description("Retrieve commits between timespan with users")
		Response(OK, CollectionOf(CommitMedia))
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("repositories", func() {
	DefaultMedia(RepositoryMedia)
	BasePath("/repositories")
	Action("show", func() {
		Routing(
			GET("/:repoID"),
		)
		Description("Retrieve repository from db")
		Params(func() {
			Param("repoID", Integer, "Repository ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Routing(
			GET("/list"),
		)
		Payload(ListPayload)
		Description("Retrieve all repositories in DB")
		Response(OK, CollectionOf(CommitMedia))
		Response(BadRequest, ErrorMedia)
	})
})
