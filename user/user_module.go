package user

import (
	"context"

	pb "github.com/ebrickdev/ebrick/example/user/proto"
	"github.com/ebrickdev/ebrick/logger"
	"github.com/ebrickdev/ebrick/module"
	egrpc "github.com/ebrickdev/ebrick/transport/grpc"
	"github.com/ebrickdev/ebrick/transport/httpserver"
	"google.golang.org/grpc"
)

type User struct {
	log logger.Logger
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
	return nil
}

// Name implements module.Module.
func (c *User) Name() string {
	return "User"
}

// Start implements module.Module.
func (c *User) Start(ctx context.Context) error {
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
