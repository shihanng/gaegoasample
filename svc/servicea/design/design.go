package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"
import _ "goa.design/goa/codegen/generator"

var _ = API("servicea", func() {
	Title("The service a")
})

var Info = ResultType("application/gaegoasample.info", func() {
	Description("Info about service a")
	Attributes(func() {
		Attribute("id", String, "ID of the service")
		Attribute("service_name", String, "Service's name")
		Attribute("version", String, "Service's version")
	})
})

var _ = Service("api", func() {
	Description("API of the service")

	HTTP(func() {
		Path("/api")
	})

	Method("info", func() {
		Description("Show info of the service")
		Payload(Empty)
		Result(Info)
		HTTP(func() {
			GET("/info")
			Response(StatusOK)
		})
	})
})
