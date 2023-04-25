package main

import "github.com/izaakdale/distcache/internal/app"

func main() {
	app.Must(app.New()).Run()
}
