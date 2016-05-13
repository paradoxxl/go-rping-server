package Components

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Commands,
	},
	Route{
		"SimplePing",
		"GET",
		"/SimplePing/{ipaddr}/{times}",
		SimplePing,
	},
	Route{
		"SimplePing",
		"GET",
		"/SimplePing/{ipaddr}",
		SimplePing,
	},
}