package main

import (
	"github.com/ccardenosa/front-back-app/backend"
)

var beConfig = backend.Config{
	ListenUri:        "0.0.0.0:28081",
}

func main() {

	backend.StartBackend(beConfig)
}
