package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ListPayload = Type("ListPayload", func() {
	Attribute("from", DateTime)
	Attribute("till", DateTime)
	Attribute("limit", Integer)
})