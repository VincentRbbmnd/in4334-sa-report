package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("commits", func() {
	DefaultMedia(CommitMedia)
	BasePath("/commits")
	Action("show", func() {
		Routing(
			GET("/:commitID"),
		)
		Description("Retrieve commit from db")
		Params(func() {
			Param("commitID", Integer, "Commit ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Routing(
			POST("/list"),
		)
		Payload(ListPayload)
		Description("Retrieve commits between timespan with users")
		Response(OK, CollectionOf(CommitMedia))
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("repositories", func() {
	DefaultMedia(RepositoryMedia)
	BasePath("/projects")
	Action("list", func() {
		Routing(
			GET("/list"),
		)
		Description("Retrieve all projects")
		Response(OK, CollectionOf(RepositoryMedia))
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
