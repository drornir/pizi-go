package main

import (
	"context"

	"github.com/drornir/pizi-go/app"
)

func main() {
	app, err := app.NewLocalLibSQLApp("./bin/db/")
	if err != nil {
		panic(err)
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
