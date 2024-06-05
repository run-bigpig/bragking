package app

import (
	"context"
	"embed"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/run-bigpig/bragking/backend/config"
	"github.com/run-bigpig/bragking/backend/model"
	"github.com/run-bigpig/bragking/backend/router"
	"github.com/run-bigpig/bragking/backend/task"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var configPath = flag.String("f", "config.yaml", "config file path")

func Run(static embed.FS) {
	flag.Parse()
	ctx := context.TODO()
	config.Init(*configPath)
	model.Init(ctx)
	runTask()
	app := fiber.New()
	subFS, err := fs.Sub(static, "frontend")
	if err != nil {
		log.Fatalln(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(subFS),
	}))
	router.NewRouter(app)
	app.Listen(":8080")
}

func runTask() {
	args := os.Args
	if len(args) == 2 {
		if args[1] == "task" {
			task.NewTask(context.TODO()).StartUp()
		}
		os.Exit(0)
	}
}
