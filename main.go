package main

import (
	"context"

	"github.com/ebrickdev/ebrick"
	"github.com/ebrickdev/example/user"
	"github.com/ebrickdev/extensions/v1/db/postgresql"
	"github.com/ebrickdev/extensions/v1/logger/zap"
	redisstream "github.com/ebrickdev/extensions/v1/messaging/redis-stream"
)

func main() {

	app := ebrick.NewApplication(
		ebrick.WithEventBus(redisstream.Init()),
		ebrick.WithLogger(zap.Init()),
		ebrick.WithDB(postgresql.Init()),
	)

	app.RegisterModules(context.Background(), &user.User{})
	app.Start(context.Background())

}
