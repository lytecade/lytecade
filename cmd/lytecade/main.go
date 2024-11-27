package main

import (
	"github.com/lytecade/lytecade/internal/routes"
)

func main() {
    routes.RouteInit()
    routes.RouteSites()
    routes.RouteListen()
}

