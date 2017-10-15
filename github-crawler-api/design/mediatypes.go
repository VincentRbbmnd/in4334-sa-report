package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CommitMedia = MediaType("application/vnd.commit+json", func() {
	Description("Commit data")
	Attributes(func() {
		Attribute("id", Integer, "ID of the commit in the database")
		Attribute("sha", String, "Unique identifier of the commit")
		Attribute("authorOfCommit", GHUserMedia, "Owner of the commit", func() {
			Metadata("struct:field:name", "author")
			Metadata("struct:tag:json", "author")
		})
		Attribute("timestamp", DateTime, "Time the commit happened")
		Attribute("message", String, "Message of the commit")
		Required("id", "sha", "authorOfCommit", "timestamp")
	})
	View("default", func() {
		Attribute("id")
		Attribute("authorOfCommit")
		Attribute("sha")
		Attribute("timestamp")
		Attribute("message")
	})
})

var RepositoryMedia = MediaType("application/vnd.repository+json", func() {
	Description("Commit data")
	Attributes(func() {
		Attribute("id", Primitive(Integer), "ID of the commit in the database")
		Attribute("owner", String, "Name of the owner of the repository")
		Attribute("org", Boolean, "If owner is an organization")
		Attribute("project_id", Number, "Time the commit happened")
		Attribute("full_name", String, "Full name of the repo")
		Required("id", "owner", "full_name", "project_id")
	})
	View("default", func() {
		Attribute("id")
		Attribute("owner")
		Attribute("full_name")
		Attribute("project_id")
		Attribute("org")
	})
})

var GHUserMedia = MediaType("application/vnd.ghuser+json", func() {
	Description("GH user data")
	Attributes(func() {
		Attribute("id", Integer, "ID of the user in the database")
		Attribute("login", String, "Unique username of the user")
		Attribute("type", String, "Type of the user")
		Attribute("locationForUser", LocationMedia, "Location of the user", func() {
			Metadata("struct:field:name", "location")
			Metadata("struct:tag:json", "location")
		})
		Required("id", "login", "type")
	})
	View("default", func() {
		Attribute("id")
		Attribute("login")
		Attribute("type")
		Attribute("locationForUser")
	})
})

var LocationMedia = MediaType("application/vnd.location+json", func() {
	Description("Location as geocoordinates")
	Attributes(func() {
		Attribute("id", Integer, "ID of the location in the database")
		Attribute("lat", Number, "coordinates lat")
		Attribute("lng", Number, "coordinates lng")
		Required("id", "lat", "lng")
	})
	View("default", func() {
		Attribute("id")
		Attribute("lat")
		Attribute("lng")
	})
})
