package regmap

import (
	"yzuinfra/atlas/entities"
)

var AgentMap = make(map[string]*entities.Agent)

func AddAgent(id string, name string) {
	AgentMap[id] = &entities.Agent{IP: id, Name: name}
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
