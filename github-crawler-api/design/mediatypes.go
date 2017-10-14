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
		Attribute("author", GHUserMedia, "Owner of the commit")
		Attribute("timestamp", DateTime, "Time the commit happened")
		Attribute("message", String, "Message of the commit")
		Required("id", "sha", "author", "timestamp")
	})
	View("default", func() {
		Attribute("id")
		Attribute("author")
		Attribute("sha")
		Attribute("timestamp")
		Attribute("message")
	})
})

var GHUserMedia = MediaType("application/vnd.ghuser+json", func() {
	Description("GH user data")
	Attributes(func() {
		Attribute("id", Integer, "ID of the user in the database")
		Attribute("login", String, "Unique username of the user")
		Attribute("type", String, "Type of the user")
		Attribute("location", LocationMedia, "Location of the user")
		Required("id", "login", "type")
	})
	View("default", func() {
		Attribute("id")
		Attribute("login")
		Attribute("type")
		Attribute("location")
	})
})

var LocationMedia = MediaType("application/vnd.location+json", func() {
	Description("Location as geocoordinates")
	Attributes(func() {
		Attribute("id", Integer, "ID of the location in the database")
		Attribute("lat", String, "coordinates lat")
		Attribute("lng", String, "coordinates lng")
		Required("id", "lat", "lng")
	})
	View("default", func() {
		Attribute("id")
		Attribute("lat")
		Attribute("lng")
	})
})

var RepositoryMedia = MediaType("application/vnd.repository+json", func() {
	Description("Repository data")
	Attributes(func() {
		Attribute("id", Integer, "ID of the commit in the database")
		Attribute("owner", String, "Owner of repo")
		Attribute("org", Boolean, "Is repo of organisation")
		Attribute("project_id", Integer, "Repository id")
		Attribute("user_type", String, "Type of owner of repo")
		Attribute("full_name", String, "Repository full name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("owner")
		Attribute("org")
		Attribute("project_id")
		Attribute("user_type")
		Attribute("full_name")
	})
})
