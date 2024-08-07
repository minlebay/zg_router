package grpc_client

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math"
	"sync"
	"time"
	"zg_router/pkg/message_v1"
)

type Client struct {
	Done              chan struct{}
	Logger            *zap.Logger
	Config            *Config
	GrpcClientPool    map[string]message.MessageRouterClient
	ConnectionPool    map[string]*grpc.ClientConn
	ActiveConnections map[string]int
	ConnectionsLock   sync.Mutex
	wg                sync.WaitGroup
}

func NewClient(logger *zap.Logger, config *Config) *Client {
	return &Client{
		Done:              make(chan struct{}),
		Logger:            logger,
		Config:            config,
		GrpcClientPool:    make(map[string]message.MessageRouterClient),
		ConnectionPool:    make(map[string]*grpc.ClientConn),
		ActiveConnections: make(map[string]int),
	}
}

func (c *Client) StartClient() {
	go func() {
		for _, server := range c.Config.ProcessingServersList {
			grpcTarget := fmt.Sprintf("%s", server)

			conn, err := grpc.NewClient(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				c.Logger.Fatal(err.Error())
			}

			c.ConnectionPool[server] = conn
			c.GrpcClientPool[server] = message.NewMessageRouterClient(conn)
			c.ActiveConnections[server] = 0
		}

		for {
			select {
			case <-c.Done:
				return
			default:
				continue
			}
		}
	}()
}

func (c *Client) StopClient() {
	for _, conn := range c.ConnectionPool {
		conn.Close()
	}

	c.Done <- struct{}{}
	c.Logger.Info("Client stopped")
}

func (c *Client) SendMessage(msg *message.Message, server string) {

	if srv, ok := c.GrpcClientPool[server]; !ok {
		c.Logger.Error("server not found")
		return
	} else {
		c.ConnectionsLock.Lock()
		c.ActiveConnections[server] = c.ActiveConnections[server] + 1
		c.Logger.Info("connections on servers", zap.Any("connections", c.ActiveConnections))
		c.ConnectionsLock.Unlock()
		resp, err := srv.ReceiveMessage(context.Background(), msg)
		if err != nil {
			c.Logger.Error("error sending message: ", zap.Error(err))
			return
		}

		// Simulate processing time
		time.Sleep(5 * time.Second)

		c.ConnectionsLock.Lock()
		c.ActiveConnections[server] = c.ActiveConnections[server] - 1
		c.ConnectionsLock.Unlock()
		c.Logger.Info("message sent", zap.Bool("success", resp.Success))
	}
}

func (c *Client) GetLeastLoadedServer() string {
	c.ConnectionsLock.Lock()
	defer c.ConnectionsLock.Unlock()

	if len(c.ActiveConnections) == 0 {
		return ""
	}

	minConnections := math.MaxInt32
	var leastLoadedServer string
	for server, connections := range c.ActiveConnections {
		if connections < minConnections {
			minConnections = connections
			leastLoadedServer = server
		}
	}

	return leastLoadedServer
}
