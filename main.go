package main

import (
	"log"
	"lookupStock/components"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &components.Hello{})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("lookup-stock"),
	})

	if err != nil {
		log.Fatal(err)
	}
}
