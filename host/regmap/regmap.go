package regmap

import (
	"yzuinfra/atlas/entities"
)

var AgentMap = make(map[string]*entities.Agent)

func AddAgent(agent *entities.Agent) {
	AgentMap[agent.ID] = agent
}

func RemoveAgent(id string) {
	delete(AgentMap, id)
}

func GetAgents() []*entities.Agent {
	agents := make([]*entities.Agent, 0, len(AgentMap))
	for _, agent := range AgentMap {
		agents = append(agents, agent)
	}
	return agents
}

func GetAgentsWithService(service string) []*entities.Agent {
	agents := make([]*entities.Agent, 0, len(AgentMap))
	for _, agent := range AgentMap {
		if agent.Service == service {
			agents = append(agents, agent)
		}
	}
	return agents
}

func GetAgentsFromHostname(hostname string) []*entities.Agent {
	agents := make([]*entities.Agent, 0, len(AgentMap))
	for _, agent := range AgentMap {
		if agent.Hostname == hostname {
			agents = append(agents, agent)
		}
	}
	return agents
}
