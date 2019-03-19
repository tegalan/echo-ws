package main

import (
	"echo-ws/app"
	"fmt"
	"log"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(time.RFC3339) + " " + string(bytes))
}

func main() {
	app := app.App{}

	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	app.Initialize()

	app.Run()
}
