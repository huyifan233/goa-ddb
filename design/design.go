package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")
	HTTP(func(){
		Path("/resumes")
	})
	// Method describes a service method (endpoint)
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})

	Method("addresume", func(){

		Payload(ArrayOf(Resume))
		Result(ArrayOf(Int))
		HTTP(func(){
			POST("/addresume")
			MultipartRequest()
		})
	})

	Method("list", func(){

		Result(CollectionOf(StoredResume))
		HTTP(func(){
			GET("/getresumes")
			Response(StatusOK)
		})

	})


})



var Resume = Type("Resume", func(){
	Attribute("name", String, "Name in the resume")
	Attribute("experience", ArrayOf(Experience), "Experience section in the resume")
	Attribute("education", ArrayOf(Education), "Education section in the resume")
	Required("name")
})

var Experience = Type("Experience", func() {
	Attribute("company", String, "Name of the company")
	Attribute("role", String, "Name of the role in the company")
	Attribute("duration", Int, "Duration (in years) in the company")
	Required("company", "role", "duration")
})

var Education = Type("Education", func() {
	Attribute("institution", String, "Name of the institution")
	Attribute("major", String, "Major name")
	Required("institution", "major")
})



var StoredResume = ResultType("application/vnd.goa.resume", func(){

	TypeName("StoredResume")
	Attributes(func() {
		Extend(Resume)
		Attribute("id", Int, "ID of the resume")
		Attribute("created_at", String, "Time when resume was created")
		Required("id", "name", "experience", "education", "created_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("experience")
		Attribute("education")
		Attribute("created_at")
	})

})












