package host

import (
	"log"
	"net"
	"time"
	"yzuinfra/atlas/entities"
	"yzuinfra/atlas/host/regmap"
	"yzuinfra/atlas/protos"

	"google.golang.org/grpc"
)

type BiRPCServer struct {
	protos.BiRPCServer
}

func (s *BiRPCServer) RegisterAgent(stream protos.BiRPC_RegisterAgentServer) error {
	firstMsg, err := stream.Recv()
	if err != nil {
		return err
	}

	agentID := firstMsg.Service + "-" + firstMsg.Version + "-" + firstMsg.Hostname

	regmap.AddAgent(&entities.Agent{
		ID:       agentID,
		IP:       firstMsg.Ip,
		Hostname: firstMsg.Hostname,
		Service:  firstMsg.Service,
		Version:  firstMsg.Version,
	})
	log.Printf("Agent registered: %s %s", firstMsg.Ip, agentID)

	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				log.Printf("Agent %s disconnected", agentID)
				regmap.RemoveAgent(agentID)
				return
			}
		}
	}()

	for {
		if err := stream.Send(&protos.RegisterResponse{Registered: true}); err != nil {
			log.Printf("Failed to send keepalive to %s: %v", agentID, err)
			regmap.RemoveAgent(agentID)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func RunGRPCHost(host string, port string) {
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterBiRPCServer(s, &BiRPCServer{})
	log.Println("GRPC server starting")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
