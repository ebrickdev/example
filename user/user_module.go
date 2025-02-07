package user

import (
	"context"
	"time"

	"github.com/ebrickdev/ebrick/logger"
	"github.com/ebrickdev/ebrick/messaging"
	"github.com/ebrickdev/ebrick/module"
	egrpc "github.com/ebrickdev/ebrick/transport/grpc"
	"github.com/ebrickdev/ebrick/transport/httpserver"
	pb "github.com/ebrickdev/example/user/proto"
	"google.golang.org/grpc"
)

type User struct {
	log      logger.Logger
	eventbus messaging.EventBus
}

// Dependencies implements module.Module.
func (c *User) Dependencies() []string {
	return []string{"dependency1", "dependency2"}
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
	// Perform initialization tasks here
	c.log = options.Logger
	c.eventbus = options.EventBus
	return nil
}

// Name implements module.Module.
func (c *User) Name() string {
	return "User"
}

// Start implements module.Module.
func (c *User) Start(ctx context.Context) error {

	// Subscribe to an event
	c.log.Info("Subscribing to user.started event")
	if err := c.eventbus.Subscribe("user.started", func(ctx context.Context, event messaging.Event) {
		c.log.Info("Received user.started event", logger.Any("event", event))
	}); err != nil {
		c.log.Error("Failed to subscribe to user.started event", logger.Error(err))
	}

	// Publish an event
	c.log.Info("Publishing user.started event")
	if err := c.eventbus.Publish(context.Background(), messaging.Event{
		ID:          "12",
		Source:      "user",
		SpecVersion: "1.0",
		Type:        "user.started",
		Data: map[string]any{
			"message": "User module started",
		},
		Time: time.Time{},
	}); err != nil {
		c.log.Error("Failed to publish user.started event", logger.Error(err))
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

func (c *User) RegisterRoutes(router httpserver.RouterGroup) {
	router.GET("/customers", func(ctx httpserver.Context) {
		ctx.JSON(200, "GET /customers")
	})

	router.POST("/customers", func(ctx httpserver.Context) {
		ctx.JSON(200, "POST /customers")
	})
}

func (c *User) RegisterGRPCServices(s egrpc.GRPCServer) {
	// Register gRPC services here
	s.RegisterService(func(s *grpc.Server) {
		pb.RegisterUserServer(s, NewUserServiceServer(c.log))
	})
}
