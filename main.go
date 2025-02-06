package main

import (
	"context"

	"github.com/ebrickdev/ebrick"
	"github.com/ebrickdev/ebrick/example/user"
)

func main() {

	app := ebrick.NewApplication()

	app.RegisterModules(context.Background(), &user.User{})
	app.Start(context.Background())

}
