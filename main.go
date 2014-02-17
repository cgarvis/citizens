package main

var config Config

func init() {
	config = LoadConfig()
}

func main() {
	RunServer()
}
