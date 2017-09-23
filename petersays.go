package main

import "github.com/stacknowledge/petersays.io/component"

func main() {
	microservice := new(component.Engine)
	microservice.Boot()
}
