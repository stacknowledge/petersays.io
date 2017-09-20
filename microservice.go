package main

import "microservice/delivery"

func main() {
	application := new(delivery.Engine)
	application.Bootstrap()
}
