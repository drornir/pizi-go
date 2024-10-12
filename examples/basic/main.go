package main

import (
	"fmt"

	"github.com/drornir/pizi-go/app"
)

func main() {
	app := newApp()
	fmt.Println(app)
}

func newApp() app.App {
	return app.App{}
}
