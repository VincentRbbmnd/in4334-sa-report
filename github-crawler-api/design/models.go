package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

// var JSON gorma.FieldType = "[]byte"

var _ = StorageGroup("GHAPI", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")

		//User model
		Model("User", func() {
			RendersTo(GHUserMedia)
			Description("Github user model in DB")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("login", gorma.String)
			Field("github_user_id", gorma.BigInteger)
			Field("type", gorma.String)
			Field("raw", gorma.String, func() {
				SQLTag("type:jsonb")
			})
			Field("location_checked", gorma.Boolean)
			Field("location_id", gorma.Integer)
		})

		// Commit model
		Model("Commit", func() {
			RendersTo(CommitMedia)
			Description("Github commit model in DB")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("message", gorma.String)
			Field("sha", gorma.String)
			Field("author_id", gorma.BigInteger)
			Field("committer_id", gorma.BigInteger)
			Field("repository_id", gorma.BigInteger)
			Field("raw", gorma.String, func() {
				SQLTag("type:jsonb")
			})
		})

		// Location model
		Model("Location", func() {
			Description("Location model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("point", gorma.String, func() {
				SQLTag("type:geometry(Point,4326)")
			})
			Field("user_id", gorma.BigInteger)
		})
	})
})