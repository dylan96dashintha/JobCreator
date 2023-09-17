package main

import (
	"github.com/JobCreator/app/container"
	"github.com/JobCreator/app/http"
)

func main() {
	container.Resolve()
	http.InitServer()
}
