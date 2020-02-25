package main

import (
	"github.com/longguikeji/arkfbp-go-examples/flows/flow1"
	"github.com/longguikeji/arkfbp-go-examples/server"
)

// Routes ...
func Routes() []server.Route {
	var routes = []server.Route{
		server.Route{
			Name:    "flows.flow1",
			Pattern: "/hello",
			Handler: flow1.New(),
		},
	}

	return routes
}
