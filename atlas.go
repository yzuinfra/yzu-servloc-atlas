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
	serverPort := flag.String("port", "8000", "Port to run the server on")
	agentServerURL := flag.String("server_url", "", "Run atlas agent server, specify host server location.")
	agentName := flag.String("name", "Agent", "Agent name")
	flag.Parse()

	if !*runServer && *agentServerURL == "" {
		log.Fatalf("either -s for running the host or -server_url with -name for the agent is required")
	}

	if *runServer {
		log.Println("Running server")
		go host.RunWebHost()
		host.RunGRPCHost(*serverPort) // blocking
	}

	if *agentServerURL != "" {
		if *agentName == "" {
			log.Fatalf("agent name is required")
		}
		ip, err := ipify.GetIp()
		if err != nil {
			log.Fatalf("could not get ip: %v", err)
		} else {
			agent.RegisterAgent(*agentServerURL, &entities.Agent{IP: ip, Name: *agentName})
			log.Println("Agent registered")
		}
	}
}
