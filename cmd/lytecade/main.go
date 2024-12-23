package main

import (
    "github.com/lytecade/lytecade/internal/routes"
    "fmt"
)

func main() {
    routes.RouteSites()
    routes.RouteInit()
    fmt.Println(routes.GetRouteSites())
    routes.RouteListen()
}

