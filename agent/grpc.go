package agent

import (
	"context"
	"log"
	"time"
	"yzuinfra/atlas/entities"
	"yzuinfra/atlas/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BiRPCClient struct {
	protos.BiRPCClient
}

func RegisterAgent(serverURL string, agent *entities.Agent) {
	conn, err := grpc.Dial(serverURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := protos.NewBiRPCClient(conn)
	stream, err := c.RegisterAgent(context.Background())
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}

	err = stream.Send(&protos.RegisterRequest{Ip: agent.IP, Hostname: agent.Hostname, Service: agent.Service, Version: agent.Version})
	if err != nil {
		log.Fatalf("failed to send registration: %v", err)
	}

	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				log.Printf("Server disconnected: %v", err)
				return
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		err := stream.Send(&protos.RegisterRequest{Ip: agent.IP, Hostname: agent.Hostname, Service: agent.Service, Version: agent.Version})
		if err != nil {
			log.Printf("Failed to send heartbeat: %v", err)
			return
		}
	}
}
