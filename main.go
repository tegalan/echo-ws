package main

import "echo-ws/app"

func main() {
	app := app.App{}

	app.Initialize()

	app.Run()
}
