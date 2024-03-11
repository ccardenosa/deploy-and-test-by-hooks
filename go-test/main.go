package main

import "github.com/ccardenosa/deploy-and-test-by-hooks/go-test/backend"

var beConfig = backend.Config{
	ListenUri: "0.0.0.0:28081",
}

func main() {

	backend.StartBackend(beConfig)
}
