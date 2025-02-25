package user

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ebrickdev/ebrick"
	"github.com/ebrickdev/ebrick/logger"
	"github.com/ebrickdev/ebrick/module"
	"github.com/ebrickdev/ebrick/security/auth"
	"github.com/ebrickdev/ebrick/transport/http"
	pb "github.com/ebrickdev/example/user/proto"
	"google.golang.org/grpc"
)

type User struct {
}

// Dependencies implements module.Module.
func (c *User) Dependencies() []string {
	return []string{"dependency1"}
}

// Description implements module.Module.
func (c *User) Description() string {
	return "User module handles customer related operations."
}

// Id implements module.Module.
func (c *User) Id() string {
	return "customer-module"
}

// Initialize implements module.Module.
func (c *User) Initialize(ctx context.Context, options *module.Options) error {
	c.registerApiRoutes()
	// c.registerGRPCServices()
	return nil
}

// Name implements module.Module.
func (c *User) Name() string {
	return "User"
}

// Start implements module.Module.
func (c *User) Start(ctx context.Context) error {

	// Subscribe to an event
	ebrick.Logger.Info("Subscribing to user.started event")
	if err := ebrick.EventBus.Subscribe("user.started", func(ctx context.Context, event cloudevents.Event) {
		ebrick.Logger.Info("Received user.started event", logger.Any("event", event))
	}); err != nil {
		ebrick.Logger.Error("Failed to subscribe to user.started event", logger.Error(err))
	}

	// Publish an event
	ebrick.Logger.Info("Publishing user.started event")

	event := cloudevents.NewEvent()
	event.SetID("12")
	event.SetSource("user")
	event.SetType("user.started")
	event.SetSpecVersion("1.0")
	event.SetTime(time.Now())
	event.SetData(cloudevents.ApplicationJSON, map[string]any{
		"message": "User module started",
	})

	if result := ebrick.EventBus.Publish(context.Background(), "user.started", event); cloudevents.IsUndelivered(result) {
		ebrick.Logger.Error("Failed to publish user.started event", logger.Error(result))
	}
	// Perform start tasks here
	return nil

}

// Stop implements module.Module.
func (c *User) Stop(ctx context.Context) error {
	// Perform stop tasks here
	return nil
}

// Version implements module.Module.
func (c *User) Version() string {
	return "v1.0.0"
}

func (c *User) registerApiRoutes() {
	auMiddle := auth.NewAuthMiddleware(ebrick.AuthManager, ebrick.Logger)
	router := ebrick.HTTPServer.Engine().Group("protected")
	{
		router.Use(auMiddle.TokenAuth())
		router.GET("/customers", func(ctx *http.Context) {
			ctx.JSON(200, "GET /customers")
		})

		router.POST("/customers", func(ctx *http.Context) {
			ctx.JSON(200, "POST /customers")
		})
	}
}

func (c *User) registerGRPCServices() {
	// Register gRPC services here
	ebrick.GRPCServer.RegisterService(func(s *grpc.Server) {
		pb.RegisterUserServer(s, NewUserServiceServer(ebrick.Logger))
	})
}
