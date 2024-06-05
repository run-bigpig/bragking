package main

import (
	"embed"
	"github.com/run-bigpig/bragking/backend/app"
)

//go:embed frontend
var frontend embed.FS

func main() {
	app.Run(frontend)
}
