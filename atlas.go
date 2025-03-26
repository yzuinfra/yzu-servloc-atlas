package main

import (
	"flag"
	"log"
	"yzuinfra/atlas/agent"
	"yzuinfra/atlas/entities"
	"yzuinfra/atlas/host"

	"github.com/nettempo/go-ipify"
)

func main() {
	runServer := flag.Bool("s", false, "Run atlas host server")
	serverWebHost := flag.String("web_host", "0.0.0.0", "Host to listen on")
	serverWebPort := flag.String("web_port", "8000", "Port to run the web server on")
	serverGRPCPort := flag.String("grpc_port", "8001", "Port to run the gRPC server on")
	agentServerURL := flag.String("server_url", "127.0.0.1:8001", "Run atlas agent server, specify host server location.")
	agentHostname := flag.String("hostname", "Agent", "Agent name")
	agentService := flag.String("service", "atlas", "Agent service name")
	agentVersion := flag.String("version", "1.0", "Agent version")
	flag.Parse()

	if !*runServer && (*agentServerURL == "" || *agentHostname == "" || *agentService == "" || *agentVersion == "") {
		flag.PrintDefaults()
	}

	if *runServer {
		log.Println("Running server")
		host.RunWebHost(*serverWebHost, *serverWebPort)
		host.RunGRPCHost(*serverGRPCPort) // blocking
		log.Fatalln("Server stopped")     // should never happen
	}

	if *agentServerURL != "" {
		ip, err := ipify.GetIp()
		if err != nil {
			log.Fatalf("could not get ip: %v", err)
		} else {
			agent.RegisterAgent(*agentServerURL, &entities.Agent{
				IP:       ip,
				Hostname: *agentHostname,
				Service:  *agentService,
				Version:  *agentVersion,
			})
			log.Println("Agent registered")
		}
	}
}
